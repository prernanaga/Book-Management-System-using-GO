package routes

import (
	"bookmanagement/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// For creating a new instance
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	// For getting all the instances
	router.HandleFunc("/book/", controllers.GetAllBook).Methods("GET")
	// For getting only one specified instance
	router.HandleFunc("/book/{bookId}", controllers.GetBookbyID).Methods("GET")
	// Updating a instance
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// Deleting an instance
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
