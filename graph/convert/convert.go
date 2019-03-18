package convert

import (
	"relation-graph/graph/modelBase"
	"relation-graph/graph/util"
	"relation-graph/graph/modelRelation"
)

func GetDataFromModelRelations(allCreateFileLink []modelRelation.CreateFileLink, allClickFileLink []modelRelation.ClickFileLink, allCreateGroupShareLink []modelRelation.CreateGroupShareLink, allClickGroupShareLink []modelRelation.ClickGroupShareLink) ([]modelBase.User, []modelBase.Group, []modelBase.File) {
	var allUser []modelBase.User
	var allGroup []modelBase.Group
	var allFile []modelBase.File



	set := util.NewSet()
	for _, cfl := range allCreateFileLink {
		set.Add(cfl.User)
		set.Add(cfl.File)
	}
	for _, cfl := range allClickFileLink {
		set.Add(cfl.User)
		set.Add(cfl.File)
	}
	for _, cgsl := range allCreateGroupShareLink {
		set.Add(cgsl.User)
		set.Add(cgsl.Group)
	}
	for _, cgsl := range allClickGroupShareLink {
		set.Add(cgsl.User)
		set.Add(cgsl.Group)
	}
	allUser, allGroup, allFile = set.GetResult()
	return allUser, allGroup, allFile
}