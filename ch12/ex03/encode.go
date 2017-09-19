// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習12-3: float, complex, bool, interface の実装をします。
// 練習12-6: ゼロ値の場合はエンコーディングしないように修正します。
// 練習12-13: フィールドタグを処理するように修正します。

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {

			// 練習12-6
			// ゼロ値かどうか比較、ゼロ値でなければエンコーディングしない
			if reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {
				continue
			}

			// 練習12-13: フィールドタグを見て、そのタグ名でエンコードします。
			name := getFieldName(v.Type().Field(i))
			if name == "-" {
				// 値を完全無視する
				continue
			}

			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Bool:
		fmt.Fprintf(buf, "%v", v.Bool())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%s", strconv.FormatFloat(v.Float(), 'f', 1000, 64))

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%s, %s)", strconv.FormatFloat(real(v.Complex()), 'f', 1000, 64), strconv.FormatFloat(imag(v.Complex()), 'f', 1000, 64))

	case reflect.Interface:
		fmt.Fprintf(buf, "#I(\"%s\" ", reflect.ValueOf(v.Interface()).Type())

		encode(buf, reflect.ValueOf(v.Interface()))
		fmt.Fprintf(buf, ")")
	default: // chan, func
		// chan はシステムコールに関わるため、シリアライズできない。
		// func は関数であり、シリアライズできない。
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
