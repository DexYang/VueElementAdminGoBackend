package util

import (
	"fmt"
	"github.com/json-iterator/go"
	"reflect"
	"strconv"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Mapping(origin interface{}, target interface{}) error {
	userJson, err := json.Marshal(origin)

	if err != nil {
		return err
	}

	err = json.Unmarshal(userJson, &target)

	if err != nil {
		return err
	}

	return nil
}


func Reflect(originPointer interface{}, targetPointer interface{}) {
	originValue := reflect.ValueOf(originPointer).Elem() // 源值

	targetType := reflect.TypeOf(targetPointer).Elem() // 目标类型
	targetValue := reflect.ValueOf(targetPointer).Elem() // 目标值域

	for i := 0; i < targetType.NumField(); i++ { // 循环遍历目标类型
		field := targetValue.Field(i) // 目标值域i

		name := targetType.Field(i).Name // 目标类型i的名字name

		value := originValue.FieldByName(name) // 根据目标类型名name获取在origin中获取源值
		if value.Kind() == reflect.Invalid { // 若源值中不存在则跳过
			continue
		}

		kind := field.Kind() // 目标值域的类型

		if kind == reflect.Ptr {
			// 获取对应字段的kind
			kind = field.Type().Elem().Kind()
		}
		switch kind {
		case reflect.Uint:
			res, _ := strconv.ParseUint(fmt.Sprint(value), 10, 64)
			targetValue.Field(i).SetUint(res)
		case reflect.Int:
			res, _ := strconv.ParseInt(fmt.Sprint(value), 10, 64)
			targetValue.Field(i).SetInt(res)
		case reflect.String:
			targetValue.Field(i).SetString(fmt.Sprint(value))
		case reflect.Struct:
		}

	}
}
