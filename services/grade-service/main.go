package main

import (
	"context"
	"database/sql"
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

	r.Group(func(r chi.Router) {
		r.Use(authMid.AuthMiddleware)

		r.Get("/student_grades", func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value("custom_claims").(models.User)

			if user.Role != "student" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(models.GeneralResponse{
					Message: "Cannot view grades as you are a faculty!",
				})
			}

			conn, err := database.ConnectDB()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			sqlQuery := `
				SELECT 
						c.course_name,
						c.course_code,
						c.available,
						e.status AS enrollment_status,
						e.grade
				FROM 
						courses.enrollments e
				JOIN 
						courses.courses c ON e.course_id = c.id
				WHERE 
						e.user_id = $1 AND e.status = 'done'
				ORDER BY 
						c.course_name;
			`

			rows, queryErr := conn.Query(context.Background(), sqlQuery, user.ID)

			if queryErr != nil {
				http.Error(w, queryErr.Error(), http.StatusBadRequest)
				return
			}

			var userGrades []models.Grade
			var userGrade models.Grade
			_, forErr := pgx.ForEachRow(rows, []any{&userGrade.CourseName, &userGrade.CourseCode, &userGrade.Available, &userGrade.EnrollmentStatus, &userGrade.Grade}, func() error {
				userGrades = append(userGrades, userGrade)
				return nil
			})

			if forErr != nil {
				http.Error(w, forErr.Error(), http.StatusBadRequest)
				return
			}

			json.NewEncoder(w).Encode(models.GradeResponse{
				Grades: userGrades,
			})
		})

		r.Get("/faculty_grades", func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value("custom_claims").(models.User)

			if user.Role != "faculty" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(models.GeneralResponse{
					Message: "You can't view your classmates' grades!",
				})
				return
			}

			conn, err := database.ConnectDB()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			sqlQuery := `
				SELECT 
						u.id AS user_id,
						u.username,
						c.course_name,
						c.course_code,
						c.available,
						e.status AS enrollment_status,
						e.grade
				FROM 
						auth.users u
				JOIN 
						courses.enrollments e ON u.id = e.user_id
				JOIN 
						courses.courses c ON e.course_id = c.id
				ORDER BY 
						u.id,
						c.course_name;
			`

			rows, queryErr := conn.Query(context.Background(), sqlQuery)

			if queryErr != nil {
				http.Error(w, queryErr.Error(), http.StatusBadRequest)
				return
			}
			var studentGrades []models.StudentGrades

			var studentId, currentStudent int
			var username, course_name, course_code, enrollment_status, currentUsername string
			var available bool
			var grade sql.NullFloat64
			var grades []models.Grade

			_, forErr := pgx.ForEachRow(rows, []any{&studentId, &username, &course_name, &course_code, &available, &enrollment_status, &grade}, func() error {
				if currentStudent != studentId && currentStudent != 0 {
					userGrade := models.StudentGrades{
						Username: currentUsername,
						UserID:   currentStudent,
						Grades:   grades,
					}
					studentGrades = append(studentGrades, userGrade)
					grades = []models.Grade{}
				}

				currentStudent = studentId
				currentUsername = username

				newGrade := -1.0
				if grade.Valid {
					newGrade = grade.Float64
				}

				grades = append(grades, models.Grade{
					CourseName:       course_name,
					CourseCode:       course_code,
					Available:        available,
					EnrollmentStatus: enrollment_status,
					Grade:            float32(newGrade),
				})

				return nil
			})

			if forErr != nil {
				http.Error(w, forErr.Error(), http.StatusBadRequest)
				return
			}

			json.NewEncoder(w).Encode(models.StudentGradesResponse{
				StudentGrades: studentGrades,
			})
		})

		r.Post("/grade", func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value("custom_claims").(models.User)

			if user.Role != "faculty" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(models.GeneralResponse{
					Message: "Please don't try to cheat with your grades.",
				})
			}

			var gradeChangeReq models.ChangeGradeRequest
			json.NewDecoder(r.Body).Decode(&gradeChangeReq)

			sqlQuery := `
				UPDATE courses.enrollments
				SET grade = $1 
				WHERE user_id = $2
				AND course_id = (
						SELECT id 
						FROM courses.courses 
						WHERE course_code = $3 
				);
			`

			conn, err := database.ConnectDB()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			_, err = conn.Exec(context.Background(), sqlQuery, gradeChangeReq.Grade, gradeChangeReq.UserID, gradeChangeReq.CourseCode)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			json.NewEncoder(w).Encode(models.GeneralResponse{
				Message: "Successfully uploaded student grades",
			})
		})
	})

	fmt.Println("Grade service starting on :8080")
	http.ListenAndServe(":"+port, r)
}
