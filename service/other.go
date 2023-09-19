package service

import (
	"os"
)

type OtherService struct {
}

func (s OtherService) LoadReadme() (string, error) {
	b, err := os.ReadFile("README.md")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
