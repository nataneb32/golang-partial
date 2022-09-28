package usage

import "partial/usage/entity"

func FindUser(id uint) *entity.User {
 	user := users[id]
	return &user
}
