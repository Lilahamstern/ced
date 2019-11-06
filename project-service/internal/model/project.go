package model

type Project struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Model string `json:"model"`
}

type Projects []Project
