package entity

type Store map[string]string
type SessionID string

type Session struct {
	ID     SessionID
	Values Store
}
