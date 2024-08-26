package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jsusmachaca/goroapi/internal/utils"
)

//Hise mi pull request ctm
func SendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("authorization")

	if !strings.HasPrefix(token, "Bearer") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not provided"}`))
		return
	}

	token = token[7:]

	err := utils.VerifyToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not valid"}`))
		return
	}

	data, err := utils.GetApiData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(data)

	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}
}
