package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"log"
	"net/http"

	"example.com/database"
	"example.com/models"
	"github.com/go-chi/chi"
)

// GetDB returns a database object
func GetDB() (*sql.DB) {
  db,err := database.ConnectDB()
  if err!=nil{
      log.Fatal(err)
  }
  return db
}
//insert project
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

	db := GetDB()
	defer db.Close()

	// Prepare the SQL statements
	projectQuery := "INSERT INTO projects (title, date, sapnumber, notes, branchId, statusId, serviceId) VALUES (?, ?, ?, ?, ?, ?, ?)"
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
	_, err = db.Exec(projectQuery, project.Title, project.Date, project.SAPNumber, project.Notes, project.BranchID, project.StatusID, project.Service.ID)
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


func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := GetDB()
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
	projectQuery := "UPDATE projects SET title = ?, sapnumber = ?, notes = ?, branchId = ?, statusId = ?, serviceId = ? WHERE id = ?"
	serviceQuery := "INSERT INTO services (id, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE name = ?"

	// Update project data in the database
	_, err = db.Exec(projectQuery, project.Title, project.SAPNumber, project.Notes, project.BranchID, project.StatusID, project.Service.ID, project.ID)
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

// GetProject = Select Project API
func GetProject(w http.ResponseWriter, r *http.Request) {
  var response models.Response
  var arrProject []models.Project

  db := GetDB()
  defer db.Close()

  rows, err := db.Query("SELECT p.id, p.title,  p.sapnumber, p.notes, p.branchid, p.statusid, s.id, s.name "+
      "FROM projects p "+
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
      err = rows.Scan(&project.ID, &project.Title, &project.SAPNumber, &project.Notes, &project.BranchID, &project.StatusID, &service.ID, &service.Name)

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
func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	idStr := chi.URLParam(r, "id")
  id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		
		return
	}

	db := GetDB()
	defer db.Close()

	project := models.Project{}
	err = db.QueryRow("SELECT p.id, p.title, p.date, p.sapnumber, p.notes, p.branchId, p.statusId, p.serviceId, s.name FROM projects p JOIN services s ON p.serviceId = s.id WHERE p.id = ?", id).
		Scan(&project.ID, &project.Title, &project.Date, &project.SAPNumber, &project.Notes, &project.BranchID, &project.StatusID, &project.Service.ID, &project.Service.Name)
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

	db := GetDB()
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
	response.Message = "Delete data successfully"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
