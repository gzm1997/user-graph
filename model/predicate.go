package model


type Predicate int

const (
	ShareFolder Predicate = iota // value --> 0
	WriteNum                 // value --> 1
	LastTime                 // value --> 2
	Name                     // value --> 3
	Share
	Create
	HasMember
	Permission
	GroupId
	LinkCreateTime
	LinkType
	Link
	SaveTime
)


func (this Predicate) String() string {
	switch this {
	case ShareFolder:
		return "ShareFolder"
	case WriteNum:
		return "WriteNum"
	case LastTime:
		return "LastTime"
	case Name:
		return "Name"
	case Share:
		return "Share"
	case Create:
		return "Create"
	case HasMember:
		return "HasMember"
	case Permission:
		return "Permission"
	case GroupId:
		return "GroupId"
	case LinkCreateTime:
		return "LinkCreateTime"
	case LinkType:
		return "LinkType"
	case SaveTime:
		return "SaveTime"
	default:
		return "Unknow"
	}
}