package repositories

import (
	"errors"
	"math"
	"strconv"

	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo_Casts struct {
	*sqlx.DB
}

func NewCasts(db *sqlx.DB) *Repo_Casts {
	return &Repo_Casts{db}
}

func (r *Repo_Casts) Get_Data(data *models.Casts, page string, limit string) (*config.Result, error) {
	datas := []models.Casts{}
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

	count_data := r.Get_Count_Data()

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

	r.Select(&datas, `SELECT * FROM public.casts LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(datas) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: datas, Meta: metas}, nil
}

func (r *Repo_Casts) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.casts WHERE id_cast=$1", id)
	return count_data
}

func (r *Repo_Casts) Get_Count_Data() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.casts")
	return id
}

func (r *Repo_Casts) Insert_Data(data *models.Casts) (string, error) {
	query := `INSERT INTO public.casts(
			name_cast
		)VALUES(
			:name_cast
		);`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "add cast data successful", nil
}
func (r *Repo_Casts) Update_Data(data *models.Casts) (string, error) {
	query := `UPDATE public.casts SET
			name_cast=:name_cast,
			updated_at=now()
			WHERE id_cast=:id_cast;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update cast data successful", nil
}
func (r *Repo_Casts) Delete_Data(data *models.Casts) (string, error) {
	query := `DELETE FROM public.casts WHERE id_cast=:id_cast;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete cast data successful", nil
}
