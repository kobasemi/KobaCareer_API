package http

import (
	"KobaCareer_API/adapter/presenter"
	"KobaCareer_API/packages/context"
	"KobaCareer_API/resource/request"
	"KobaCareer_API/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type internship struct {
	inputFactory   usecase.InternshipInputFactory
	outputFactory  func(c *gin.Context) usecase.InternshipOutputPort
	internshipRepo usecase.InternshipRepository
}

func NewInternship(inputFactory usecase.InternshipInputFactory, outputFactory presenter.InternshipOutputFactory) {
	router := gin.Default()

	handler := internship{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	internshipsGroup := router.Group("/internships")
	internshipsGroup.POST("create", CreateAdapter(handler))
	internshipsGroup.GET("", handler.GetAll)
	internshipsGroup.GET(":id", handler.GetByID)
	internshipsGroup.PUT(":id", handler.Update)
	internshipsGroup.DELETE(":id", handler.Delete)
}

func bind(c *gin.Context, request interface{}) (ok bool) {
	if err := c.BindJSON(request); err != nil {
		c.Status(http.StatusBadRequest)
		return false
	} else {
		return true
	}
}

func (i internship) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreateInternship

	if !bind(c, &req) {
		return nil
	}

	outputPort := i.outputFactory(c)
	inputPort := i.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}
