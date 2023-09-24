package model

import "time"

const (
	ExecutionLogSuccess = 1
	ExecutionLogFail    = 0
)

type ExecutionLog struct {
	Id          *uint     `json:"id,omitempty"`
	Source      string    `json:"source,omitempty"`
	Database    string    `json:"database,omitempty"`
	Sql         string    `json:"sql,omitempty"`
	Status      int       `json:"status,omitempty"`
	Cost        float64   `json:"cost,omitempty"`
	Ip          string    `json:"ip"`
	UserAgent   string    `json:"user_agent"`
	CreatedTime time.Time `json:"created_time,omitempty"`
}

func NewExecutionLog() *ExecutionLog {
	return &ExecutionLog{
		Id:          nil,
		Source:      "",
		Sql:         "",
		Status:      ExecutionLogSuccess,
		Cost:        0,
		CreatedTime: time.Now(),
	}
}
