package entities

type Student struct {
	StudentID   string  `json:"studentID"`
	StudentName string  `json:"studentName"`
	DOB         string  `json:"dob"`
	Branch      string  `json:"branch"`
	Comp        Company `json:"comp"`
}
