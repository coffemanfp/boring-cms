package auth

// Auth represents authentication credentials.
type Auth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
