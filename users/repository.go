package users

import "gorm.io/gorm"

// interface
type Repository interface {
	Save(user User) (User, error)
}

// struct
type repository struct {
	// save to db
	db *gorm.DB
}

// declation to repository
func NewRepository(db *gorm.DB) *repository {
	// & == Pointer
	return &repository{db}
}

// function
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
