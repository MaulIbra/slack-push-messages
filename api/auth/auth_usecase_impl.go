package auth

type authUseCase struct {
	authRepo IAuthRepository
}

func (a authUseCase) GetToken(ip string) (string, error) {
	return a.authRepo.GenerateJWT(ip)
}

func NewAuthUseCase(authRepo IAuthRepository) IAuthUseCase {
	return &authUseCase{authRepo: authRepo}
}
