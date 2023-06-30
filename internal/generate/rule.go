package generate

import "strings"

const (
	Parameter       = "parameter"
	Required        = "required"
	AutoGenGormTag  = "autogengormtag"
	EnableValidator = "enablevalidator"
	NameStyle       = "namestyle"
)

const (
	SnakeCase = "snake_cake"
	CamelCase = "camelCase"
)

// Rule is field tag. Example: `rule:"Required;AutoGenGormTag;NameStyle:snake_case;EnableValidator"`
// Except rule tag any others tag will be as fields tag. for example: you use http parameter validator
type Rule struct {
	Parameter bool
	// Required: represent this field is required parameter
	Required bool

	// AutoGenGormTag: generator will auto codeBuild gorm's tag.
	// if not specify this rule and not specify gorm's or others orm's tag the field's tag will be empty
	AutoGenGormTag bool

	// EnableValidator: enable http parameter validator, but you must specify validator tag.
	// if you have no this validator tag EnableValidator is invalid
	EnableValidator bool

	// NameStyle: you can specify the json tag naming style as snake_case or camelCase or you can directly specify json tag
	NameStyle string
}

func buildRuleFromTags(tags map[string]Value) *Rule {
	if val, ok := tags["rule"]; ok {
		return ParseRule(val)
	}
	return &Rule{NameStyle: SnakeCase}
}

// ParseRule parse rule from tag Value
func ParseRule(rules Value) *Rule {
	var rule = &Rule{NameStyle: SnakeCase}
	splitRules := strings.Split(string(rules), ";")
	if len(splitRules) > 0 {
		for _, r := range splitRules {
			if strings.Contains(r, ":") {
				split := strings.Split(r, ":")
				style := split[1]
				if strings.ToLower(split[0]) == NameStyle && (style == SnakeCase || style == CamelCase) {
					rule.NameStyle = style
				}
			}

			if strings.ToLower(r) == Parameter {
				rule.Parameter = true
			}

			if strings.ToLower(r) == Required {
				rule.Required = true
			}

			if strings.ToLower(r) == AutoGenGormTag {
				rule.AutoGenGormTag = true
			}

			if strings.ToLower(r) == EnableValidator {
				rule.EnableValidator = true
			}
		}
	}

	return rule
}
