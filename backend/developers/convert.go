package developers

import (
	"github.com/giornetta/devcv/devcv"
	"github.com/giornetta/devcv/proto"
)

// toDeveloper is an helper function which converts a proto.Developer to a devcv.Developer
func toDeveloper(p *proto.Developer) *devcv.Developer {
	links := make([]devcv.Link, len(p.Links))
	for i, l := range p.Links {
		links[i] = devcv.Link{
			Title: l.Title,
			URL:   l.Url,
		}
	}
	projects := make([]devcv.Project, len(p.Projects))
	for i, p := range p.Projects {
		projects[i] = devcv.Project{
			Title: p.Title,
			Link:  p.Link,
			Stack: p.Stack,
			Scope: p.Scope,
		}
	}
	skillgroups := make([]devcv.SkillGroup, len(p.SkillGroups))
	for i, sg := range p.SkillGroups {
		skills := make([]devcv.Skill, len(sg.Skills))
		skillgroups[i] = devcv.SkillGroup{
			Title: sg.Title,
		}
		for j, s := range sg.Skills {
			skills[j] = devcv.Skill{
				Title:      s.Title,
				Experience: int(s.Experience),
			}
		}
		skillgroups[i].Skills = skills
	}

	return &devcv.Developer{
		Username:     p.Username,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Speciality:   p.Speciality,
		Timezone:     p.Timezone,
		Introduction: p.Introduction,
		City:         p.City,
		Languages:    p.Languages,
		Links:        links,
		Projects:     projects,
		SkillGroups:  skillgroups,
	}
}

// toProto is an helper function which converts a devcv.Developer to a proto.Developer
func toProto(d *devcv.Developer) *proto.Developer {
	links := make([]*proto.Link, len(d.Links))
	for i, l := range d.Links {
		links[i] = &proto.Link{
			Title: l.Title,
			Url:   l.URL,
		}
	}
	projects := make([]*proto.Project, len(d.Projects))
	for i, p := range d.Projects {
		projects[i] = &proto.Project{
			Title: p.Title,
			Link:  p.Link,
			Stack: p.Stack,
			Scope: p.Scope,
		}
	}
	skillgroups := make([]*proto.SkillGroup, len(d.SkillGroups))
	for i, sg := range d.SkillGroups {
		skills := make([]*proto.Skill, len(sg.Skills))
		skillgroups[i] = &proto.SkillGroup{
			Title: sg.Title,
		}
		for j, s := range sg.Skills {
			skills[j] = &proto.Skill{
				Title:      s.Title,
				Experience: uint32(s.Experience),
			}
		}
		skillgroups[i].Skills = skills
	}

	return &proto.Developer{
		Username:     d.Username,
		FirstName:    d.FirstName,
		LastName:     d.LastName,
		Speciality:   d.Speciality,
		Timezone:     d.Timezone,
		Introduction: d.Introduction,
		City:         d.City,
		Languages:    d.Languages,
		Links:        links,
		Projects:     projects,
		SkillGroups:  skillgroups,
	}
}
