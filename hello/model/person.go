package model

import "github.com/lib/pq"

type PersonCreateRequest struct {
    Id        int64          `json:"id"`
    Name      string         `json:"name"`
    Arm       *string        `json:"arm"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt int64          `json:"created_at"`
    UpdatedAt *int64         `json:"updated_at"`
}
type PersonUpdateRequest struct {
    Id        int64          `json:"id"`
    Name      string         `json:"name"`
    Arm       *string        `json:"arm"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt int64          `json:"created_at"`
    UpdatedAt *int64         `json:"updated_at"`
}
type PersonListRequest struct {
    Id        int64          `json:"id"`
    Name      string         `json:"name"`
    Arm       *string        `json:"arm"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt int64          `json:"created_at"`
    UpdatedAt *int64         `json:"updated_at"`
    Index     int            `json:"index"`
    Size      int            `json:"size"`
}
type PersonListResponse struct {
    Total int           `json:"total"`
    List  []*PersonInfo `json:"list"`
}
type PersonInfoRequest struct {
    Id int64 `json:"id"`
}
type PersonInfo struct {
    Id        int64          `json:"id"`
    Name      string         `json:"name"`
    Arm       string         `json:"arm"`
    Age       int            `json:"age"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt int64          `json:"created_at"`
    UpdatedAt int64          `json:"updated_at"`
}
type PersonDeleteRequest struct {
    Id int64 `json:"id"`
}
