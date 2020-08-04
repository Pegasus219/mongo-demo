package db

import (
	"gopkg.in/mgo.v2"
	"mongo-demo/db/mongo"
)

var MongoConfig = &mongo.Config{
	MgoUrl:  "wanglei:PA9k9DPuy778gx8k@localhost:27017", //mongo地址
	MgoName: "assessment",                               //库名
}

var mongoSession *mgo.Session

func newMongoSession() (*mgo.Session, error) {
	if mongoSession == nil {
		var err error
		mongoSession, err = mgo.Dial(MongoConfig.MgoUrl)
		if err != nil {
			return nil, err
		}
		mongoSession.SetMode(mgo.Strong, false)
	}

	return mongoSession.Copy(), nil
}

func RunMongoTask(task func(*mongo.Mongo) error) error {
	session, err := newMongoSession()
	defer session.Close()

	if err != nil {
		return err
	}
	return task(mongo.NewMongo(session.DB(MongoConfig.MgoName)))
}
