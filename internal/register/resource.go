package register

import "encoding/json"

type Resource struct {
	Parent bool   `json:"parent"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Path   string `json:"path"`
}

func (p *Resource) ToJsonBytes() []byte {
	res, _ := json.Marshal(p)
	return res
}
