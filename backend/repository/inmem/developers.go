package inmem

import (
	"errors"

	"github.com/giornetta/devcv/devcv"
)

type developerRepository struct {
	m map[string]*devcv.Developer
}

// NewDeveloperRepository returns an in-memory implementation of devcv.DeveloperRepository
func NewDeveloperRepository() devcv.DeveloperRepository {
	return &developerRepository{
		m: map[string]*devcv.Developer{},
	}
}

// Lookup returns all the public information about a developer
func (r *developerRepository) Lookup(username string) (*devcv.Developer, error) {
	dev, ok := r.m[username]
	if !ok {
		return nil, errors.New("developer not found")
	}

	dev.Hash = ""

	return dev, nil
}

// Create creates a new developer
func (r *developerRepository) Create(username, firstName, lastName, password string) error {
	if _, ok := r.m[username]; ok {
		return errors.New("developer already exists")
	}

	r.m[username] = &devcv.Developer{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Hash:      password,
	}

	return nil
}

// Update updates information about a developer
// Update can update every field except the Username.
func (r *developerRepository) Update(developer *devcv.Developer) error {
	if _, ok := r.m[developer.Username]; !ok {
		return errors.New("developer not found")
	}

	r.m[developer.Username] = developer

	return nil
}

// Delete deletes a developer.
func (r *developerRepository) Delete(username string) error {
	if _, ok := r.m[username]; !ok {
		return errors.New("developer not found")
	}

	delete(r.m, username)

	return nil
}

// GetHash returns the hashed password of an user, used for authentication.
func (r *developerRepository) GetHash(username string) (string, error) {
	dev, ok := r.m[username]
	if !ok {
		return "", errors.New("developer not found")
	}

	return dev.Hash, nil
}
