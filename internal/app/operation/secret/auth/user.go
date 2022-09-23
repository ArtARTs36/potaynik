package auth

type User struct {
	IPAddress string
}

func NewUser(addr string) User {
	return User{IPAddress: addr}
}
