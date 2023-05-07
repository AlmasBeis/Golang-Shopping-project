package repository

import (
	model "final-project/pkg/models"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(category model.Category) error
	GetAll() ([]model.Category, error)
	GetById(id int) (model.Category, error)
	Update(category model.Category, id int) error
	Delete(id int) error
}

type CommentRepo interface {
	Create(comment model.Comment) error
	GetAll() ([]model.Comment, error)
	GetById(id int) (model.Comment, error)
	GetByItem(itemId int) ([]model.Comment, error)
	GetByUser(userId int) ([]model.Comment, error)
	Update(comment model.Comment, id int) error
	Delete(id int) error
}

type PurchaseRepo interface {
	Create(purchase model.Purchase) error
	GetAll() ([]model.Purchase, error)
	GetById(id int) (model.Purchase, error)
	GetByItem(itemId int) ([]model.Purchase, error)
	GetByUser(userId int) ([]model.Purchase, error)
	Update(purchase model.Purchase, id int) error
	Delete(id int) error
}

type ItemRepo interface {
	Create(item model.Item) error
	GetAll() ([]model.Item, error)
	GetById(id int) (model.Item, error)
	SortByRating(sort string) ([]model.Item, error)
	SortByPrice(sort string) ([]model.Item, error)
	Update(item model.Item, id int) error
	SetRating(rating float32, id int) error
	Delete(id int) error
}

type UserRepo interface {
	GetBalance(id int) (float32, error)
	Withdraw(id int, price float32) error
	Deposit(id int, dep float32) error
}

type AuthorizationRepo interface {
	Register(user model.User) error
	Login(username, password string) (model.User, error)
}

type Repository struct {
	CategoryRepo
	CommentRepo
	PurchaseRepo
	ItemRepo
	UserRepo
	AuthorizationRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{NewCategoryCrud(db),
		NewCommentCrud(db),
		NewPurchaseCrud(db),
		NewItemCrud(db),
		NewUserCrud(db),
		NewAuthorizationCrud(db),
	}
}
