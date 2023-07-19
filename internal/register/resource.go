package register

import "encoding/json"

type Resource struct {
	Parent bool
	Name   string
	Type   string
	Path   string
}

func (p *Resource) ToJsonBytes() []byte {
	res, _ := json.Marshal(p)
	return res
}
