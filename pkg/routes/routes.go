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

	// Routes for project information
	/**
	 * @swagger
	 * /projects:
	 *   get:
	 *     summary: Get all projects
	 *     tags:
	 *       - Projects
	 *     responses:
	 *       200:
	 *         description: Success
	 */
	r.Get("/projects", controller.GetProject)

	/**
	 * @swagger
	 * /projects:
	 *   post:
	 *     summary: Create a new project
	 *     tags:
	 *       - Projects
	 *     responses:
	 *       200:
	 *         description: Success
	 */
	r.Post("/projects", controller.InsertProject)

	/**
	 * @swagger
	 * /projects/{id}:
	 *   get:
	 *     summary: Get a project by ID
	 *     tags:
	 *       - Projects
	 *     parameters:
	 *       - in: path
	 *         name: id
	 *         required: true
	 *         description: ID of the project
	 *         schema:
	 *           type: integer
	 *     responses:
	 *       200:
	 *         description: Success
	 */
	r.Get("/projects/{id}", controller.GetProjectByID)

	/**
	 * @swagger
	 * /projects/{id}:
	 *   put:
	 *     summary: Update a project by ID
	 *     tags:
	 *       - Projects
	 *     parameters:
	 *       - in: path
	 *         name: id
	 *         required: true
	 *         description: ID of the project
	 *         schema:
	 *           type: integer
	 *     responses:
	 *       200:
	 *         description: Success
	 */
	r.Put("/projects/{id}", controller.UpdateProject)

	/**
	 * @swagger
	 * /projects/{id}:
	 *   delete:
	 *     summary: Delete a project by ID
	 *     tags:
	 *       - Projects
	 *     parameters:
	 *       - in: path
	 *         name: id
	 *         required: true
	 *         description: ID of the project
	 *         schema:
	 *           type: integer
	 *     responses:
	 *       200:
	 *         description: Success
	 */
	r.Delete("/projects/{id}", controller.DeleteProject)

	// Routes for FileUpload
	/**
	 * @swagger
	 * /upload:
	 *   post:
	 *     summary: Upload a file
	 *     tags:
	 *       - FileUpload
	 *     responses:
	 *       200:
	 *         description: Success
	 */
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
