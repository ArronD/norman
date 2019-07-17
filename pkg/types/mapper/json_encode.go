package mapper

import (
	"encoding/json"
	"strings"

	"github.com/rancher/norman/pkg/types"
	convert2 "github.com/rancher/norman/pkg/types/convert"
	values2 "github.com/rancher/norman/pkg/types/values"
)

type JSONEncode struct {
	Field            string
	IgnoreDefinition bool
	Separator        string
}

func (m JSONEncode) FromInternal(data map[string]interface{}) {
	if v, ok := values2.RemoveValue(data, strings.Split(m.Field, m.getSep())...); ok {
		obj := map[string]interface{}{}
		if err := json.Unmarshal([]byte(convert2.ToString(v)), &obj); err == nil {
			values2.PutValue(data, obj, strings.Split(m.Field, m.getSep())...)
		} else {
			log.Errorf("Failed to unmarshal json field: %v", err)
		}
	}
}

func (m JSONEncode) ToInternal(data map[string]interface{}) error {
	if v, ok := values2.RemoveValue(data, strings.Split(m.Field, m.getSep())...); ok && v != nil {
		if bytes, err := json.Marshal(v); err == nil {
			values2.PutValue(data, string(bytes), strings.Split(m.Field, m.getSep())...)
		}
	}
	return nil
}

func (m JSONEncode) getSep() string {
	if m.Separator == "" {
		return "/"
	}
	return m.Separator
}

func (m JSONEncode) ModifySchema(s *types.Schema, schemas *types.Schemas) error {
	if m.IgnoreDefinition {
		return nil
	}

	return ValidateField(m.Field, s)
}
