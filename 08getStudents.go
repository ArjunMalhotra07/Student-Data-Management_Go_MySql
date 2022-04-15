package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type mssg struct {
// 	Message string `json:"message"`
// }

// var messages = []mssg{
// 	{Message: "hello"},
// }

func Greet(c *gin.Context) {

	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data")

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.Id, &alb.Cgpa, &alb.StudentId, &alb.FatherName, &alb.MotherName, &alb.StudentName, &alb.City); err != nil {
			return
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, students)
}
