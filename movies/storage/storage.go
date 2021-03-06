package storage

import "github.com/irahardianto/monorepo-microservices/movies/model"

type Storage interface {
	GetAll() []model.Movie
	Create(movie *model.Movie) error
	GetById(id string) (movie model.Movie, err error)
	Delete(id string) error
	Ping() error
}
