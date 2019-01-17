package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// NewDB connects to PostgreSQL, create tables if they don't exist and returns a connection that can be injected into Repositories
func NewDB(host string, port int, dbname, user, password string) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	q := `
		CREATE TABLE IF NOT EXISTS developers (
			username TEXT PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			
			pass TEXT NOT NULL,
		
			speciality TEXT DEFAULT '',
			timezone TEXT DEFAULT '',
			introduction TEXT DEFAULT '',
			city TEXT DEFAULT '',
			languages TEXT DEFAULT '',
		
			created_at TIMESTAMP DEFAULT now(),
			updated_at TIMESTAMP DEFAULT now()
		);
		
		CREATE TABLE IF NOT EXISTS projects (
			project_id SERIAL PRIMARY KEY,
			developer_username TEXT NOT NULL REFERENCES developers(username) ON DELETE CASCADE,
			title TEXT NOT NULL,
			link TEXT DEFAULT '',
			stack TEXT DEFAULT '',
			scope TEXT DEFAULT '',
		
			project_position INT NOT NULL,
		
			created_at TIMESTAMP DEFAULT now(),
			updated_at TIMESTAMP DEFAULT now(),
		
			CONSTRAINT proj_pos UNIQUE(developer_username, project_position)
		);
		
		CREATE TABLE IF NOT EXISTS skillgroups (
			skillgroup_id SERIAL PRIMARY KEY,
			developer_username TEXT NOT NULL REFERENCES developers(username) ON DELETE CASCADE,
			title TEXT NOT NULL,
			skillgroup_position INT NOT NULL,
		
			CONSTRAINT group_pos UNIQUE(developer_username, skillgroup_position)
		);
		
		CREATE TABLE IF NOT EXISTS skills (
			skill_id SERIAL PRIMARY KEY,
			skillgroup_id INTEGER NOT NULL REFERENCES skillgroups(skillgroup_id) ON DELETE CASCADE,
			title TEXT NOT NULL,
			experience SMALLINT,
			skill_position INT NOT NULL,
		
			CONSTRAINT skill_pos UNIQUE(skillgroup_id, skill_position)
		);
		
		CREATE TABLE IF NOT EXISTS links (
			link_id SERIAL PRIMARY KEY,
			developer_username TEXT NOT NULL REFERENCES developers(username) ON DELETE CASCADE,
			title TEXT NOT NULL,
			url TEXT NOT NULL,
			link_position INT NOT NULL,
		
			CONSTRAINT link_pos UNIQUE(developer_username, link_position)
		);	
	`

	_, err = db.Exec(q)
	if err != nil {
		return nil, err
	}

	return db, nil
}
