package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var ErrDataType = errors.New("不是期待的数据类型")

// ToString 将所有类型都输出成String的方法.
func ToString(value interface{}) string {
	res := ""
	if value == nil {
		return res
	}

	v := reflect.ValueOf(value)

	switch value.(type) {
	case float32, float64:
		res = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		return res
	case int, int8, int16, int32, int64:
		res = strconv.FormatInt(v.Int(), 10)
		return res
	case uint, uint8, uint16, uint32, uint64:
		res = strconv.FormatUint(v.Uint(), 10)
		return res
	case string:
		res = v.String()
		return res
	case []byte:
		res = string(v.Bytes())
		return res
	default:
		newValue, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		res = string(newValue)
		return res
	}
}

// ToJSON 接口类型转换为json字符串.
func ToJSON(value interface{}) (string, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// StructToMap 结构类型转Map类型.
func StructToMap(value interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%w: %T不受支持， 需要struct或者指向struct的指针", ErrDataType, value)
	}

	res := make(map[string]interface{})

	fieldNum := t.NumField()
	pattern := `^[A-Z]`
	regex := regexp.MustCompile(pattern)
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag.Get("json")
		if regex.MatchString(name) && tag != "" {
			res[tag] = v.Field(i).Interface()
		}
	}

	return res, nil
}

func ToInt(value interface{}) (int, error) {
	v, err := ToInt64(value)
	if err != nil {
		return 0, err
	}
	s := ToString(v)
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func ToInt64(value interface{}) (int64, error) {
	v := reflect.ValueOf(value)

	var res int64
	err := fmt.Errorf("%w: %T不是受支持的转换类型", ErrDataType, value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = v.Int()
		return res, nil
	case uint, uint8, uint16, uint32, uint64:
		res = int64(v.Uint())
		return res, nil
	case float32, float64:
		res = int64(v.Float())
		return res, nil
	case string:
		res, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			res = 0
		}
		return res, err
	default:
		return res, err
	}
}
