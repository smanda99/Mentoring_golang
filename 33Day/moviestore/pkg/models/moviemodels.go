package models

import (
	"github.com/Manda-supraja26/moviestore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name   string  `gorm:"" json:"name"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})

}

func (m *Movie) CreateMovie() *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMovies() []Movie {
	var movie []Movie
	db.Find(&movie)
	return movie
}

func GetMovieByID(ID int) (Movie, *gorm.DB) {
	var movie Movie
	db.Where("ID=?", ID).Find(&movie)
	return movie, db
}

func DeleteMovie(ID int) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}
