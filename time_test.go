package pastemystgo

import (
	"testing"
)

func TestTimeExpiresInProperly(t *testing.T) { 
	tests := []struct {
		createdAt uint64
		expiresIn ExpiresIn
		expected  uint64
	}{
		{
			createdAt: 1615242814,
            expiresIn: TwoHours,
            expected: 1615250014,
		},
		{
			createdAt: 1615121479,
			expiresIn: OneDay,
			expected: 1615207879,
		},
		{
			createdAt: 1615297946,
			expiresIn: OneWeek,
			expected: 1615902746,
		},
		{
			createdAt: 1588441258,
			expiresIn: OneWeek,
			expected: 1589046058,
		},
	}

	for _, tt := range tests { 
		value, _ := ExpiresInToUnixTime(tt.createdAt, tt.expiresIn)
		if value != tt.expected { 
			t.Errorf("Time format incorrect %d", value)
		}
	}
}