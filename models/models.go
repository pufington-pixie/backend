package models

import "time"

type Project struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	SAPNumber string    `json:"sapnumber"`
	Notes     string    `json:"notes"`
	BranchID  int       `json:"branchId"`
	StatusID  int       `json:"statusId"`
	Service   Service   `json:"services"`
}

type Service struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}
