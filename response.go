package discord

type MeResponse struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	Avatar            string `json:"avatar"`
	Discriminator     string `json:"discriminator"`
	PublicFlags       int    `json:"public_flags"`
	Flags             int    `json:"flags"`
	PurchasedFlags    int    `json:"purchased_flags"`
	PremiumUsageFlags int    `json:"premium_usage_flags"`
	Banner            string `json:"banner"`
	BannerColor       *int   `json:"banner_color,omitempty"`
	AccentColor       *int   `json:"accent_color,omitempty"`
	Bio               string `json:"bio"`
	Locale            string `json:"locale"`
	NSFWAllowed       bool   `json:"nsfw_allowed"`
	MFAEnabled        bool   `json:"mfa_enabled"`
	PremiumType       int    `json:"premium_type"`
	Email             string `json:"email"`
	Verified          bool   `json:"verified"`
	Phone             string `json:"phone"`
}
