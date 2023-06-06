package controller

import (
	"encoding/csv"
	// "io"
	// "os"
	// "path/filepath"

	"net/http"

	"example.com/database"
)
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

	// Read the CSV file
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := database.GetDB()
    defer db.Close()

		// // Save a copy of the CSV file in the assets folder
		// dst, err := os.Create("assets/" + filepath.Base(handler.Filename))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// defer dst.Close()
	
		// _, err = file.Seek(0, io.SeekStart) // Reset file position to the beginning
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	
		// _, err = io.Copy(dst, file)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

	// Prepare the SQL statement for bulk insertion
	stmt, err := db.Prepare("INSERT INTO datapoints (EquipID, System,EquipType,Descriptor) VALUES (?, ?, ?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Insert each row into the database
	for _, row := range records {
		_, err := stmt.Exec(row[0], row[1], row[2],row[4]) // Modify column names based on your CSV structure
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Return a response
	w.Write([]byte("CSV file uploaded and saved to the database successfully!"))
}

