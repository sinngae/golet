package gomapping

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func SliceToStruct(src []interface{}, dest interface{}, tag string) error {
	stt := reflect.TypeOf(dest).Elem()
	stv := reflect.ValueOf(dest).Elem()

	if stt.Kind() != reflect.Struct {
		return fmt.Errorf("dest[%T] should be pointer to struct", dest)
	}

	// TODO check required statement
	//if len(mapping) != stt.NumField() {
	//	return fmt.Errorf("field num dest[%T] should be equal with mapping", dest)
	//}

	for i := 0; i < stt.NumField(); i++ {
		ft := stt.Field(i)
		fv := stv.Field(i)
		tagStr := ft.Tag.Get(tag)
		ftt := ft.Type.Kind()

		if tagStr == "" {
			continue
		}

		idx := -1
		tagStrItemList := strings.Split(tagStr, ",")
		for _, item := range tagStrItemList {
			fieldKV := strings.Split(item, ":")
			if fieldKV[0] == "idx" {
				idxVal, err := strconv.Atoi(fieldKV[1])
				if err != nil {
					return fmt.Errorf("parse tag[%s] idx failed, err=%+v", tagStr, err)
				}
				if idxVal < 0 {
					return fmt.Errorf("tag[%s] idx[%d] invalid", tagStr, idx)
				}
				idx = idxVal
				break
			}
		}
		if idx < 0 {
			//FIXME
			continue
		}
		if idx > len(src) {
			return fmt.Errorf("idx[%d] is out of src", idx)
		}
		val := src[idx]

		vt := reflect.TypeOf(val).Kind()
		if vt != ftt {
			return fmt.Errorf("idx[%d] unmatch, struct[%+v] but map[%+v]", idx, ftt, vt)
		}

		//fv.Set(reflect.ValueOf(val))
		switch ftt {
		case reflect.Int:
			rVal := val.(int)
			fv.SetInt(int64(rVal))
		case reflect.Int8:
			rVal := val.(int8)
			fv.SetInt(int64(rVal))
		case reflect.Int16:
			rVal := val.(int16)
			fv.SetInt(int64(rVal))
		case reflect.Int32:
			rVal := val.(int32)
			fv.SetInt(int64(rVal))
		case reflect.Int64:
			rVal := val.(int64)
			fv.SetInt(rVal)
		case reflect.Uint:
			rVal := val.(uint)
			fv.SetUint(uint64(rVal))
		case reflect.Uint8:
			rVal := val.(uint8)
			fv.SetUint(uint64(rVal))
		case reflect.Uint16:
			rVal := val.(uint16)
			fv.SetUint(uint64(rVal))
		case reflect.Uint32:
			rVal := val.(uint32)
			fv.SetUint(uint64(rVal))
		case reflect.Uint64:
			rVal := val.(uint64)
			fv.SetUint(rVal)
		case reflect.Float32:
			rVal := val.(float32)
			fv.SetFloat(float64(rVal))
		case reflect.Float64:
			rVal := val.(float64)
			fv.SetFloat(rVal)
		case reflect.String:
			rVal := val.(string)
			fv.SetString(rVal)
		case reflect.Bool:
			rVal := val.(bool)
			fv.SetBool(rVal)
		default:
			return fmt.Errorf("unsupported field[%s] type[%s]", ft.Name, ft.Type.Name())
		}
	}
	return nil
}

func MapToStruct(src map[string]interface{}, dest interface{}, tag string) error {
	stt := reflect.TypeOf(dest).Elem()
	stv := reflect.ValueOf(dest).Elem()

	if stt.Kind() != reflect.Struct {
		return fmt.Errorf("dest[%T] should be pointer to struct", dest)
	}

	// TODO check required statement
	//if len(mapping) != stt.NumField() {
	//	return fmt.Errorf("field num dest[%T] should be equal with mapping", dest)
	//}

	for i := 0; i < stt.NumField(); i++ {
		ft := stt.Field(i)
		fv := stv.Field(i)
		tagStr := ft.Tag.Get(tag)
		ftt := ft.Type.Kind()

		fieldName := ""
		tagStrItemList := strings.Split(tagStr, ",")
		for _, item := range tagStrItemList {
			fieldKV := strings.Split(item, ":")
			if fieldKV[0] == "field" {
				fieldName = fieldKV[1]
			}
		}
		val, ok := src[fieldName]
		if !ok {
			return fmt.Errorf("field[%s] not found", fieldName)
		}
		vt := reflect.TypeOf(val).Kind()
		if vt != ftt {
			return fmt.Errorf("field[%s] unmatch, struct[%+v] but map[%+v]", fieldName, ftt, vt)
		}

		//fv.Set(reflect.ValueOf(val))
		switch ftt {
		case reflect.Int:
			rVal := val.(int)
			fv.SetInt(int64(rVal))
		case reflect.Int8:
			rVal := val.(int8)
			fv.SetInt(int64(rVal))
		case reflect.Int16:
			rVal := val.(int16)
			fv.SetInt(int64(rVal))
		case reflect.Int32:
			rVal := val.(int32)
			fv.SetInt(int64(rVal))
		case reflect.Int64:
			rVal := val.(int64)
			fv.SetInt(rVal)
		case reflect.Uint:
			rVal := val.(uint)
			fv.SetUint(uint64(rVal))
		case reflect.Uint8:
			rVal := val.(uint8)
			fv.SetUint(uint64(rVal))
		case reflect.Uint16:
			rVal := val.(uint16)
			fv.SetUint(uint64(rVal))
		case reflect.Uint32:
			rVal := val.(uint32)
			fv.SetUint(uint64(rVal))
		case reflect.Uint64:
			rVal := val.(uint64)
			fv.SetUint(rVal)
		case reflect.Float32:
			rVal := val.(float32)
			fv.SetFloat(float64(rVal))
		case reflect.Float64:
			rVal := val.(float64)
			fv.SetFloat(rVal)
		case reflect.String:
			rVal := val.(string)
			fv.SetString(rVal)
		case reflect.Bool:
			rVal := val.(bool)
			fv.SetBool(rVal)
		default:
			return fmt.Errorf("unsupported field[%s] type[%s]", ft.Name, ft.Type.Name())
		}
	}
	return nil
}
