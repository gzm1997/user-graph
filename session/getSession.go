package session

import (
	"gopkg.in/mgo.v2"
	"relation-graph/graphRelation/createTriple/conf"
)

var DB *mgo.Session


func init()  {
	dbUrl := conf.GetDbUrl()
	var err error
	DB, err = mgo.Dial(dbUrl)
	//defer DB.Close()
	if err != nil {
		panic(err)
	}
}

func GetSession() *mgo.Session {
	sessionCopy := DB.Copy()
	//defer sessionCopy.Close()
	return sessionCopy
}