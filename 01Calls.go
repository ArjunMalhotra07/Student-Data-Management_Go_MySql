package main

import (
	"fmt"
	"log"
)

func calls(i int) {
	f := fmt.Println
	var err2 error
	var students []StudentData
	var searchkey string
	var cgpa float64
	var id int64

	if i == 1 {
		f("Enter Fathers Name")
		fmt.Scan(&searchkey)
		students, err2 = searchStudent(searchkey, "fatherName", i, 0)

	} else if i == 2 {
		f("Enter Mothers Name")
		fmt.Scan(&searchkey)
		students, err2 = searchStudent(searchkey, "motherName", i, 0)

	} else if i == 3 {
		f("Enter Student Name")
		fmt.Scan(&searchkey)
		students, err2 = searchStudent(searchkey, "studentName", i, 0)
	} else if i == 5 {
		f("Enter City Name")
		fmt.Scan(&searchkey)
		students, err2 = searchStudent(searchkey, "city", i, 0)
		/*********************/
	} else if i == 6 {
		f("Enter Student ID")
		fmt.Scan(&id)
		students, err2 = searchStudent(" ", "studentId", i, id)
	} else if i == 4 {
		f("Enter CGPA ")
		fmt.Scan(&cgpa)
		students, err2 = studentByCGPA(cgpa)
		f("Students with CGPA less than ", cgpa)
	} else if i == 7 {
		students, err2 = searchStudent("", "", 7, 0)
	}
	if err2 != nil {
		log.Fatal(err2)
	}
	f()
	f("Found : ")
	for i := 0; i < len(students); i++ {
		f(students[i])
	}

	f()
}
