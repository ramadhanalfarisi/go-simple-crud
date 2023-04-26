package repository

import (
	"database/sql"
	"log"

	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type MangaRepository struct {
	Db *sql.DB
}

func NewMangaRepository(db *sql.DB) MangaRepositoryInterface {
	return &MangaRepository{Db: db}
}

// DeleteManga implements MangaRepositoryInterface
func (m *MangaRepository) DeleteManga(id uint) bool {
	_, err := m.Db.Exec("DELETE FROM manga WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GetAllManga implements MangaRepositoryInterface
func (m *MangaRepository) GetAllManga() []model.Manga {
	query, err := m.Db.Query("SELECT * FROM manga")
	if err != nil {
		log.Println(err)
		return nil
	}
	var mangas []model.Manga
	if query != nil {
		for query.Next() {
			var (
				id       uint
				title    string
				genre    string
				volumes  uint8
				chapters uint16
				author   string
			)
			err := query.Scan(&id, &title, &genre, &volumes, &chapters, &author)
			if err != nil {
				log.Println(err)
			}
			manga := model.Manga{Id: id, Title: title, Genre: genre, Volumes: volumes, Chapters: chapters, Author: author}
			mangas = append(mangas, manga)
		}
	}
	return mangas
}

// GetOneManga implements MangaRepositoryInterface
func (m *MangaRepository) GetOneManga(id uint) model.Manga {
	query, err := m.Db.Query("SELECT * FROM manga WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return model.Manga{}
	}
	var manga model.Manga
	if query != nil {
		for query.Next() {
			var (
				id       uint
				title    string
				genre    string
				volumes  uint8
				chapters uint16
				author   string
			)
			err := query.Scan(&id, &title, &genre, &volumes, &chapters, &author)
			if err != nil {
				log.Println(err)
			}
			manga = model.Manga{Id: id, Title: title, Genre: genre, Volumes: volumes, Chapters: chapters, Author: author}
		}
	}
	return manga
}

// InsertManga implements MangaRepositoryInterface
func (m *MangaRepository) InsertManga(post model.PostManga) bool {
	stmt, err := m.Db.Prepare("INSERT INTO manga(title, genre, volumes, chapters, author) VALUES ($1,$2,$3,$4,$5)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Title, post.Genre, post.Volumes, post.Chapters, post.Author)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// UpdateManga implements MangaRepositoryInterface
func (m *MangaRepository) UpdateManga(id uint, post model.PostManga) model.Manga {
	_, err := m.Db.Exec("UPDATE manga SET title = $1, genre = $2, volumes = $3, chapters = $4, author = $5 WHERE id = $6", post.Title, post.Genre, post.Volumes, post.Chapters, post.Author, id)
	if err != nil {
		log.Println(err)
		return model.Manga{}
	}
	return m.GetOneManga(id)
}
