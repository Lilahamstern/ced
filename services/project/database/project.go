package database

import (
	"fmt"
	"time"
)

// Project struct
type Project struct {
	PID       string    `json:"pid"`
	OID       string    `json:"oid"`
	Name      string    `json:"name"`
	Desc      string    `json:"description"`
	Manager   string    `json:"manager"`
	Client    string    `json:"client"`
	Sector    string    `json:"sector"`
	Version   int       `json:"version,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Save project ot database, reutrn error or project if succeeded
func (p *Project) Save() (Project, error) {
	stmt := `INSERT INTO project(pid, oid, name, description, manager, client, sector) 
VALUES ($1, $2, $3, $4, $5, $6, $7);
`
	var id = 0
	err := db.QueryRow(stmt, p.PID, p.OID, p.Name, p.Desc, p.Manager, p.Client, p.Sector).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	return Project{}, nil
}
