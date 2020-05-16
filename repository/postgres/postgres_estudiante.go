package postgres

import (
	"errors"
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

	r, err := stmt.Exec(e.Name, e.Age, e.Active) // this statement is use in insert, delete and update
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("No rows affected")
	}
	return nil
}
