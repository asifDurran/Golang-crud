package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(){
	var err error
	dsn := "host=dumbo.db.elephantsql.com user=phrwhwvk password=4IfKanSrPe5bJwME86Zfd6515viHhh6N dbname=phrwhwvk port=5432 sslmode=disable"
	
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Database")
	}
}