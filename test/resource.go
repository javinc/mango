package test

import (
	"strconv"
	"strings"
	"time"

	r "github.com/dancannon/gorethink"
	"github.com/javinc/mango/module"
)

const tableName = "test"

// Resource test
type Resource struct {
}

// Object test
type Object struct {
	ID          string    `gorethink:"id,omitempty" json:"id,omitempty"`
	Title       string    `gorethink:"title,omitempty" json:"title,omitempty"`
	Description string    `gorethink:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time `gorethink:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorethink:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Option test
type Option struct {
	Filter Object
	Slice  string
	Sort   string
	Search string
}

func init() {
	r.TableCreate(tableName).Run(module.RSession)
}

// Find test
func (t *Resource) Find(o Option) ([]Object, error) {
	data := []Object{}

	q := r.Table(tableName)

	// slicing
	if o.Slice != "" {
		slice := strings.Split(o.Slice, ",")
		start, _ := strconv.Atoi(slice[0])
		end, _ := strconv.Atoi(slice[1])
		q = q.Slice(start, end)
	}

	// sorting
	if o.Sort != "" {
		sort := strings.Split(strings.ToLower(o.Sort), ",")
		if len(sort) == 2 && sort[1] == "desc" {
			q = q.OrderBy(r.Desc(sort[0]))
		} else {
			q = q.OrderBy(sort[0])
		}
	}

	// filtering
	q = q.Filter(o.Filter)

	res, err := q.Run(module.RSession)
	if err != nil {
		return data, err
	}

	err = res.All(&data)
	if err != nil {
		return data, err
	}

	return data, err
}

// Get test
func (t *Resource) Get(id string) (Object, error) {
	data := Object{}

	res, err := r.Table(tableName).Get(id).Run(module.RSession)
	if err != nil {
		return data, err
	}

	err = res.One(&data)
	if err != nil {
		return data, err
	}

	return data, err
}

// Create test
func (t *Resource) Create(p Object) (Object, error) {
	// meta data
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	// insert to database
	_, err := r.Table(tableName).Insert(p).RunWrite(module.RSession)
	if err != nil {
		return Object{}, err
	}

	return p, err
}

// Update test
func (t *Resource) Update(p Object, id string) (Object, error) {
	// check item if exists
	_, err := t.Get(id)
	if err != nil {
		return Object{}, err
	}

	// update meta data
	p.UpdatedAt = time.Now()

	// insert to database
	_, err = r.Table(tableName).Get(id).Update(p).RunWrite(module.RSession)
	if err != nil {
		return Object{}, err
	}

	return p, err
}

// Remove test
func (t *Resource) Remove(id string) (bool, error) {
	// check item if exists
	_, err := t.Get(id)
	if err != nil {
		return false, err
	}

	// insert to database
	_, err = r.Table(tableName).Get(id).Delete().RunWrite(module.RSession)
	if err != nil {
		return false, err
	}

	return true, nil
}
