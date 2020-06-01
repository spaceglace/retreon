package config

import (
	"fmt"
)

func GetGameMode() string {
	return settings.Mode
}

func GetRefresh() int {
	return settings.Refresh
}

func SetGameMode(mode string) error {
	if mode != "DateEarned" && mode != "DateEarnedHardcore" {
		return fmt.Errorf("Mode '%s' is not supported", mode)
	}
	settings.Mode = mode
	return save()
}

func SetRefresh(refresh int) error {
	if refresh < 0 || refresh > 6 {
		return fmt.Errorf("Refresh '%d' is not within bounds 0-6", refresh)
	}
	settings.Refresh = refresh
	return save()
}
