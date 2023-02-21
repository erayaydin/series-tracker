package auth

type Service interface {
	Login(username string, password string) bool
	Register(username string, password string, passwordConfirmation string, name string) bool
}

type authService struct {
	repo Repository
}

func AuthService(repo Repository) Service {
	return &authService{
		repo: repo,
	}
}

func (svc *authService) Login(username string, password string) bool {
	return svc.repo.Login(&Login{Username: username, Password: password})
}

func (svc *authService) Register(username string, password string, passwordConfirmation string, name string) bool {

	if password != passwordConfirmation {
		return false
	}
	
	return svc.repo.Register(&Register{Name: name, Username: username, Password: password})
}
