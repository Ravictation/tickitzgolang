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

type Repo_Movies struct {
	*sqlx.DB
}

type Regencys struct {
	Regency string `db:"regency" json:"regency"`
}

type Set_dates struct {
	Set_date string `db:"set_date" json:"set_date"`
}

type Times struct {
	Time_schedule string `db:"time_schedule" json:"time_schedule"`
}

func NewMovies(db *sqlx.DB) *Repo_Movies {
	return &Repo_Movies{db}
}

func (r *Repo_Movies) Get_Data(data *models.Movies, page string, limit string, search string, orderby string, by_genre string, date string) (*config.Result, error) {
	var list_movies []models.Movies
	movies_data := models.Movies{}
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

	count_data := r.Get_Count_Data(search, by_genre, date)

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

	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(m.title) like LOWER('%s')`, "%"+search+"%")
	}

	join_genre := ""
	if by_genre == "" {
		by_genre = ""
	} else {
		join_genre = fmt.Sprint(` left join movies_genres mg on mg.id_movie = m.id_movie `)
		by_genre = fmt.Sprintf(` AND mg.id_genre='%s'`, by_genre)
	}

	if orderby == "" {
		orderby = ""
	} else {
		orderby = fmt.Sprintf(` ORDER BY m.%s`, orderby)
	}

	if date == "" {
		date = ""
	} else {
		dates := strings.Split(date, "-")
		date = fmt.Sprintf(` AND EXTRACT(MONTH FROM m.release_date)='%s' AND EXTRACT(YEAR FROM m.release_date)='%s'`, dates[1], dates[0])
	}

	q := fmt.Sprintf(`select m.id_movie, m.title, m.release_date, m.duration_hour, m.duration_minute, m.synopsis, m.image, m.cover_image, m.created_at, m.updated_at from public.movies m %s WHERE TRUE %s %s %s %s LIMIT %d OFFSET %d`, join_genre, date, by_genre, search, orderby, limit_int, offset)
	rows, err := r.Queryx(r.Rebind(q))
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var list_movies_casts []models.Casts
		var list_movies_genres []models.Genres
		var list_movies_directors []models.Directors
		err := rows.StructScan(&movies_data)
		if err != nil {
			log.Fatalln(err)
		}
		movies_data.Release_date = &strings.Split(*movies_data.Release_date, "T")[0]
		rows, _ := r.Queryx("select c.* from movies_casts mc left join casts c on mc.id_cast = c.id_cast where mc.id_movie = $1", movies_data.Id_movie)
		for rows.Next() {
			var movies_casts models.Casts
			err := rows.Scan(&movies_casts.Id_cast, &movies_casts.Name_cast, &movies_casts.Created_at, &movies_casts.Updated_at)
			if err != nil {
				log.Fatalln(err)
			}
			// fmt.Println(movies_casts)
			list_movies_casts = append(list_movies_casts, movies_casts)
		}
		rows1, _ := r.Queryx("select g.* from movies_genres mg left join genres g on mg.id_genre=g.id_genre  where mg.id_movie=$1", movies_data.Id_movie)
		for rows1.Next() {
			var movies_genres models.Genres
			err1 := rows1.Scan(&movies_genres.Id_genre, &movies_genres.Name_genre, &movies_genres.Created_at, &movies_genres.Updated_at)
			if err1 != nil {
				log.Fatalln(err1)
			}
			// fmt.Println(movies_genres)
			list_movies_genres = append(list_movies_genres, movies_genres)
		}
		rows2, _ := r.Queryx("select d.* from movies m left join directors d on m.id_director=d.id_director where m.id_movie=$1", movies_data.Id_movie)
		for rows2.Next() {
			var movies_directors models.Directors
			err2 := rows2.Scan(&movies_directors.Id_director, &movies_directors.Name_director, &movies_directors.Created_at, &movies_directors.Updated_at)
			if err2 != nil {
				log.Fatalln(err2)
			}
			// fmt.Println(movies_genres)
			list_movies_directors = append(list_movies_directors, movies_directors)
		}

		movies_data.Casts = list_movies_casts
		movies_data.Genres = list_movies_genres
		movies_data.Directors = list_movies_directors

		// fmt.Println(movies_data)
		list_movies = append(list_movies, movies_data)
	}
	rows.Close()
	if len(list_movies) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: list_movies, Meta: metas}, nil
}

func (r *Repo_Movies) Get_Data_by_Id(data *models.Movies) (*config.Result, error) {
	var list_movies []models.Movies
	movies_data := models.Movies{}
	q := fmt.Sprintf(`select m.id_movie, m.title, m.release_date, m.duration_hour, m.duration_minute, m.synopsis, m.image, m.cover_image, m.created_at, m.updated_at from public.movies m WHERE m.id_movie='%s'`, data.Id_movie)
	rows, err := r.Queryx(r.Rebind(q))
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var list_movies_casts []models.Casts
		var list_movies_genres []models.Genres
		var list_movies_directors []models.Directors
		err := rows.StructScan(&movies_data)
		if err != nil {
			log.Fatalln(err)
		}
		movies_data.Release_date = &strings.Split(*movies_data.Release_date, "T")[0]
		regencys := []Regencys{}
		err = r.Select(&regencys, "select distinct regency from schedules s where s.id_movie =$1", movies_data.Id_movie)
		if err != nil {
			log.Fatalln(err)
		}
		movies_data.Locations = append(movies_data.Locations, "all")
		for i := range regencys {
			movies_data.Locations = append(movies_data.Locations, regencys[i].Regency)
		}

		set_dates := []Set_dates{}
		err = r.Select(&set_dates, "select distinct set_date from schedules s where s.id_movie =$1", movies_data.Id_movie)
		if err != nil {
			log.Fatalln(err)
		}
		movies_data.Set_dates = append(movies_data.Set_dates, "all")
		for i := range set_dates {
			movies_data.Set_dates = append(movies_data.Set_dates, strings.Split(set_dates[i].Set_date, "T")[0])
		}

		times := []Times{}
		err = r.Select(&times, "select distinct time_schedule from times_schedules ts left join schedules s on ts.id_schedule =s.id_schedule where s.id_movie =$1", movies_data.Id_movie)
		if err != nil {
			log.Fatalln(err)
		}
		movies_data.Times = append(movies_data.Times, "all")
		for i := range times {
			movies_data.Times = append(movies_data.Times, strings.Split(strings.Split(times[i].Time_schedule, "T")[1], "Z")[0])
		}

		rows, _ := r.Queryx("select c.* from movies_casts mc left join casts c on mc.id_cast = c.id_cast where mc.id_movie = $1", movies_data.Id_movie)
		for rows.Next() {
			var movies_casts models.Casts
			err := rows.Scan(&movies_casts.Id_cast, &movies_casts.Name_cast, &movies_casts.Created_at, &movies_casts.Updated_at)
			if err != nil {
				log.Fatalln(err)
			}
			// fmt.Println(movies_casts)
			list_movies_casts = append(list_movies_casts, movies_casts)
		}
		rows1, _ := r.Queryx("select g.* from movies_genres mg left join genres g on mg.id_genre=g.id_genre  where mg.id_movie=$1", movies_data.Id_movie)
		for rows1.Next() {
			var movies_genres models.Genres
			err1 := rows1.Scan(&movies_genres.Id_genre, &movies_genres.Name_genre, &movies_genres.Created_at, &movies_genres.Updated_at)
			if err1 != nil {
				log.Fatalln(err1)
			}
			// fmt.Println(movies_genres)
			list_movies_genres = append(list_movies_genres, movies_genres)
		}
		rows2, _ := r.Queryx("select d.* from movies m left join directors d on m.id_director=d.id_director where m.id_movie=$1", movies_data.Id_movie)
		for rows2.Next() {
			var movies_directors models.Directors
			err2 := rows2.Scan(&movies_directors.Id_director, &movies_directors.Name_director, &movies_directors.Created_at, &movies_directors.Updated_at)
			if err2 != nil {
				log.Fatalln(err2)
			}
			// fmt.Println(movies_genres)
			list_movies_directors = append(list_movies_directors, movies_directors)
		}

		movies_data.Casts = list_movies_casts
		movies_data.Genres = list_movies_genres
		movies_data.Directors = list_movies_directors

		// fmt.Println(movies_data)
		list_movies = append(list_movies, movies_data)
	}
	rows.Close()
	if len(list_movies) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: list_movies}, nil
}

func (r *Repo_Movies) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.movies WHERE id_movie=$1", id)
	return count_data
}

func (r *Repo_Movies) Get_Count_Data(search string, by_genre string, date string) int {
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(title) like LOWER('%s')`, "%"+search+"%")
	}

	join_genre := ""
	if by_genre == "" {
		by_genre = ""
	} else {
		join_genre = fmt.Sprint(` left join movies_genres mg on mg.id_movie = m.id_movie `)
		by_genre = fmt.Sprintf(` AND mg.id_genre='%s'`, by_genre)
	}

	if date == "" {
		date = ""
	} else {
		dates := strings.Split(date, "-")
		date = fmt.Sprintf(` AND EXTRACT(MONTH FROM m.release_date)='%s' AND EXTRACT(YEAR FROM m.release_date)='%s'`, dates[1], dates[0])
	}

	var id int
	q := fmt.Sprintf(`SELECT count(*) FROM public.movies m %s WHERE TRUE %s %s %s`, join_genre, search, by_genre, date)
	r.Get(&id, r.Rebind(q))
	return id
}

func (r *Repo_Movies) Insert_Data(data *models.Moviesset) (string, error) {
	if data.Image == "" {
		data.Image = "https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png"
	}
	if data.Cover_image == "" {
		data.Cover_image = "https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png"
	}
	tx := r.MustBegin()
	var new_id string
	tx.Get(&new_id, "select gen_random_uuid()")
	data.Id_movie = new_id
	tx.NamedExec(`INSERT INTO public.movies (id_movie, id_director, title, release_date, duration_hour, duration_minute, synopsis, image, cover_image) VALUES(:id_movie, :id_director, :title, :release_date, :duration_hour, :duration_minute, :synopsis, :image, :cover_image);`, data)
	if data.Casts[0] != "" {
		for i := range data.Casts {
			tx.MustExec("INSERT INTO public.movies_casts (id_movie, id_cast) VALUES ($1, $2);", &new_id, &data.Casts[i])
		}
	}
	if data.Genres[0] != "" {
		for i := range data.Genres {
			tx.MustExec("INSERT INTO public.movies_genres (id_movie, id_genre) VALUES ($1, $2);", &new_id, &data.Genres[i])
		}
	}
	if data.Locations != "" {
		arr_loc := strings.Split(strings.Replace(data.Locations, ", ", ",", -1), ",")
		for i := range arr_loc {
			// fmt.Println(arr_loc[i])
			if data.Premiers[0] != "" {
				for ii := range data.Premiers {
					// fmt.Println(data.Premiers[ii])
					var new_id_schedule string
					tx.Get(&new_id_schedule, "select gen_random_uuid()")
					tx.MustExec("INSERT INTO public.schedules (id_schedule,id_movie, id_premier,regency,price,set_date) VALUES ($1, $2,$3,$4,$5,$6);", &new_id_schedule, &new_id, data.Premiers[ii], arr_loc[i], data.Price, data.Set_date)
					if data.Times[0] != "" {
						for iii := range data.Times {
							tx.MustExec("INSERT INTO public.times_schedules (id_schedule,time_schedule) VALUES ($1, $2);", &new_id_schedule, data.Times[iii])
						}
					}
				}
			}
			// fmt.Println("")
		}
	}
	tx.Commit()

	return "add movie data successful", nil
}
func (r *Repo_Movies) Update_Data(data *models.Moviesset) (string, error) {
	if data.Image == "" {
		data.Image = "https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png"
	}
	if data.Cover_image == "" {
		data.Cover_image = "https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png"
	}
	var id string
	id = data.Id_movie

	tx := r.MustBegin()
	tx.NamedExec(`UPDATE public.movies SET title=:title, id_director=:id_director, release_date=:release_date, duration_hour=:duration_hour, duration_minute=:duration_minute, synopsis=:synopsis, image=:image, cover_image=:cover_image WHERE id_movie=:id_movie;`, data)
	tx.MustExec(`DELETE FROM public.times_schedules WHERE id_schedule in (select id_schedule from public.schedules where id_movie=$1) `, &id)
	tx.MustExec(`DELETE FROM public.schedules WHERE id_movie=$1;`, &id)
	tx.MustExec(`DELETE FROM public.movies_casts WHERE id_movie=$1;`, &id)
	tx.MustExec(`DELETE FROM public.movies_genres WHERE id_movie=$1;`, &id)
	if data.Casts[0] != "" {
		for i := range data.Casts {
			tx.MustExec("INSERT INTO public.movies_casts (id_movie, id_cast) VALUES ($1, $2);", &id, &data.Casts[i])
		}
	}
	if data.Genres[0] != "" {
		for i := range data.Genres {
			tx.MustExec("INSERT INTO public.movies_genres (id_movie, id_genre) VALUES ($1, $2);", &id, &data.Genres[i])
		}
	}
	if data.Locations != "" {
		arr_loc := strings.Split(strings.Replace(data.Locations, ", ", ",", -1), ",")
		for i := range arr_loc {
			// fmt.Println(arr_loc[i])
			if data.Premiers[0] != "" {
				for ii := range data.Premiers {
					// fmt.Println(data.Premiers[ii])
					var new_id_schedule string
					tx.Get(&new_id_schedule, "select gen_random_uuid()")
					tx.MustExec("INSERT INTO public.schedules (id_schedule,id_movie, id_premier,regency,price,set_date) VALUES ($1, $2,$3,$4,$5,$6);", &new_id_schedule, &id, data.Premiers[ii], arr_loc[i], data.Price, data.Set_date)
					if data.Times[0] != "" {
						for iii := range data.Times {
							tx.MustExec("INSERT INTO public.times_schedules (id_schedule,time_schedule) VALUES ($1, $2);", &new_id_schedule, data.Times[iii])
						}
					}
				}
			}
			// fmt.Println("")
		}
	}
	tx.Commit()

	return "update movie data successful", nil
}
func (r *Repo_Movies) Delete_Data(data *models.Movies, data2 *models.Movies_Casts, data3 *models.Movies_Genres) (string, error) {
	tx := r.MustBegin()
	_, err4 := tx.NamedExec(`DELETE FROM public.times_schedules WHERE id_schedule in (select id_schedule from public.schedules where id_movie=:id_movie)`, data)
	if err4 != nil {
		return "", err4
	}
	_, err5 := tx.NamedExec(`DELETE FROM public.schedules WHERE id_movie=:id_movie;`, data)
	if err5 != nil {
		return "", err5
	}
	_, err1 := tx.NamedExec(`DELETE FROM public.movies_genres WHERE id_movie=:id_movie;`, data2)
	if err1 != nil {
		return "", err1
	}
	_, err3 := tx.NamedExec(`DELETE FROM public.movies_casts WHERE id_movie=:id_movie;`, data3)
	if err3 != nil {
		return "", err3
	}
	_, err2 := tx.NamedExec(`DELETE FROM public.movies WHERE id_movie=:id_movie;`, data)
	if err2 != nil {
		return "", err2
	}
	tx.Commit()
	return "delete movie data successful", nil
}
