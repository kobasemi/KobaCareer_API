package domain

import (
	"KobaCareer_API/resource/request"
	"time"
)

type Internships struct {
	ID          uint      `json:"id" gorm:"primary_key;auto_increment"`
	Company     string    `json:"company"`
	Title       string    `json:"title"`
	Salary      uint      `json:"salary"`
	Period      string    `json:"period"`
	Select      string    `json:"select"`
	Deadline    string    `json:"deadline"`
	Contributor string    `json:"contributor"`
	Detail      string    `json:"detail"`
	FutureJob   string    `json:"future_job"`
	Flow        string    `json:"flow"`
	Method      string    `json:"method"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewInternship(dto *request.CreateInternship) (*Internships, error) {
	var internships = Internships{
		Company:     dto.Company,
		Title:       dto.Title,
		Salary:      dto.Salary,
		Period:      dto.Period,
		Select:      dto.Select,
		Deadline:    dto.Deadline,
		Contributor: dto.Contributor,
		Detail:      dto.Detail,
		FutureJob:   dto.FutureJob,
		Flow:        dto.Flow,
		Method:      dto.Method,
	}
	return &internships, nil
}
