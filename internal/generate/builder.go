package generate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

const (
	CreateRequest = "CreateRequest"
	UpdateRequest = "UpdateRequest"
	DeleteRequest = "DeleteRequest"
	InfoRequest   = "InfoRequest"
	InfoResponse  = "InfoResponse"
	ListRequest   = "ListRequest"
	ListResponse  = "ListResponse"
)

func buildTemplateField(project string, structName StructName, n *ast.StructType) *TemplateField {
	var (
		fields = make([]*Field, 0)
	)
	for _, f := range n.Fields.List {
		tags := parseStructTags(f.Tag)
		var rule = buildRuleFromTags(tags)
		field := &Field{
			Name:      f.Names[0].Name,
			SnakeName: Camel2Case(f.Names[0].Name),
			Type:      processFieldType(f.Type),
			Rule:      rule,
			Tags:      buildTagsFromTagsText(tags),
		}
		fields = append(fields, field)
	}

	return &TemplateField{
		Fields:      fields,
		StructName:  structName,
		ProjectName: project,
	}
}

func codeBuild(genTypes []GenType, n ast.Node, tmpl *TemplateField) map[PathName]*bytes.Buffer {
	var codes = map[PathName]*bytes.Buffer{}
	for _, genType := range genTypes {
		if genType == ModelName {
			codes[ModelName] = buildModelCode(n, tmpl)
		}
		if genType == EntityName {
			codes[EntityName] = buildEntityCode(n, tmpl)
		}
	}
	return codes
}

func buildEntityCode(n ast.Node, tmpl *TemplateField) *bytes.Buffer {
	var (
		packageName = "entity"
		entity      = token.NewFileSet()
	)

	decls := []ast.Decl{
		buildImport(tmpl),
		buildEntityStruct(n, tmpl),
		buildEntityMethod(n),
	}

	file := &ast.File{
		Name:  ast.NewIdent(packageName),
		Decls: decls,
	}
	var buf = bytes.NewBuffer(nil)
	config := printer.Config{Mode: printer.UseSpaces, Tabwidth: 4}
	config.Fprint(buf, entity, file)

	return buf
}

func buildModelCode(n ast.Node, tmpl *TemplateField) *bytes.Buffer {
	var (
		model       = token.NewFileSet()
		packageName = "model"
	)
	decls := []ast.Decl{
		buildImport(tmpl),
		buildCreateRequest(n, tmpl),
		buildUpdateRequest(n, tmpl),
		buildListRequest(n, tmpl),
		buildListResponse(n, tmpl),
		buildInfoRequest(n, tmpl),
		buildInfoResponse(n, tmpl),
		buildDeleteRequest(n, tmpl),
	}

	file := &ast.File{
		Name:  ast.NewIdent(packageName),
		Decls: decls,
	}
	var buf = bytes.NewBuffer(nil)
	config := printer.Config{Mode: printer.UseSpaces, Tabwidth: 4}
	config.Fprint(buf, model, file)
	return buf
}

func buildEntityStruct(n ast.Node, tmpl *TemplateField) ast.Decl {
	var fields = make([]*ast.Field, 0)

	switch node := n.(type) {
	case *ast.TypeSpec:
		if s, ok := node.Type.(*ast.StructType); ok {
			for _, f := range s.Fields.List {
				fields = append(fields, buildField(f, EntityName, tmpl))
			}
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(string(tmpl.StructName)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}
	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildEntityMethod(n ast.Node) ast.Decl {
	switch node := n.(type) {
	case *ast.TypeSpec:
		name := node.Name
		method := &ast.FuncDecl{
			Name: ast.NewIdent("String"),
			Recv: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent(string(strings.ToLower(name.Name)[0]))},
						Type: &ast.StarExpr{
							X: &ast.Ident{
								Name: name.Name,
							},
						},
					},
				},
			},
			Type: &ast.FuncType{
				Results: &ast.FieldList{
					List: []*ast.Field{
						{
							Type: ast.NewIdent("string"),
						},
					},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", Camel2Case(name.Name)),
							},
						},
					},
				},
			},
		}
		return method
	}
	return nil
}

func buildImport(tmpl *TemplateField) ast.Decl {
	specs := make([]ast.Spec, 0, len(tmpl.Imports))

	for _, i := range tmpl.Imports {
		spec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: i,
			},
		}
		specs = append(specs, spec)
	}
	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs,
	}
}

func buildDeleteRequest(n ast.Node, tmpl *TemplateField) *ast.GenDecl {
	var field *ast.Field
	for _, f := range tmpl.Fields {
		if f.Name == "Id" || f.Name == "ID" {
			field = &ast.Field{
				Names: []*ast.Ident{{Name: f.Name}},
				Type:  ast.NewIdent(f.Type),
			}
			field.Tag = buildTag(f.Name, ModelName, tmpl.Fields)
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(fmt.Sprintf("%s%s", tmpl.StructName, DeleteRequest)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: []*ast.Field{field}},
		},
	}

	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildInfoRequest(n ast.Node, tmpl *TemplateField) *ast.GenDecl {
	var field *ast.Field
	for _, f := range tmpl.Fields {
		if f.Name == "Id" || f.Name == "ID" {
			field = &ast.Field{
				Names: []*ast.Ident{{Name: f.Name}},
				Type:  ast.NewIdent(f.Type),
			}
			field.Tag = buildTag(f.Name, ModelName, tmpl.Fields)
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(fmt.Sprintf("%s%s", tmpl.StructName, InfoRequest)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: []*ast.Field{field}},
		},
	}

	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildInfoResponse(n ast.Node, tmpl *TemplateField) *ast.GenDecl {
	var fields = make([]*ast.Field, 0)

	switch node := n.(type) {
	case *ast.TypeSpec:
		if s, ok := node.Type.(*ast.StructType); ok {
			for _, f := range s.Fields.List {
				fieldName := f.Names[0].Name
				fields = append(fields, f)
				f.Tag = buildTag(fieldName, ModelName, tmpl.Fields)
			}
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(fmt.Sprintf("%s%s", tmpl.StructName, InfoResponse)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}
	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildListRequest(n ast.Node, tmpl *TemplateField) *ast.GenDecl {
	var fields = make([]*ast.Field, 0)

	switch node := n.(type) {
	case *ast.TypeSpec:
		if s, ok := node.Type.(*ast.StructType); ok {
			for _, f := range s.Fields.List {
				rule := tmpl.getFieldRule(f.Names[0].Name)
				if rule.Parameter {
					field := buildField(f, ModelName, tmpl)
					if !rule.Required {
						field = rebuildFieldAsStar(field, tmpl)
					}
					fields = append(fields, field)
				}
			}
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(fmt.Sprintf("%s%s", tmpl.StructName, ListRequest)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}
	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildListResponse(n ast.Node, tmpl *TemplateField) *ast.GenDecl {
	var fields = make([]*ast.Field, 0)

	switch node := n.(type) {
	case *ast.TypeSpec:
		if _, ok := node.Type.(*ast.StructType); ok {
			total := &ast.Field{
				Names: []*ast.Ident{ast.NewIdent("Total")},
				Type:  ast.NewIdent("int64"),
				Tag:   buildTag("Total", ModelName, tmpl.Fields),
			}
			list := &ast.Field{
				Names: []*ast.Ident{ast.NewIdent("List")},
				Type: &ast.ArrayType{
					Elt: ast.NewIdent(node.Name.Name),
				},
				Tag: buildTag("List", ModelName, tmpl.Fields),
			}
			fields = append(fields, total, list)
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(fmt.Sprintf("%s%s", tmpl.StructName, ListResponse)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}
	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildUpdateRequest(n ast.Node, tmpl *TemplateField) *ast.GenDecl {
	var fields = make([]*ast.Field, 0)

	switch node := n.(type) {
	case *ast.TypeSpec:
		if s, ok := node.Type.(*ast.StructType); ok {
			for _, f := range s.Fields.List {
				rule := tmpl.getFieldRule(f.Names[0].Name)
				if rule.Parameter {
					field := buildField(f, ModelName, tmpl)
					if !rule.Required {
						field = rebuildFieldAsStar(field, tmpl)
					}
					fields = append(fields, field)
				}
			}
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(fmt.Sprintf("%s%s", tmpl.StructName, UpdateRequest)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}
	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildCreateRequest(n ast.Node, tmpl *TemplateField) *ast.GenDecl {
	var fields = make([]*ast.Field, 0)

	switch node := n.(type) {
	case *ast.TypeSpec:
		if s, ok := node.Type.(*ast.StructType); ok {
			for _, f := range s.Fields.List {
				rule := tmpl.getFieldRule(f.Names[0].Name)
				if rule.Parameter {
					field := buildField(f, ModelName, tmpl)
					if !rule.Required {
						field = rebuildFieldAsStar(field, tmpl)
					}
					fields = append(fields, field)
				}
			}
		}
	}

	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(fmt.Sprintf("%s%s", tmpl.StructName, CreateRequest)),
		Type: &ast.StructType{
			Fields: &ast.FieldList{List: fields},
		},
	}
	return &ast.GenDecl{
		Tok:   token.STRUCT,
		Specs: []ast.Spec{typeSpec},
	}
}

func buildField(f *ast.Field, path PathName, tmpl *TemplateField) *ast.Field {
	field := &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(f.Names[0].Name)},
		Type:  f.Type,
		Tag:   buildTag(f.Names[0].Name, path, tmpl.Fields),
	}
	return field
}

func rebuildFieldAsStar(f *ast.Field, tmpl *TemplateField) *ast.Field {
	switch fType := f.Type.(type) {
	case *ast.MapType:
		return f
	case *ast.ArrayType:
		return f
	case *ast.StarExpr:
		return f
	case *ast.SelectorExpr:
		newType := transType(fmt.Sprintf("%s.%s", fType.X, fType.Sel))
		if newType == nil {
			return f
		}
		f.Type = newType
		return f
	}
	f.Type = &ast.StarExpr{
		X: f.Type,
	}
	return f
}

func transType(tName string) ast.Expr {
	var res ast.Expr
	switch tName {
	case "pq.StringArray":
		res = &ast.ArrayType{
			Elt: ast.NewIdent("string"),
		}
	case "pq.Float32Array":
		res = &ast.ArrayType{
			Elt: ast.NewIdent("float32"),
		}
	case "pq.Float64Array":
		res = &ast.ArrayType{
			Elt: ast.NewIdent("float64"),
		}
	case "pq.int32Array":
		res = &ast.ArrayType{
			Elt: ast.NewIdent("int32"),
		}
	case "pq.int64Array":
		res = &ast.ArrayType{
			Elt: ast.NewIdent("int64"),
		}
	default:
		return nil
	}
	return res
}
