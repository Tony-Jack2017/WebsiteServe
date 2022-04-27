package read_config

//type Config struct {
//	application Application
//	logger Logger
//	jwt Jwt
//	database Database
//}

type Config struct {
	Application
	Logger
	Jwt
	Database
}

type Application struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Logger struct {
	Path string `yaml:"path"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
	Expire int `yaml:"expire"`
}

type Database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}