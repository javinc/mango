package service

import r "github.com/javinc/puto/user/resource"

// Create service
func Create(m *r.Model) (r.Model, error) {
	row, err := r.Create(m)
	if err != nil {
		return r.Model{}, err
	}

	return row, err
}