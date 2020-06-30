package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Mongo struct {
		db *mgo.Database
	}
	Config struct {
		MgoUrl     string
		MgoName    string
		MgoMaxOpen int
	}
)

func NewMongo(db *mgo.Database) *Mongo {
	return &Mongo{db}
}

func (mongo *Mongo) GetDB() *mgo.Database {
	return mongo.db
}

func (mongo *Mongo) Find(c string, query interface{}, result interface{}) error {
	return mongo.db.C(c).Find(query).All(result)
}

func (mongo *Mongo) FindSort(c string, query interface{}, sort string, result interface{}) error {
	return mongo.db.C(c).Find(query).Sort(sort).All(result)
}

func (mongo *Mongo) FindProject(c string, query interface{}, projection interface{}, result interface{}) error {
	return mongo.db.C(c).Find(query).Select(projection).All(result)
}

func (mongo *Mongo) FindSortProject(c string, query interface{}, sort string, projection interface{},
	result interface{}) error {
	return mongo.db.C(c).Find(query).Sort(sort).Select(projection).All(result)
}

func (mongo *Mongo) Count(c string, query interface{}) (int, error) {
	return mongo.db.C(c).Find(query).Count()
}

func (mongo *Mongo) FindSortPage(c string, query interface{}, sort string, pageSize, pageNum int,
	result interface{}) error {
	return mongo.db.C(c).Find(query).Sort(sort).Skip(pageSize * (pageNum - 1)).Limit(pageSize).All(result)
}

func getFieldsSelector(fields ...string) bson.M {
	result := make(bson.M, len(fields))
	for _, field := range fields {
		result[field] = 1
	}
	return result
}

func (mongo *Mongo) FindField(c string, query interface{}, result interface{}, fields ...string) error {
	return mongo.db.C(c).Find(query).Select(getFieldsSelector(fields...)).All(result)
}

func (mongo *Mongo) FindOne(c string, query interface{}, result interface{}) error {
	return mongo.db.C(c).Find(query).One(result)
}

func (mongo *Mongo) FindOneProject(c string, query interface{}, projection interface{}, result interface{}) error {
	return mongo.db.C(c).Find(query).Select(projection).One(result)
}

func (mongo *Mongo) Update(c string, selector interface{}, update interface{}) error {
	return mongo.db.C(c).Update(selector, update)
}

func (mongo *Mongo) Insert(c string, docs ...interface{}) error {
	return mongo.db.C(c).Insert(docs...)
}

func (mongo *Mongo) UpsertId(c string, id interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return mongo.db.C(c).UpsertId(id, update)
}

func (mongo *Mongo) Close() {
	mongo.db.Session.Close()
}
