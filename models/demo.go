package models

import (
	"gopkg.in/mgo.v2/bson"
	"mongo-demo/db"
	"mongo-demo/db/mongo"
)

type TestUser struct {
	Id        int    `bson:"id" json:"id"`                         // 用户ID
	RealName  string `bson:"realName" json:"realName"`             // 真实姓名
	Sex       string `bson:"sex" json:"sex"`                       // 性别
	Birthday  int    `bson:"birthday" json:"birthday"`             // 生日
	Role      string `bson:"role" json:"role"`                     //角色
	OrgId     int    `bson:"orgId" json:"orgId"`                   // 机构ID
	GradeId   int    `bson:"gradeId,omitempty" json:"gradeId"`     // 年级ID
	ClassId   int    `bson:"classId,omitempty" json:"classId"`     // 班级ID
	ClassName string `bson:"className,omitempty" json:"className"` // 班级名称
}

type Test struct {
	Id   bson.ObjectId `bson:"_id"`  // 测试ID
	User *TestUser     `bson:"user"` // 用户信息

	ActivityId    bson.ObjectId `bson:"activityId,omitempty"`    // 活动ID
	NewActivityId int           `bson:"newActivityId,omitempty"` // 新活动ID
	Platform      string        `bson:"platform"`                //来源
	EventsCount   int           `bson:"eventsCount"`             //提交事件次数
	EventsToken   string        `bson:"eventsToken"`             //事件token
	ScaleId       bson.ObjectId `bson:"scaleId"`                 // 量表ID
	VersionId     bson.ObjectId `bson:"versionId,omitempty"`     // 量表版本ID

	ResetStatus bool      `bson:"resetStatus"` //重测状态  false 未重测   true 已重测
	Duration    int       `bson:"duration"`    // 测试时长
	Ext         *bson.Raw `bson:"ext,omitempty"`

	SubmitAt int `bson:"submitAt,omitempty"` // 提交时间
	CreateAt int `bson:"createAt"`           // 创建时间
	UpdateAt int `bson:"updateAt"`           // 最后更新时间
}

const TABLE = "test" //mongoDB表名

// equal和in查询
func GetUserActivities(userId, orgId int, activityIds []int) ([]*Test, error) {
	var result []*Test
	query := bson.M{
		"user.orgId":    orgId,
		"newActivityId": mongo.In(activityIds),
		"user.id":       userId,
	}
	err := db.RunMongoTask(func(db *mongo.Mongo) error {
		return db.Find(TABLE, query, &result)
	})
	return result, err
}

// between查询
func GetActivitiesDuring(userId, orgId int, startAt, endAt int64) ([]*Test, error) {
	var result []*Test
	query := bson.M{
		"user.orgId": orgId,
		"user.id":    userId,
		"submitAt":   mongo.Between(startAt, endAt),
	}
	err := db.RunMongoTask(func(db *mongo.Mongo) error {
		// -submitAt表示按字段倒序
		return db.FindSortPage(TABLE, query, "-submitAt", 4, 2, &result)
	})
	return result, err
}

// update
func UpdateResetStatus(userId, orgId int, resetStatus bool) error {
	query := bson.M{
		"user.orgId": orgId,
		"user.id":    userId,
	}
	update := mongo.Set("resetStatus", resetStatus)
	err := db.RunMongoTask(func(db *mongo.Mongo) error {
		return db.Update(TABLE, query, update)
	})
	return err
}

// distinct aggregate
func DistinctRoles(orgId int) (roles []string, err error) {
	query := bson.M{
		"user.orgId": orgId,
	}
	var aggData []*mongo.DistinctValue
	err = db.RunMongoTask(func(db *mongo.Mongo) error {
		aggs := mongo.NewMongoAggs().Match(query).Distinct("user.role")
		return db.Aggregate(TABLE, aggs, &aggData)
	})
	for _, v := range aggData {
		roles = append(roles, v.Id.(string))
	}
	return
}

// avg value aggregate
func GetAvgDurationByOrg() (resMap map[int]float64, err error) {
	var aggData []*mongo.GroupAggValue
	err = db.RunMongoTask(func(db *mongo.Mongo) error {
		aggs := mongo.NewMongoAggs().Avg("duration", "user.orgId")
		return db.Aggregate(TABLE, aggs, &aggData)
	})
	resMap = map[int]float64{}
	for _, v := range aggData {
		k := v.Id.(int)
		resMap[k] = v.Value
	}
	return
}
