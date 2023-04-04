package service

import (
	"context"
	"{{.Mod}}/internal/entity"

	"golang.org/x/xerrors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type {{.Class}}Service interface {
	Get{{.Class}}List(ctx context.Context, pageIndex int, pageSize int) (list []*entity.{{.Class}}, total int64, err error)
	Get{{.Class}}All(ctx context.Context) (list []*entity.{{.Class}}, err error)
	Get{{.Class}}Detail(ctx context.Context, id uint) ({{.Camel}} *entity.{{.Class}}, err error)
	Create{{.Class}}(ctx context.Context, {{.Camel}} *entity.{{.Class}}) (err error)
	Update{{.Class}}(ctx context.Context, {{.Camel}} *entity.{{.Class}}) (err error)
	Delete{{.Class}}(ctx context.Context, id uint) (err error)
}

type {{.Camel}}Service struct {
	db *gorm.DB
}

func (s *{{.Camel}}Service) Get{{.Class}}List(ctx context.Context, pageIndex int, pageSize int) ({{.Camel}}List []*entity.{{.Class}}, total int64, err error) {
	query := s.db.WithContext(ctx).Model(&entity.{{.Class}}{})
	err = query.Count(&total).
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Order("id desc").
		Find(&{{.Camel}}List).Error

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *{{.Camel}}Service) Get{{.Class}}All(ctx context.Context) ({{.Camel}}List []*entity.{{.Class}}, err error) {
	err = s.db.WithContext(ctx).Model(&entity.{{.Class}}{}).
		Order("id desc").
		Find(&{{.Camel}}List).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *{{.Camel}}Service) Get{{.Class}}Detail(ctx context.Context, id uint) ({{.Camel}} *entity.{{.Class}}, err error) {
	err = s.db.WithContext(ctx).Model(&entity.{{.Class}}{}).Where("id = ?", id).Find(&{{.Camel}}).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *{{.Camel}}Service) Create{{.Class}}(ctx context.Context, {{.Camel}} *entity.{{.Class}}) (err error) {
	err = s.db.WithContext(ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create({{.Camel}}).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *{{.Camel}}Service) Update{{.Class}}(ctx context.Context, {{.Camel}} *entity.{{.Class}}) (err error) {
	err = s.db.WithContext(ctx).Model(&entity.{{.Class}}{}).Where("id = ?", {{.Camel}}.Id).Updates({{.Camel}}).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *{{.Camel}}Service) Delete{{.Class}}(ctx context.Context, id uint) (err error) {
	err = s.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.{{.Class}}{}).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func New{{.Class}}Service(db *gorm.DB) {{.Class}}Service {
	return &{{.Camel}}Service{db: db}
}
