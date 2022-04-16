package main

import (
	"fmt"
)

func studentByCGPA(cgpa float64) ([]StudentData, error) {
	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data WHERE cgpa = ?", cgpa)
	if err != nil {
		return nil, fmt.Errorf("error1: no student found with cgpa = %.2f, %v,  ", cgpa, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.StudentId, &alb.StudentName, &alb.FatherName, &alb.MotherName, &alb.Cgpa, &alb.City); err != nil {
			return nil, fmt.Errorf("error2: no student found with cgpa = %.2f, %v ", cgpa, err)
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error3 : no student found with cgpa = %.2f, %v,  ", cgpa, err)
	}
	return students, nil
}
