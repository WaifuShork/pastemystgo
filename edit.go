package pastemystgo

type Edit struct {
	Id       string   `json:"_id"`
	EditId   string   `json:"editId"`
	EditType uint64   `json:"editType"`
	Metadata []string `json:"metadata"`
	Edit     string   `json:"edit"`
	EditedAt uint64   `json:"editedAt"`
}