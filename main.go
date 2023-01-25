package main

import (
	"fmt"
	"log"
	"net/http"
	"webapi/configs"
	"webapi/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Configuracoes Carregadas")

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)
	log.Println("Rotas Criadas")

	log.Println("Servidor rodando na porta 9000")
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
