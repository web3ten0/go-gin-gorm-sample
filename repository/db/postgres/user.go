package postgres

import (
	"web3ten0/go-gin-gorm-sample/domain"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepo) GetAllByName(name string) ([]*domain.User, error) {
	var users []*domain.User
	result := r.db.Where("name LIKE ?", "%"+name+"%").Find(&users)
	return users, result.Error
}
