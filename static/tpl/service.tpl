package service

import (
	"juno/internal/model"
	
	"golang.org/x/xerrors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GreetingService interface {
	GetList(pageIndex int, pageSize int) (greetingList []*model.Greeting, total int64, err error)
	GetAll() (greetingList []*model.Greeting, err error)
	GetDetail(id uint) (greeting *model.Greeting, err error)
	Create(greeting *model.Greeting) (err error)
	Update(greeting *model.Greeting) (err error)
	Delete(id uint) (err error)
}

type greetingService struct {
	db *gorm.DB
}

func (s *greetingService) GetList(pageIndex int, pageSize int) (greetingList []*model.Greeting, total int64, err error) {
	query := s.db.Model(&model.Greeting{})
	err = query.Count(&total).
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Order("created_at desc,id desc").
		Find(&greetingList).Error

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *greetingService) GetAll() (greetingList []*model.Greeting, err error) {
	err = s.db.Model(&model.Greeting{}).
		Order("created_at desc,id desc").
		Find(&greetingList).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *greetingService) GetDetail(id uint) (greeting *model.Greeting, err error) {
	var ma model.Greeting
	err = s.db.Model(&model.Greeting{}).Where("id = ?", id).Find(&ma).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ma.Id != 0 {
		greeting = &ma
	}

	return
}

func (s *greetingService) Create(greeting *model.Greeting) (err error) {
	err = s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(greeting).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *greetingService) Update(greeting *model.Greeting) (err error) {

	err = s.db.Model(&model.Greeting{}).Where("id = ?", greeting.Id).Updates(greeting).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *greetingService) Delete(id uint) (err error) {
	err = s.db.Where("id =?", id).Delete(&model.Greeting{}).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func NewGreetingService(db *gorm.DB) GreetingService {
	return &greetingService{db: db}
}
