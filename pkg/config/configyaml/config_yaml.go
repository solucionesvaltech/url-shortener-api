package configyaml

import "url-shortener/pkg/config"

func NewConfigYaml() (*config.Config, error) {
	l := &config.Loader{
		ConfName:  "config",
		ConfType:  "yaml",
		Directory: "../config/configyaml",
	}
	return l.LoadConfig()
}
