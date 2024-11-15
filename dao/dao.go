package dao

import (
	"context"
	"fmt"
	"pkg/conf"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlClient   *gorm.DB
	MongoDBClient *mongo.Client
	ctx           = context.Background()
)

func init() {
	mongoDB()
	mySql()
}
func mongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://" + conf.MongoDBAddr + ":" + conf.MongoDBPort)
	var err error
	MongoDBClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("mongodb connect error")
		panic(err)
	}
}
func mySql() {
	var err error
	gsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	MysqlClient, err = gorm.Open(mysql.Open(gsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect error")
		panic(err)
	}
	if MysqlClient.Error != nil {
		fmt.Printf("database error %v", MysqlClient.Error)
	}
	db, _ := MysqlClient.DB()
	db.SetMaxIdleConns(10)  // 链接池
	db.SetMaxOpenConns(100) // 打开最大连接
	db.SetConnMaxLifetime(time.Hour)
}
func InsertDocument(db, collection string, document interface{}) error {
	Collection := MongoDBClient.Database(db).Collection(collection)
	_, err := Collection.InsertOne(ctx, document)
	return err
}
func QueryDocument(db, collection string, filter bson.M) ([]bson.M, error) {
	Collection := MongoDBClient.Database(db).Collection(collection)
	cur, err := Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = cur.All(ctx, &results)
	return results, err
}
func UpdateDocument(db, collection string, filter, update bson.M) error {
	Collection := MongoDBClient.Database(db).Collection(collection)
	_, err := Collection.UpdateMany(ctx, filter, bson.M{
		"$set": update,
	})
	return err
}
func DeleteDocuments(db, collection string, filter bson.M) error {
	Collection := MongoDBClient.Database(db).Collection(collection)
	_, err := Collection.DeleteMany(ctx, filter)
	return err
}
