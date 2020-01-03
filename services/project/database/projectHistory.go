package database

// ProjectHistory struct
type ProjectHistory struct {
	PID     string
	OID     string `json:"uid"`
	Name    string `json:"name" gorm:"unique"`
	Desc    string `json:"desc"`
	Manager string `json:"manager"`
	Client  string `json:"client"`
	Sector  string `json:"sector"`
	Version int    `json:"version"`
}
