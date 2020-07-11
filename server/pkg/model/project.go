package model

type Project struct {
	ID        int64
	CreatedAt string
	UpdatedAt string
}

type CreateProject struct {
	ID int64
}

type DeleteProject struct {
}
