package configs

type (
	Config struct {
		LogLevel    string      `yaml:"log-level"`
		Database    Database    `yaml:"database"`
		RestServer  RestServer  `yaml:"rest-server"`
		GrpcService GrpcService `yaml:"grpc-service"`
	}

	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}

	RestServer struct {
		Port string `yaml:"port"`
	}

	GrpcService struct {
		Port string `yaml:"port"`
	}
)
