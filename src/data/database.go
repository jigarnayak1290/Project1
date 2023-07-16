package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

var dbObj *sql.DB

// InitializeDB initializes the database connection
func InitializeDB() error {
	// Open the database connection

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	dbObj = db

	fmt.Println("Successfully connected!")
	return nil
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return dbObj
}

// GetVessal returns the Vessal by NAVVS codes
func GetVessal(NACCS_Code string) error {
	//sqlStatement := `SELECT vessel_name, naccs_code, owner_name, modified_person_name, note FROM vessel WHERE naccs_code=$1;`
	sqlStatement := `SELECT vessel_name, naccs_code FROM vessel WHERE naccs_code=$1;`

	var vessel_name string
	var naccs_code string
	row := dbObj.QueryRow(sqlStatement, NACCS_Code)
	switch err := row.Scan(&naccs_code, &vessel_name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(naccs_code, vessel_name)
	default:
		fmt.Println("Panic No rows were returned!")
		//panic(err)

	}
	return nil
	//  println("Your result :", result)
	//println("Your result :")
}
