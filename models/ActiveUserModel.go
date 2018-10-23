package models

// An ActiveUserModel represents the return type from the
// active user service
type ActiveUserModel struct {
	Username           string  `json:"username"`
	Name               string  `json:"name"`
	CheckInTime        int64   `json:"checkInTime"`
	CheckOutTime       int64   `json:"checkOutTime"`
	Courses            Courses `json:"courses"`
	ProblemDescription string  `json:"problemDescription"`
	TutorUsername      string  `json:"tutorUsername"`
	TutorName          string  `json:"tutorName"`
	RoomID             string  `json:"roomId"`
}
