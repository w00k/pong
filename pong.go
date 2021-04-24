package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App export
type App struct {
	Router *mux.Router
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("pongHandler - request")
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "message": "pong" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (app *App) initialiseRoutes() {
	log.Println("start - pong service up")
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/pong", pongHandler)
}

func (app *App) run() {
	log.Fatal(http.ListenAndServe(":8081", app.Router))
}

func main() {
	app := App{}
	app.initialiseRoutes()
	app.run()
}
