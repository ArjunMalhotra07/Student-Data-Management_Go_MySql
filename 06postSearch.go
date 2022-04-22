package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func searchByColumn(c *gin.Context) {
	var searchData StudentData
	var students []StudentData
	if err1 := c.BindJSON(&searchData); err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	rows, err := db.Query(
		`SELECT * FROM student_Data
		WHERE studentName = ? 
		AND city = ?
		AND fatherName = ?
		AND motherName = ?
		AND cgpa = ?
		AND studentId = ?`, searchData.StudentName, searchData.City, searchData.FatherName, searchData.MotherName, searchData.Cgpa, searchData.StudentId)
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
	x := searchData.Cgpa
	fmt.Println(x)
	c.IndentedJSON(http.StatusOK, students)
	// c.IndentedJSON(http.StatusOK, searchData.FatherName)
	// c.IndentedJSON(http.StatusOK, searchData)

}
