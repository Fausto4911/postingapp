package model

import "time"

type Estudiante struct {
	ID        int
	Name      string
	Age       int16 // int16 es la contraparte de postgres smallint
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EstudianteRepository interface {
	Store(estudiante Estudiante) error
	GetAll(estudiante Estudiante) ([]Estudiante, error)
	GetById(id int) (Estudiante, error)
	Delete(id int) error
	Update(id int) error
}
