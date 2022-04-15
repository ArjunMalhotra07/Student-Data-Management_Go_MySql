package main

import (
	"fmt"
)

func studentByName(studentName string) ([]StudentData, error) {
	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data WHERE studentName = ?", studentName)
	if err != nil {
		return nil, fmt.Errorf("error1: no student found with studentName = %s, %v,  ", studentName, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.Id, &alb.Cgpa, &alb.StudentId, &alb.FatherName, &alb.MotherName, &alb.StudentName, &alb.City); err != nil {
			return nil, fmt.Errorf("error2: no student found with studentName = %s, %v ", studentName, err)
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error3 : no student found with studentName = %s, %v,  ", studentName, err)
	}
	return students, nil
}
