package uspsgo

import (
	"net/url"
	"reflect"
)

func toParams(data any) url.Values {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			panic("nil pointer passed to toParams")
		}
		v = v.Elem()
	}
	typeOfS := v.Type()

	params := url.Values{}
	for i := range v.NumField() {
		field := v.Field(i)
		tag := typeOfS.Field(i).Tag.Get("json")
		if tag != "" {
			str := field.String()
			if str != "" {
				params.Add(tag, str)
			}
		}
	}
	return params
}
