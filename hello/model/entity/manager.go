package entity

import (
    "time"
    "github.com/lib/pq"
)

type Manager struct {
    Id        int64          `json:"id" gorm:"column:id;primary_key"`
    Name      string         `json:"name" gorm:"column:name;type:varchar(255);not null"`
    Arm       string         `json:"name" gorm:"column:arm;type:varchar(255);not null"`
    Age       int            `gorm:"column:age;type:integer;not null" json:"age"`
    Phones    pq.StringArray `gorm:"column:phones;type:varchar[];not null" json:"phones"`
    CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null" json:"created_at"`
    UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null" json:"updated_at"`
}

func (m *Manager) TableName() string {
    return "manager"
}
