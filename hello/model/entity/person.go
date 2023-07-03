package entity

import (
    "time"
    "github.com/lib/pq"
)

type Person struct {
    Id        int64          `json:"id"`
    Name      string         `json:"name" gorm:"type:VERCHAR(255);NotNull"`
    Arm       string         `json:"name" gorm:"type:VERCHAR(255);NotNull"`
    Age       int            `gorm:"type:INT;NotNull" json:"age"`
    Phones    pq.StringArray `gorm:"type:[]VERCHAR;NotNull" json:"phones"`
    CreatedAt time.Time      `gorm:"type:TIMESTEMP;NotNull" json:"created_at"`
    UpdatedAt time.Time      `gorm:"type:TIMESTEMP;NotNull" json:"updated_at"`
}

func (p *Person) String() string {
    return "person"
}
