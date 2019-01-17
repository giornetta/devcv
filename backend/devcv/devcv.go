package devcv

import "time"

// A Developer represents all the information about a developer stored in a database
type Developer struct {
	Username string `db:"username"`
	Hash     string `db:"pass"`

	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`

	Speciality   string `db:"speciality"`
	Timezone     string `db:"timezone"`
	Introduction string `db:"introduction"`
	City         string `db:"city"`
	Languages    string `db:"languages"`

	Links       []Link       `db:"-"`
	Projects    []Project    `db:"-"`
	SkillGroups []SkillGroup `db:"-"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// A Project is owned by an user
type Project struct {
	Title string `db:"title"`
	Link  string `db:"link"`
	Stack string `db:"stack"`
	Scope string `db:"scope"`
}

// A SkillGroup is a group of skills related to each other identified by a title
type SkillGroup struct {
	Title  string  `db:"title"`
	Skills []Skill `db:"-"`
}

// A Skill is an individual skill of a developer and it can have an experience level
type Skill struct {
	Title      string `db:"title"`
	Experience int    `db:"experience"`
}

// A Link is a link to a developer's social network profile/website/blog/etc
type Link struct {
	Title string `db:"title"`
	URL   string `db:"url"`
}

// The DeveloperRepository interface provides access to developers
type DeveloperRepository interface {
	// Lookup returns all the public information about a developer
	Lookup(username string) (*Developer, error)

	// Create creates a new developer
	Create(username, firstName, lastName, password string) error

	// Update updates information about a developer
	// Update can update every field except the Username.
	Update(developer *Developer) error

	// Delete deletes a developer.
	Delete(username string) error

	// GetHash returns the hashed password of an user, used for authentication.
	GetHash(username string) (string, error)
}
