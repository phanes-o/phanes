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
	Name   string
	Values []Value
}

const (
	Column        = "column"
	Type          = "type"
	Size          = "size"
	PrimaryKey    = "primaryKey"
	Unique        = "unique"
	NotNull       = "NotNull"
	Index         = "index"
	AutoIncrement = "autoIncrement"
)

type Value string

func NewTag(name string) *Tag {
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
	buf.WriteString(tag.Name)
	buf.WriteString(":")
	buf.WriteString("\"")
	for i, v := range tag.Values {
		if i > 0 {
			buf.WriteString(";")
		}
		buf.WriteString(string(v))
	}
	buf.WriteString("\"")
	buf.WriteString("`")
	return buf.String()
}

func mergeTags(name PathName, tags ...*Tag) string {
	if len(tags) == 1 {
		if (name == EntityName && strings.Contains(tags[0].Name, "validat")) || (name == ModelName && strings.Contains(tags[0].Name, "orm")) {
			return ""
		}
		return tags[0].String()
	}
	var exists = make(map[string]struct{})
	var str string
	for i, t := range tags {
		tag := t
		if _, ok := exists[t.Name]; ok || t.Name == "" {
			continue
		}
		if (name == EntityName && strings.Contains(t.Name, "validat")) || (name == ModelName && strings.Contains(t.Name, "orm")) {
			continue
		}

		if i == 0 {
			str = strings.TrimRight(tag.String(), "`")
		} else {
			str += " " + strings.Trim(tag.String(), "`")
		}
		exists[t.Name] = struct{}{}
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
		tag := NewTag(k).AddValue(v)
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
						tags = append(f.Tags, buildEntityTag(f))
					}
				case ModelName:

				}
			}
			if len(f.Tags) == 0 {
				switch path {
				case EntityName:
					if f.Rule.AutoGenGormTag {
						tags = append(f.Tags, buildEntityTag(f))
					}
				case ModelName:
					if f.Rule.EnableValidator {
						tags = append(f.Tags, buildModelTag(f))
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
		if t.Name == s {
			return true
		}
	}
	return false
}

func buildModelTag(f *Field) *Tag {
	return buildJsonSnakeCodeTag(f.Name)
}

func buildJsonSnakeCodeTag(fieldName string) *Tag {
	return NewTag("json").AddValue(Value(Camel2Case(fieldName)))
}

func buildEntityTag(f *Field) *Tag {
	var (
		gormTag *Tag
	)
	if f.Name == "Id" || f.Name == "ID" {
		gormTag = NewTag("gorm").AddValue(PrimaryKey).AddValue(Unique).AddValue(AutoIncrement).AddValue(Value(fmt.Sprintf("%s:%s", Type, "BIGINT")))
	}

	switch f.Type {
	case "string":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "VERCHAR(255)"))).AddValue(NotNull)
	case "int":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "INT"))).AddValue(NotNull)
	case "int32":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "BIGINT"))).AddValue(NotNull)
	case "int64":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "BIGINT"))).AddValue(NotNull)
	case "pq.StringArray":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "[]VERCHAR"))).AddValue(NotNull)
	case "pq.Float32Array":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "[]FLOAT"))).AddValue(NotNull)
	case "pq.Float64Array":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "[]FLOAT"))).AddValue(NotNull)
	case "pq.Int32Array":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "[]INT"))).AddValue(NotNull)
	case "pq.Int64Array":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "[]BIGINT"))).AddValue(NotNull)
	case "time.Time":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "TIMESTEMP"))).AddValue(NotNull)
	case "*time.Time":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "TIMESTEMP"))).AddValue(NotNull)
	default:
		fmt.Fprintf(os.Stderr, "\033[31mWRAN: Unknown field type \033[m\n")
		return &Tag{}
	}
	return gormTag
}
