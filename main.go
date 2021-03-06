package main

import (
	"go-swan/common/constants"
	"go-swan/config"
	"go-swan/database"
	"go-swan/logs"
	"go-swan/routers/commonRouters"
	"go-swan/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jinzhu/gorm"
)

func main() {
	db := initMethod()

	defer func() {
		err := db.Close()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()

	go service.FindMiners()
	createGinServer()
	//test.Test()
}

func createGinServer() {
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := r.Group("/api/v1")
	commonRouters.HostManager(v1.Group(constants.URL_HOST_GET_COMMON))
	err := r.Run(":" + strconv.Itoa(config.GetConfig().Port))
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func initMethod() *gorm.DB {
	db := database.Init()
	config.InitConfig("")
	return db
}
