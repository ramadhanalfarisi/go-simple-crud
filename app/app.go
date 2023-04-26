package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-crud/controller"
	"github.com/ramadhanalfarisi/go-simple-crud/db"
)

type App struct {
	DB *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection(){
	db := db.Connectdb()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewMangaController(a.DB)
	r.POST("/manga", controller.InsertManga)
	r.GET("/manga", controller.GetAllManga)
	r.GET("/manga/:id", controller.GetOneManga)
	r.PUT("/manga/:id", controller.UpdateManga)
	r.DELETE("/manga/:id", controller.DeleteManga)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
