package models

type User struct {
	Subject        string `mapstructure:"sub"`
	Email          string `mapstructure:"email"`
	EmailVerified  bool   `mapstructure:"email_verified"`
	NotValidBefore int64  `mapstructure:"nbf"`
	Name           string `mapstructure:"name"`
	Picture        string `mapstructure:"picture"`
	GivenName      string `mapstructure:"given_name"`
	FamilyName     string `mapstructure:"family_name"`
	Locale         string `mapstructure:"locale"`
	IssuedAt       int64  `mapstructure:"iat"`
	ExpiresAt      int64  `mapstructure:"exp"`
}

var users []*User = []*User{}

func AddUser(user *User) {
	users = append(users, user)
}

func GetUser(sub string) *User {
	for _, user := range users {
		if user.Subject == sub {
			return user
		}
	}
	return nil
}
