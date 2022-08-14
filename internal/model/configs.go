package model

type Config struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	Database  DB     `json:"db"`
	SecretKey string `json:"secret_key"`
}

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}
