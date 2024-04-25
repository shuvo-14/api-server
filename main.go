package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/shuvo-14/api-server/api"
	"github.com/shuvo-14/api-server/auth"
	"github.com/shuvo-14/api-server/db"
	"log"
	"net/http"
)

func main() {
	db.Init()
	r := chi.NewRouter()
	r.Post("/login", auth.Login)
	r.Post("/logout", auth.LogOut)
	r.Group(func(chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Get("/", api.GetAllBooks)
			r.Get("/{id}", api.GetOneBook)
			r.Group(func(r chi.Router) {
				// need to add authentication
				r.Use(jwtauth.Verifier(db.TokenAuth))
				r.Use(jwtauth.Authenticator(db.TokenAuth))

				r.Post("/", api.NewBook)
				r.Put("/{id}", api.UpdateBook)
				r.Delete("/{id}", api.DeleteBook)
			})

		})
		r.Route("/authors", func(r chi.Router) {
			r.Get("/", api.GetAllAuthors)
			r.Get("/{id}", api.GetOneAuthor)
		})
	})

	fmt.Println("Listening and Serving to 8000")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
