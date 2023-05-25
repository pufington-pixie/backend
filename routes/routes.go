package routers

import (
	"net/http"

	"example.com/controllers"
	"github.com/go-chi/chi"
	
)

// SetRoutes sets up the routing for the API
func SetRoutes() {
  r := chi.NewRouter()
  

  
  // Routes
  r.Get("/projects", controller.GetProject)
  r.Post("/project", controller.InsertProject)
  r.Get("/project/{id}", controller.GetProjectByID)
  r.Put("/project/{id}", controller.UpdateProject)
  r.Delete("/project/{id}", controller.DeleteProject)
  
  // Serve
  http.ListenAndServe(":1234", r)
}
