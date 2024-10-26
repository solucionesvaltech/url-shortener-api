package config

import (
	"os"
	"url-shortener/pkg/log"
)

const (
	EnvironmentName  = "STAGE"
	NamespaceName    = "NAMESPACE"
	ClusterName      = "CLUSTER"
	LocalEnvironment = "local"
	SecretsPath      = "/config"
)

type RouteConfig struct {
	Method string `json:"method" validate:"required"`
	Path   string `json:"path" validate:"required"`
}

type Routes struct {
	Health  RouteConfig `mapstructure:"health" validate:"required"`
	Create  RouteConfig `mapstructure:"create" validate:"required"`
	Get     RouteConfig `mapstructure:"get" validate:"required"`
	Update  RouteConfig `mapstructure:"update" validate:"required"`
	Toggle  RouteConfig `mapstructure:"toggle" validate:"required"`
	Details RouteConfig `mapstructure:"details" validate:"required"`
}

type ServerConfig struct {
	Port           string `json:"port" validate:"required"`
	TimeoutMinutes int64  `json:"timeoutMinutes" validate:"required"`
	Routes         Routes `mapstructure:"routes"`
}

type DynamoConfig struct {
	TableName string `mapstructure:"tableName" validate:"required"`
	Endpoint  string `mapstructure:"endpoint" validate:"required"`
	Region    string `mapstructure:"region" validate:"required"`
}

type RedisCondig struct {
	Address           string `mapstructure:"address" validate:"required"`
	Password          string `mapstructure:"password" validate:"required"`
	DB                int    `mapstructure:"db"`
	ExpirationMinutes int    `mapstructure:"expiration" validate:"required"`
}

type DatabasesConfig struct {
	DynamoDB DynamoConfig `mapstructure:"dynamoDB" validate:"required"`
	Redis    RedisCondig  `mapstructure:"redis" validate:"required"`
}

type Config struct {
	AppName         string          `json:"appName" validate:"required"`
	Namespace       string          `json:"namespace" validate:"required"`
	Cluster         string          `json:"cluster" validate:"required"`
	Stage           string          `json:"stage" validate:"required"`
	LogLevel        log.LogLevel    `json:"logLevel" validate:"required"`
	ServerConfig    ServerConfig    `mapstructure:"server" validate:"required"`
	DatabasesConfig DatabasesConfig `mapstructure:"database" validate:"required"`
}

func (cfg *Config) SetOnFlyVariables() {
	cfg.Namespace = getEnvWithDefault(NamespaceName, LocalEnvironment)
	cfg.Stage = getEnvWithDefault(EnvironmentName, LocalEnvironment)
	cfg.Cluster = getEnvWithDefault(ClusterName, LocalEnvironment)
}

func IsLocalEnvironment() (local bool) {
	environment := os.Getenv(EnvironmentName)
	local = environment == LocalEnvironment || environment == ""
	return
}

func getEnvWithDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
