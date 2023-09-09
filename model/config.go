package model

type ConnectionConfig struct {
	Id       *uint   `db:"id" json:"id,omitempty"`
	Type     *string `db:"type" json:"type,omitempty" binding:"required"`
	Name     *string `db:"name" json:"name,omitempty" binding:"required"`
	Host     *string `db:"host" json:"host,omitempty" binding:"required"`
	Port     *string `db:"port" json:"port,omitempty" binding:"required"`
	Username *string `db:"username" json:"username,omitempty" binding:"required"`
	Password *string `db:"password" json:"password,omitempty" binding:"required"`
	DbFile   *string `db:"db_file" json:"dbFile,omitempty"`
}

func NewConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{
		Id:       nil,
		Type:     nil,
		Name:     nil,
		Host:     nil,
		Port:     nil,
		Username: nil,
		Password: nil,
		DbFile:   nil,
	}
}
