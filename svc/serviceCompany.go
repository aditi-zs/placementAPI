package svc

import (
	"placementAPI/store"
)

type CompanyServiceHandler struct {
	datastore store.CompanyStorer
}

func NewCompanyServiceHandler(company store.CompanyStorer) CompanyServiceHandler {
	return CompanyServiceHandler{datastore: company}
}
