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

// pongHandler ::: controlador del servicio pong
func pongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("pongHandler - request")
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "message": "pong" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

// respondWithJSON :: agrega la cabecera de JSON a la respuesta
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// initialiseRoutes ::: contiene el controlador y path de servicio ping
func (app *App) initialiseRoutes() {
	log.Println("start - pong service up")
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/pong", pongHandler)
}

// run ::: levanta el servidor en puerto indicado
func (app *App) run() {
	log.Fatal(http.ListenAndServe(":8081", app.Router))
}

// main ::: inicio del servicio ping
func main() {
	app := App{}
	app.initialiseRoutes()
	app.run()
}
