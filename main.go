package main

import (
	// "01StudentData/routes"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type StudentData struct {
	StudentId   int64
	StudentName string
	FatherName  string
	MotherName  string
	Cgpa        float64
	City        string
}

func main() {
	f := fmt.Println

	//  Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "studentInfo",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	// Get a database handle.

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	f("Connected!")
	f("*** Press To Search Student Bio Data Using *** ")
	f("1  Father's Name")
	f("2  Mother's Name")
	f("3  Students's Name")
	f("4  CGPA")
	f("5  City")
	f("6  Students Id")
	f("7  Every Student")
	var i int
	fmt.Scan(&i)
	calls(i)
	route := gin.Default()
	route.GET("/everyStudent", getStudents)
	route.POST("/everyStudent", postStudents)
	route.Run("localhost:8080")

}

func calls(i int) {
	f := fmt.Println
	var err2 error
	var students []StudentData
	var s string
	var cgpa int
	if i == 1 {
		f("Enter Fathers Name")
		fmt.Scan(&s)
		students, err2 = studentByFather(s)
	} else if i == 2 {
		f("Enter Mothers Name")
		fmt.Scan(&s)
		students, err2 = studentByMother(s)

	} else if i == 3 {
		f("Enter Student Name")
		fmt.Scan(&s)
		students, err2 = studentByName(s)
	} else if i == 4 {
		f("Enter CGPA ")
		fmt.Scan(&cgpa)
		students, err2 = studentByCGPA(cgpa)
		f("Students with CGPA less than ", cgpa)
	} else if i == 5 {
		f("Enter City Name")
		fmt.Scan(&s)
		students, err2 = studentByCity(s)
	} else if i == 6 {
		f("Enter Student ID")
		fmt.Scan(&s)
		students, err2 = studentById(s)
	} else if i == 7 {
		students, err2 = all()
	}
	if err2 != nil {
		log.Fatal(err2)
	}
	f()
	for i := 0; i < len(students); i++ {
		f(students[i])
	}
}
