package entity

import (
    "time"
    "github.com/lib/pq"
)

type Person struct {
    Id        int64          `gorm:"column:id;primary_key" json:"id"`
    Name      string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
    Age       int            `gorm:"column:age;type:integer;not null" json:"age"`
    Phones    pq.StringArray `gorm:"column:phones;type:varchar[];not null" json:"phones"`
    CreatedAt int64          `validate:"required;" gorm:"column:created_at;type:timestamp with time zone;not null" json:"created_at"`
    OrderTime int64          `gorm:"column:order_time;type:timestamp with time zone;not null" json:"order_time"`
    UpdatedAt int64          `gorm:"column:updated_at;type:timestamp with time zone;not null" json:"updated_at"`
}

func (p *Person) TableName() string {
    return "person"
}
