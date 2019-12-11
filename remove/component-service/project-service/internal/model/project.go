package model

type Project struct {
	Id     string `json:"id"`
	Client string `json:"client"`
	Sector string `json:"sector"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Desc   string `json:"desc"`
	State  string `json:"state"`
}

type Projects []Project
