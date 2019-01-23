package repository

import (
	"github.com/giornetta/devcv/devcv"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

// NewDevelopers returns a PostgreSQL implementation of devcv.DeveloperRepository
func NewDevelopers(db *sqlx.DB) devcv.DeveloperRepository {
	return &repository{db}
}

// Create implements devcv.DeveloperRepository
func (r *repository) Create(username, firstName, lastName, password string) error {
	q := `
		insert into developers(username, first_name, last_name, pass)
		values($1, $2, $3, $4);
	`

	_, err := r.db.Exec(q, username, firstName, lastName, password)
	return err
}

// Lookup implements devcv.DeveloperRepository
func (r *repository) Lookup(username string) (*devcv.Developer, error) {
	var dev devcv.Developer

	// Get Developer
	q := `
		select username, first_name, last_name, speciality, timezone, introduction, city, languages, created_at, updated_at
		from developers
		where username = $1;
	`

	if err := r.db.Get(&dev, q, username); err != nil {
		return nil, err
	}

	// Get Skills
	q = `
		select sg.title, s.title, s.experience
		from skillgroups as sg
		join skills as s on sg.skillgroup_id = s.skillgroup_id
		where sg.developer_username = $1
		order by sg.skillgroup_position, s.skill_position;
	`

	rows, err := r.db.Query(q, username)
	if err != nil {
		return &dev, err
	}

	var sg devcv.SkillGroup
	for rows.Next() {
		var sgTitle string
		var s devcv.Skill
		rows.Scan(&sgTitle, &s.Title, &s.Experience)

		// if it's a new skillgroup, but not the first
		if sgTitle != sg.Title && sg.Title != "" {
			// append it to skillgroups
			dev.SkillGroups = append(dev.SkillGroups, sg)
			sg = devcv.SkillGroup{}
		}

		sg.Title = sgTitle
		sg.Skills = append(sg.Skills, s)
	}
	dev.SkillGroups = append(dev.SkillGroups, sg)

	// Get Links
	q = `
		select title, url
		from links
		where developer_username = $1
		order by link_position;
	`

	if err := r.db.Select(&dev.Links, q, username); err != nil {
		return &dev, err
	}

	// Get Projects
	q = `
		select title, link, stack, scope
		from projects
		where developer_username = $1
		order by project_position;
	`

	if err := r.db.Select(&dev.Projects, q, username); err != nil {
		return &dev, err
	}

	return &dev, nil
}

// Update implements devcv.DeveloperRepository
func (r *repository) Update(d *devcv.Developer) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Update Developer
	q := `
		update developers
		set 
			first_name = $2,
			last_name = $3,
			speciality = $4,
			timezone = $5,
			introduction = $6,
			city = $7,
			languages = $8,
			updated_at = now()
		where username = $1;
	`

	_, err = tx.Exec(q, d.Username, d.FirstName, d.LastName, d.Speciality, d.Timezone, d.Introduction, d.City, d.Languages)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update Developer's Links
	q = `
		delete from links
		where developer_username = $1 and link_position > $2;
	`

	_, err = tx.Exec(q, d.Username, len(d.Links))
	if err != nil {
		tx.Rollback()
		return err
	}

	q = `
		insert into links(developer_username, title, url, link_position)
		values ($1, $2, $3, $4)
		on conflict on constraint link_pos
		do
			update
			set
				title = EXCLUDED.title,
				url = EXCLUDED.url;
	`

	for i, l := range d.Links {
		// Upsert each link in the list
		_, err = tx.Exec(q, d.Username, l.Title, l.URL, i+1)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update Skills
	q = `
		delete from skillgroups
		where developer_username = $1 and skillgroup_position > $2
	`

	_, err = tx.Exec(q, d.Username, len(d.SkillGroups))
	if err != nil {
		tx.Rollback()
		return err
	}

	sgq := `
		insert into skillgroups(developer_username, title, skillgroup_position)
		values ($1, $2, $3)
		on conflict on constraint group_pos
		do
			update
			set
				title = EXCLUDED.title
		returning skillgroup_id;
	`

	sq := `
		insert into skills(skillgroup_id, title, experience, skill_position)
		values ($1, $2, $3, $4)
		on conflict on constraint skill_pos
		do
			update
			set
				title = EXCLUDED.title,
				experience = EXCLUDED.experience;
	`

	for i, sg := range d.SkillGroups {
		var sgID uint
		if err := tx.QueryRow(sgq, d.Username, sg.Title, i+1).Scan(&sgID); err != nil {
			tx.Rollback()
			return err
		}

		q = `
			delete from skills
			where skillgroup_id = $1 and skill_position > $2
		`

		_, err = tx.Exec(q, sgID, len(sg.Skills))
		if err != nil {
			tx.Rollback()
			return err
		}

		for j, s := range sg.Skills {
			_, err := tx.Exec(sq, sgID, s.Title, s.Experience, j+1)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

// Delete implements devcv.DeveloperRepository
func (r *repository) Delete(username string) error {
	q := `
		delete from developers
		where username = $1;
	`

	_, err := r.db.Exec(q, username)
	return err
}

// GetHash implements devcv.DeveloperRepository
func (r *repository) GetHash(username string) (string, error) {
	q := `
		select pass
		from developers
		where username = $1
	`

	var pass string
	err := r.db.Get(&pass, q, username)

	return pass, err
}
