package types

import "gopkg.in/mgo.v2/bson"

// FixerIn ...
/* Example JSON:
{
	"base":"EUR",
    "date":"2017-10-24",
    "rates":{
        "AUD":1.5117,
        "BGN":1.9558,
		.....
		"ZAR":16.14
    }
}
*/
type FixerIn struct {
	ID        bson.ObjectId      `json:"id" bson:"_id,omitempty"`
	Base      string             `bson:"base"`
	Timestamp string             `bson:"timestamp,omitempty"`
	Date      string             `bson:"date"`
	Rates     map[string]float64 `bson:"rates"`
}
