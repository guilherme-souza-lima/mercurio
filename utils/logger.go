package utils

import (
	myLib_entity "github.com/guilherme-souza-lima/solar-system/entity"
	myLib_logger "github.com/guilherme-souza-lima/solar-system/logger"
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

func LoggerWriting(level, msg string) {
	log := myLib_logger.GetInstance().Logger
	switch level {
	case "Error":
		log.Error(msg)
	case "Warn":
		log.Warn(msg)
	default:
		log.Info(msg)
	}
}
