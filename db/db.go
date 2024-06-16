package db

import (
	"database/sql"
	"fmt"
	"restapi/utils"
	"restapi/log"

	_ "github.com/go-sql-driver/mysql"
)

type Database_Connection struct{
	Db *sql.DB
}


var DB *Database_Connection


func init() {
	var connection Database_Connection

	dbData:= utils.DB_CONNCTION

	connStr:=fmt.Sprintf("%s:%s@tcp(%s:%d)/",dbData.Db_User,dbData.Db_Pass,dbData.Db_Host,dbData.Db_Port)

	db, err := sql.Open("mysql", connStr)
	if err!=nil{log.Log.Fatal("unable to connect to database",err.Error())}


	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS "+dbData.Db_Name)
	if err!=nil{log.Log.Fatal("unable to create databse",err.Error())}
	
	_, err = db.Exec("USE "+dbData.Db_Name)
	if err!=nil{log.Log.Fatal("unable to use database "+dbData.Db_Name,err.Error())}

	connection.Db=db

	DB = &connection

	//for creating required tables
	DB.CreateTables()
}


func (db *Database_Connection) CreateTables() {

	//person table
	_,err:=DB.Db.Exec(`CREATE TABLE IF NOT EXISTS person(id int not null key auto_increment, name varchar(255), age int);`)
	if err!=nil{log.Log.Error("unable to create person table",err.Error())}

	//phone table
	_,err=DB.Db.Exec(`CREATE TABLE IF NOT EXISTS phone(id int not null key auto_increment, number varchar(255), person_id int);`)
	if err!=nil{log.Log.Error("unable to create phone table",err.Error())}
	
	//Address table
	_,err=DB.Db.Exec(`CREATE TABLE IF NOT EXISTS address(id int not null key auto_increment, city varchar(255), state varchar(255),street1 varchar(255), street2 varchar(255), zip_code varchar(255));`)
	if err!=nil{log.Log.Error("unable to create address table",err.Error())}
	
	//Address Join table
	_,err=DB.Db.Exec(`CREATE TABLE IF NOT EXISTS address_join(id int not null key auto_increment, person_id int, address_id int);`)
	if err!=nil{log.Log.Error("unable to create address_join table",err.Error())}

}
