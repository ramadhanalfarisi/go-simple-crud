package controller

import "github.com/gin-gonic/gin"

type MangaControllerInterface interface {
	InsertManga(*gin.Context)
	GetAllManga(*gin.Context)
	GetOneManga(*gin.Context)
	UpdateManga(*gin.Context)
	DeleteManga(*gin.Context)
}