package pastemystgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Returns the created at time in unix format. Will error out and return 0 if the
// request was malformed.
//
// Returns:
//  (uint64, error)
func (c *Client) ExpiresInToUnixTime(createdAt uint64, expires ExpiresIn) (uint64, error) {
	expiresIn := c.getExpiresInString(expires)
	url := TimeExpiresInToUnix + fmt.Sprintf("?createdAt=%d&expiresIn=%s", createdAt, expiresIn)

	response, err := http.Get(url)
	if err != nil {
		return 0, newError(err)
	}

	// Read the responses body to get the raw text 
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil { 
		return 0, newError(err)
	}

	// Pattern of the value to locate from the response.Body bytes 
	var pattern map[string]float64
	err = c.deserializeJson(bytes, &pattern)
	if err != nil { 
		return 0, newError(err)
	}

	return uint64(pattern["result"]), nil
}

// Gets the string value of ExpiresIn input. 
//
// Will return "" if the method is unable to locate your request. 
//
// Returns:
//  (string)
func (c *Client) getExpiresInString(expiresIn ExpiresIn) string { 
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