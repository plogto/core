package constants

type ConnectionResult string

const (
	Followers ConnectionResult = "followers"
	Following ConnectionResult = "following"
	Requests  ConnectionResult = "requests"
)
