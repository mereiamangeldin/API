package toDo

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID        int        `gorm:"column:id"`
	Message   string     `json:"message"`
	Complete  bool       `json:"complete"`
	DeletedAt *time.Time `gorm:"Index"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	Updated   *time.Time `gorm:"column:updated"`
}

type Getter interface {
	GetAll() []Task
}

type Adder interface {
	Add(task Task)
}

type Repo struct {
	Tasks []Task
}

func New() *Repo {
	return &Repo{
		Tasks: []Task{},
	}
}

func (r *Repo) Add(task Task) {
	r.Tasks = append(r.Tasks, task)
}

func (r *Repo) GetAll() []Task {
	return r.Tasks
}
