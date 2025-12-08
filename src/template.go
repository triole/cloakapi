package main

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

type tVarMap map[string]string

func (kc tKC) execTemplate(tplStr string) {
	tmpl := template.Must(
		template.New("new.tmpl").Parse(tplStr),
	)

	var varMap tVarMap
	for _, user := range kc.API.Users {
		varMap = kc.convertUserToMap(user)
		remID, remIDP := kc.getFedID(*user.Username)
		varMap["remote_id"] = remID
		varMap["remote_idp"] = remIDP

		buf := &bytes.Buffer{}
		err := tmpl.Execute(buf, varMap)
		if err != nil {
			panic(err)
		}
		s := buf.String()
		s = strings.ReplaceAll(s, "\n", "")
		s = strings.TrimSpace(s)
		if s != "" {
			fmt.Printf("%s\n", buf.String())
		}
	}
}

func (kc tKC) convertUserToMap(user *gocloak.User) map[string]string {
	return convertStructToMap(user)
}

func ConvertFederatedIdentityToMap(fedID *gocloak.FederatedIdentityRepresentation) map[string]string {
	return convertStructToMap(fedID)
}

func convertStructToMap(s any) (result tVarMap) {
	result = make(tVarMap)

	if s == nil {
		return result
	}

	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
		t = v.Type()
	} else if v.Kind() == reflect.Ptr && v.IsNil() {
		return result
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Name

		var value string
		if field.IsValid() {
			if field.Kind() == reflect.Ptr && !field.IsNil() {
				switch field.Elem().Kind() {
				case reflect.String:
					value = field.Elem().String()
				case reflect.Bool:
					value = fmt.Sprintf("%t", field.Elem().Bool())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					value = fmt.Sprintf("%d", field.Elem().Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					value = fmt.Sprintf("%d", field.Elem().Uint())
				case reflect.Float32, reflect.Float64:
					value = fmt.Sprintf("%f", field.Elem().Float())
				default:
					value = fmt.Sprintf("%v", field.Elem().Interface())
				}
			} else if field.Kind() == reflect.String {
				value = field.String()
			} else {
				switch field.Kind() {
				case reflect.Bool:
					value = fmt.Sprintf("%t", field.Bool())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					value = fmt.Sprintf("%d", field.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					value = fmt.Sprintf("%d", field.Uint())
				case reflect.Float32, reflect.Float64:
					value = fmt.Sprintf("%f", field.Float())
				default:
					value = fmt.Sprintf("%v", field.Interface())
				}
			}
		}
		if value == "\u003cnil\u003e" {
			value = ""
		}
		result[strings.ToLower(fieldName)] = value
	}
	return result
}
