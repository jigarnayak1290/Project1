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

var db *sql.DB

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
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return nil
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

func getVesselByNaccsCode(NACCS_Code string) error {

	sqlStatement := `SELECT vessel_name, naccs_code, owner_name, modified_person_name, note FROM vessel WHERE naccs_code=$1;`

	result := db.QueryRow(sqlStatement, NACCS_Code)
	println("Your result :", result)

	return nil
}
