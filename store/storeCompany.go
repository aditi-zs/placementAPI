package store

import "database/sql"

type CompanyStorer struct {
	db *sql.DB
}

func NewCompanyStorer(d *sql.DB) CompanyStorer {
	return CompanyStorer{db: d}
}
