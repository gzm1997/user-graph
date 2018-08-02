package session

import (
	"gopkg.in/mgo.v2"
	"relation-graph/graphRelation/createTriple/conf"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
)

var DB *mgo.Session
var store *cayley.Handle


func init()  {
	dbUrl := conf.GetDbUrl()
	var err error
	DB, err = mgo.Dial(dbUrl)
	//defer DB.Close()
	if err != nil {
		panic(err)
	}
	err = graph.InitQuadStore("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	store, err = cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
}

func GetSession() *mgo.Session {
	sessionCopy := DB.Copy()
	//defer sessionCopy.Close()
	return sessionCopy
}

func GetGraph() *cayley.Handle {
	return store
}