package domain

import (
	"time"
)

type Internships struct {
	ID          uint      `json:"id" gorm:"primary_key;auto_increment"`
	Company     string    `json:"company"`
	Title       string    `json:"title"`
	Salary      uint      `json:"salary"`
	Period      string    `json:"period"`
	Selection   string    `json:"select"`
	Deadline    string    `json:"deadline"`
	Contributor string    `json:"contributor"`
	Detail      string    `json:"detail"`
	FutureJob   string    `json:"future_job"`
	Flow        string    `json:"flow"`
	Method      string    `json:"method"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InternshipsResponse struct {
	ID          uint      `json:"id" gorm:"primary_key;auto_increment"`
	Company     string    `json:"company"`
	Title       string    `json:"title"`
	Salary      uint      `json:"salary"`
	Period      string    `json:"period"`
	Selection   string    `json:"select"`
	Deadline    string    `json:"deadline"`
	Contributor string    `json:"contributor"`
	Detail      string    `json:"detail"`
	FutureJob   string    `json:"future_job"`
	Flow        string    `json:"flow"`
	Method      string    `json:"method"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
