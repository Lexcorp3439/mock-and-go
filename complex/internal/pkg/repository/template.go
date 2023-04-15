package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"heisenbug/complex/internal/pkg/model"
	"heisenbug/complex/internal/pkg/store"
	desc "heisenbug/complex/pkg/api/complex"
	desc2 "heisenbug/complex/pkg/api/public"
)

// Template repo interface
type Template interface {
	GetTemplate(ctx context.Context, templateName string) (*model.Template, error)
	GetTemplatesList(ctx context.Context, filter *desc.GetTemplateListRequest_Filter, offset, limit *uint64) ([]model.Template, error)
	CreateTemplate(ctx context.Context, template *model.Template) error
	UpdateTemplate(ctx context.Context, template *model.Template) error
	DeleteTemplate(ctx context.Context, template *model.Template) error
	CountTemplates(ctx context.Context, filter *desc.GetTemplateListRequest_Filter) (uint64, error)

	GetBindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.TemplateBinding, error)
	GetBindByID(ctx context.Context, id *uuid.UUID) (*model.TemplateBinding, error)
	BindPhoneWithTemplateName(ctx context.Context, bind *model.TemplateBinding) error
	UpdateBindStatusByID(ctx context.Context, id *uuid.UUID, status desc2.BindingStatus) error
	MarkOldBindsAsExpired(ctx context.Context, phoneNumber string) error
}

type templateRepo struct {
	store *store.Storage
}

func NewTemplate(db *store.Storage) Template {
	return &templateRepo{store: db}
}

// GetTemplate - получение шаблона по имени
func (r *templateRepo) GetTemplate(ctx context.Context, templateName string) (*model.Template, error) {
	query := store.PSQL().
		Select("*").
		From(model.TemplateTableName).
		Where("name = ?", templateName)

	record := model.Template{}
	err := r.store.Getx(ctx, query,
		&record.Name,
		&record.Status,
		&record.Description,
		&record.CreatedAt,
		&record.UpdatedAt,
	)
	if err != nil {
		return nil, store.DbError(err)
	}
	return &record, nil
}

// GetTemplatesList - получение всех шаблонов
func (r *templateRepo) GetTemplatesList(ctx context.Context, filter *desc.GetTemplateListRequest_Filter, offset, limit *uint64) ([]model.Template, error) {
	conds := r.applyGetTemplateListFilters(filter)

	query := store.PSQL().
		Select("*").
		From(model.TemplateTableName).
		OrderBy("updated_at DESC").
		Where(conds)

	if limit != nil {
		query = query.Limit(*limit)
	}

	if offset != nil {
		query = query.Offset(*offset)
	}

	templates, err := store.Select[model.Template](ctx, r.store, query)
	if err != nil {
		return nil, store.DbError(err)
	}
	return templates, nil
}

// CreateTemplate - создание фантомного Tsys шаблона
func (r *templateRepo) CreateTemplate(ctx context.Context, template *model.Template) error {
	attrs := map[string]interface{}{
		"status":      template.Status,
		"description": template.Description,
		"name":        template.Name,
	}

	if !template.CreatedAt.IsZero() {
		attrs["created_at"] = template.CreatedAt
	}
	if !template.UpdatedAt.IsZero() {
		attrs["updated_at"] = template.UpdatedAt
	}

	query := store.PSQL().
		Insert(model.TemplateTableName).
		SetMap(attrs).
		Suffix("RETURNING created_at, updated_at")

	err := r.store.Getx(ctx, query, &template.CreatedAt, &template.UpdatedAt)
	if err != nil {
		return store.DbError(err)
	}
	return nil
}

// UpdateTemplate - обновление шаблона
func (r *templateRepo) UpdateTemplate(ctx context.Context, template *model.Template) error {
	attrs := map[string]interface{}{
		"status": template.Status,
	}
	if template.Description != nil {
		attrs["description"] = template.Description
	}

	query := store.PSQL().
		Update(model.TemplateTableName).
		Where("name = ?", template.Name).
		SetMap(attrs)

	err := r.store.Execx(ctx, query)
	if err != nil {
		return store.DbError(err)
	}
	return nil
}

// DeleteTemplate - удаление шаблона
func (r *templateRepo) DeleteTemplate(ctx context.Context, template *model.Template) error {
	query := store.PSQL().
		Delete(model.TemplateTableName).
		Where("name = ?", template.Name)

	err := r.store.Execx(ctx, query)

	if err != nil {
		return store.DbError(err)
	}
	return nil
}

// CountTemplates - получение числа шаблонов
func (r *templateRepo) CountTemplates(ctx context.Context, filter *desc.GetTemplateListRequest_Filter) (uint64, error) {
	conds := r.applyGetTemplateListFilters(filter)

	query := store.PSQL().
		Select("COUNT(*)").From(model.TemplateTableName).
		Where(conds)

	var count uint64
	err := r.store.Getx(ctx, query, &count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *templateRepo) applyGetTemplateListFilters(filter *desc.GetTemplateListRequest_Filter) sq.And {
	conds := make(sq.And, 0, 6)
	if filter == nil {
		return conds
	}
	if filter.CreatedAtFrom != nil {
		conds = append(conds, sq.GtOrEq{"created_at": filter.CreatedAtFrom.AsTime()})
	}
	if filter.CreatedAtTo != nil {
		conds = append(conds, sq.LtOrEq{"created_at": filter.CreatedAtTo.AsTime()})
	}
	if filter.UpdatedAtFrom != nil {
		conds = append(conds, sq.GtOrEq{"updated_at": filter.UpdatedAtFrom.AsTime()})
	}
	if filter.UpdatedAtTo != nil {
		conds = append(conds, sq.LtOrEq{"updated_at": filter.UpdatedAtTo.AsTime()})
	}
	if filter.TemplateName != nil {
		conds = append(conds, sq.Like{"name": "%" + *filter.TemplateName + "%"})
	}
	if filter.TemplateDescription != nil {
		conds = append(conds, sq.Like{"description": *filter.TemplateDescription + "%"})
	}
	return conds
}

// GetBindByID - получение бинда по ID
func (r *templateRepo) GetBindByID(ctx context.Context, id *uuid.UUID) (*model.TemplateBinding, error) {
	query := store.PSQL().
		Select("*").
		From(model.TemplateBindingTableName).
		Where("id = ?", id)

	var template model.TemplateBinding
	err := r.store.Getx(ctx, query, &template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// GetBindByPhoneNumber - получение бинда по номеру телефона
func (r *templateRepo) GetBindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.TemplateBinding, error) {
	query := store.PSQL().
		Select("*").
		From(model.TemplateBindingTableName).
		Where("phone_number = ?", phoneNumber).
		Where("status = ?", desc.BindingStatus_CREATED).
		OrderBy("created_at DESC").
		Limit(1)

	var template model.TemplateBinding
	err := r.store.Getx(ctx, query,
		&template.ID,
		&template.Name,
		&template.PhoneNumber,
		&template.Status,
		&template.CreatedAt,
		&template.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// BindPhoneWithTemplateName - биндинг телефона с шаблоном
func (r *templateRepo) BindPhoneWithTemplateName(ctx context.Context, bind *model.TemplateBinding) error {
	attrs := map[string]interface{}{
		"name":         bind.Name,
		"phone_number": bind.PhoneNumber,
	}
	if bind.Status != nil {
		attrs["status"] = bind.Status
	}

	query := store.PSQL().
		Insert(model.TemplateBindingTableName).
		SetMap(attrs).
		Suffix("RETURNING id, status, created_at, updated_at")

	err := r.store.Getx(ctx, query, &bind.ID, &bind.Status, &bind.CreatedAt, &bind.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBindStatusByID - обновление статуса бинда
func (r *templateRepo) UpdateBindStatusByID(ctx context.Context, id *uuid.UUID, status desc2.BindingStatus) error {
	attrs := map[string]interface{}{
		"status": status,
	}
	query := store.PSQL().
		Update(model.TemplateBindingTableName).
		Where("id = ?", id).
		SetMap(attrs)

	err := r.store.Execx(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// MarkOldBindsAsExpired - маркание старых биндов как протухшие
func (r *templateRepo) MarkOldBindsAsExpired(ctx context.Context, phoneNumber string) error {
	attrs := map[string]interface{}{
		"status": desc.BindingStatus_EXPIRED,
	}
	query := store.PSQL().
		Update(model.TemplateBindingTableName).
		Where("phone_number = ?", phoneNumber).
		SetMap(attrs)

	err := r.store.Execx(ctx, query)
	return err
}
