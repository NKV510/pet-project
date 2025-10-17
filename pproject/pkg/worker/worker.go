package worker

import (
	"math/rand"
	"strconv"
	"time"
)

type List struct {
	worker map[string]Worker
}

func NewList() *List {
	return &List{
		worker: make(map[string]Worker),
	}
}

type Worker struct {
	Name        string `json:"name"`
	SecondName  string `json:"secondName"`
	WorkStatus  bool
	StartTime   time.Time
	EndWorkTime time.Time
}

func NewWorker(name string, secondName string) Worker {
	return Worker{
		Name:       name,
		SecondName: secondName,
		WorkStatus: false,
	}
}

func (l *List) AddWorker(name string, secondName string) string {
	workerNew := NewWorker(name, secondName)
	id := strconv.Itoa(rand.Intn(1000000))
	l.worker[string(id)] = workerNew
	return id
}

func (l *List) StartWork(id string) Worker {
	user := l.worker[id]
	user.WorkStatus = true
	user.StartTime = time.Now()
	user.EndWorkTime = time.Time{}
	l.worker[id] = user
	return user
}

func (l *List) EndWork(id string) Worker {
	u := l.worker[id]
	u.WorkStatus = false
	u.EndWorkTime = time.Now()
	l.worker[id] = u
	return u
}

func (l List) GetAllWorkers() any {
	return l.worker
}

func (l *List) DeleteWorker(id string) {
	delete(l.worker, id)
}
