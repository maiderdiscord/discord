package discord

import "github.com/go-utils/cont"

type Platform int

const PlatformLinux Platform = iota

func SupportedPlatform(platform Platform) bool {
	supportedPlatforms := []Platform{PlatformLinux}

	return cont.Contains(supportedPlatforms, platform)
}
