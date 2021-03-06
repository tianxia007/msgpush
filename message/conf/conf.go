package conf

import (
	ini "gopkg.in/ini.v1"
)

var Conf *Config

type Config struct {
	GrpcConf
	BaseConf
	RedisConf
	LogConf
}

// protobuf grpc config
type GrpcConf struct {
	GrpcHost                  string `ini:"GrpcHost"`
	GrpcPort                  string `ini:"GrpcPort"`
	GrpcMaxConnectIdleSec     int    `ini:"GrpcMaxConnectIdleSec"`
	GrpcMaxConnectAgeSec      int    `ini:"GrpcMaxConnectAgeSec"`
	GrpcMaxConnectAgeGraceSec int    `ini:"GrpcMaxConnectAgeGraceSec"`
	GrpcTimeSec               int    `ini:"GrpcTimeSec"`
	GrpcTimeTimeoutSec        int    `ini:"GrpcTimeTimeoutSec"`
}

type BaseConf struct {
	StorageType string `ini:"StorageType"`
}

type RedisConf struct {
	HostPort  string `ini:"HostPort"`
	Proto     string `ini:"Proto"`
	MaxIdle   int    `ini:"MaxIdle"`
	MaxActive int    `ini:"MaxActive"`
	AUTH      string `ini:"AUTH"`
}

// Log config
type LogConf struct {
	LogPath  string `ini:"LogPath"`
	LogLevel string `ini:"LogLevel"`
}

func InitConfig(confPath *string) (*Config, error) {
	Conf = new(Config)
	if err := ini.MapTo(Conf, *confPath); err != nil {
		return nil, err
	}

	return Conf, nil
}
