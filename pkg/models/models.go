package models

import "time"

// Project represents a project.
//
// swagger:model
type Project struct {
	// The unique identifier of the project.
	//
	// required: true
	// example: 1
	ID int `json:"id"`

	// The name of the project.
	//
	// required: true
	// example: Project 1
	Name string `json:"name"`

	// The title of the project.
	//
	// required: true
	// example: Title 1
	Title string `json:"title"`

	// The date of the project.
	//
	// required: true
	// format: date
	// example: 2023-06-07
	Date time.Time `json:"date"`

	// The SAP number of the project.
	//
	// required: true
	// example: SAP12345
	SAPNumber string `json:"sapNumber"`

	// Additional notes for the project.
	//
	// required: false
	// example: Some notes about the project.
	Notes string `json:"notes"`

	// The branch ID associated with the project.
	//
	// required: true
	// example: 1
	BranchID int `json:"branchId"`

	// The status ID of the project.
	//
	// required: true
	// example: 1
	StatusID int `json:"statusId"`

	// The service associated with the project.
	//
	// required: true
	Service Service `json:"services"`
}

// Service represents a service.
//
// swagger:model
type Service struct {
	// The ID of the service.
	//
	// required: true
	// example: 1
	ID int `json:"serviceId"`

	// The name of the service.
	//
	// required: true
	// example: Service 1
	Name string `json:"serviceName"`
}

// Response represents a generic API response.
//
// swagger:model
type Response struct {
	// The status code of the response.
	//
	// required: true
	// example: 200
	Status int `json:"status"`

	// The message associated with the response.
	//
	// required: true
	// example: Success
	Message string `json:"message"`

	// The data payload of the response.
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

