package repository

type BaseRepository[T any] interface {
	Create(entity *T) error
	GetByID(id int) (*T, error)
	Update(entity *T) error
	Delete(id int) error
	List() ([]*T, error)
}
