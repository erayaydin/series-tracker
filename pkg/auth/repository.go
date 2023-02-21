package auth

type Repository interface {
	Login(*Login) bool
	Register(*Register) bool
}
