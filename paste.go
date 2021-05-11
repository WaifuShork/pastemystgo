package main

type Paste struct {
	Id        string    `json:"_id"`
	OwnerId   string    `json:"ownerId"`
	Title     string    `json:"title"`
	CreatedAt uint64    `json:"createdAt"`
	ExpiresIn ExpiresIn `json:"expiresIn"`
	DeletesAt uint64    `json:"deletesAt"`
	Stars     uint64    `json:"stars"`
	IsPrivate bool      `json:"isPrivate"`
	IsPublic  bool      `json:"isPublic"`
	Tags      []string  `json:"tags"`
	Pasties   []Pasty   `json:"pasties"`
	Edits     []Edit    `json:"edits"`
}

// TODO: idfk what to do with this lol
type ExpiresIn int

const (
	Never ExpiresIn = iota
	OneHour
	TwoHours
	TenHours
	OneDays
	TwoDays
	OneWeek
	OneMonth
	OneYear
)