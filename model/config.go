package model

type Config struct {
	Id           *uint   `json:"id,omitempty"`
	Type         *string `json:"type,omitempty" binding:"required"`
	Name         *string `json:"name,omitempty" binding:"required"`
	Host         *string `json:"host,omitempty" binding:"required"`
	Port         *string `json:"port,omitempty" binding:"required"`
	Username     *string `json:"username,omitempty" binding:"required"`
	Password     *string `json:"password,omitempty" binding:"required"`
	Charset      *string `json:"charset,omitempty"`
	Timeout      *uint64 `json:"timeout,omitempty"`
	ReadTimeout  *uint64 `json:"read_timeout,omitempty"`
	WriteTimeout *uint64 `json:"write_timeout,omitempty"`
	DbFile       *string `json:"dbFile,omitempty"`
}

func NewConfig() *Config {
	return &Config{
		Id:           nil,
		Type:         nil,
		Name:         nil,
		Host:         nil,
		Port:         nil,
		Username:     nil,
		Password:     nil,
		Charset:      nil,
		Timeout:      nil,
		ReadTimeout:  nil,
		WriteTimeout: nil,
		DbFile:       nil,
	}
}

func (c Config) TableName() string {
	return "sys_config"
}
