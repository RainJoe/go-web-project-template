package listing

// Service provides user listing operations.
type Service interface {
	GetAllUsers() ([]User, error)
}

// Repository provides access to user repository.
type Repository interface {
	// GetAllUsers returns all users
	GetAllUsers() ([]User, error)
}

type service struct {
	uR Repository
}

// NewService creates an listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetAllUsers returns all users
func (s *service) GetAllUsers() ([]User, error) {
	return s.uR.GetAllUsers()
}
