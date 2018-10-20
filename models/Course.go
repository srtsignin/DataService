package models

// A Course represents a single course
type Course struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Department  string `json:"department"`
	QueryString string `json:"queryString"`
}

// A Courses is an array of multiple Course objects
type Courses []Course
