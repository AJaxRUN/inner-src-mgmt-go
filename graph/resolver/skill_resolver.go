package resolver

import (
	"context"
	"github.com/cassini-Inner/inner-src-mgmt-go/graph/model"
)

func (r *skillResolver) CreatedBy(ctx context.Context, obj *model.Skill) (*model.User, error) {
	return r.usersRepo.GetById(obj.CreatedBy)
}
