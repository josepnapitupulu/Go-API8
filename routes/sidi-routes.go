package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Sidi/controllers"
)

var RegisterSidisRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/sidi/", controllers.CreateSidi).Methods("POST")
	router.HandleFunc("/sidi/", controllers.GetSidi).Methods("GET")
	router.HandleFunc("/sidi/{sidiId}", controllers.GetSidiById).Methods("GET")
	router.HandleFunc("/sidi/{sidiId}", controllers.UpdateSidi).Methods("PUT")
	router.HandleFunc("/sidi/{sidiId}", controllers.DeleteSidi).Methods("DELETE")
}
