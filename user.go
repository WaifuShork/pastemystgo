package pastemystgo

type User struct {
	Id              string `json:"_id"`
	Username        string `json:"username"`
	AvatarUrl       string `json:"avatarUrl"`
	DefaultLang     string `json:"defaultLang"`
	PublicProfile   bool   `json:"publicProfile"`
	SupporterLength uint64 `json:"supporterLength"`
	Contributor     bool   `json:"contributor"`
}