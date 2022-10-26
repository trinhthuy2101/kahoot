package user

type Repository interface {
	Login(request *User) bool
	Register(request *User) bool
}
type repo struct {
	db Database
}

func (r *repo) Login(request *User) bool {
	return r.db.Login(request)
}

func (r *repo) Register(request *User) bool {
	return r.db.Register(request)
}
func NewRepository(db Database) Repository {
	return &repo{
		db: db,
	}
}
