package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	EXISTS = "$exists"
	TYPE   = "$type"

	MOD = "mod"

	NE  = "$ne"
	GT  = "$gt"
	GTE = "$gte"
	LT  = "$lt"
	LTE = "$lte"
	IN  = "$in"
	NIN = "$nin"

	AND = "$and"
	NOT = "$not"
	NOR = "$nor"
	OR  = "$or"

	ELEM_MATCH = "$elemMatch"
)

func Merge(operators ...bson.M) bson.M {
	result := bson.M{}
	for _, operator := range operators {
		for key, value := range operator {
			result[key] = value
		}
	}
	return result
}

func And(operators ...bson.M) bson.M {
	return bson.M{
		AND: operators,
	}
}

func Not(value interface{}) bson.M {
	return bson.M{
		NOT: value,
	}
}

func Nor(operators ...bson.M) bson.M {
	return bson.M{
		NOR: operators,
	}
}

func Or(operators ...bson.M) bson.M {
	return bson.M{
		OR: operators,
	}
}

func NotEqual(value interface{}) bson.M {
	return bson.M{
		NE: value,
	}
}

func GreaterThan(value interface{}) bson.M {
	return bson.M{
		GT: value,
	}
}

func GreaterThanOrEqual(value interface{}) bson.M {
	return bson.M{
		GTE: value,
	}
}

func LessThan(value interface{}) bson.M {
	return bson.M{
		LT: value,
	}
}

func LessThanOrEqual(value interface{}) bson.M {
	return bson.M{
		LTE: value,
	}
}

func Between(from, to interface{}) bson.M {
	return bson.M{
		GTE: from,
		LTE: to,
	}
}

func In(value interface{}) bson.M {
	return bson.M{
		IN: value,
	}
}

func InValues(value ...interface{}) bson.M {
	return In(value)
}

func NotIn(value interface{}) bson.M {
	return bson.M{
		NIN: value,
	}
}

func Exists() bson.M {
	return bson.M{
		EXISTS: true,
	}
}

func NotExists() bson.M {
	return bson.M{
		EXISTS: false,
	}
}

func Type(bsonType int) bson.M {
	return bson.M{
		TYPE: bsonType,
	}
}

func Mod(divisor, reminder float64) bson.M {
	return bson.M{
		MOD: []float64{divisor, reminder},
	}
}

func ElementMatch(m bson.M) bson.M {
	return bson.M{
		ELEM_MATCH: m,
	}
}
