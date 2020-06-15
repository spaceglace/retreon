package config

import "fmt"

func GetLayouts() map[string]string {
	if len(settings.Layouts) == 0 {
		AddLayout(
			"default",
			`{"name":"default","width":300,"height":600,"background":"#1E1E1E","auto":false,"interval":30,"widgets":[]}`,
		)
	}
	return settings.Layouts
}

func AddLayout(name string, serialized string) error {
	// check if the name is already in use
	if _, ok := settings.Layouts[name]; ok {
		return fmt.Errorf("Name '%s' is already in use", name)
	}

	settings.Layouts[name] = serialized
	return save()
}

func RemoveLayout(name string) error {
	delete(settings.Layouts, name)
	return save()
}

func UpdateLayout(name string, serialized string) error {
	settings.Layouts[name] = serialized
	return save()
}

func GetActiveLayout() string {
	if settings.Layout == "" {
		return "default"
	}
	return settings.Layout
}

func SetActiveLayout(name string) error {
	// ensure layout exists
	if _, ok := settings.Layouts[name]; !ok {
		return fmt.Errorf("Layout '%s' does not exist", name)
	}
	settings.Layout = name
	return save()
}
