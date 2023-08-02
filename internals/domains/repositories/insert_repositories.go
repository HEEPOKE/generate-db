package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type InsertRepository struct {
	conn    *gorm.DB
	mongoDB *mongo.Database
}

func NewInsertRepository(conn *gorm.DB, mongoDB *mongo.Database) *InsertRepository {
	return &InsertRepository{
		conn:    conn,
		mongoDB: mongoDB,
	}
}