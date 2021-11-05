package converting

import "fmt"

/**
Iatosa converts interface array to string array
*/
func Iatosa(arr []interface{}) []string {
	s := make([]string, len(arr))
	for i, v := range arr {
		s[i] = fmt.Sprint(v)
	}
	return s
}

/**
Itosa converts interface to string array
*/
func Itosa(arr interface{}) []string {
	return Iatosa(arr.([]interface{}))
}
