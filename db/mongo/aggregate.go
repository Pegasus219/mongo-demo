package mongo

import "gopkg.in/mgo.v2/bson"

const (
	MATCH = "$match"
	GROUP = "$group"
	SORT  = "$sort"

	SUM = "$sum"
	AVG = "$avg"
	MIN = "$min"
	MAX = "$max"
)

type (
	MongoAggs struct {
		match bson.M
		group bson.M
		sort  bson.M
	}

	DistinctValue struct {
		Id interface{} `bson:"_id"`
	}

	GroupAggValue struct {
		Id    interface{} `bson:"_id"`
		Value float64     `bson:"value"`
	}
)

func NewMongoAggs() *MongoAggs {
	return new(MongoAggs)
}

func groupBy(fields ...string) interface{} {
	if len(fields) > 1 {
		ret := bson.M{}
		for _, v := range fields {
			ret[v] = "$" + v
		}
		return ret

	} else if len(fields) == 1 {
		return "$" + fields[0]

	} else {
		return nil
	}
}

func (m *MongoAggs) Match(query interface{}) *MongoAggs {
	m.match = bson.M{
		MATCH: query,
	}
	return m
}

func (m *MongoAggs) Sort(key string, sort int8) *MongoAggs {
	m.sort = bson.M{
		SORT: bson.M{
			key: sort,
		},
	}
	return m
}

func (m *MongoAggs) Distinct(fields ...string) *MongoAggs {
	m.group = bson.M{
		GROUP: bson.M{
			"_id": groupBy(fields...),
		},
	}
	return m
}

func (m *MongoAggs) CountBy(fields ...string) *MongoAggs {
	m.group = bson.M{
		GROUP: bson.M{
			"_id": groupBy(fields...),
			"value": bson.M{
				SUM: 1,
			},
		},
	}
	return m
}

func (m *MongoAggs) Sum(dataField string, groupFields ...string) *MongoAggs {
	m.group = bson.M{
		GROUP: bson.M{
			"_id": groupBy(groupFields...),
			"value": bson.M{
				SUM: "$" + dataField,
			},
		},
	}
	return m
}

func (m *MongoAggs) Avg(dataField string, groupFields ...string) *MongoAggs {
	m.group = bson.M{
		GROUP: bson.M{
			"_id": groupBy(groupFields...),
			"value": bson.M{
				AVG: "$" + dataField,
			},
		},
	}
	return m
}

func (m *MongoAggs) Min(dataField string, groupFields ...string) *MongoAggs {
	m.group = bson.M{
		GROUP: bson.M{
			"_id": groupBy(groupFields...),
			"value": bson.M{
				MIN: "$" + dataField,
			},
		},
	}
	return m
}

func (m *MongoAggs) Max(dataField string, groupFields ...string) *MongoAggs {
	m.group = bson.M{
		GROUP: bson.M{
			"_id": groupBy(groupFields...),
			"value": bson.M{
				MAX: "$" + dataField,
			},
		},
	}
	return m
}

func (m *MongoAggs) build() []bson.M {
	ret := []bson.M{}
	if m.match != nil {
		ret = append(ret, m.match)
	}

	ret = append(ret, m.group)

	if m.sort != nil {
		ret = append(ret, m.sort)
	}
	return ret
}
