package models

import(
	"database/sql"
)

type EnvData struct{
	Db_User string
	Db_Pass string
	Db_Host string
	Db_Port uint16
	Db_Name string
	Db *sql.DB
}

type Person struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}


