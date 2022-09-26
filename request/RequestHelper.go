package request

type User struct {
	ID              string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Username        string `param:"username" query:"username" form:"username" json:"username" xml:"username" validate:"required"`
	Email           string `param:"email" query:"email" form:"email" json:"email" xml:"email" validate:"required"`
	EmailVerifiedAt string `param:"email_verified_at" query:"email_verified_at" form:"email_verified_at" json:"email_verified_at" xml:"email_verified_at" validate:"required"`
	BackgroundColor string `param:"bg_color" query:"bg_color" form:"bg_color" json:"bg_color" xml:"bg_color" validate:"required"`
	TextColor       string `param:"text_color" query:"text_color" form:"text_color" json:"text_color" xml:"text_color" validate:"required"`
}
