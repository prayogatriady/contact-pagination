package repository

import (
	"github.com/prayogatriady/golang-rest-pagination/internal/app/model"
	"gorm.io/gorm"
)

type ContactRepository interface {
	GetContactList(pagination *model.Pagination) (contacts []*model.Contact, totalRows int64, err error) 
	GetContact(contactId int) (contact *model.Contact, err error)
	CreateContact(contact *model.Contact) (err error)
	UpdateContact(contact *model.Contact) (err error)
	DeleteContact(contactId int) (err error)
}

type contactRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{
		DB: db,
	}
}

// Get contact list per request page alongside with limit
func (repo *contactRepository) GetContactList(pagination *model.Pagination) (contacts []*model.Contact, totalRows int64, err error) {

	err = repo.DB.Model(&contacts).Count(&totalRows).Error
	err = repo.DB.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&contacts).Error
	return
}

// Get contact by id
func (repo *contactRepository) GetContact(contactId int) (contact *model.Contact, err error) {

	err = repo.DB.Model(contact).Where("id = ?", contactId).Find(&contact).Error
	return
}

// Create an new contact
func (repo *contactRepository) CreateContact(contact *model.Contact) (err error) {

	err = repo.DB.Model(contact).Create(&contact).Error
	return
}

// Update an existing contact
func (repo *contactRepository) UpdateContact(contact *model.Contact) (err error) {

	err = repo.DB.Model(contact).Where("id = ?", contact.ID).Updates(&contact).Error
	return
}

// Delete an existing contact by updating deleted_at column
func (repo *contactRepository) DeleteContact(contactId int) (err error) {

	var contact *model.Contact
	
	err = repo.DB.Where("id = ?", contactId).Delete(&contact).Error
	return
}
