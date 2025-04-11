package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ethea2/Distributed-Fault-Tolerance/gateway/internal/handlers"
	"github.com/Ethea2/Distributed-Fault-Tolerance/gateway/internal/proxy"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("DB_HOST") == "" {
		_ = godotenv.Load()
	}
	port := os.Getenv("PORT")
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	courseServiceURL := os.Getenv("COURSE_SERVICE_URL")
	gradeServiceURL := os.Getenv("GRADE_SERVICE_URL")

	if authServiceURL == "" || courseServiceURL == "" || gradeServiceURL == "" {
		log.Fatal("Service URLs must be provided")
	}

	authProxy := proxy.NewServiceProxy(authServiceURL)
	courseProxy := proxy.NewServiceProxy(courseServiceURL)
	gradeProxy := proxy.NewServiceProxy(gradeServiceURL)

	authHandler := handlers.NewAuthHandler(authProxy)
	courseHandler := handlers.NewCourseHandler(courseProxy)
	gradeHandler := handlers.NewGradeHandler(gradeProxy)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gateway service is healthy"))
	})

	r.Route("/api", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHandler.Login)
		})

		r.Route("/course", func(r chi.Router) {
			r.Get("/", courseHandler.GetCourses)
			r.Get("/available", courseHandler.GetAvailableCourses)
			r.Post("/", courseHandler.CreateCourse)
			r.Post("/enroll", courseHandler.EnrollCourse)
		})

		r.Route("/grade", func(r chi.Router) {
			r.Get("/", gradeHandler.GetGrades)
			r.Get("/all_students", gradeHandler.GetStudentGrades)
			r.Post("/", gradeHandler.GradeStudent)
		})
	})

	fmt.Println("Gateway service starting on :" + port)
	http.ListenAndServe(":"+port, r)
}
