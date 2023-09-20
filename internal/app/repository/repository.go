package repository

type Repository interface {
	Create(object any) (int error)
	Read(index int) (any error)
	Update(index int, object any) (any error)
	Delete(index int) error
}
