package models

import (
	"math/big"
	"time"
)

// An Checkoff represents the type being sent to
// Long Term Storage
type Checkoff struct {
	StudentUsername    string    `json:"studentUsername"`
	StudentFirstName   string    `json:"studentFirstName"`
	StudentLastName    string    `json:"studentLastName"`
	CheckInTime        time.Time `json:"checkInTime"`
	CheckOutTime       time.Time `json:"checkOutTime"`
	Duration           big.Int   `json:"duration"`
	Courses            Courses   `json:"courses"`
	ProblemDescription string    `json:"problemDescription"`
	TutorUsername      string    `json:"tutorUsername"`
	TutorFirstName     string    `json:"tutorFirstName"`
	TutorLastName      string    `json:"tutorLastName"`
	RoomID             string    `json:"roomId"`
	ID                 string    `json:"_id"`
}
