package migrate

import (
	"KobaCareer_API/db"
	"KobaCareer_API/domain"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&domain.Internships{})
}
