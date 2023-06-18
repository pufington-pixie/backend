package routers

import (
	"log"
	"net/http"

	// Import the generated Swagger docs
	"github.com/go-chi/chi"
	controller "github.com/pufington-pixie/haver/pkg/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/pdrum/swagger-automation/docs" 
	_ "github.com/pufington-pixie/haver/docs"
	
	
)

// SetRoutes sets up the routing for the API
func SetRoutes() {
	r := chi.NewRouter()

	
	r.Get("/projects", controller.GetProject)

	
	r.Post("/projects", controller.InsertProject)

	
	r.Get("/projects/{id}", controller.GetProjectByID)

	
	r.Put("/projects/{id}", controller.UpdateProject)


	r.Delete("/projects/{id}", controller.DeleteProject)

	
	r.Post("/upload/{id}", controller.UploadHandler)

	// Swagger UI route
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))


	// Start the HTTP server
	port := ":8080"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
