package api

import (
    "net/http"
    
    "postwoman/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    
    e := handlers.ConfigGlobalHandler()

    e = handlers.TemplateHandler()
    e = handlers.UserHandler()
    e = handlers.RequestHandler()

    e.ServeHTTP(w, r)
}
