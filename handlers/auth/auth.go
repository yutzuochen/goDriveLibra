package auth

type Manager interface {
	Login(username, pwd string) (string, error)
}
