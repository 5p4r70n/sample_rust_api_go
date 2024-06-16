package controllers

import (
	"encoding/json"
	"restapi/db"
	"restapi/log"
	"restapi/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetPersonById(c *gin.Context) {
	db := db.DB.Db
	
	person_id := c.Param("person_id")
	
	
	// var person models.Person
	query:= 	`select name,number,city,state,street1,street2,zip_code from person as p 
				inner join phone as ph on ph.person_id=p.id 
				inner join address_join as aid on aid.person_id = p.id 
				inner join address as a on a.id = aid.address_id
				WHERE p.id = ?;
				`
	rows:= db.QueryRow(query,person_id)
	
	var person models.Person

	err:=rows.Scan(&person.Name, &person.PhoneNumber, &person.City,&person.State,&person.Street1,&person.Street2,&person.ZipCode)
	
	if err==sql.ErrNoRows{
		c.JSON(404, "Person not found")
		return
	}else if err!=nil && err!=sql.ErrNoRows{
		log.Log.Error("unable to scan rows",err.Error())
		c.JSON(500, "Something Buggy")		
		return
	}

	json_data,err:=json.Marshal(person)
	if err!=nil{
		log.Log.Error("unable to marshal json",err.Error())
		c.JSON(500, "Something Buggy")		
		return
	}

	c.JSON(200, string(json_data))
	return
}




func InsertData(c *gin.Context) {
	db := db.DB.Db

	
	var person models.Person
	err:=c.BindJSON(&person)
	if err!=nil{
		log.Log.Error("unable to bind json",err.Error())
		c.JSON(500, "Something Buggy")		
		return
	}
	
	query:=fmt.Sprintf( `INSERT INTO person (name, age) VALUES ('%s', NULL);`,person.Name)

	txn,err:=db.Exec(`INSERT INTO person (name, age) VALUES (?, NULL);`,person.Name)
	if err!=nil{log.Log.Error("unable to insert person to db ",err.Error(),"Query: ",query)}
	newPersonID,err:=txn.LastInsertId()
	if err!=nil{log.Log.Error("unable to get last insert id",err.Error(),"Query: ",query)}
	
	_,err=db.Exec(`INSERT INTO phone (number, person_id) VALUES (?, ?)`,person.PhoneNumber,newPersonID)
	if err!=nil{log.Log.Error("unable to insert phone to db ",err.Error(),"Query: ",query)}

	txn,err=db.Exec(`INSERT INTO address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)`,person.City,person.State,person.Street1,person.Street1,person.ZipCode)
	if err!=nil{log.Log.Error("unable to insert address to db ",err.Error(),"Query: ",query)}
	newAddressID,err:=txn.LastInsertId()
	if err!=nil{log.Log.Error("unable to get last insert id",err.Error(),"Query: ",query)}
	
	_,err=db.Exec(`INSERT INTO address_join (person_id, address_id) VALUES (?, ?);`,newPersonID,newAddressID)
	if err!=nil{log.Log.Error("unable to insert address_join to db ",err.Error(),"Query: ",query)}
	
	if err!=nil{
		c.JSON(500, "Something Buggy")		
		return
	}

	c.JSON(200, string("Successfully inserted"))
	return
}
