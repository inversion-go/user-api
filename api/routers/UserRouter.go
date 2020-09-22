package routers

import (
	"github.com/gorilla/mux"
	"github.com/inversion-go/alejoab12/user-api/api/handler"
	"log"
	"net/http"
)

func RouterUserHandler() {
	userRouter := mux.NewRouter().StrictSlash(true)
	userRouter.HandleFunc("/user/findById", handler.UserById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", userRouter))
}
