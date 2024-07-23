package router

import (
	controller "github.com/Aviral0702/mongoApi/contoller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/create", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/update/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/delete/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/delete-all", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
