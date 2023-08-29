package repository

import (
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"gorm.io/gorm"
)

type ContactRepository interface {
	Paginate(offset int, limit int, sort string) ([]*model.Contact, error)
	Count(value interface{}) (int64, error)
}

type contactRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{
		DB: db,
	}
}

func (repo *contactRepository) Paginate(offset int, limit int, sort string) ([]*model.Contact, error) {
	var contacts []*model.Contact

	if err := repo.DB.Offset(offset).Limit(limit).Order(sort).Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func (repo *contactRepository) Count(value interface{}) (int64, error) {
	var totalRows int64
	if err := repo.DB.Model(value).Count(&totalRows).Error; err != nil {
		return 0, err
	}

	return totalRows, nil
}
