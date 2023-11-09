package usecase

import (
	"KobaCareer_API/domain"
	"KobaCareer_API/packages/context"
	"KobaCareer_API/resource/request"
)

type InternshipInputPort interface {
	Create(ctx context.Context, req request.CreateInternship) error
	GetByID(ctx context.Context, id uint) error
	GetAll(ctx context.Context) error
	Update(ctx context.Context, req request.UpdateInternship) error
	Delete(ctx context.Context, id uint) error
}

type InternshipOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.Internships) error
	GetAll(res []*domain.Internships) error
	Update(res *domain.Internships) error
	Delete() error
}

type InternshipRepository interface {
	Create(ctx context.Context, internship *domain.Internships) (uint, error)
	GetByID(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]*domain.Internships, error)
	Update(ctx context.Context, internship *domain.Internships) error
	Delete(ctx context.Context, id uint) error
}

type internship struct {
	outputPort     InternshipOutputPort
	internshipRepo InternshipRepository
}

type InternshipInputFactory func(outputPort InternshipOutputPort) InternshipInputPort

func NewInternshipInputFactory(ir InternshipRepository) InternshipInputFactory {
	return func(o InternshipOutputPort) InternshipInputPort {
		return &internship{
			outputPort:     o,
			internshipRepo: ir,
		}
	}
}

func (i internship) Create(ctx context.Context, req request.CreateInternship) error {
	newInternship, err := domain.NewInternship(req)
	if err != nil {
		return err
	}

}

func (i internship) GetByID(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (i internship) GetAll(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (i internship) Update(ctx context.Context, req request.UpdateInternship) error {
	//TODO implement me
	panic("implement me")
}

func (i internship) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
