package user

type UserRepository interface {
	Create()
	Update()
	Delete()
	GetById()
	GetUsers()
}
