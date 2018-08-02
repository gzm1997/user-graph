package result

import (
	"relation-graph/graphRelation/createTriple/session"
	"relation-graph/graphRelation/createTriple/find"
)




func EncodeConnectorToArray(userid int) ([]int, [][numOfRelation]int) {
	result := make(map[int](*[numOfRelation]int))
	store := session.GetGraph()
	index := 0
	someOneInvitedBy, someOneInvite := find.FindInviteRelevant(store, userid)
	for k := range someOneInvitedBy {
		_, ok := result[k]; if !ok {
			result[k] = &[numOfRelation]int{}
		}
		//fmt.Println(someOneInvitedBy[k])
		result[k][index] = someOneInvitedBy[k]
	}
	index++
	for k := range someOneInvite {
		_, ok := result[k]; if !ok {
			result[k] = &[numOfRelation]int{}
		}
		result[k][index] = someOneInvite[k]
	}
	index++
	createShareFileLinkTo, clickShareFileLinkFrom := find.FindShareRelevant(store, userid)
	for k := range createShareFileLinkTo {
		_, ok := result[k]; if !ok {
			result[k] = &[numOfRelation]int{}
		}
		result[k][index] = createShareFileLinkTo[k]
	}
	index++
	for k := range clickShareFileLinkFrom {
		_, ok := result[k]; if !ok {
			result[k] = &[numOfRelation]int{}
		}
		result[k][index] = clickShareFileLinkFrom[k]
	}
	index++
	groupmemberRelvent := find.FindGroupMemberRelvent(store, userid)
	for k := range groupmemberRelvent {
		_, ok := result[k]; if !ok {
			result[k] = &[numOfRelation]int{}
		}
		result[k][index] = groupmemberRelvent[k]
	}
	index++
	a, b, c, d := find.FindShareSameFileRelevant(store, userid)
	for g := range a {
		group := a[g]
		for k := range group {
			_, ok := result[k]; if !ok {
				result[k] = &[numOfRelation]int{}
			}
			result[k][index] = group[k]
		}
	}
	index++
	for g := range b {
		group := b[g]
		for k := range group {
			_, ok := result[k]; if !ok {
				result[k] = &[numOfRelation]int{}
			}
			result[k][index] = group[k]
		}
	}
	index++
	for g := range c {
		group := c[g]
		for k := range group {
			_, ok := result[k]; if !ok {
				result[k] = &[numOfRelation]int{}
			}
			result[k][index] = group[k]
		}
	}
	index++
	for g := range d {
		group := d[g]
		for k := range group {
			_, ok := result[k]; if !ok {
				result[k] = &[numOfRelation]int{}
			}
			result[k][index] = group[k]
		}
	}
	index++
	r := find.FindGroupMemberBecauseOfShareFileRelevant(store, userid)
	for g := range r {
		group := r[g]
		for k := range group {
			_, ok := result[k]; if !ok {
				result[k] = &[numOfRelation]int{}
			}
			result[k][index] = group[k]
		}
	}
	var userIds []int
	var returnResult [][numOfRelation]int
	for user := range result {
		userIds = append(userIds, user)
		returnResult = append(returnResult, *result[user])
	}
	return userIds, returnResult
}

func Calculate(arr [][numOfRelation]int) []float64 {
	result := []float64{}
	for _, row := range arr {
		sum := 0.0
		for i := 0; i < numOfRelation; i++ {
			sum += float64(row[i]) * Weight[i]
		}
		result = append(result, sum)
	}
	return result
}