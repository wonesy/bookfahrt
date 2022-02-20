package auth

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	InvitationID string `json:"invitation,omitempty"`
}
