package student

import (
	"time"
)

type Student struct {
	uuid string

	firstName string
	lastName  string

	dateOfBirth time.Time

	averageScore uint
	address      string
}

//
//func NewStudent(uuid string, firstName string, lastName string, dateOfBirth time.Time, averageScore uint, address string) (*Student, error) {
//	if uuid == "" {
//		return nil, errors.New("empty uuid")
//	}
//
//	if firstName == "" {
//		return nil, errors.New("empty firstName")
//	}
//
//	if lastName == "" {
//		return nil, errors.New("empty lastName")
//	}
//
//	if averageScore < 0 || averageScore > 10 {
//		return nil, errors.New("invalid score")
//	}
//
//	return &Student{
//		uuid:         uuid,
//		firstName:    firstName,
//		lastName:     lastName,
//		dateOfBirth:  dateOfBirth,
//		averageScore: averageScore,
//		address:      address,
//	}, nil
//}
