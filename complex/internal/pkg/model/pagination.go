package model

const (
	// PerPageDefault дефолтное значение для пагинации
	PerPageDefault = 25
	// PerPageMax максимально-возможное кол-во записей на странице
	PerPageMax = 250
)

// Pagination пагинация
type Pagination struct {
	Page    uint64
	PerPage uint64
}

// GetPerPage возвращает количество строк на странице
func (g *Pagination) GetPerPage() uint64 {
	if g.PerPage <= 0 {
		return PerPageDefault
	}

	return g.PerPage
}

// GetOffset возвращает номер строки, с которой надо начинать выборку
func (g *Pagination) GetOffset() uint64 {
	if g.Page <= 0 {
		return 0
	}

	return (g.Page - 1) * g.GetPerPage()
}
