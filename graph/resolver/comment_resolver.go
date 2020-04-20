package resolver

import (
	"context"
	gqlmodel "github.com/cassini-Inner/inner-src-mgmt-go/graph/model"
)

func (r *commentResolver) CreatedBy(ctx context.Context, obj *gqlmodel.Comment) (*gqlmodel.User, error) {
	user, err := r.UsersRepo.GetById(obj.CreatedBy)
	if err != nil {
		return nil, err
	}
	var gqlUser gqlmodel.User
	gqlUser.MapDbToGql(*user)
	return &gqlUser, nil
}

