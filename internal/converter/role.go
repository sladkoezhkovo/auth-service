package converter

import (
	api "github.com/sladkoezhkovo/auth-service/api/auth"
	"github.com/sladkoezhkovo/auth-service/internal/entity"
)

func RoleFromEntityToDto(e *entity.Role) *api.Role {
	return &api.Role{
		Id:        e.Id,
		Name:      e.Name,
		Authority: e.Authority,
	}
}

func RoleFromDtoToEntity(d *api.Role) *entity.Role {
	return &entity.Role{
		Id:        d.Id,
		Name:      d.Name,
		Authority: d.Authority,
	}
}
