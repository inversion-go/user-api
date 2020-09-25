package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/inventario-go/alejoab12/user-api/api/models"
	"github.com/inventario-go/alejoab12/user-api/api/services"

	"net/http"
)

func UserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	id, ok := mux.Vars(r)["id"]
	fmt.Println(id)
	if !ok {
		w.WriteHeader(404)
	} else {
		fmt.Println(id)
		json.NewEncoder(w).Encode(services.FindById(id))
	}

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	services.CreateUser(&user)
	w.WriteHeader(201)
}
