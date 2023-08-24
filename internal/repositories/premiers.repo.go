package repositories

import (
	"errors"
	"math"
	"strconv"

	"github.com/Ravictation/tickitzgolang/config"
	"github.com/Ravictation/tickitzgolang/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repo_Premiers struct {
	*sqlx.DB
}

func NewPremiers(db *sqlx.DB) *Repo_Premiers {
	return &Repo_Premiers{db}
}

func (r *Repo_Premiers) Get_Data(data *models.Premiers, page string, limit string) (*config.Result, error) {
	datas := []models.Premiers{}
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

	r.Select(&datas, `SELECT * FROM public.premiers LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(datas) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: datas, Meta: metas}, nil
}

func (r *Repo_Premiers) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.premiers WHERE id_premier=$1", id)
	return count_data
}

func (r *Repo_Premiers) Get_Count_Data() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.premiers")
	return id
}

func (r *Repo_Premiers) Insert_Data(data *models.Premiers) (string, error) {
	if data.Image == "" {
		data.Image = "https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png"
	}

	query := `INSERT INTO public.premiers(
			name_premier,
			image,
			count_row_seat,
			count_col_seat
		)VALUES(
			:name_premier,
			:image,
			:count_row_seat,
			:count_col_seat
		);`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "add premier data successful", nil
}
func (r *Repo_Premiers) Update_Data(data *models.Premiers) (string, error) {
	if data.Image == "" {
		data.Image = "https://res.cloudinary.com/deaia7unw/image/upload/v1692809165/no-product-image-400x400_hrg7mo.png"
	}

	query := `UPDATE public.premiers SET
			name_premier=:name_premier,
			image=:image,
			count_row_seat=:count_row_seat,
			count_col_seat=:count_col_seat,
			updated_at=now()
			WHERE id_premier=:id_premier;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update premier data successful", nil
}
func (r *Repo_Premiers) Delete_Data(data *models.Premiers) (string, error) {
	query := `DELETE FROM public.premiers WHERE id_premier=:id_premier;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete premier data successful", nil
}
