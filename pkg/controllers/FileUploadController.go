package controller

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/pufington-pixie/haver/pkg/database"
	"github.com/pufington-pixie/haver/pkg/models"
	"github.com/pufington-pixie/haver/utils"
)

// UploadHandler uploads a project by it's ID.
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
	err := r.ParseMultipartForm(32 << 20) 
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest, err.Error())
		return
	}

	// Get the file from the request
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	// Create a  file name for storing the uploaded file
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
	destinationPath := filepath.Join("uploads", fileName)

	// Create the uploads directory if it doesn't exist
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError, err.Error())
		return
	}

	// Create the destination file on the server
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError, err.Error())
		return
	}
	defer destinationFile.Close()

	// Copy the contents of the uploaded file to the destination file
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError, err.Error())
		return
	}

	// Read the CSV file
	csvFile, err := os.Open(destinationPath)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError, err.Error())
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError, err.Error())
		return
	}

	db := database.GetDB()
	defer db.Close()

	// Get the project ID from the projects table
	projectIDStr := chi.URLParam(r, "id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest, err.Error())
		return
	}


	headers := records[0]
	for i, header := range records[0] {
		// Replace spaces with underscores and remove any other unwanted characters
		header = strings.ReplaceAll(header, " ", "")
		header = strings.ReplaceAll(header, ".", "") 
		header = strings.ReplaceAll(header, "?", "") 
		headers[i] = header
	}
	// Prepare the SQL statement for bulk insertion
	columnCount := len(records[0]) 
	placeholders := strings.TrimRight(strings.Repeat("?, ", columnCount), ", ")
	
	// Escape column names
	escapedHeaders := make([]string, len(headers))
	for i, h := range headers {
		escapedHeaders[i] = "`" + h + "`"
	}
	
	insertQuery := fmt.Sprintf("INSERT INTO datapoints (projectID, %s) VALUES (?, %s)", strings.Join(escapedHeaders, ", "), placeholders)
		// Prepare the SQL statement
		stmt, err := db.Prepare(insertQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()
	
		// Insert each row into the database
		for _, row := range records {
			values := make([]interface{}, 0, len(row)+1)
			values = append(values, projectID)
	
			for _, col := range row {
				values = append(values, col)
			}
	
			_, err := stmt.Exec(values...)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	
		response := models.Response{
			Status:  http.StatusOK,
			Message: "Data inserted successfully",
			Data:    nil,
		}
	
		utils.SendJSONResponse(w, response, http.StatusOK)
	}
	

