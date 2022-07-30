package utils

import "github.com/gofrs/uuid"

func GeneratorUUid() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err.Error()
	}
	return uuid.String()
}
