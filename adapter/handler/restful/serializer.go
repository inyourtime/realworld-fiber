package restful

import "realworld-go-fiber/core/domain"

type User struct {
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

type UserResponse struct {
	User User `json:"user"`
}

func serializeUser(arg domain.User) User {
	return User{
		Email:    arg.Email,
		Username: arg.Username,
		Bio:      arg.Bio,
		Image:    arg.Image,
		Token:    arg.Token,
	}
}

type Profile struct {
	Username  string  `json:"username"`
	Bio       *string `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
}

func serializeProfile(arg domain.User) Profile {
	return Profile{
		Username:  arg.Username,
		Image:     arg.Image,
		Bio:       arg.Bio,
		Following: arg.IsFollowed,
	}
}
