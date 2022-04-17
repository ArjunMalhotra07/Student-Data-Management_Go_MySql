package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStudents(c *gin.Context) {

	var students []StudentData
	rows, err := db.Query("SELECT * FROM student_Data")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var student StudentData
		if err := rows.Scan(&student.StudentId, &student.StudentName, &student.FatherName, &student.MotherName, &student.Cgpa, &student.City); err != nil {
			return
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}
