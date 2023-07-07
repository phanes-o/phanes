package generate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"
)

type Tag struct {
	Name   TagType
	Values []Value
}

type TagType string

const (
	JsonTag     TagType = "json"
	GormTag     TagType = "gorm"
	ValidateTag TagType = "binding"
)

const (
	Type          = "type"
	PrimaryKey    = "primary_key"
	Unique        = "unique"
	UniqueIndex   = "unique_index"
	NotNull       = "not null"
	AutoIncrement = "AUTO_INCREMENT"
)

func NewTypeVarchar(size int64) Value {
	return Value(fmt.Sprintf("type:varchar(%d)", size))
}

func NewIndex(name string) Value {
	return Value(fmt.Sprintf("index:%s", name))
}
func NewSize(size int64) Value {
	return Value(fmt.Sprintf("size:%d", size))
}
func NewColumn(name string) Value {
	return Value(fmt.Sprintf("column:%s", name))
}

type Value string

func NewTag(name TagType) *Tag {
	return &Tag{
		Name: name,
	}
}

func (tag *Tag) AddValue(v Value) *Tag {
	tag.Values = append(tag.Values, v)
	return tag
}

func (tag *Tag) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("`")
	buf.WriteString(string(tag.Name))
	buf.WriteString(":")
	buf.WriteString("\"")
	for i, v := range tag.Values {
		if i > 0 {
			switch tag.Name {
			case GormTag:
				buf.WriteString(";")
			case ValidateTag:
				buf.WriteString(",")
			}
		}
		buf.WriteString(string(v))
	}
	buf.WriteString("\"")
	buf.WriteString("`")
	return buf.String()
}

func mergeTags(name PathName, tags ...*Tag) string {
	if len(tags) == 1 {
		if (name == EntityName && strings.Contains(string(tags[0].Name), string(ValidateTag))) || (name == ModelName && strings.Contains(string(tags[0].Name), string(GormTag))) {
			return ""
		}
		return tags[0].String()
	}
	var exists = make(map[string]struct{})
	var str string
	if len(tags) == 1 {
		return tags[0].String()
	}

	for _, t := range tags {
		if _, ok := exists[string(t.Name)]; ok || t.Name == "" {
			continue
		}
		if (name == EntityName && strings.Contains(string(t.Name), string(ValidateTag))) || (name == ModelName && strings.Contains(string(t.Name), string(GormTag))) {
			continue
		}

		if len(str) == 0 {
			str = strings.TrimRight(t.String(), "`")
		} else {
			str += " " + strings.Trim(t.String(), "`")
		}
		exists[string(t.Name)] = struct{}{}
	}
	if len(str) > 0 {
		return str + "`"
	}
	return ""
}

func buildTagsFromTagsText(texts map[string]Value) []*Tag {
	var tags = make([]*Tag, 0)
	for k, v := range texts {
		if k == "rule" {
			continue
		}
		tag := NewTag(TagType(k)).AddValue(v)
		tags = append(tags, tag)
	}
	return tags
}

func parseStructTags(tag *ast.BasicLit) map[string]Value {
	tags := make(map[string]Value)

	if tag != nil {
		tagValue := strings.Trim(tag.Value, "`")
		parts := strings.Split(tagValue, " ")
		for _, part := range parts {
			if !strings.Contains(part, ":") {
				continue
			}
			keyValue := strings.SplitN(part, ":", 2)
			key := strings.Trim(keyValue[0], "\"")
			value := strings.Trim(keyValue[1], "\"")
			tags[key] = Value(value)
		}
	}

	return tags
}

// @params
// name: field name
// path: generate type
func buildTag(name string, path PathName, fields []*Field) *ast.BasicLit {
	var (
		tag  string
		tags = make([]*Tag, 0, 0)
	)
	for _, f := range fields {
		if name == f.Name {
			if len(f.Tags) >= 1 {
				switch path {
				case EntityName:
					if f.Rule.AutoGenGormTag && !containsTagType(f.Tags, "gorm") {
						tags = append(f.Tags, buildGormTag(f))
					}
				case ModelName:

				}
			}
			if len(f.Tags) == 0 {
				switch path {
				case EntityName:
					if f.Rule.AutoGenGormTag {
						tags = append(f.Tags, buildGormTag(f))
					}
				case ModelName:
					if f.Rule.EnableValidator {
						tags = append(f.Tags, buildValidatorTag(f))
					}
				}
			}
		}
		if !containsTagType(tags, "json") {
			tags = append(tags, buildJsonSnakeCodeTag(name))
		}
		tag = mergeTags(path, tags...)
	}
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: tag,
	}
}

func containsTagType(tags []*Tag, s string) bool {
	for _, t := range tags {
		if string(t.Name) == s {
			return true
		}
	}
	return false
}

func buildValidatorTag(f *Field) *Tag {
	if f.Rule.EnableValidator {
		if !containsTagType(f.Tags, string(ValidateTag)) {
			tag := NewTag(ValidateTag)
			if f.Rule.Parameter && f.Rule.Required {
				if f.Name == "Id" || f.Name == "ID" {
					return &Tag{}
				}
				tag.AddValue("required")
			} else {
				return &Tag{}
			}
			return tag
		}
	}
	return &Tag{}
}

func buildJsonSnakeCodeTag(fieldName string) *Tag {
	return NewTag(JsonTag).AddValue(Value(Camel2Case(fieldName)))
}

func buildGormTag(f *Field) *Tag {
	var (
		gormTag = NewTag(GormTag).AddValue(NewColumn(f.SnakeName))
	)
	if f.Name == "Id" || f.Name == "ID" {
		gormTag.AddValue(PrimaryKey)
	} else {
		switch f.Type {
		case "string":
			gormTag.AddValue(NewTypeVarchar(255)).AddValue(NotNull)
		case "int":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "integer"))).AddValue(NotNull)
		case "int32":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "bigint"))).AddValue(NotNull)
		case "int64":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "bigint"))).AddValue(NotNull)
		case "pq.StringArray":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "varchar[]"))).AddValue(NotNull)
		case "pq.Float32Array":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "float4[]"))).AddValue(NotNull)
		case "pq.Float64Array":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "float8[]"))).AddValue(NotNull)
		case "pq.Int32Array":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "bigint[]"))).AddValue(NotNull)
		case "pq.Int64Array":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "bigint[]"))).AddValue(NotNull)
		case "time.Time":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "timestamp with time zone"))).AddValue(NotNull)
		case "*time.Time":
			gormTag.AddValue(Value(fmt.Sprintf("%s:%s", Type, "timestamp with time zone"))).AddValue(NotNull)
		default:
			fmt.Fprintf(os.Stderr, "\033[31mWRAN: Unknown field type \033[m\n")
			return &Tag{}
		}
	}

	return gormTag
}
