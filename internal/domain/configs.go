package domain

type Config struct {
	Sever    Server `json:"server"`
	Database DB     `json:"db"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}
