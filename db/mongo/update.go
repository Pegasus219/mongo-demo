package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	SET        = "$set"
	UNSET      = "$unset"
	INC        = "$inc"
	PUSH       = "$push"
	ADD_TO_SET = "$addToSet"
	EACH       = "$each"
)

func Set(key string, value interface{}) bson.M {
	return bson.M{
		SET: bson.M{
			key: value,
		},
	}
}

func SetMany(m bson.M) bson.M {
	return bson.M{
		SET: m,
	}
}

func Inc(key string, value interface{}) bson.M {
	return bson.M{
		INC: bson.M{
			key: value,
		},
	}
}

func IncMany(m bson.M) bson.M {
	return bson.M{
		INC: m,
	}
}

func Unset(keys ...string) bson.M {
	result := bson.M{}
	for _, key := range keys {
		result[key] = ""
	}
	return bson.M{
		UNSET: result,
	}
}

func Push(key string, value interface{}) bson.M {
	return bson.M{
		PUSH: bson.M{
			key: value,
		},
	}
}

func PushMany(m bson.M) bson.M {
	return bson.M{
		PUSH: m,
	}
}

func AddToSet(key string, value interface{}) bson.M {
	return bson.M{
		ADD_TO_SET: bson.M{
			key: value,
		},
	}
}

func AddToSetMany(m bson.M) bson.M {
	return bson.M{
		ADD_TO_SET: m,
	}
}

func Each(elements ...interface{}) bson.M {
	return bson.M{
		EACH: elements,
	}
}
