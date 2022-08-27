package utils

import (
	myLib_entity "github.com/guilherme-souza-lima/solar-system/entity"
	myLib_logger "github.com/guilherme-souza-lima/solar-system/logger"
	"ssMercurio/entities"
	"time"
)

func MappingLoggerElasticNew(statusCode int, url, message, idUser, nameUser, nickUser string) myLib_entity.History {
	return myLib_entity.History{
		ID:         GeneratorUUid(),
		URL:        url,
		StatusCode: statusCode,
		Message:    message,
		CreateAt:   time.Now(),
		User: myLib_entity.User{
			ID:   idUser,
			Name: nameUser,
			Nick: nickUser,
		},
	}
}

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
