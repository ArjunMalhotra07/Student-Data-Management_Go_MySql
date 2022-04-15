package main

import "errors"

func all() ([]StudentData, error) {
	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data")
	if err != nil {
		return nil, errors.New("error1: empty database")
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.Id, &alb.Cgpa, &alb.StudentId, &alb.FatherName, &alb.MotherName, &alb.StudentName, &alb.City); err != nil {
			return nil, errors.New("error2: empty database")
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("error3 : empty database")
	}
	return students, nil
}
