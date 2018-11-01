package models

import (
	"strconv"
	"strings"
)

// A Course represents a single course
type Course struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Department  string `json:"department"`
	QueryString string `json:"queryString"`
}

// A Courses is an array of multiple Course objects
type Courses []Course

// Delimited returns a delimited copy of the Course object without the QueryString
/* The fields are in the order:
Name
Number
Department
*/
func (course Course) Delimited(delimiter rune) string {
	var builder strings.Builder
	builder.WriteString(course.Name)
	builder.WriteRune(delimiter)
	builder.WriteString(strconv.Itoa(course.Number))
	builder.WriteRune(delimiter)
	builder.WriteString(course.Department)
	return builder.String()
}

// Delimited Returns the Courses in a Delimited Format
func (courses Courses) Delimited(delimiter rune) string {
	var builder strings.Builder
	for index, course := range courses {
		builder.WriteString(course.Delimited(delimiter))
		if index != len(courses)-1 {
			builder.WriteRune(delimiter)
		}
	}
	return builder.String()
}
