package entity

import "github.com/google/uuid"

const (
	LinesCntCap  = 99
	LineTotalCap = 1000
)

type OrderEntity struct {
	ID          string
	Description string
	OwnerID     int64
	Lines       []OrderLineEntity
	Status      string
	CreatedTime string
	UpdatedTime string
}

type OrderLineEntity struct {
	ID        string
	ProductID string
	Quantity  float64
	Price     float64
	Total     float64
	Comment   string
}

func (en *OrderEntity) Create() {
	en.complete()
	en.ruleCheck()
}

func (en *OrderEntity) complete() {
	en.Status = "open"
	en.ID = uuid.New().String()

	for _, line := range en.Lines {
		line.complete()
	}
}

func (line *OrderLineEntity) complete() {
	line.Total = line.Quantity * line.Price
}

func (en *OrderEntity) ruleCheck() {
	lines := en.Lines
	if len(lines) > LinesCntCap {
	}

	for _, line := range lines {
		if line.Total > LineTotalCap {
		}

	}
}
