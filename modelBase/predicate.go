package modelBase

import (
	"fmt"
	"relation-graph/graphRelation/createTriple/modelRelation"
)

type Predicate int

const (
	GroupId Predicate = iota // value --> 0
	LinkCreateTime
	LinkType
	Link
	SaveTime
	UserName
	GroupName
	FileName
	CreateGroupLink
	ClickGroupLink
	CreateTime
	ClickTime
)


func (this Predicate) String() string {
	switch this {
	case GroupId:
		return "GroupId"
	case LinkCreateTime:
		return "LinkCreateTime"
	case LinkType:
		return "LinkType"
	case SaveTime:
		return "SaveTime"
	case UserName:
		return "username"
	case GroupName:
		return "groupname"
	case FileName:
		return "filename"
	case CreateGroupLink:
		return "create_group_link"
	case ClickGroupLink:
		return "click_group_link"
	case CreateTime:
		return "create_time"
	case ClickTime:
		return "click_time"
	default:
		return "Unknow"
	}
}

func UserId_FileId_Permission(userid int, fileid int, permission modelRelation.FileLinkPermission) string {
	return fmt.Sprintf("userid_fileid_%s_%s_%s", userid, fileid, permission.String())
}

func UserId_GroupId(userid int, groupid int) string {
	return fmt.Sprintf("userid_groupid_%s_%s", userid, groupid)
}



type FileLinkPredicate int

const (
	CreateWriteFileLink FileLinkPredicate = iota
	CreateReadFileLink
	ClickWriteFileLink
	ClickReadFileLink
)

func (this FileLinkPredicate) String() string {
	switch this {
	case CreateWriteFileLink:
		return "create_write_file_link"
	case CreateReadFileLink:
		return "create_read_file_link"
	case ClickWriteFileLink:
		return "click_write_file_link"
	case ClickReadFileLink:
		return "click_read_file_link"
	default:
		return "Unknow"
	}
}



