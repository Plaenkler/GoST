package model

import "time"

type Article struct {
	ID             string  `gorm:"primaryKey"`
	PurchasePrice  float64 `gorm:"not null"`
	Name           string  `gorm:"not null"`
	ManufacturerID uint64
	Manufacturer   Manufacturer
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Agent struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DetectedError struct {
	ID          string `gorm:"primaryKey"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Manufacturer struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomerError struct {
	ID          string `gorm:"primaryKey"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Technician struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WorkshopInvoice struct {
	ID               string `gorm:"primaryKey"`
	ReferedArticleID string `gorm:"not null"`
	ReferedArticle   Article
	Warranty         bool    `gorm:"not null"`
	SellPrice        float64 `gorm:"not null"`
	DeliveryCost     float64 `gorm:"not null"`
	PaymentCondition uint64  `gorm:"not null"`
	AgentID          uint64
	Agent            Agent
	TechnicianID     uint64
	Technician       Technician
	Date             time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type WorkshopInvoiceDetectedError struct {
	WorkshopInvoiceID string
	WorkshopInvoice   WorkshopInvoice `gorm:"not null"`
	DetectedErrorID   string
	DetectedError     DetectedError `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type WorkshopInvoiceCustomerError struct {
	WorkshopInvoiceID string
	WorkshopInvoice   WorkshopInvoice `gorm:"not null"`
	CustomerErrorID   string
	CustomerError     CustomerError `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type WorkshopInvoicePosition struct {
	ID                uint64  `gorm:"primaryKey"`
	PositionType      string  `gorm:"not null"`
	SellPrice         float64 `gorm:"not null"`
	Discount          float64 `gorm:"not null"`
	Amount            uint64  `gorm:"not null"`
	WorkshopInvoiceID string
	WorkshopInvoice   WorkshopInvoice `gorm:"not null"`
	ArticleID         string
	Article           Article `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
