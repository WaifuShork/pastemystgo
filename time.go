package pastemystgo

// ExpiresInToUnixTime returns the created at time in unix format. Will error out and return 0 if the
// request was malformed.
//
// Params:
//  (createdAt uint64, expires ExpiresIn)
//
// Returns:
//	(uint64, error)
func (c *Client) ExpiresInToUnixTime(createdAt uint64, expires ExpiresIn) (uint64, error) {
	expiresIn := c.getExpiresInString(expires)
	url := TimeExpiresInToUnix(createdAt, expiresIn)
	var pattern map[string]uint64

	err := c.get(url, &pattern)
	if err != nil {
		return 0, err
	}

	return pattern["result"], nil
}

// Gets the string value of ExpiresIn input. 
//
// Returns:
//  (string)
//
// Remarks: will return "" if the method is unable to locate your request.
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