package find

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"reflect"
	"relation-graph/graphRelation/createTriple/model"
)

//根据id查找names
func FindNameById(store *cayley.Handle, id int) (string, error) {
	p := cayley.StartPath(store, quad.Int(id)).Out(quad.String(model.Name.String()))
	var name string
	err := p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value)
		//fmt.Println("nav", nativeValue)
		if reflect.TypeOf(nativeValue).Kind() == reflect.String {
			name = nativeValue.(string)
			//fmt.Println("this is name", name)
		}
	})
	//fmt.Println("return name", name)
	return name, err
}