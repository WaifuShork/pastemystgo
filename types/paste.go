package types

type Paste struct {
	Id        string
	OwnerId   string
	Title     string
	CreatedAt uint64
	ExpiresIn ExpiresIn
	DeletesAt uint64
	Stars     uint64
	IsPrivate bool
	IsPublic  bool
	Tags      []string
	Pasties   []Pasty
	Edits     []Edit
}