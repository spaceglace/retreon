package config

import (
	"fmt"
)

type Profile struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func GetProfiles() map[string]Profile {
	return settings.Profiles
}

func AddProfile(name, key string) error {
	settings.Profiles[name] = Profile{
		Name: name,
		Key:  key,
	}
	return SetActiveProfile(name)
}

func RemoveProfile(name string) error {
	delete(settings.Profiles, name)
	return save()
}

func GetActiveProfile() string {
	return settings.Profile
}

func SetActiveProfile(name string) error {
	// ensure profile exists
	if _, ok := settings.Profiles[name]; !ok {
		return fmt.Errorf("Profile '%s' does not exist", name)
	}
	settings.Profile = name
	return save()
}
