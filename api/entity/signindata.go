package entity

const (
	GroupNameKey = "GroupName"
	UserIDKey    = "UserID"
	UserNameKey  = "UserName"
)

type SigninData struct {
	GroupName string
	UserID    string
	UserName  string
}
