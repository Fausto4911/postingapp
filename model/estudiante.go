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
