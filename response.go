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

type GetInviteResponse struct {
	Code      string      `json:"code"`
	Type      int         `json:"type"`
	ExpiresAt interface{} `json:"expires_at"`
	Guild     struct {
		ID                string      `json:"id"`
		Name              string      `json:"name"`
		Splash            string      `json:"splash"`
		Banner            interface{} `json:"banner"`
		Description       interface{} `json:"description"`
		Icon              string      `json:"icon"`
		Features          []string    `json:"features"`
		VerificationLevel int         `json:"verification_level"`
		VanityURLCode     interface{} `json:"vanity_url_code"`
		WelcomeScreen     struct {
			Description     string `json:"description"`
			WelcomeChannels []struct {
				ChannelID   string      `json:"channel_id"`
				Description string      `json:"description"`
				EmojiID     interface{} `json:"emoji_id"`
				EmojiName   string      `json:"emoji_name"`
			} `json:"welcome_channels"`
		} `json:"welcome_screen"`
		Nsfw      bool `json:"nsfw"`
		NsfwLevel int  `json:"nsfw_level"`
	} `json:"guild"`
	Channel struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"channel"`
	Inviter struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Avatar        string `json:"avatar"`
		Discriminator string `json:"discriminator"`
		PublicFlags   int    `json:"public_flags"`
	} `json:"inviter"`
}
