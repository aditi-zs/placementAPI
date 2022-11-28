package store

import "database/sql"

type StudentStorer struct {
	db *sql.DB
}

func NewStudentStorer(d *sql.DB) StudentStorer {
	return StudentStorer{db: d}
}
