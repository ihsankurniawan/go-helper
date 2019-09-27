package helper

import (
	"reflect"
	"strings"
)

/* 
	Ex:
	var a int
	var b []int

	a = 1
	b = []int{1, 2, 3}

	InArray(a, b) // return true, 0

	Could be used to check other type as well, such as string, uint, other basic types
*/
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(array)

			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
					index = i
					exists = true
					return
				}
			}
	}

	return
}

/* 
	Same like InArray function but not returning index
	Could be used to check other type as well, such as string, uint, other basic types
*/
func InArrayNoIndex(val interface{}, array interface{}) bool {
	exists := false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return exists
			}
		}
	}

	return exists
}

/* 
	Same like InArray function but only for string, faster than InArray
*/
func InArrayContains(val string, checkArray []string) bool {
	for _, check := range checkArray  {
		if strings.Contains(val, check) {
			return true
		}
	}
	return false
}