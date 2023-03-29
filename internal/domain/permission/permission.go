package permission

type Permission struct {
	id          string
	name        string
	description string
}

func (p *Permission) ID() string {
	return p.id
}

func (p *Permission) Name() string {
	return p.name
}

func (p *Permission) Description() string {
	return p.description
}
