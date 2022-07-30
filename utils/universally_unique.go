package utils

import "github.com/gofrs/uuid"

func GeneratorUUid() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
