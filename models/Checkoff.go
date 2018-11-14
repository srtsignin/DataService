package models

import (
	"log"
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
	id                 string
}

// GetID returns the _id of the Checkoff Object
func (c Checkoff) GetID() string {
	return c.id
}

// CreateCheckoff creates a Checkoff object from an
// ActiveUserModel changing all relevant fields
func CreateCheckoff(activeUserModel ActiveUserModel) Checkoff {
	checkoff := Checkoff{
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
		id:                 strconv.FormatInt(activeUserModel.CheckInTime, 10) + activeUserModel.Username,
	}
	log.Printf("Converted to: %v\n", checkoff)
	return checkoff
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

// Delimited returns the struct in a delimited form using the provided rune
/* The order of the fields printed are:
StudentUsername
StudentFirstName
StudentLastName
CheckInTime
CheckOutTime
Duration
ProblemDescription
TutorUsername
TutorFirstName
TutorLastName
RoomID
id
Courses
*/
func (c Checkoff) Delimited(delimiter rune) string {
	var builder strings.Builder

	builder.WriteString(c.StudentUsername)
	builder.WriteRune(delimiter)
	builder.WriteString(c.StudentFirstName)
	builder.WriteRune(delimiter)
	builder.WriteString(c.StudentLastName)
	builder.WriteRune(delimiter)
	builder.WriteString(c.CheckInTime.String())
	builder.WriteRune(delimiter)
	builder.WriteString(c.CheckOutTime.String())
	builder.WriteRune(delimiter)
	builder.WriteString(strconv.FormatInt(c.Duration, 10))
	builder.WriteRune(delimiter)
	builder.WriteString(escape(c.ProblemDescription))
	builder.WriteRune(delimiter)
	builder.WriteString(c.TutorUsername)
	builder.WriteRune(delimiter)
	builder.WriteString(c.TutorFirstName)
	builder.WriteRune(delimiter)
	builder.WriteString(c.TutorLastName)
	builder.WriteRune(delimiter)
	builder.WriteString(c.Courses.Delimited(delimiter))

	return builder.String()
}

func escape(description string) string {
	description = strings.Replace(description, "\"", "\"\"", -1)
	description = strings.Replace(description, "\n", " ", -1)
	return "\"\"" + description + "\"\""
}
