package devcv

import "time"

// A Developer represents all the information about a developer stored in a database
type Developer struct {
	Username string `json:"username" db:"username"`
	Hash     string `json:"-" db:"pass"`

	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`

	Speciality   string `json:"speciality" db:"speciality"`
	Timezone     string `json:"timezone" db:"timezone"`
	Introduction string `json:"introduction" db:"introduction"`
	City         string `json:"city" db:"city"`
	Languages    string `json:"languages" db:"languages"`

	Links       []Link       `json:"links" db:"-"`
	Projects    []Project    `json:"projects" db:"-"`
	SkillGroups []SkillGroup `json:"skill_groups" db:"-"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// A Project is owned by an user
type Project struct {
	Title string `json:"title" db:"title"`
	Link  string `json:"link" db:"link"`
	Stack string `json:"stack" db:"stack"`
	Scope string `json:"scope" db:"scope"`
}

// A SkillGroup is a group of skills related to each other identified by a title
type SkillGroup struct {
	Title  string  `json:"title" db:"title"`
	Skills []Skill `json:"skills" db:"-"`
}

// A Skill is an individual skill of a developer and it can have an experience level
type Skill struct {
	Title      string `json:"title" db:"title"`
	Experience int    `json:"experience" db:"experience"`
}

// A Link is a link to a developer's social network profile/website/blog/etc
type Link struct {
	Title string `json:"title" db:"title"`
	URL   string `json:"url" db:"url"`
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
