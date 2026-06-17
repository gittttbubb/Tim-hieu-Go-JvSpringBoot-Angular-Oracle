package authz

import "go-rbac-system/internal/constants"

type dataScopeResolver struct{}

type DataScopeResolver interface {
	IsAllowed(
		scope string,
		isOwner bool,
		isSameTeam bool,
	) bool
}

func NewDataScopeResolver() DataScopeResolver {
	return &dataScopeResolver{}
}

func (r *dataScopeResolver) IsAllowed(scope string, isOwner bool, isSameTeam bool,) bool {
	switch scope {
		case constants.DataScopeAll:
			return true
		case constants.DataScopeTeam:
			return isSameTeam
		case constants.DataScopeOwn:
			return isOwner
		default:
			return false
	}
}