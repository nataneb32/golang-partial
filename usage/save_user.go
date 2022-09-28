package usage

import "partial/usage/entity"

func SaveUser(u *entity.User) {
	users[u.ID] = *u
}
