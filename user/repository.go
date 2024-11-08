package user

type repository interface {
	Save(user User) (User, error)
}
