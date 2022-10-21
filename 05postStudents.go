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
	c.IndentedJSON(200, "Successfully Added Student Data")
}

func add(c *gin.Context) {

	var newStudent numbers
	if err1 := c.BindJSON(&newStudent); err1 != nil {
		log.Fatal(err1)
		return
	}

	c.JSON(200, newStudent.A+newStudent.B)
}

func add1(c *gin.Context) {

	var newStudent numbers1
	if err1 := c.BindJSON(&newStudent); err1 != nil {
		log.Print("Error -- ")
		log.Fatal(err1)
		log.Print("-- ")
		return
	}

	c.JSON(200, newStudent.A-newStudent.B.C)
}
