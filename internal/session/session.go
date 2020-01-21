package session

type UserSession struct {
	Token string
}

var CurrentUser = &UserSession{}
