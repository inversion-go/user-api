package handler

import (
	"encoding/json"
	"github.com/inversion-go/alejoab12/user-api/api/services"

	"net/http"
)

func UserById(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.FindById("2"))

}
