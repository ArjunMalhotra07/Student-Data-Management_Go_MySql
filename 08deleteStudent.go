package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteByColumn(c *gin.Context) {
	var searchData StudentID

	if err1 := c.BindJSON(&searchData); err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	_, err := db.Query(
		`DELETE FROM student_Data
		WHERE 
		studentId = ?
		`, searchData.StudentId)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, "Deleted Student")

}
