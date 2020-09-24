package routers

import (
	"github.com/gorilla/mux"
	"github.com/inventario-go/alejoab12/user-api/api/handler"
	"github.com/inventario-go/alejoab12/user-api/api/repository"
	"log"
	"net/http"
)

func RouterUserHandler() {
	userRouter := mux.NewRouter().StrictSlash(true)
	userRouter.HandleFunc("/api/user/{id}", handler.UserById).Methods("GET")
	userRouter.HandleFunc("/api/user", handler.CreateUser).Methods("POST")
	repository.InitialDB()
	log.Fatal(http.ListenAndServe(":8082", userRouter))
}
