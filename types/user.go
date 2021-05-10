package types

type User struct {
	Id              string
	Username        string
	AvatarUrl       string
	DefaultLang     string
	PublicProfile   bool
	SupporterLength uint64
	Contributor     bool
}