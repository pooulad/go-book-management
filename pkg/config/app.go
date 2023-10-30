package config

var (
	db *gorm.DB
)

type MysqlConfig struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Name      string `json:"name"`
	Port      string `json:"port"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parseTime"`
	Loc       string `json:"loc"`
}

