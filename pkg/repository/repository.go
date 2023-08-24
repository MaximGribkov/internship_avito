package repository

type Logics interface {
	//CreateUser
}

type Repository struct {
	Logics
}

func NewRepository() *Repository {
	return &Repository{}
}
