package discord

import "net/http"

func GetHeaders(token string, platform Platform) http.Header {
	headers := make(http.Header)

	// common headers
	headers.Add("Content-Type", "application/json")
	headers.Add("Authorization", token)

	if userAgent := GetUserAgent(platform); len(userAgent) > 0 {
		headers.Add("User-Agent", userAgent)
	}

	if superProperties := GetSuperProperties(platform); len(superProperties) > 0 {
		headers.Add("X-Super-Properties", superProperties)
	}

	return headers
}

func GetUserAgent(platform Platform) string {
	// TODO: バージョンを埋め込めるようにする
	switch platform {
	case PlatformLinux:
		return "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.16 Chrome/91.0.4472.164 Electron/13.4.0 Safari/537.36"
	}

	return ""
}

func GetSuperProperties(platform Platform) string {
	// TODO: 構造体からJSON文字列を生成するようにする
	switch platform {
	case PlatformLinux:
		return "eyJvcyI6IkxpbnV4IiwiYnJvd3NlciI6IkRpc2NvcmQgQ2xpZW50IiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X3ZlcnNpb24iOiIwLjAuMTYiLCJvc192ZXJzaW9uIjoiNS4xMS4wLTQzLWdlbmVyaWMiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwid2luZG93X21hbmFnZXIiOiJVbml0eSx1YnVudHUiLCJkaXN0cm8iOiJVYnVudHUgMjAuMDQuMyBMVFMiLCJjbGllbnRfYnVpbGRfbnVtYmVyIjoxMDkxODMsImNsaWVudF9ldmVudF9zb3VyY2UiOm51bGx9"
	}

	return ""
}
