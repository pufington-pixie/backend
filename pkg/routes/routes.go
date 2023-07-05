package routers

import (
	"log"
	"net/http"

	// Import the generated Swagger docs
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/pdrum/swagger-automation/docs"
	_ "github.com/pufington-pixie/haver/docs"
	controller "github.com/pufington-pixie/haver/pkg/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetRoutes sets up the routing for the API
func SetRoutes() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER","Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		})
	r.Use(cors.Handler) 

	
	r.Get("/api/projects", controller.GetProject)

	
	r.Post("/api/projects", controller.InsertProject)

	
	r.Get("/api/projects/{id}", controller.GetProjectByID)

	
	r.Put("/api/projects/{id}", controller.UpdateProject)


	r.Delete("/api/projects/{id}", controller.DeleteProject)

	
	r.Post("/api/upload/{id}", controller.UploadHandler)

	r.Get("/api/data/{id}",controller.GetData)

	// Swagger UI route
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), 
	))


	// Start the HTTP server
	port := ":8080"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
