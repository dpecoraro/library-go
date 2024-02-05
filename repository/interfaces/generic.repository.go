package repository

type Repository[T any] interface {
	Add(entity *T)
	Update(entity *T)
	Delete(id uint)
	Get(id uint)
	List()
}
