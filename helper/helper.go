package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response_struct struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func Response(res Response_struct, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Erro ao fazer marshal no json ", err)
		Response(Response_struct{Error: "something went wrong"}, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		fmt.Println("erro ao enviar resposta: ", err)
		return
	}
}
