package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Could not load .env file")
	}

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Could not get Port from .env file")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Links"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Post("/signUp",HandleRegisteredUser)
	v1Router.Post("/login",login)
	v1Router.Post("/logout", Auth(AuthHandler(HandleUserLogout)))

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	log.Print("Server is running on PORT:",portString)

	err = srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}