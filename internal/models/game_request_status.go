// models/game_request_status.go
package models

// Ensure this type is not redeclared elsewhere in the project
type GameRequestStatus string

const (
    Pending  GameRequestStatus = "pending"
    Accepted GameRequestStatus = "accepted"
    Rejected GameRequestStatus = "rejected"
)
