package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/auth"
	"github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/database"
	"github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Grade service is healthy"))
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		loginReq := &models.LoginModel{}

		conn, connErr := database.ConnectDB()

		if connErr != nil {
			panic(connErr)
		}

		err := json.NewDecoder(r.Body).Decode(&loginReq)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = conn.QueryRow(context.Background(), "SELECT id, username, role FROM auth.users WHERE password=$1", loginReq.Password).Scan(&user.ID, &user.Username, &user.Role)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, tokenerr := auth.SignToken(user)
		if tokenerr != nil {
			http.Error(w, tokenerr.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(models.ResponseToken{
			Token:   token,
			Message: "Success",
		})
	})

	fmt.Println("auth service starting on :8080")

	http.ListenAndServe(":"+port, r)
}
