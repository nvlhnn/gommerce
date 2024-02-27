package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nvlhnn/gommerce/internal/api/routes"
	"github.com/nvlhnn/gommerce/internal/cache"
	"github.com/nvlhnn/gommerce/internal/config"
	"github.com/nvlhnn/gommerce/internal/db"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(config config.Config) *ServerHTTP {

	// init redis 
	cache, err := cache.CreateClient(config)
	if err != nil {
		log.Fatal("cannot connect to redis: ", err)
	}

	// init db
	db, dbErr := db.ConnectDatabase(config)
	if dbErr != nil {
		log.Fatal("cannot connect to db: ", dbErr)
	}


	engine := gin.New()
	engine.Use(gin.Logger())


	routes.SetupRoutes(engine, db, cache)

	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":3000")
	if err != nil {
		log.Fatal("gin engine couldn't start")
	}
}