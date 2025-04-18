package utils

import (
	"github.com/google/uuid"
	"github.com/lucsky/cuid"
)

func GenerateAPIKey() string {
	uuidV6, err := uuid.NewV6()

	if err != nil {
		panic(err)
	}

	apiKey := "" + "pay_" + cuid.New() + "@" + cuid.New() + uuidV6.String()
	return apiKey

}
