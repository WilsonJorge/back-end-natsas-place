package user

import (
	"fmt"
	"natsas-place/repository"
)

func DeleteUser(id string) error {
	err := repository.DeleteUserById(id)
	if err != nil {
		return fmt.Errorf("ERROR WHEN DELETE USER %w", err)
	}

	return nil

}
