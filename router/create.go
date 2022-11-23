package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gravitalia/AirBack/model"
)

func Post(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(model.Message{Error: true, Message: "test"})

	fmt.Fprintf(w, string(res))
}
