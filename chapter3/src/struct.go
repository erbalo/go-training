package src

import (
	"errors"
	"time"
)

type Cache struct {
	m map[string]interface{}
}

func (c *Cache) Set(k string, v interface{}) {
	if c.m == nil {
		c.m = make(map[string]interface{}, 0)
	}
	c.m[k] = v
}

type Customer struct {
	name string
	age  int
}

func NewCustomer(name string, age int) (*Customer, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}
	if age < 0 {
		return nil, errors.New("age should be greater than 0")
	}
	return &Customer{name: name, age: age}, nil
}

func (c *Customer) Age() int {
	return c.age
}

func (c *Customer) SetAge(age int) error {
	if age < 0 {
		return errors.New("age should be greater than 0")
	}
	c.age = age
	return nil
}

type BaseEntity struct {
	id      int
	created time.Time
}

func (b *BaseEntity) ID() int {
	return b.id
}

func NewBaseEntity(id int, created time.Time) BaseEntity {
	return BaseEntity{id: id, created: created}
}

type Customer1 struct {
	BaseEntity
	name string
}

func (c *Customer1) SetCreated(created time.Time) {
	c.created = created
}

func NewCustomer1(id int, created time.Time, name string) *Customer1 {
	return &Customer1{
		BaseEntity: NewBaseEntity(id, created),
		name:       name,
	}
}
