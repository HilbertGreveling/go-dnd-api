package policies

import "github.com/hilbertgreveling/dnd-character-api/models"

type UserPolicy interface {
	CanView(userID int, user *models.User) bool
	CanEdit(userID int, user *models.User) bool
	CanDelete(userID int, user *models.User) bool
}

type UserPolicyImpl struct{}

func NewUserPolicy() UserPolicy {
	return &UserPolicyImpl{}
}

func (p *UserPolicyImpl) CanView(userID int, user *models.User) bool {
	return true
}

func (p *UserPolicyImpl) CanEdit(userID int, user *models.User) bool {
	return user.ID == userID
}

func (p *UserPolicyImpl) CanDelete(userID int, user *models.User) bool {
	return user.ID == userID
}
