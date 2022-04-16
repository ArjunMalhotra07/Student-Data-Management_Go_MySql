package main

import (
	"fmt"
)

func studentByCity(city string) ([]StudentData, error) {
	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data WHERE city = ?", city)
	if err != nil {
		return nil, fmt.Errorf("error1: no student found with city = %s, %v,  ", city, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.StudentId, &alb.StudentName, &alb.FatherName, &alb.MotherName, &alb.Cgpa, &alb.City); err != nil {
			return nil, fmt.Errorf("error2: no student found with city = %s, %v ", city, err)
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error3 : no student found with city = %s, %v,  ", city, err)
	}
	return students, nil
}
