package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type InsertRepository struct {
	conn                *gorm.DB
	mongoDB             *mongo.Database
	utilitiesRepository *UtilitiesRepository
}

func NewInsertRepository(conn *gorm.DB, mongoDB *mongo.Database, utilitiesRepository *UtilitiesRepository) *InsertRepository {
	return &InsertRepository{
		conn:                conn,
		mongoDB:             mongoDB,
		utilitiesRepository: utilitiesRepository,
	}
}

// func (i *InsertRepository) InsertDataBaseWithKey(key string, typeDatabase enums.Database) error {
// 	if typeDatabase == enums.MONGO {

// 	} else {
// 		getData, err := i.utilitiesRepository.CheckKeyData(key)

// 		if err != nil {
// 			return err
// 		}
// 	}
// }
