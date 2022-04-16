package main

import (
	"fmt"
)

func studentByMother(input string) ([]StudentData, error) {
	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data WHERE motherName = ?", input)

	if err != nil {
		return nil, fmt.Errorf("error1: no student found with motherName = %s, %v,  ", input, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.StudentId, &alb.StudentName, &alb.FatherName, &alb.MotherName, &alb.Cgpa, &alb.City); err != nil {
			return nil, fmt.Errorf("error2: no student found with motherName = %s, %v ", input, err)
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error3 : no student found with motherName = %s, %v,  ", input, err)
	}
	return students, nil
}
