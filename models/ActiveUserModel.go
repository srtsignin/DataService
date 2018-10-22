package models

import "math/big"

// An ActiveUserModel represents the return type from the
// active user service
type ActiveUserModel struct {
	Username           string  `json:"username"`
	Name               string  `json:"name"`
	CheckInTime        big.Int `json:"checkInTime"`
	CheckOutTime       big.Int `json:"checkOutTime"`
	Courses            Courses `json:"courses"`
	ProblemDescription string  `json:"problemDescription"`
	TutorUsername      string  `json:"tutorUsername"`
	TutorName          string  `json:"tutorName"`
	RoomID             string  `json:"roomId"`
}
