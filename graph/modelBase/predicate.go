package modelBase

import (
	"fmt"
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
	CreateTimeWrite
	CreateTimeRead
	ClickTimeWrite
	ClickTimeRead
	CreateGroupShareTime
	ClickGroupShareTime
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
	case CreateTimeWrite:
		return "create_time_write"
	case CreateTimeRead:
		return "create_time_read"
	case ClickTimeWrite:
		return "click_time_write"
	case ClickTimeRead:
		return "click_time_read"
	case CreateGroupShareTime:
		return "create_group_share_time"
	case ClickGroupShareTime:
		return "click_group_share_time"
	default:
		return "Unknow"
	}
}

func UserId_FileId(userid int, fileid int) string {
	return fmt.Sprintf("userid_fileid_%s_%s", userid, fileid)
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



