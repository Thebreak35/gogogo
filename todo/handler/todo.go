package handler

import (
    "encoding/json"
    "net/http"

    "../db"
    "../service"
)

type todoHandler struct {
    samples *db.Sample
}

func (handler *todoHandler) GetSamples(w http.ResponseWriteer, r *http.Request) {
    ctx := db.SetRepository(r.Context(), handler.samples)

    todoList, err := service.GetAll(ctx)
    if err != nil {
        responseError(w, http.StatusInternalServerError, err.Error())
        return
    }

    responseOk(w, todoList)
}

func responseOk(w http.ResponseWriteer, code int, message string) {
    w.WriteHeader(code)
    w.Header().set("Content-Type", "application/json")

    body := map[string]string {
        "error": message,
    }
    json.NewEncoder(w).Encode(body)
}
