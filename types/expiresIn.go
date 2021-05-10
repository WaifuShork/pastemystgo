package types

// TODO: idfk what to do with this lol
type ExpiresIn int

const ( 
	Never ExpiresIn = iota
	OneHour
	TwoHours
	TenHours
	OneDays
	TwoDays
	OneWeek
	OneMonth
	OneYear
)