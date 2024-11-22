package config

type Configuration struct {
	AddonsDirectory string `json:"addons_directory"`
	GameVersion     string `json:"game_version"` // retail, classic etc. for now we want to basically only support retail
	// fuck classic andys
}
