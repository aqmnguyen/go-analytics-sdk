package utils

import "reflect"

func ValidateRequired(event interface{}) []string {
	var missing []string
	v := reflect.ValueOf(event)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("required")
		jsonTag := t.Field(i).Tag.Get("json")

		if tag == "true" && field.IsZero() {
			// Use JSON tag instead of Go field name
			if jsonTag != "" {
				missing = append(missing, jsonTag)
			} else {
				missing = append(missing, t.Field(i).Name)
			}
		}
	}

	return missing
}

func GetStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
