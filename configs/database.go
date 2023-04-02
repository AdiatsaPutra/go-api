package configs

import (
	"log"
	// "os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	// dsn := os.Getenv("DSN")
	DB, err = gorm.Open(mysql.Open("st7oind5rm2sirynqq8f:pscale_pw_E3tx8CE9BZbw5JquQBnFzVXfEpcykIVPtt4ekYfou42@tcp(aws.connect.psdb.cloud)/goapi?tls=true"), &gorm.Config{})

	if err != nil {
		log.Fatal("Cant connect to database")
	}

	log.Println("Success to connect")
}
