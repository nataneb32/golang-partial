package usage

import (
	"partial/usage/entity"
	"partial"
)

func UpdateUser(puser partial.Partial[entity.User]) error {
	var user entity.User
	err := puser.Apply(&user)
	if err != nil {
		return err
	}

	user = *FindUser(user.ID)

	puser.Apply(&user)
	if err != nil {
		return err
	}

	SaveUser(&user)
	if err != nil {
		return err
	}

	return nil
}
