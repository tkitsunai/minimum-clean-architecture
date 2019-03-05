package convert

import "github.com/tkitsunai/minimum-clean-architecture/domain"

type UserJSON struct {
	Name string
}

func ToUserJSON(users domain.Users) []UserJSON {
	res := []UserJSON{}
	for _, user := range users {
		res = append(res, UserJSON{
			Name: user.Name,
		})
	}
	return res
}
