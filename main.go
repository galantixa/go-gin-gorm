package go_gin_gorm

import (
	"github.com/galantixa/go-gin-gorm/database"
	"github.com/galantixa/go-gin-gorm/routes"
)

func main() {
	database.SetupDb()

	r := routes.SetupRouter()

	r.Run(":8080")
}
