package svc

import "placementAPI/store"

type StudentServiceHandler struct {
	datastore store.StudentStorer
}

func NewStudentServiceHandler(student store.StudentStorer) StudentServiceHandler {
	return StudentServiceHandler{datastore: student}
}
