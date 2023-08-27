package repositories

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo_Times_Schedules struct {
	*sqlx.DB
}

func NewTimesSchedules(db *sqlx.DB) *Repo_Times_Schedules {
	return &Repo_Times_Schedules{db}
}

func (r *Repo_Times_Schedules) Get_Data(data *models.Times_Scheduless, page string, limit string, location_schedule string, time string, date string, movie string) (*config.Result, error) {
	var list_times_schedules []models.Times_Scheduless
	times_schedules_data := models.Times_Scheduless{}
	var metas config.Metas
	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 5
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := r.Get_Count_Data(location_schedule, time, date, movie)

	if count_data <= 0 {
		metas.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			metas.Next = ""
		} else {
			metas.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		metas.Prev = ""
	} else {
		metas.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		metas.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		metas.Last_page = ""
	}

	if count_data != 0 {
		metas.Total_data = strconv.Itoa(count_data)
	} else {
		metas.Total_data = ""
	}

	if location_schedule == "" || location_schedule == "all" {
		location_schedule = ""
	} else {
		location_schedule = fmt.Sprintf(` AND s.regency='%s'`, location_schedule)
	}
	if time == "" || time == "all" {
		time = ""
	} else {
		time = fmt.Sprintf(` AND ts.time_schedule='%s'`, time)
	}
	if date == "" || date == "all" {
		date = ""
	} else {
		date = fmt.Sprintf(` AND s.set_date='%s'`, date)
	}

	if movie == "" {
		movie = ""
	} else {
		movie = fmt.Sprintf(` AND m.id_movie='%s'`, movie)
	}

	q := fmt.Sprintf(`select ts.id_time_schedule,s.id_premier, ts.time_schedule, s.regency,s.price, s.set_date,m.id_movie,m.title,m.image as image_movie,m.release_date,p.name_premier, p.image as image_premier,p.count_row_seat,p.count_col_seat
	from times_schedules ts 
	left join schedules s on ts.id_schedule = s.id_schedule 
	left join movies m on s.id_movie = m.id_movie
	left join premiers p on s.id_premier = p.id_premier WHERE TRUE %s %s %s %s LIMIT %d OFFSET %d`, location_schedule, time, date, movie, limit_int, offset)
	rows, err := r.Queryx(r.Rebind(q))
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var list_movies_genres []models.Genres
		err := rows.StructScan(&times_schedules_data)
		if err != nil {
			log.Fatalln(err)
		}
		times_schedules_data.Release_date = &strings.Split(*times_schedules_data.Release_date, "T")[0]
		times_schedules_data.Set_date = strings.Split(times_schedules_data.Set_date, "T")[0]
		times_schedules_data.Time_schedule = strings.Split(strings.Split(times_schedules_data.Time_schedule, "T")[1], "Z")[0]
		rows1, _ := r.Queryx("select g.* from movies_genres mg left join genres g on mg.id_genre=g.id_genre  where mg.id_movie=$1", times_schedules_data.Id_movie)
		for rows1.Next() {
			var movies_genres models.Genres
			err1 := rows1.Scan(&movies_genres.Id_genre, &movies_genres.Name_genre, &movies_genres.Created_at, &movies_genres.Updated_at)
			if err1 != nil {
				log.Fatalln(err1)
			}
			// fmt.Println(movies_genres)
			list_movies_genres = append(list_movies_genres, movies_genres)
		}
		times_schedules_data.Genres = list_movies_genres

		// fmt.Println(times_schedules_data)
		list_times_schedules = append(list_times_schedules, times_schedules_data)
	}
	rows.Close()
	if len(list_times_schedules) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: list_times_schedules, Meta: metas}, nil
}

func (r *Repo_Times_Schedules) Get_Data_by_Id(data *models.Times_Scheduless) (*config.Result, error) {
	var list_times_schedules []models.Times_Scheduless
	times_schedules_data := models.Times_Scheduless{}
	q := fmt.Sprintf(`select ts.id_time_schedule, ts.time_schedule, s.regency,s.price, s.set_date,m.id_movie,m.title,m.image as image_movie,m.release_date,p.name_premier, p.image as image_premier,p.count_row_seat,p.count_col_seat
	from times_schedules ts 
	left join schedules s on ts.id_schedule = s.id_schedule 
	left join movies m on s.id_movie = m.id_movie
	left join premiers p on s.id_premier = p.id_premier WHERE TRUE AND ts.id_time_schedule='%s'`, data.Id_time_schedule)
	rows, err := r.Queryx(r.Rebind(q))
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var list_movies_genres []models.Genres
		err := rows.StructScan(&times_schedules_data)
		if err != nil {
			log.Fatalln(err)
		}
		times_schedules_data.Release_date = &strings.Split(*times_schedules_data.Release_date, "T")[0]
		times_schedules_data.Set_date = strings.Split(times_schedules_data.Set_date, "T")[0]
		times_schedules_data.Time_schedule = strings.Split(strings.Split(times_schedules_data.Time_schedule, "T")[1], "Z")[0]
		rows1, _ := r.Queryx("select g.* from movies_genres mg left join genres g on mg.id_genre=g.id_genre  where mg.id_movie=$1", times_schedules_data.Id_movie)
		for rows1.Next() {
			var movies_genres models.Genres
			err1 := rows1.Scan(&movies_genres.Id_genre, &movies_genres.Name_genre, &movies_genres.Created_at, &movies_genres.Updated_at)
			if err1 != nil {
				log.Fatalln(err1)
			}
			// fmt.Println(movies_genres)
			list_movies_genres = append(list_movies_genres, movies_genres)
		}
		times_schedules_data.Genres = list_movies_genres

		// fmt.Println(times_schedules_data)
		list_times_schedules = append(list_times_schedules, times_schedules_data)
	}
	rows.Close()
	if len(list_times_schedules) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: list_times_schedules}, nil
}

func (r *Repo_Times_Schedules) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, `select count(*) 
	from times_schedules ts 
	left join schedules s on ts.id_schedule = s.id_schedule 
	left join movies m on s.id_movie = m.id_movie
	left join premiers p on s.id_premier = p.id_premier
	where m.id_movie =$1`, id)
	return count_data
}

func (r *Repo_Times_Schedules) Get_Count_Data(location_schedule string, time string, date string, movie string) int {
	var id int

	if location_schedule == "" || location_schedule == "all" {
		location_schedule = ""
	} else {
		location_schedule = fmt.Sprintf(` AND s.regency='%s'`, location_schedule)
	}
	if time == "" || time == "all" {
		time = ""
	} else {
		time = fmt.Sprintf(` AND ts.time_schedule='%s'`, time)
	}
	if date == "" || date == "all" {
		date = ""
	} else {
		date = fmt.Sprintf(` AND s.set_date='%s'`, date)
	}

	if movie == "" {
		movie = ""
	} else {
		movie = fmt.Sprintf(` AND m.id_movie='%s'`, movie)
	}

	q := fmt.Sprintf(`select count(*) 
	from times_schedules ts 
	left join schedules s on ts.id_schedule = s.id_schedule 
	left join movies m on s.id_movie = m.id_movie
	left join premiers p on s.id_premier = p.id_premier WHERE TRUE %s %s %s %s`, location_schedule, time, date, movie)
	r.Get(&id, r.Rebind(q))
	return id
}
