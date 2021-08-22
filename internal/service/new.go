package service

import (
	"github.com/go-juno/juno/internal/model"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NewService interface {
	GetList(pageIndex int, pageSize int) (newList []*model.New, total int64, err error)
	GetAll() (newList []*model.New, err error)
	GetDetail(id uint) (new *model.New, err error)
	Create(new *model.New) (err error)
	Update(new *model.New) (err error)
	Delete(id uint) (err error)
}

type newService struct {
	db *gorm.DB
}

func (s *newService) GetList(pageIndex int, pageSize int) (newList []*model.New, total int64, err error) {
	query := s.db.Model(&model.New{})
	err = query.Count(&total).
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Order("created_at desc,id desc").
		Find(&newList).Error

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *newService) GetAll() (newList []*model.New, err error) {
	err = s.db.Model(&model.New{}).
		Order("created_at desc,id desc").
		Find(&newList).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *newService) GetDetail(id uint) (new *model.New, err error) {
	var ma model.New
	err = s.db.Model(&model.New{}).Where("id = ?", id).Find(&ma).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ma.Id != 0 {
		new = &ma
	}

	return
}

func (s *newService) Create(new *model.New) (err error) {
	err = s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(new).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *newService) Update(new *model.New) (err error) {

	err = s.db.Model(&model.New{}).Where("id = ?", new.Id).Updates(new).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *newService) Delete(id uint) (err error) {
	err = s.db.Where("id =?", id).Delete(&model.New{}).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func NewNewService(db *gorm.DB) NewService {
	return &newService{db: db}
}
