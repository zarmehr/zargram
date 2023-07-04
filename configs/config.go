package configs

import "fmt"

type DatabaseConnConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (d DatabaseConnConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe",
		d.Host, d.Port, d.User, d.Password, d.DbName)
}
