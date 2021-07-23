package usecase

import "16_api_crud/entities"

// Reader interface
type Reader interface {
	//Get(id entities.Id) (*entities.Product, error)
	//Search(query string) ([]*entities.Product, error)

	List() ([]*entities.Product, error)
}

// Writer interface
type Writer interface {
	//Create(e *entities.Product) (entities.Id, error)
	//Update(e *entities.Product) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// Getter interface
type Getter interface {
	GetAll() ([]*entities.Product, error)
}
