package register

var tmpl = `
package assistant

import (
	"context"
	"time"

	"auth/config"
	"auth/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const RegisterKeyPrefix = "/phanes/register_resource"

func init() {
	Register(Resource)
}

type resource struct{}

var Resource = &resource{}

func (r *resource) Init() func() {
	var (
		err       error
		cli       *clientv3.Client
		endpoints = make([]string, 0)
	)

	if config.EtcdAddr == "" {
		endpoints = []string{"localhost:2379"}
	} else {
		endpoints = append(endpoints, config.EtcdAddr)
	}

	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second * 3,
	})
	utils.Throw(err)
	utils.Throw(r.register(cli))

	return func() {}
}

func (r *resource) register(cli *clientv3.Client) error {
	var (
		err error
	)
	
	{{ range $v := .Fields }}
		{{.Name}} := {{ .Value | keepBytesType }}

		if len({{.Name}}) > 0 {
			if _, err = cli.KV.Put(context.Background(), "{{.Key}}", string({{.Name}})); err != nil {
				return err
			}
		}
	{{ end }}
	return nil
}
`
