// Copyright © 2017 Yuki Nagahara
// フィールドタグ周辺ツール

package params

import (
	"reflect"
	"strings"
)

func getFieldTags(sf reflect.StructField) (fieldtags []string) {
	fieldInfo := sf      // a reflect.StructField
	tag := fieldInfo.Tag // a reflect.StructTag
	fieldtag := tag.Get("http")
	if fieldtag == "" {
		fieldtag = strings.ToLower(fieldInfo.Name)
	}

	return strings.Split(fieldtag, ",")
}

func nameByTag(name string, v reflect.Value) (fieldname string) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		fieldtag := tag.Get("http")
		if fieldtag == name {
			return v.Type().Field(i).Name
		}
	}

	return name
}
