package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"url-shortener/pkg/log"
)

type Loader struct {
	ConfType  string
	ConfName  string
	Directory string
}

func (l *Loader) LoadConfig() (*Config, error) {
	log.InitLogger()
	var conf Config
	filepath := getConfigPath(l.Directory)

	viper.SetConfigType(l.ConfType)
	viper.SetConfigName(l.ConfName)
	viper.AddConfigPath(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		return &conf, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return &conf, err
	}

	conf.SetOnFlyVariables()

	validate := validator.New()
	err = validate.Struct(&conf)
	if err != nil {
		return &conf, err
	}
	log.Log.SetLevel(log.GetLogLevel(conf.LogLevel))

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(&conf)
		if err != nil {
			log.Log.Infof("an error occur while deserializing new config: %v\n", err)
			return
		}

		err = validate.Struct(&conf)
		if err != nil {
			log.Log.Infof("an error occur while validating new config: %v", err)
			return
		}

		log.Log.Infof("config updated: %v", conf.AppName)
		log.Log.SetLevel(log.GetLogLevel(conf.LogLevel))
	})

	return &conf, nil
}

func getConfigPath(localConfigPath string) string {
	if IsLocalEnvironment() {
		var (
			_, directory, _, _ = runtime.Caller(1)
		)
		return path.Join(path.Dir(directory), localConfigPath)
	}
	return SecretsPath
}
