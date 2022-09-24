package response

type User struct {
	ID              string `json:"id"` 
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	BackgroundColor string `json:"bg_color"`
	TextColor       string `json:"text_color"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}
