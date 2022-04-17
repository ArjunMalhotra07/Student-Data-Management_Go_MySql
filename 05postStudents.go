package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func postStudents(c *gin.Context) {

	var newStudent StudentData
	if err1 := c.BindJSON(&newStudent); err1 != nil {
		return
	}

	_, err := db.Exec(`INSERT INTO 
	student_Data(studentId, studentName, fatherName, motherName, cgpa, city) 
	VALUES (?,?,?,?,?,?)`,
		newStudent.StudentId,
		newStudent.StudentName,
		newStudent.FatherName,
		newStudent.MotherName,
		newStudent.Cgpa, newStudent.City)

	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(200, "Successfully Added Entries")
}
