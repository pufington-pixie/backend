package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pufington-pixie/haver/pkg/database"
	"github.com/pufington-pixie/haver/pkg/models"
)

// InsertProject inserts a new project.
// @Summary Insert a new project
// @Description Insert a new project into the database
// @Tags projects
// @Accept json
// @Produce json
// @Param project body models.Project true "Project to be inserted"
// @Success 200 {object} models.Project
// @Failure 404 {} string "User not found"
// @Failure 500 {object} models.Response
// @Router /api/projects [post]
func InsertProject(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	// Parse JSON request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
		return
	}

	// Define a struct to hold the JSON data
	var project models.Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		log.Print(err)
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
		return
	}

	db := database.GetDB()
	defer db.Close()

	// Prepare the SQL statements
	projectQuery := "INSERT INTO projects (id, name, title, date, sapnumber, notes, branchId, statusId, serviceId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	serviceQuery := "INSERT INTO services (id, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE name = ?"

	// Insert or update the service in the services table
	_, err = db.Exec(serviceQuery, project.Service.ID, project.Service.Name, project.Service.Name)
	if err != nil {
		log.Print(err)
		response.Status = 500
		response.Message = "Internal Server Error"
		json.NewEncoder(w).Encode(response)
		return
	}

	// Insert project into the projects table
	_, err = db.Exec(projectQuery, project.ID,project.Name, project.Title, project.Date, project.SAPNumber, project.Notes, project.BranchID, project.StatusID, project.Service.ID)
	if err != nil {
		log.Print(err)
		response.Status = 500
		response.Message = "Internal Server Error"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = 200
	response.Message = "Insert data successfully"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// UpdateProject updates an existing project.
// @Summary Update an existing project
// @Description Update an existing project in the database
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body models.Project true "Project object to be updated"
// @Success 200 {object} models.Project
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/projects/{id} [put]
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := database.GetDB()
	defer db.Close()

	// Read JSON request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		response.Status = 400
		response.Message = "Bad Request"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Parse JSON request body
	var project models.Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		log.Fatal(err.Error())
		response.Status = 400
		response.Message = "Bad Request"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Prepare the SQL statements
	projectQuery := "UPDATE projects SET  name = ?, title = ?, sapnumber = ?, notes = ?, branchId = ?, statusId = ?, serviceId = ? WHERE id = ?"
	serviceQuery := "INSERT INTO services (id, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE name = ?"

	// Update project data in the database
	_, err = db.Exec(projectQuery,project.Name, project.Title, project.SAPNumber, project.Notes, project.BranchID, project.StatusID, project.Service.ID, project.ID)
	if err != nil {
		log.Print(err)
		response.Status = 500
		response.Message = "Internal Server Error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Update or insert the service data in the database
	_, err = db.Exec(serviceQuery, project.Service.ID, project.Service.Name, project.Service.Name)
	if err != nil {
		log.Print(err)
		response.Status = 500
		response.Message = "Internal Server Error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = 200
	response.Message = "Update data successfully"
	fmt.Print("Update data in the database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// GetProject returns the list of projects.
// @Summary Get the list of projects
// @Description Get the list of projects from the database
// @Tags projects
// @Produce json
// @Success 200 {object} models.Project
// @Failure 500 {object} models.Response
// @Router /api/projects [get]
func GetProject(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	var arrProject []models.Project

	db := database.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT p.id, p.name, p.title,  p.sapnumber, p.notes, p.branchid, p.statusid, s.id, s.name " +
		"FROM projects p " +
		"JOIN services s ON p.serviceid = s.id")
	if err != nil {
		log.Print(err)
		response.Status = 500
		response.Message = "Internal Server Error"
		json.NewEncoder(w).Encode(response)
		return
	}

	for rows.Next() {
		var project models.Project
		var service models.Service
		err = rows.Scan(&project.ID, &project.Name,&project.Title, &project.SAPNumber, &project.Notes, &project.BranchID, &project.StatusID, &service.ID, &service.Name)

		if err != nil {
			log.Fatal(err.Error())
			response.Status = 500
			response.Message = "Internal Server Error"
			json.NewEncoder(w).Encode(response)
			return
		}

		project.Service = service
		arrProject = append(arrProject, project)
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrProject

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// GetProjectByID returns a specific project by its ID.
// @Summary Get a project by ID
// @Description Get a project from the database by its ID
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} models.Project
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/projects/{id} [get]
func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)

		return
	}

	db := database.GetDB()
	defer db.Close()

	project := models.Project{}
	err = db.QueryRow("SELECT p.id, p.name, p.title, p.date, p.sapnumber, p.notes, p.branchId, p.statusId, p.serviceId, s.name FROM projects p JOIN services s ON p.serviceId = s.id WHERE p.id = ?", id).
		Scan(&project.ID, &project.Name,&project.Title, &project.Date, &project.SAPNumber, &project.Notes, &project.BranchID, &project.StatusID, &project.Service.ID, &project.Service.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No Project with that ID.")
			response.Status = 404
			response.Message = "Project not found"
		} else {
			log.Fatal(err.Error())
			response.Status = 500
			response.Message = "Internal Server Error"
		}
	} else {
		response.Status = 200
		response.Message = "Success"
		response.Data = project
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// DeleteProject deletes a project by its ID.
// @Summary Delete a project by ID
// @Description Delete a project from the database by its ID
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/projects/{id} [delete]
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
		return
	}

	db := database.GetDB()
	defer db.Close()

	_, err = db.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		log.Print(err)
		response.Status = 500
		response.Message = "Internal Server Error"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = 200
	response.Message = "Project deleted successfully"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
