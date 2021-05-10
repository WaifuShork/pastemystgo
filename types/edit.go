package types

type Edit struct {
	Id       string
	EditId   string
	EditType uint64
	Metadata []string
	Edit     string
	EditedAt uint64
}

type EditType int

const (
	Title EditType = iota
	PastyTitle
	PastyLanguage
	PastyContent
	PastyAdded
	PastyRemoved
)