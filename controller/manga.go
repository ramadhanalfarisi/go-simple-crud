package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
	"github.com/ramadhanalfarisi/go-simple-crud/repository"
)

type MangaController struct {
	Db *sql.DB
}

func NewMangaController(db *sql.DB) MangaControllerInterface {
	return &MangaController{Db: db}
}

// DeleteManga implements MangaControllerInterface
func (m *MangaController) DeleteManga(c *gin.Context) {
	DB := m.Db
	var uri model.MangaUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewMangaRepository(DB)
	delete := repository.DeleteManga(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete manga successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete manga failed"})
		return
	}
}

// GetAllManga implements MangaControllerInterface
func (m *MangaController) GetAllManga(c *gin.Context) {
	DB := m.Db
	repository := repository.NewMangaRepository(DB)
	get := repository.GetAllManga()
	if (get != nil) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get manga successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "manga not found"})
		return
	}
}

// GetOneManga implements MangaControllerInterface
func (m *MangaController) GetOneManga(c *gin.Context) {
	DB := m.Db
	var uri model.MangaUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewMangaRepository(DB)
	get := repository.GetOneManga(uri.ID)
	if (get != model.Manga{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get manga successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "manga not found"})
		return
	}
}

// InsertManga implements MangaControllerInterface
func (m *MangaController) InsertManga(c *gin.Context) {
	DB := m.Db
	var post model.PostManga
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewMangaRepository(DB)
	insert := repository.InsertManga(post)
	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
		return
	}
}

// UpdateManga implements MangaControllerInterface
func (m *MangaController) UpdateManga(c *gin.Context) {
	DB := m.Db
	var post model.PostManga
	var uri model.MangaUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewMangaRepository(DB)
	update := repository.UpdateManga(uri.ID,post)
	if (update != model.Manga{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update manga successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update manga failed"})
		return
	}
}
