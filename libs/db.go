package libs

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Connect() (db *gorm.DB, err error) {

	err = godotenv.Load("../.env")

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	USER := os.Getenv("POSTGRES_USER")
	PASS := os.Getenv("POSTGRES_PASSWORD")
	DBNAME := os.Getenv("POSTGRES_DBNAME")

	CONNECT := "host=" + os.Getenv("POSTGRES_DBHOST") +
		" port=5432" +
		" user=" + USER +
		" dbname=" + DBNAME +
		" password=" + PASS +
		" sslmode=disable"

	db, err = gorm.Open("postgres", CONNECT)

	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}
