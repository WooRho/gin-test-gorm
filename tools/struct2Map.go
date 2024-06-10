package tools

import (
	"reflect"
	"strings"
)

func Structure2ModifyMap(o interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	ForeachField(o, func(field reflect.StructField, v interface{}) {

		column := GetTagKeyField(field, "gorm", "column")
		if column == "" {
			return
		}
		r[column] = v
	})

	FilterFieldOnUpdate(r)
	return r
}

// 更新时过滤掉的字段
func FilterFieldOnUpdate(m map[string]interface{}) {
	filter := []string{
		"id",
		"create_time",
	}

	for i := range filter {
		delete(m, filter[i])
	}
}

func ForeachField(o interface{}, f func(field reflect.StructField, value interface{})) {
	if o == nil {
		return
	}

	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)

		if !IsCapitalHeader(tf.Name) {
			continue
		}

		if tf.Type.Kind() == reflect.Ptr {
			if tf.Type.Elem().Kind() == reflect.Struct {
				vok := reflect.New(tf.Type.Elem()).Interface()
				ForeachField(vok, f)
			}
		} else if tf.Type.Kind() == reflect.Struct {
			ForeachField(vf.Interface(), f)
		}

		column := GetTagKeyField(tf, "gorm", "column")
		if column == "" {
			continue
		}

		f(tf, vf.Interface())
	}
}

func GetTagKeyField(f reflect.StructField, key, field string) string {
	tag := f.Tag.Get(key)
	if tag == "-" {
		return ""
	}
	kvs := strings.Split(tag, ";")
	r := ""
	for i := range kvs {
		kv := strings.Split(kvs[i], ":")
		if len(kv) >= 1 {
			if kv[0] == field {
				if len(kv) >= 2 {
					return kv[1]
				}
			}
		}

	}
	return r
}

// 是否为大写开头
func IsCapitalHeader(s string) bool {
	if len(s) == 0 {
		return false
	}
	head := s[:1]
	t := []rune(head)
	if t[0] >= 65 && t[0] < 91 {
		return true
	} else {
		return false
	}
}
