package models

import (
	"net/url"
	"reflect"
	"strings"
)

type QueryField struct {
	Text    string
	Operand *string
}

type Query struct {
	Color         []QueryField `scryfall:"c"`
	ColorIdentity []QueryField `scryfall:"id"`
	Type          []QueryField `scryfall:"t"`
	Keyword       []QueryField `scryfall:"kw"`
}

func (q Query) Format(text *string) string {
	var sb strings.Builder
	if text != nil {
		sb.WriteString(*text)
	}

	refFields := reflect.TypeOf(q)
	refValues := reflect.ValueOf(q)
	for i := 0; i < refFields.NumField(); i += 1 {
		field := refFields.Field(i)
		values := refValues.Field(i).Interface().([]QueryField)

		for _, value := range values {
			if sb.Len() != 0 {
				sb.WriteByte('+')
			}
			sb.WriteString(field.Tag.Get("scryfall"))
			if value.Operand != nil {
				sb.WriteString(*value.Operand + value.Text)
				continue
			}
			sb.WriteString(":" + value.Text)
		}
	}

	return url.QueryEscape(sb.String())
}
