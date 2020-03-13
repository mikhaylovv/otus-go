package config

import "encoding/json"

type Config struct {
	HttpListen string `json:"http_listen"` // - ip и port на котором должен слушать web-сервер
	LogFile string `json:"log_file"` //- путь к файлу логов
	LogLevel string `json:"log_level"` //- уровень логирования (error / warn / info / debug)
}

func ParseConfig(cf []byte) (*Config, error) {
	cfg := &Config{}
	err := json.Unmarshal(cf, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}