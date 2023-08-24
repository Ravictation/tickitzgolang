package repositories

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo_Movies struct {
	*sqlx.DB
}

func NewMovies(db *sqlx.DB) *Repo_Movies {
	return &Repo_Movies{db}
}

func (r *Repo_Movies) Get_Data(data *models.Movies, page string, limit string, search string, orderby string) (*config.Result, error) {
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

	count_data := r.Get_Count_Data(search)

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
		search = fmt.Sprintf(` AND LOWER(title) like LOWER('%s')`, "%"+search+"%")
	}
	if orderby == "" {
		orderby = ""
	} else {
		orderby = fmt.Sprintf(` ORDER BY %s`, orderby)
	}
	q := fmt.Sprintf(`select id_movie, title, release_date, duration_hour, duration_minute, synopsis, image, cover_image, created_at, updated_at from public.movies WHERE TRUE %s %s LIMIT %d OFFSET %d`, search, orderby, limit_int, offset)
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

func (r *Repo_Movies) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.movies WHERE id_movie=$1", id)
	return count_data
}

func (r *Repo_Movies) Get_Count_Data(search string) int {
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(name_product) like LOWER('%s')`, "%"+search+"%")
	}
	var id int
	q := fmt.Sprintf(`SELECT count(*) FROM public.movies WHERE TRUE %s`, search)
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
	tx.Commit()

	return "update movie data successful", nil
}
func (r *Repo_Movies) Delete_Data(data *models.Movies, data2 *models.Movies_Casts, data3 *models.Movies_Genres) (string, error) {
	tx := r.MustBegin()
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
