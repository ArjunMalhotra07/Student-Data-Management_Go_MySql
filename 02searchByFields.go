package main

import (
	"database/sql"
	"fmt"
)

func searchStudent(searchkey, colmName string, index int, studentId int64) ([]StudentData, error) {
	var students []StudentData
	var err error
	var rows *sql.Rows

	if index == 1 {
		rows, err = db.Query("SELECT * FROM student_Data WHERE fatherName = ?", searchkey)
	} else if index == 2 {
		rows, err = db.Query("SELECT * FROM student_Data WHERE motherName = ?", searchkey)
	} else if index == 3 {
		rows, err = db.Query("SELECT * FROM student_Data WHERE studentName = ?", searchkey)
	} else if index == 5 {
		rows, err = db.Query("SELECT * FROM student_Data WHERE city = ?", searchkey)
	} else if index == 6 {
		rows, err = db.Query("SELECT * FROM student_Data WHERE studentId = ?", studentId)
	} else if index == 7 {
		rows, err = db.Query("SELECT * FROM student_Data")
	}

	if err != nil {
		return nil, fmt.Errorf("error1: no student found with %s = %s, %v,  ", colmName, searchkey, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.StudentId, &alb.StudentName, &alb.FatherName, &alb.MotherName, &alb.Cgpa, &alb.City); err != nil {
			return nil, fmt.Errorf("error2: no student found with %s = %s, %v,  ", colmName, searchkey, err)
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error3 : no student found with %s = %s, %v,  ", colmName, searchkey, err)
	}
	return students, nil
}
