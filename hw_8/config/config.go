package config

import "encoding/json"

// Config - structure with microservice configuration
type Config struct {
	HTTPLiten string `json:"http_listen"` // - ip и port на котором должен слушать web-сервер
	LogFile   string `json:"log_file"`    //- путь к файлу логов
	LogLevel  string `json:"log_level"`   //- уровень логирования (error / warn / info / debug)
}

// ParseConfig - parse row JSON into Config struct, u sing Unmarshal
func ParseConfig(cf []byte) (*Config, error) {
	cfg := &Config{}
	err := json.Unmarshal(cf, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}