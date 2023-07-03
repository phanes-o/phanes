package mapping

import (
	"hello/model"
	"hello/model/entity"
)

// ManagersEntityToDto entity data transfer
func ManagersEntityToDto(managers []*entity.Manager) []*model.ManagerInfo {
	out := make([]*model.ManagerInfo, 0, len(managers))
	for _, c := range managers {
		out = append(out, ManagerEntityToDto(c))
	}
	return out
}

// ManagerEntityToDto entity data transfer
func ManagerEntityToDto(e *entity.Manager) *model.ManagerInfo {
	return &model.ManagerInfo{

		Id: e.Id,

		Name: e.Name,

		Arm: e.Arm,

		Age: e.Age,

		Phones: e.Phones,

		CreatedAt: e.CreatedAt.Unix(),

		UpdatedAt: e.UpdatedAt.Unix(),
	}
}
