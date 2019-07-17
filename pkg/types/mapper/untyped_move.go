package mapper

import (
	"strings"

	"github.com/rancher/norman/pkg/types"
	values2 "github.com/rancher/norman/pkg/types/values"
)

type UntypedMove struct {
	From, To  string
	Separator string
}

func (m UntypedMove) FromInternal(data map[string]interface{}) {
	if v, ok := values2.RemoveValue(data, strings.Split(m.From, m.getSep())...); ok {
		values2.PutValue(data, v, strings.Split(m.To, m.getSep())...)
	}
}

func (m UntypedMove) ToInternal(data map[string]interface{}) error {
	if v, ok := values2.RemoveValue(data, strings.Split(m.To, m.getSep())...); ok {
		values2.PutValue(data, v, strings.Split(m.From, m.getSep())...)
	}

	return nil
}

func (m UntypedMove) getSep() string {
	if m.Separator == "" {
		return "/"
	}
	return m.Separator
}

func (m UntypedMove) ModifySchema(s *types.Schema, schemas *types.Schemas) error {
	return nil
}
