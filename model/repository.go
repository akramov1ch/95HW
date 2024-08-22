package model

import "time"

type RepositoryEvent struct {
	ID        uint      `gorm:"primaryKey"`
	Action    string    `json:"action"`
	RepoName  string    `json:"repository.name"`
	RepoOwner string    `json:"repository.owner.login"`
	CreatedAt time.Time
}
