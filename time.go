package pastemystgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExpiresInToUnixTime(createdAt uint64, expires ExpiresIn) (uint64, error) {
	expiresIn := GetExpiresInString(expires)
	url := TimeExpiresInToUnix + fmt.Sprintf("?createdAt=%d&expiresIn=%s", createdAt, expiresIn)

	response, err := http.Get(url)
	if err != nil {
		return 0, sadness("%v", err)
	}

	// Read the responses body to get the raw text 
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil { 
		return 0, sadness("%v", err)
	}

	var pattern map[string]float64

	err = json.Unmarshal(bytes, &pattern)
	if err != nil { 
		return 0, sadness("%v", err)
	}

	return uint64(pattern["result"]), nil

	//return &language, nil
}

func GetExpiresInString(expiresIn ExpiresIn) string { 
	switch expiresIn {
	case Never:
		return "never"
	case OneHour:
		return "1h"
	case TwoHours:
		return "2h"
	case TenHours:
		return "10h"
	case OneDay:
		return "1d"
	case TwoDays:
		return "2d"
	case OneWeek:
		return "1w"
	case OneMonth:
		return "1m"
	case OneYear:
		return "1y"
	default:
		return ""
	}
}