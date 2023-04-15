package model

import (
	"time"

	"github.com/google/uuid"
	desc "heisenbug/complex/pkg/api/complex"
)

// TemplateBindingTableName - название таблицы в бд
const TemplateBindingTableName = "template_binding"

// TemplateBinding - таблица шаблонов
type TemplateBinding struct {
	ID          *uuid.UUID          `db:"id"`           // ID сообщения в формате UUID
	Name        string              `db:"name"`         // имя шаблона
	PhoneNumber string              `db:"phone_number"` // номер телефона
	Status      *desc.BindingStatus `db:"status"`       // Статус бинда
	CreatedAt   time.Time           `db:"created_at"`   // Дата создания
	UpdatedAt   time.Time           `db:"updated_at"`   // Дата обновления
}
