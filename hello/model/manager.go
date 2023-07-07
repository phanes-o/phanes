package model

import "github.com/lib/pq"

type ManagerCreateRequest struct {
    Id        int64          `json:"id"`
    Name      string         `binding:"required" json:"name"`
    Arm       *string        `json:"arm"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt *int64         `json:"created_at"`
    UpdatedAt *int64         `json:"updated_at"`
}
type ManagerUpdateRequest struct {
    Id        int64          `json:"id"`
    Name      string         `binding:"required" json:"name"`
    Arm       *string        `json:"arm"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt *int64         `json:"created_at"`
    UpdatedAt *int64         `json:"updated_at"`
}
type ManagerListRequest struct {
    Id        int64          `json:"id"`
    Name      string         `binding:"required" json:"name"`
    Arm       *string        `json:"arm"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt *int64         `json:"created_at"`
    UpdatedAt *int64         `json:"updated_at"`
    Index     int            `json:"index"`
    Size      int            `json:"size"`
}
type ManagerListResponse struct {
    Total int            `json:"total"`
    List  []*ManagerInfo `json:"list"`
}
type ManagerInfoRequest struct {
    Id int64 `json:"id"`
}
type ManagerInfo struct {
    Id        int64          `json:"id"`
    Name      string         `binding:"required" json:"name"`
    Arm       string         `json:"arm"`
    Age       int            `json:"age"`
    Phones    pq.StringArray `json:"phones"`
    CreatedAt int64          `json:"created_at"`
    UpdatedAt int64          `json:"updated_at"`
}
type ManagerDeleteRequest struct {
    Id int64 `json:"id"`
}
