package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pufington-pixie/haver/pkg/database"
	"github.com/pufington-pixie/haver/pkg/models"
	"github.com/pufington-pixie/haver/utils"
)

type DataPoint struct {
	ID         int    `json:"id"`
	EquipID    string `json:"EquipID"`
	System     string `json:"System"`
	EquipType  string `json:"EquipType"`
	PointName  string `json:"point_name"`
	PointType  string `json:"point_type"`
	Descriptor string `json:"descriptor"`
}

func GetData(w http.ResponseWriter, r *http.Request) {
	projectIDStr := chi.URLParam(r, "id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest, "Bad Request")
		return
	}

	db := database.GetDB()
	defer db.Close()

	// Query the database to retrieve data

	rows, err := db.Query("SELECT id,EquipId, System, EquipType, Descriptor, PointName, PointType FROM datapoints WHERE projectID = ?", projectID)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	rows.Next()
	defer rows.Close()

	// Create a slice to store the retrieved data
	data := []DataPoint{}

	// Iterate over the rows and populate the data slice
	for rows.Next() {

		var dp DataPoint

		err := rows.Scan(&dp.ID,&dp.EquipID, &dp.System, &dp.EquipType, &dp.Descriptor, &dp.PointName, &dp.PointType)
		if err != nil {
			utils.HandleError(w, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		data = append(data, dp)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Data retrieved successfully",
		Data:    data,
	}

	utils.SendJSONResponse(w, response, http.StatusOK)

}
