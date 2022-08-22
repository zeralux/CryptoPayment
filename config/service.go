package config

type Service struct {
	port int `mapstructure:"port"`
}

func (s Service) Port() int {
	return s.port
}
