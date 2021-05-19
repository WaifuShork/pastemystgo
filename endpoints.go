package pastemystgo

import (
	"fmt"
	"net/url"
)

// All the pastemyst API endpoints that can be accessed
const (
	EndpointBase  = "https://paste.myst.rs/api/v2/"
	EndpointData  = EndpointBase + "data/"  // Value: https://paste.myst.rs/api/v2/data/
	EndpointTime  = EndpointBase + "time/"  // Value: https://paste.myst.rs/api/v2/time/
	EndpointUser  = EndpointBase + "user/"  // Value: https://paste.myst.rs/api/v2/user/

	EndpointSelfUser       = EndpointUser + "self/"      // Value: https://paste.myst.rs/api/v2/user/self/
	EndpointSelfUserPastes = EndpointSelfUser + "pastes" // Value: https://paste.myst.rs/api/v2/user/self/pastes
)

var (
	// EndpointPaste - Value: https://paste.myst.rs/api/v2/paste/{pasteId}
	EndpointPaste = func(pasteId string) string {
		return fmt.Sprintf("%spaste/%s", EndpointBase, pasteId)
	}
	// DataLanguageByName - Value:
	// https://paste.myst.rs/api/v2/data/language?name={name}
	DataLanguageByName = func(name string) string {
		return fmt.Sprintf("%slanguage?name=%s", EndpointData, url.QueryEscape(name))
	}

	// DataLanguageByExt - Value:
	// https://paste.myst.rs/api/v2/data/languageExt?extension={extension}
	DataLanguageByExt = func(extension string) string {
		return fmt.Sprintf("%slanguageExt?extension=%s", EndpointData, url.QueryEscape(extension))
	}

	// TimeExpiresInToUnix - Value:
	// https://paste.myst.rs/api/v2/time/expiresInToUnixTime/?createdAt={createdAt}&expiresIn={expiresIn}
	TimeExpiresInToUnix = func(createdAt uint64, expires string) string {
		return fmt.Sprintf("%sexpiresInToUnixTime?createdAt=%d&expiresIn=%s", EndpointTime, createdAt, expires)
	}
)