package flatten

import "reflect"

func Flatten(nestedList interface{}) []interface{} {
	return flatten(nestedList, []interface{}{})
}

func flatten(nestedList interface{}, flattenList []interface{}) []interface{} {
	for _, value := range nestedList.([]interface{}) {
		if value == nil {
			continue
		}

		if reflect.TypeOf(value).String() == "int" {
			flattenList = append(flattenList, value.(int))
		} else if reflect.TypeOf(value).String() == "[]interface {}" {
			flattenList = flatten(value.(interface{}), flattenList)
		}
	}

	return flattenList
}
