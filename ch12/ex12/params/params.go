// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習12-12: フィールドタグを拡張し、妥当性がチェックできるようにします。

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	// ErrZip は郵便番号の形式が正しくない
	ErrZip = fmt.Errorf("Invalid ZipCode")
	// ErrMailAddress はメールアドレスの形式が正しくありません。
	ErrMailAddress = fmt.Errorf("Invalid MailAddress")
	// ErrCreditCardNumber はクレジットカード番号の形式が正しくありません。
	ErrCreditCardNumber = fmt.Errorf("Invalid CreditCard Number")

	validationFuncs = map[string]func(string) error{
		"zipcode":     vzipcode,
		"mailaddress": vmailaddress,
		"creditcard":  vcreditcard,
	}
)

// Pack は与えられた構造体からリクエストURLを作成し返却します。
func Pack(url string, ptr interface{}) (request string, err error) {
	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		nametags := strings.Split(name, ",")
		if len(nametags) > 1 {
			str, err := valueToString(v.Field(i))
			if err != nil {
				log.Printf("%s: %v", name, err)
			}

			if err = validationFuncs[nametags[1]](str); err != nil {
				return "", err
			}
		}

		fields[nametags[0]] = v.Field(i)
	}

	request = url + "?"
	for name, f := range fields {
		if f.Kind() == reflect.Slice {
			for i := 0; f.Len() > i; i++ {
				if str, err := valueToString(f.Index(i)); err != nil {
					log.Printf("%s: %v", name, err)
				} else {
					request = request + name + "=" + str + "&"
				}
			}
		} else {
			if str, err := valueToString(f); err != nil {
				log.Printf("%s: %v", name, err)
			} else {
				request = request + name + "=" + str + "&"
			}
		}
	}

	return strings.TrimRight(request, "&"), nil
}

func valueToString(v reflect.Value) (str string, err error) {
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil

	case reflect.Int:
		return fmt.Sprint(v.Int()), nil

	case reflect.Bool:
		return fmt.Sprint(v.Bool()), nil

	default:
		return "", fmt.Errorf("unsupported kind %s", v.Type())
	}
}

func vzipcode(str string) error {
	reg := "^[0-9]{7}$"
	if !regexp.MustCompile(reg).Match([]byte(str)) {
		return ErrZip
	}

	return nil
}

func vmailaddress(str string) error {
	if str == "" {
		return nil
	}
	reg := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$"
	if !regexp.MustCompile(reg).Match([]byte(str)) {
		return ErrMailAddress
	}

	return nil
}

func vcreditcard(str string) error {
	if str == "" {
		return nil
	}
	reg := "^[0-9]{16}$"
	if !regexp.MustCompile(reg).Match([]byte(str)) {
		return ErrCreditCardNumber
	}

	return nil
}

//!+Unpack

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	type FieldFunc struct {
		field          reflect.Value
		validationFunc func(string) error
	}

	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fieldfuncs := make(map[string]FieldFunc)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldtags := getFieldTags(v.Type().Field(i))
		fieldfunc := FieldFunc{}
		if len(fieldtags) > 1 {
			fieldfunc.validationFunc = validationFuncs[fieldtags[1]]
		}
		fieldfunc.field = v.Field(i)

		fieldfuncs[fieldtags[0]] = fieldfunc
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fieldfuncs[name].field
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if fieldfuncs[name].validationFunc != nil {
				if err := fieldfuncs[name].validationFunc(value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate
