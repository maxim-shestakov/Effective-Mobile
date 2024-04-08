package db

//go:generate mockgen -destination=../mocks/mock_storage.go -package=mocks . Storage

type Storage interface{
	AddCar(Car) error
	DeleteCar(id string) error
	UpdateCar(Car) error
	GetCars() ([]Car, error)
	AddOwner(Owner) error
}


//swagger:model
type Car struct {
	ID      int    `json:"id" example:"3"`
	Regnum  string `json:"regnum" example:"A111AA155"`
	Mark    string `json:"mark" example:"BMW"`
	Model   string `json:"model" example:"X5"`
	Year    int    `json:"year" example:"2015"`
	OwnerID int    `json:"owner_id" example:"1"`
}

//swagger:model
type Owner struct {
	ID         int    `json:"id" example:"4"`
	Name       string `json:"name" example:"Max"`
	Surname    string `json:"surname" example:"Shestakov"`
	Patronymic string `json:"patronymic" example:"Olegovich"`
}

//swagger:model
type StatusOKMessage struct {
	Message string `json:"status" example:"ok"`
}

//swagger:model
type StatusBadRequestMessage struct {
	Message string `json:"message" example:"Error: bad request"`
}

//swagger:model
type StatusNotFoundMessage struct {
	Message string `json:"message" example:"Error: not found"`
}

//swagger:model
type StatusInternalServerErrorMessage struct {
	Message string `json:"message" example:"Error: internal server error"`
}


