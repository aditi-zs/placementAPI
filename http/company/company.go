package company

import (
	"net/http"
	"placementAPI/svc"
)

type CompanyHandler struct {
	service svc.CompanyServiceHandler
}

func NewCompanyHandler(comp svc.CompanyServiceHandler) CompanyHandler {
	return CompanyHandler{service: comp}
}

//	func (c CompanyHandler) CompHandler(w http.ResponseWriter, r *http.Request) {
//		switch r.Method {
//		case http.MethodGet:
//			c.getCompanyData(w, r)
//		case http.MethodPost:
//			c.postCompanyData(w, r)
//		case http.MethodPut:
//			c.updateCompanyData(w, r)
//		case http.MethodDelete:
//			c.deleteCompanyData(w, r)
//		default:
//			w.WriteHeader(http.StatusMethodNotAllowed)
//		}
//	}
func (c CompanyHandler) GetCompanyData(w http.ResponseWriter, r *http.Request) {

}

func (c CompanyHandler) PostCompanyData(w http.ResponseWriter, r *http.Request) {

}

func (c CompanyHandler) UpdateCompanyData(w http.ResponseWriter, r *http.Request) {

}

func (c CompanyHandler) DeleteCompanyData(w http.ResponseWriter, r *http.Request) {

}
