package pastemystgo

// ALl the pastemyst API endpoints that can be accessed
const (
	BaseEndpoint  string = `https://paste.myst.rs/api/v2/`
	DataEndpoint  string = BaseEndpoint + `data/`
	TimeEndpoint  string = BaseEndpoint + `time/`
	UserEndpoint  string = BaseEndpoint + `user/`
	PasteEndpoint string = BaseEndpoint + `paste/`
	
	SelfUserEndpoint string = BaseEndpoint + "user/self/"
	SelfUserPastesEndpoint string = SelfUserEndpoint + "pastes"

	DataLanguageByName string = DataEndpoint + `language?name=`
	DataLanguageByExt  string = DataEndpoint + `languageExt?extension=`

	TimeExpiresInToUnix string = TimeEndpoint + `expiresInToUnixTime`
)