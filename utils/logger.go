package utils

import (
	myLib_entity "github.com/guilherme-souza-lima/solar-system/entity"
	"ssMercurio/entities"
	"time"
)

func MappingLoggerElastic(statusCode int, level, message, id string) myLib_entity.MappingElastic {
	return myLib_entity.MappingElastic{
		StatusCode: statusCode,
		Level:      level,
		Message: myLib_entity.MessageElastic{
			Message: message,
			Local:   entities.NAME_SYSTEM,
		},
		Date: time.Now(),
		User: myLib_entity.UserElastic{ID: id},
	}
}
