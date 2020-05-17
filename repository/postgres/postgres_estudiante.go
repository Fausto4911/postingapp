package postgres

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"postingapp/model"
	"postingapp/repository"
)

func EstudianteCreate(e model.Estudiante) error {
	// el formato de ($1, $2, $3) es propio de postgres, la comilla ` ` son para poder usar enter en el string
	q := `INSERT INTO estudiantes (name, age, active )
          VALUES ($1, $2, $3)`
	db := repository.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(getNUllString(e.Name), getNUllInt(int64(e.Age)), e.Active) // this statement is use in insert, delete and update statement
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("No rows affected")
	}
	return nil
}

func EstudianteGetAll() (es []model.Estudiante, err error) {
	q := `SELECT id, name, age, active, created_at, updated_at
         FROM estudiantes`

	timeNull := pq.NullTime{}
	intNUll := sql.NullInt64{}
	stringNUll := sql.NullString{}
	boolNUll := sql.NullBool{}

	db := repository.GetConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := model.Estudiante{}
		err = rows.Scan(
			&e.ID,
			&stringNUll,
			&intNUll,
			&boolNUll,
			&e.CreatedAt,
			&timeNull,
		)
		if err != nil {
			return
		}

		e.UpdatedAt = timeNull.Time
		e.Name = stringNUll.String
		e.Age = int16(intNUll.Int64)
		e.Active = boolNUll.Bool

		es = append(es, e)
	}
	return es, nil
}

func EstudianteUpdate(e model.Estudiante) error {
	//q := `SELECT id, name, age, active, created_at, updated_at
	//FROM estudiantes`

	q := `UPDATE estudiantes
         SET name = $1, age = $2, active = $3, updated_at = now()
         WHERE id = $5`

	db := repository.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active, e.ID)
	if err != nil {
		return err
	}

	rows, err := r.RowsAffected()
	if rows != 1 {
		return errors.New("Error: No rows affected")
	}
	return nil
}

func EstudianteDelete (id int) error {
	q := `DELETE FROM estudiantes WHERE id = $1`

	db := repository.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	rows, err := r.RowsAffected()
	if rows != 1 {
		return errors.New("Error: No rows affected")
	}
	return nil

}

//cast zero values from struct to null values and send it to DB
func getNUllInt(i int64) (n sql.NullInt64) {
	if i == 0 {
		n.Valid = false
	} else {
		n.Valid = true
		n.Int64 = int64(i)
	}
	return n
}

func getNUllString(s string) (n sql.NullString) {
	if s == "" {
		n.Valid = false
	} else {
		n.Valid = true
		n.String = s
	}
	return n
}
