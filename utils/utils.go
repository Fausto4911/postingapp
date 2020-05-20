package utils

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
)

func GetNUllTime(t time.Time) (n pq.NullTime) {
	if t.IsZero() {
		n.Valid = false
	} else {
		n.Valid = true
		n.Time = t
	}
	return
}

func GetNUllString(s string) (n sql.NullString) {
	if s == "" {
		n.Valid = false
	} else {
		n.Valid = true
		n.String = s
	}
	return
}
func GetNUllInt(i int64) (n sql.NullInt64) {
	if i == 0 {
		n.Valid = false
	} else {
		n.Valid = true
		n.Int64 = int64(i)
	}
	return
}
