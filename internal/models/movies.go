package models

import "time"

type Movies struct {
	Id_movie        string      `db:"id_movie" form:"id_movie" json:"id_movie,omitempty" valid:"-"`
	Title           string      `db:"title" form:"title,omitempty" json:"title,omitempty" valid:"required~title is required,type(string)~title is string"`
	Release_date    *string     `db:"release_date" json:"release_date" form:"release_date" valid:"-"`
	Duration_hour   int         `db:"duration_hour" json:"duration_hour" form:"duration_hour" valid:"-"`
	Duration_minute int         `db:"duration_minute" json:"duration_minute" form:"duration_minute" valid:"-"`
	Synopsis        string      `db:"synopsis" json:"synopsis" form:"synopsis" valid:"-"`
	Image           string      `db:"image" json:"image,omitempty" valid:"-"`
	Cover_image     string      `db:"cover_image" json:"cover_image,omitempty" valid:"-"`
	Directors       []Directors `json:"directors" valid:"-"`
	Casts           []Casts     `json:"casts" valid:"-"`
	Genres          []Genres    `json:"genres" valid:"-"`
	Created_at      *time.Time  `db:"created_at" json:"created_at" valid:"-"`
	Updated_at      *time.Time  `db:"updated_at" json:"updated_at" valid:"-"`
	Locations       []string    `json:"locations,omitempty" form:"locations" valid:"-"`
	Set_dates       []string    `json:"set_dates,omitempty" form:"set_dates" valid:"-"`
	Times           []string    `json:"times,omitempty" form:"times" valid:"-"`
}

type Moviesset struct {
	Id_movie        string   `db:"id_movie" form:"id_movie" json:"id_movie,omitempty" valid:"-"`
	Title           string   `db:"title" form:"title,omitempty" json:"title,omitempty" valid:"required~title is required,type(string)~title is string"`
	Release_date    *string  `db:"release_date" json:"release_date" form:"release_date" valid:"required"`
	Duration_hour   int      `db:"duration_hour" json:"duration_hour" form:"duration_hour" valid:"required"`
	Duration_minute int      `db:"duration_minute" json:"duration_minute" form:"duration_minute" valid:"required"`
	Synopsis        string   `db:"synopsis" json:"synopsis" form:"synopsis" valid:"required"`
	Image           string   `db:"image" json:"image,omitempty" valid:"-"`
	Cover_image     string   `db:"cover_image" json:"cover_image,omitempty" valid:"-"`
	Id_director     string   `db:"id_director" json:"id_director" form:"id_director" valid:"required"`
	Casts           []string `json:"casts" form:"casts" valid:"required"`
	Genres          []string `json:"genres" form:"genres" valid:"required"`
	Locations       string   `json:"locations" form:"locations" valid:"required"`
	Price           int      `db:"price" json:"price" form:"price" valid:"required"`
	Premiers        []string `json:"premiers" form:"premiers" valid:"required"`
	Set_date        string   `db:"set_date" json:"set_date" form:"set_date" valid:"required"`
	Times           []string `json:"times" form:"times" valid:"required"`
}

type Movies_Casts struct {
	Id_movie_cast string `json:"id_movie_cast" db:"id_movie_cast"`
	Id_movie      string `json:"id_movie" db:"id_movie"`
	Id_cast       string `json:"id_cast" db:"id_cast"`
}

type Movies_Genres struct {
	Id_movie_genre string `json:"id_movie_genre" db:"id_movie_genre"`
	Id_movie       string `json:"id_movie" db:"id_movie"`
	Id_genre       string `json:"id_genre" db:"id_genre"`
}

type Times_Schedules struct {
	Id_time_schedule string `json:"id_time_schedule" db:"id_time_schedule"`
	Id_schedule      string `json:"id_schedule" db:"id_schedule"`
	Time_schedule    string `json:"time_schedule" db:"time_schedule"`
}
