package routers

import (
	"net/http"

	controller "example.com/controllers"
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	// _ "example.com/docs" // Import the generated Swagger docs
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
	 * /project:
	 *   post:
	 *     summary: Create a new project
	 *     tags:
	 *       - Projects
	 *     responses:
	 *       200:
	 *         description: Success
	 */
	r.Post("/project", controller.InsertProject)

	/**
	 * @swagger
	 * /project/{id}:
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
	r.Get("/project/{id}", controller.GetProjectByID)

	/**
	 * @swagger
	 * /project/{id}:
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
	r.Put("/project/{id}", controller.UpdateProject)

	/**
	 * @swagger
	 * /project/{id}:
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
	r.Delete("/project/{id}", controller.DeleteProject)

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
	r.Post("/upload", controller.UploadHandler)

	// Swagger UI route
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Serve
	http.ListenAndServe(":8080", r)
}
