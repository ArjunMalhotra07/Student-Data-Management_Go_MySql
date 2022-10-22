package main

// https://go.dev/doc/tutorial/database-access
import (
	// "01StudentData/routes"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
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

type StudentID struct {
	StudentId int64
}
type StudentDataUpdated struct {
	StudentId          int64
	StudentName        string
	FatherName         string
	MotherName         string
	Cgpa               float64
	City               string
	StudentIdUpdated   int64
	StudentNameUpdated string
	FatherNameUpdated  string
	MotherNameUpdated  string
	CgpaUpdated        float64
	CityUpdated        string
}

type numbers struct {
	A int64
	B int64
}
type numbers1 struct {
	A int64
	B numbers2
}
type numbers2 struct {
	D int64
	C int64
}

func main() {
	// f := fmt.Println
	//  Capture connection properties.
	/*cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "studentInfo",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())*/
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "aman:Mysql_Witcher7%@tcp(127.0.0.1:3306)/studentInfo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	route := gin.Default()
	route.GET("/everyStudent", getStudents)
	route.POST("/everyStudent", postStudents)
	route.POST("/search", searchByColumn)
	route.POST("/add", add)
	route.POST("/add1", add1)
	route.DELETE("/remove", deleteByColumn)
	route.PUT("/update", updateByColumn)
	route.Run(":8081")

}
