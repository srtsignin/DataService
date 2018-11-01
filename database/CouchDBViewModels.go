package database

import "github.com/srtsignin/data-service/models"

// AllDocsResponse represents the response from the all_docs view
type AllDocsResponse struct {
	TotalRows int  `json:"total_rows"`
	Offset    int  `json:"offset"`
	Rows      Rows `json:"rows"`
}

// A Row is a single entry in the database
type Row struct {
	ID   string          `json:"id"`
	Data models.Checkoff `json:"key"`
}

// Rows is an array of multiple Row
type Rows []Row
