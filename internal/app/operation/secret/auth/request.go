package auth

type AuthorizeRequest struct {
	UserFactorValue string
	User            User
}

func NewAuthorizeRequest(factorVal string, user User) AuthorizeRequest {
	return AuthorizeRequest{
		UserFactorValue: factorVal,
		User:            user,
	}
}
