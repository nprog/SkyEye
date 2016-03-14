package mongo

import (
	"errors"
	"github.com/nprog/IntelligentEngine/glog"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

type Event struct {
	Content string `bson:"content"`
	Time    int64  `bson:"time"`
	Type    string `bson:"type"` //string, rpc, http, exec
}

type Memory struct {
	Key   string  `bson:"key"`
	Value []Event `bson:"value"`
	Lock  bool    `bson:lock`
}

// func (m *Mongo) MemorySave(Key string, events []Event) error {

// }

// func (m *Mongo) MemoryUpdate(Key string, oldEvent Event, newEvent Event) error {

// }

// func (m *Mongo) MemoryDelete(Key string) error {

// }

func (m *Mongo) MemorySelectOneEventFromKeyRandom(Key string) (Event, error) {
	var (
		err    error
		event  Event
		events []Event
	)

	events, err = m.MemorySelectEventsFromKey(Key)
	if err != nil {
		glog.Error(err.Error())
		return event, m.Error(err)
	}

	if len(events) > 0 {
		return events[rand.Intn(len(events))], nil
	} else {
		return event, errors.New("not found")
	}
}

func (m *Mongo) MemorySelectEventsFromKey(Key string) ([]Event, error) {
	var (
		err    error
		result *Memory
	)

	op := m.Session.DB(m.Opts.DbName).C(MEMORY_COLLECTION_NAME)
	err = op.Find(bson.M{"key": Key}).One(&result)
	if err != nil {
		glog.Error(err.Error())
		return nil, m.Error(err)
	}

	return result.Value, nil
}
