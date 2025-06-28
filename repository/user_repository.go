package main

import (
	"natsas-place/database"
)

func Create() {
	return database.Instance.Create()
}
