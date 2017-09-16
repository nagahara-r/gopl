// Copyright © 2017 Yuki Nagahara
// 練習12-13: フィールドタグを見て、そのタグ名でエンコード、デコードします。

package sexpr

import "reflect"

func getFieldName(sf reflect.StructField) (fieldtag string) {
	fieldInfo := sf      // a reflect.StructField
	tag := fieldInfo.Tag // a reflect.StructTag
	fieldtag = tag.Get("sexpr")
	if fieldtag == "" {
		fieldtag = fieldInfo.Name
	}

	return
}

func nameByTag(name string, v reflect.Value) (fieldname string) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		fieldtag := tag.Get("sexpr")
		if fieldtag == name {
			return v.Type().Field(i).Name
		}
	}

	return name
}
