package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/database"
	authMid "github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/middleware"
	"github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("DB_HOST") == "" {
		_ = godotenv.Load()
	}

	port := os.Getenv("PORT")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Grade service is healthy"))
	})

	r.Get("/courses", func(w http.ResponseWriter, r *http.Request) {
		var courses []models.Course

		conn, err := database.ConnectDB()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rows, _ := conn.Query(context.Background(), "SELECT id, course_name, course_code, course_description, available FROM courses.courses")
		var course models.Course
		_, err = pgx.ForEachRow(rows, []any{&course.ID, &course.CourseName, &course.CourseCode, &course.CourseDescription, &course.Availabie}, func() error {
			courses = append(courses, course)
			return nil
		})

		json.NewEncoder(w).Encode(models.CourseResponse{
			Courses: courses,
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(authMid.AuthMiddleware)

		r.Get("/available_courses", func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value("custom_claims").(models.User)

			var courses []models.Course

			conn, err := database.ConnectDB()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			sqlQuery := `
				SELECT c.*
				FROM courses.courses c
				WHERE c.available = TRUE
				AND NOT EXISTS (
						SELECT 1
						FROM courses.enrollments e
						WHERE e.course_id = c.id
						AND e.user_id = $1 
				)
				ORDER BY c.course_code;
			`

			rows, _ := conn.Query(context.Background(), sqlQuery, user.ID)

			var course models.Course

			_, err = pgx.ForEachRow(rows, []any{&course.ID, &course.CourseName, &course.CourseCode, &course.CourseDescription, &course.Availabie}, func() error {
				courses = append(courses, course)
				return nil
			})

			json.NewEncoder(w).Encode(models.CourseResponse{
				Courses: courses,
			})
		})

		r.Post("/create_course", func(w http.ResponseWriter, r *http.Request) {
			courseReq := &models.Course{}

			err := json.NewDecoder(r.Body).Decode(&courseReq)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			conn, err := database.ConnectDB()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			user := r.Context().Value("custom_claims").(models.User)

			if user.Role == "student" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(models.GeneralResponse{
					Message: "Cannot create a course as a student.",
				})
				return
			}

			_, execErr := conn.Exec(context.Background(), `INSERT INTO courses.courses (course_name, course_code, course_description, available) VALUES ($1, $2, $3, $4)`, courseReq.CourseName, courseReq.CourseCode, courseReq.CourseDescription, courseReq.Availabie)

			if execErr != nil {
				http.Error(w, execErr.Error(), http.StatusBadRequest)
				return
			}

			json.NewEncoder(w).Encode(&models.GeneralResponse{
				Message: "Successfully created a new course!",
			})
		})

		r.Post("/enroll_course", func(w http.ResponseWriter, r *http.Request) {
			courseEnrollmentReq := &models.CourseEnrollRequest{}

			json.NewDecoder(r.Body).Decode(&courseEnrollmentReq)

			user := r.Context().Value("custom_claims").(models.User)

			conn, err := database.ConnectDB()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if user.Role == "faculty" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(models.GeneralResponse{
					Message: "Cannot enroll as a faculty...",
				})
				return
			}

			var course models.Course

			scanerr := conn.QueryRow(context.Background(), `SELECT id, course_name, course_code, course_description, available FROM courses.courses WHERE id=$1`, courseEnrollmentReq.CourseID).Scan(
				&course.ID,
				&course.CourseName,
				&course.CourseCode,
				&course.CourseDescription,
				&course.Availabie,
			)

			if scanerr != nil {
				http.Error(w, scanerr.Error(), http.StatusBadRequest)
				return
			}

			if course.Availabie == false {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(models.GeneralResponse{
					Message: "Cannot enroll an unavailable course!",
				})
			}

			_, execErr := conn.Exec(context.Background(), `INSERT INTO courses.enrollments (course_id, user_id, status) VALUES ($1, $2, $3)`, course.ID, user.ID, "enrolled")
			if execErr != nil {
				http.Error(w, execErr.Error(), http.StatusBadRequest)
				return
			}

			json.NewEncoder(w).Encode(models.GeneralResponse{
				Message: "Successfully enrolled!",
			})
		})
	})

	fmt.Println("Course service starting on :8080")
	http.ListenAndServe(":"+port, r)
}
