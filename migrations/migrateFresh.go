package migrations

import (
	"log"
	"mitramas_test/db"
	"mitramas_test/helpers"

	"github.com/gin-gonic/gin"
)

func MigrateFresh(c *gin.Context) {
	db := db.Connect()
	defer db.Close()

	sqlStatement := `DROP TABLE IF EXISTS check_outs, activities, check_ins, users;`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Printf("Failed at dropping table with err: %s", err)
		helpers.NewHandlerResponse("Failed refresh databases", nil).Failed(c)
		return
	}
	Migrate(c)
}
