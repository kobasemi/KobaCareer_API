package presenter

import (
	"KobaCareer_API/domain"
	"KobaCareer_API/usecase"
	"github.com/gin-gonic/gin"
)

type internship struct {
	c *gin.Context
}

type InternshipOutputFactory func(c *gin.Context) usecase.InternshipOutputPort

func NewInternshipOutputFactory() InternshipOutputFactory {
	return func(c *gin.Context) usecase.InternshipOutputPort {
		return &internship{c: c}
	}
}

func (i internship) Create(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (i internship) GetByID(res *domain.Internships) error {
	//TODO implement me
	panic("implement me")
}

func (i internship) GetAll(res []*domain.Internships) error {
	//TODO implement me
	panic("implement me")
}

func (i internship) Update(res *domain.Internships) error {
	//TODO implement me
	panic("implement me")
}

func (i internship) Delete() error {
	//TODO implement me
	panic("implement me")
}
