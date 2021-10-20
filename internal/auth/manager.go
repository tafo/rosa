package auth

import "github.com/tafo/rosa/internal/auth/models"

type AccountManager struct {
	repository models.Repository
	bcrypt     models.BcryptWrapper
	jwt        models.JWTWrapper
}

func NewAccountManager(repository models.Repository, bcrypt models.BcryptWrapper, jwt models.JWTWrapper) AccountManager {
	return AccountManager{
		repository: repository,
		bcrypt: bcrypt,
		jwt: jwt,
	}
}

func (am AccountManager) Register(account models.Account) (AuthResponse, error) {
	hashedPassword, err := am.bcrypt.HashPassword(account.Password)
	if err != nil {
		return AuthResponse{}, err
	}

	account.Password = hashedPassword
	if err = am.repository.Insert(&account); err != nil {
		return AuthResponse{}, err
	}

	return am.generateCredentialsForUser(account)
}

func (am AccountManager) generateCredentialsForUser(account models.Account) (AuthResponse, error) {
	token, err := am.jwt.GenerateTokenForUser(account)
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{Token: token}, nil
}