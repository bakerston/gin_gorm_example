package service

import (
	"Project01/main/dao"
	"Project01/main/models"
)

type UserService interface {
	Insert(user models.User) error
	Update(user models.User) error
	Delete(user models.User) error
	FindAll() []models.User
}

type userService struct {
	userDao dao.UserDao
}

func (service *userService) Insert(user models.User) error {
	service.userDao.Insert(user)
	return nil
}

func (service *userService) Update(user models.User) error {
	service.userDao.Update(user)
	return nil
}

func (service *userService) Delete(user models.User) error {
	service.userDao.Delete(user)
	return nil
}

func (service *userService) FindAll() []models.User {
	return service.userDao.FindAll()
}

func New(dao dao.UserDao) UserService {
	return &userService{
		userDao: dao,
	}
}
