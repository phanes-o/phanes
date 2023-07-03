package mapping

import (
	"hello/model"
	"hello/model/entity"
)

// PersonsEntityToDto entity data transfer
func PersonsEntityToDto(persons []*entity.Person) []*model.PersonInfo {
	out := make([]*model.PersonInfo, 0, len(persons))
	for _, c := range persons {
		out = append(out, PersonEntityToDto(c))
	}
	return out
}

// PersonEntityToDto entity data transfer
func PersonEntityToDto(e *entity.Person) *model.PersonInfo {
	return &model.PersonInfo{

		Id: e.Id,

		Name: e.Name,

		Arm: e.Arm,

		Age: e.Age,

		Phones: e.Phones,

		CreatedAt: e.CreatedAt.Unix(),

		UpdatedAt: e.UpdatedAt.Unix(),
	}
}
