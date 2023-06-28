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

func mergeTags(tags ...*Tag) string {
	if len(tags) == 1 {
		return tags[0].String()
	}
	var str string
	for i, t := range tags {
		tag := t
		if i == 0 {
			str = strings.TrimRight(tag.String(), "`")
		} else {
			str += " " + strings.Trim(tag.String(), "`")
		}
	}
	return str + "`"
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

func buildTag(field string, name PathName, fields []*Field) *ast.BasicLit {
	for _, f := range fields {
		if field == f.Name {
			if len(f.Tags) >= 1 {
				if f.Rule.AutoGenGormTag && !containsGormTag(f.Tags) {
					f.Tags = append(f.Tags, buildEntityTag(f))
				}
			}
			if len(f.Tags) == 0 {
				switch name {
				case EntityName:
					if f.Rule.AutoGenGormTag {
						f.Tags = append(f.Tags, buildEntityTag(f))
					}
				case ModelName:
					if f.Rule.EnableValidator {
						f.Tags = append(f.Tags, buildModelTag(f))
					}
				}
			}
			var tag string = mergeTags(f.Tags...)
			return &ast.BasicLit{
				Kind:  token.STRING,
				Value: tag,
			}
		}
	}
	return nil
}

func buildModelTag(f *Field) *Tag {

	return &Tag{}
}

func containsGormTag(tags []*Tag) bool {
	for _, t := range tags {
		if t.Name == "gorm" {
			return true
		}
	}
	return false
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
	case "pq.int32Array":
		gormTag = NewTag("gorm").AddValue(Value(fmt.Sprintf("%s:%s", Type, "[]INT"))).AddValue(NotNull)
	case "pq.int64Array":
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
