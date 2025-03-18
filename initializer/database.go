package initializer
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// ConnectionToDB is a function that connects to the database

func ConnectionToDB()  {
	var err error
	dsn :=os.Getenv("DB_URL")
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true},)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	DB = db	
}