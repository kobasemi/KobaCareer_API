package context

import (
	"KobaCareer_API/domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type Context interface {
	DB() *gorm.DB
	InternshipValidate(internship domain.Internships) error
}

type ctx struct {
	getDB func() *gorm.DB
	db    *gorm.DB
}

func New(getDB func() *gorm.DB) Context {
	return &ctx{
		getDB: getDB,
	}
}

func (c ctx) DB() *gorm.DB {
	if c.db != nil {
		return c.db
	}
	return c.getDB()
}

func (c ctx) InternshipValidate(internship domain.Internships) error {
	return validation.ValidateStruct(&internship,
		validation.Field(
			&internship.Company,
			validation.Required.Error("company is required"),
		),
		validation.Field(
			&internship.Title,
			validation.Required.Error("title is required"),
		),
		validation.Field(
			&internship.Salary,
			validation.Required.Error("salary is required"),
		),
		validation.Field(
			&internship.Period,
			validation.Required.Error("period is required"),
		),
		validation.Field(
			&internship.Select,
			validation.Required.Error("select is required"),
		),
		validation.Field(
			&internship.Deadline,
			validation.Required.Error("deadline is required"),
		),
		validation.Field(
			&internship.Contributor,
			validation.Required.Error("contributor is required"),
		),
		validation.Field(
			&internship.Detail,
			validation.Required.Error("detail is required"),
		),
		validation.Field(
			&internship.FutureJob,
			validation.Required.Error("future_job is required"),
		),
		validation.Field(
			&internship.Flow,
			validation.Required.Error("flow is required"),
		),
		validation.Field(
			&internship.Method,
			validation.Required.Error("method is required"),
		),
	)
}
