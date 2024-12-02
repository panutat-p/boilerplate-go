package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func PrintJSON(o any) {
	kind := reflect.TypeOf(o).Kind()
	switch kind {
	case reflect.String:
		s := reflect.ValueOf(o).String()
		var m = make(map[string]any)
		err := json.Unmarshal([]byte(s), &m)
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	case reflect.Slice:
		if reflect.TypeOf(o).Elem().Kind() == reflect.Uint8 {
			b := reflect.ValueOf(o).Bytes()
			var m = make(map[string]any)
			err := json.Unmarshal(b, &m)
			if err != nil {
				panic(err)
			}
			b, err = json.MarshalIndent(m, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(b))
		} else {
			b, err := json.MarshalIndent(o, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(b))
		}
	default:
		b, err := json.MarshalIndent(o, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
}
