package student

import (
	"net/http"
	"placementAPI/svc"
)

type StudentHandler struct {
	service svc.StudentServiceHandler
}

func NewStudentHandler(stu svc.StudentServiceHandler) StudentHandler {
	return StudentHandler{service: stu}
}

//func (s StudentHandler) StudentsHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		s.getStudentData(w, r)
//	case http.MethodPost:
//		s.postStudentData(w, r)
//	case http.MethodPut:
//		s.updateStudentData(w, r)
//	case http.MethodDelete:
//		s.deleteStudentData(w, r)
//	default:
//		w.WriteHeader(http.StatusMethodNotAllowed)
//	}
//}

func (s StudentHandler) GetStudentData(w http.ResponseWriter, r *http.Request) {

}

func (s StudentHandler) PostStudentData(w http.ResponseWriter, r *http.Request) {

}

func (s StudentHandler) UpdateStudentData(w http.ResponseWriter, r *http.Request) {

}

func (s StudentHandler) DeleteStudentData(w http.ResponseWriter, r *http.Request) {

}
