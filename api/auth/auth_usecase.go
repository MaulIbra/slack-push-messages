package auth

type IAuthUseCase interface {
	GetToken(ip string) (string, error)
}
