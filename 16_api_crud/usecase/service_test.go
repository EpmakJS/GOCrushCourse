package usecase

import (
	"16_api_crud/entities"
	"16_api_crud/usecase/mocks"
	"log"
	"testing"
)

func TestService_GetAll(t *testing.T) {

	repository := mocks.RepositoryMock{}
	service := NewService(&repository)

	t.Run("", func(t *testing.T) {
		products := []*entities.Product{
			{Id: 0,
				Model:   "Iphone",
				Company: "Apple",
				Price:   12000,
			},
			{Id: 1,
				Model:   "S",
				Company: "AppSsle",
				Price:   12000,
			},
		}

		repository.On("List").Return(products, nil)

		res, err := service.GetAll()
		log.Printf("%v", res)
		log.Printf("%v", err)
		repository.AssertCalled(t, "List")
	})

}
