package models

import (
	"strconv"
	"strings"
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
	Duration           int64     `json:"duration"`
	Courses            Courses   `json:"courses"`
	ProblemDescription string    `json:"problemDescription"`
	TutorUsername      string    `json:"tutorUsername"`
	TutorFirstName     string    `json:"tutorFirstName"`
	TutorLastName      string    `json:"tutorLastName"`
	RoomID             string    `json:"roomId"`
	ID                 string    `json:"_id"`
}

// CreateCheckoff creates a Checkoff object from an
// ActiveUserModel changing all relevant fields
func CreateCheckoff(activeUserModel ActiveUserModel) Checkoff {
	return Checkoff{
		StudentUsername:    activeUserModel.Username,
		StudentFirstName:   parseFirstName(activeUserModel.Name),
		StudentLastName:    parseLastName(activeUserModel.Name),
		CheckInTime:        convertLongToTime(activeUserModel.CheckInTime),
		CheckOutTime:       convertLongToTime(activeUserModel.CheckOutTime),
		Duration:           (activeUserModel.CheckOutTime - activeUserModel.CheckInTime),
		Courses:            activeUserModel.Courses,
		ProblemDescription: activeUserModel.ProblemDescription,
		TutorUsername:      activeUserModel.TutorUsername,
		TutorFirstName:     parseFirstName(activeUserModel.TutorName),
		TutorLastName:      parseLastName(activeUserModel.TutorName),
		RoomID:             activeUserModel.RoomID,
		ID:                 strconv.FormatInt(activeUserModel.CheckInTime, 10) + activeUserModel.Username,
	}
}

func parseFirstName(fullName string) string {
	var strings = strings.Split(fullName, " ")
	return strings[0]
}

func parseLastName(fullName string) string {
	var strings = strings.Split(fullName, " ")
	return strings[1]
}

func convertLongToTime(longTime int64) time.Time {
	return time.Unix(0, longTime*int64(time.Millisecond))
}
