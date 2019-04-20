package adding

// Service provides user adding operations.
type Service interface {
	AddUser(User) error
}

// Repository provides access to user repository.
type Repository interface {
	// AddUser saves a given user to the repository.
	AddUser(User) error
}

type service struct {
	uR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser adds the given user(s) to the database
func (s *service) AddUser(u User) error {
	return s.uR.AddUser(u)
}
