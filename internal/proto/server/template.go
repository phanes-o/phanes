package server

import (
	"bytes"
	"html/template"
)

var serviceTemplate = `
{{- /* delete empty line */ -}}
package service
import (
	{{- if .UseContext }}
	"context"
	{{- end }}
	{{- if .UseIO }}
	"io"
	{{- end }}
	pb "{{ .Package }}"
	{{- if .GoogleEmpty }}
	"google.golang.org/protobuf/types/known/emptypb"
	{{- end }}
)
type {{ .Service }}Service struct {}

func New{{ .Service }}Service() *{{ .Service }}Service {
	return &{{ .Service }}Service{}
}
{{- $s1 := "google.protobuf.Empty" }}
{{ range .Methods }}
{{- if eq .Type 1 }}
func (s *{{ .Service }}Service) {{ .Name }}(ctx context.Context, req {{ if eq .Request $s1 }}*emptypb.Empty
{{ else }}*pb.{{ .Request }}, resp *pb.{{ .Reply }}{{ end }}) errors {
    // todo
	return nil
}
{{- else if eq .Type 2 }}
func (s *{{ .Service }}Service) {{ .Name }}(ctx context.Context, stream pb.{{ .Service }}_{{ .Name }}Stream) errors {
	for {
		// todo
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = stream.Send(in)
		if err != nil {
			return err
		}
	}
    return nil
}
{{- else if eq .Type 3 }}
func (s *{{ .Service }}Service) {{ .Name }}(ctx context.Context, stream pb.{{ .Service }}_{{ .Name }}Stream) errors {
	for {
		// todo
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
        _ = in
		if err != nil {
			return err
		}
	}
    return stream.SendMsg(&pb.{{ .Request }}{})
}
{{- else if eq .Type 4 }}
func (s *{{ .Service }}Service) {{ .Name }}(ctx context.Context, req {{ if eq .Request $s1 }}*emptypb.Empty
{{ else }}*pb.{{ .Request }}, stream pb.{{ .Service }}_{{ .Name }}Stream{{ end }}) errors {
	for {
		err := stream.Send(&pb.{{ .Reply }}{})
		if err != nil {
			return err
		}
	}
}
{{- end }}
{{- end }}
`

type MethodType uint8

const (
	unaryType          MethodType = 1
	twoWayStreamsType  MethodType = 2 // 双向流
	requestStreamsType MethodType = 3 // 单向流，客户端流式发送
	returnsStreamsType MethodType = 4 // 单向流，服务器流式响应
)

// Service is a proto service.
type Service struct {
	Package     string
	Service     string
	Methods     []*Method
	GoogleEmpty bool

	UseIO      bool
	UseContext bool
}

// Method is a proto method.
type Method struct {
	Service string
	Name    string
	Request string
	Reply   string

	// type: unary or stream
	Type MethodType
}

func (s *Service) execute() ([]byte, error) {
	const empty = "google.protobuf.Empty"
	buf := new(bytes.Buffer)
	for _, method := range s.Methods {
		if (method.Type == unaryType && (method.Request == empty || method.Reply == empty)) ||
			(method.Type == returnsStreamsType && method.Request == empty) {
			s.GoogleEmpty = true
		}
		if method.Type == twoWayStreamsType || method.Type == requestStreamsType {
			s.UseIO = true
		}
		if method.Type == unaryType {
			s.UseContext = true
		}
	}
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
