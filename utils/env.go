package utils

import (
	"restapi/models"
	"restapi/log"
	"github.com/joho/godotenv"
  "os"
  "strconv"
  
)

var DB_CONNCTION models.EnvData

func init(){
  // Load environment variables from .env file
  err := godotenv.Load()
  if err != nil {
	  log.Log.Fatal("Error loading .env file")
  }

  // Convert string to uint
  num, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)
  if err != nil {
	  log.Log.Fatal("convert string to uint error")
  }

  DB_CONNCTION = models.EnvData{
                                Db_User: os.Getenv("DB_USER"),
                                Db_Pass: os.Getenv("DB_PASS"),
                                Db_Host: os.Getenv("DB_HOST"),
                                Db_Port: uint16(num),
                                Db_Name: os.Getenv("DB_NAME"),
                              }

}