package auth

type IAuthRepository interface {
	GenerateJWT(ip string) (string, error)
	ClaimToken(token, ip string) error
}
