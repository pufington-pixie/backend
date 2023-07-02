package controller

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/pufington-pixie/haver/pkg/database"
)

//UploadHandler uploads a project by it's ID.
// @Summary Upload CSV file and save data to the database
// @Description Uploads a CSV file, parses its content, and saves the data to the database
// @Accept multipart/form-data
// @Param file formData file true "CSV file to upload"
// @Param id path int true "Project ID"
// @Success 200 {string} string "CSV file uploaded and saved to the database successfully!"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/upload/{id} [post]
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	
	// Parse the multipart form in the request
	err := r.ParseMultipartForm(32 << 20) // 32MB max upload size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
	destinationPath := filepath.Join("uploads", fileName)

	
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer destinationFile.Close()

	
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Read the CSV file
	csvFile, err := os.Open(destinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := database.GetDB()
    defer db.Close()

		
	// Get the project ID from the projects table
	projectIDStr := chi.URLParam(r, "id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusBadRequest)
	    return
	}

	// Prepare the SQL statement for bulk insertion
	stmt, err := db.Prepare("INSERT INTO datapoints (projectID,EquipID, System,EquipType,Descriptor) VALUES (?, ?, ?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Insert each row into the database
	for _, row := range records {
		_, err := stmt.Exec(projectID,row[0], row[1], row[2],row[4]) 
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}


	w.Write([]byte("CSV file uploaded and saved to the database successfully!"))
}

