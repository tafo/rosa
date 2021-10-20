package models

type Repository interface {
	Insert(account *Account) error
}

type BcryptWrapper interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type JWTWrapper interface {
	GenerateTokenForUser(account Account) (string, error)
	ExtractUserFromToken(token string) (Account, error)
}