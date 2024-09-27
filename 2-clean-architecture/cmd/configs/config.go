package configs

type (
	Config struct {
		LogLevel string   `yaml:"log-level"`
		Database Database `yaml:"database"`
	}

	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
)
