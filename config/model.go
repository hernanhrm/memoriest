package config

type Model struct {
	AllowedOrigins []string                     `mapstructure:"allowed_origins"`
	AllowedMethods []string                     `mapstructure:"allowed_methods"`
	LogFolder      string                       `mapstructure:"log_folder"`
	Env            string                       `mapstructure:"env"`
	Databases      map[string]map[string]string `mapstructure:"databases"`
}
