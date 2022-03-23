package service

import (
	"context"
	"{mod}/internal/model"

	"golang.org/x/xerrors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type {class}Service interface {
	Get{class}List(ctx context.Context, pageIndex int, pageSize int) ({camel}List []*model.{class}, total int64, err error)
	Get{class}All(ctx context.Context) ({camel}List []*model.{class}, err error)
	Get{class}Detail(ctx context.Context, id uint) ({camel} *model.{class}, err error)
	Create{class}(ctx context.Context, {camel} *model.{class}) (err error)
	Update{class}(ctx context.Context, {camel} *model.{class}) (err error)
	Delete{class}(ctx context.Context, id uint) (err error)
}

type {camel}Service struct {
	db *gorm.DB
}

func (s *{camel}Service) Get{class}List(ctx context.Context, pageIndex int, pageSize int) ({camel}List []*model.{class}, total int64, err error) {
	query := s.db.Model(&model.{class}{})
	err = query.Count(&total).
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Order("created_at desc,id desc").
		Find(&{camel}List).Error

	if err != nil {
		err = xerrors.Errorf("%%w", err)
		return
	}
	return
}

func (s *{camel}Service) Get{class}All(ctx context.Context) ({camel}List []*model.{class}, err error) {
	err = s.db.Model(&model.{class}{}).
		Order("created_at desc,id desc").
		Find(&{camel}List).Error
	if err != nil {
		err = xerrors.Errorf("%%w", err)
		return
	}
	return
}

func (s *{camel}Service) Get{class}Detail(ctx context.Context, id uint) ({camel} *model.{class}, err error) {
	err = s.db.Model(&model.{class}{}).Where("id = ?", id).Find(&{camel}).Error
	if err != nil {
		err = xerrors.Errorf("%%w", err)
		return
	}
	return
}

func (s *{camel}Service) Create{class}(ctx context.Context, {camel} *model.{class}) (err error) {
	err = s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create({camel}).Error
	if err != nil {
		err = xerrors.Errorf("%%w", err)
		return
	}
	return
}

func (s *{camel}Service) Update{class}(ctx context.Context, {camel} *model.{class}) (err error) {

	err = s.db.Model(&model.{class}{}).Where("id = ?", {camel}.Id).Updates({camel}).Error
	if err != nil {
		err = xerrors.Errorf("%%w", err)
		return
	}
	return
}

func (s *{camel}Service) Delete{class}(ctx context.Context, id uint) (err error) {
	err = s.db.Where("id = ?", id).Delete(&model.{class}{}).Error
	if err != nil {
		err = xerrors.Errorf("%%w", err)
		return
	}
	return
}

func New{class}Service(db *gorm.DB) {class}Service {
	return &{camel}Service{db: db}
}
