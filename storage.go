package main

import (
	"errors"
	"sync"
)

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    string `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee) error
	Get(id int) (Employee, error)
	Update(id int, e *Employee)
	Delete(id int) error
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) error {
	s.Lock()

	e.Id = s.counter
	s.data[e.Id] = *e

	s.counter++

	s.Unlock()

	return nil
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	e, exists := s.data[id]

	if !exists {
		return Employee{}, errors.New("employee with such id doesn't exist")
	}

	return e, nil
}

func (s *MemoryStorage) Update(id int, newEmployee *Employee) {
	s.data[id] = *newEmployee
}

func (s *MemoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}
