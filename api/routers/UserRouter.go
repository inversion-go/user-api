package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/inversion-go/alejoab12/user-api/api/handler"
	"log"
	"net/http"
)

func RouterUserHandler() {
	fmt.Println("Entro a esta cagada")
	userRouter := mux.NewRouter().StrictSlash(true)
	userRouter.HandleFunc("/user/byId", handler.UserById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", userRouter))
}
