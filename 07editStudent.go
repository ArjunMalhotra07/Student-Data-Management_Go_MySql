package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateByColumn(c *gin.Context) {
	var searchData StudentDataUpdated

	if err1 := c.BindJSON(&searchData); err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	_, err := db.Query(
		`UPDATE student_Data 

		SET 	 
		studentId = ?,studentName = ? ,fatherName = ?,motherName = ?,cgpa = ?,city = ?

		WHERE 
		studentId = ?
		AND studentName = ? 
		AND fatherName = ?
		AND motherName = ?
		AND cgpa = ?
		AND city = ?
		`, searchData.StudentIdUpdated, searchData.StudentNameUpdated, searchData.FatherNameUpdated, searchData.MotherNameUpdated, searchData.CgpaUpdated, searchData.CityUpdated, searchData.StudentId, searchData.StudentName, searchData.FatherName, searchData.MotherName, searchData.Cgpa, searchData.City)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, "Updated Student")

}
