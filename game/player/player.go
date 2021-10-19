package player

type Player interface {
	CurrentPosition() int
	MovesPlayed() int
	NewPosition(position int)
	GetName() string
}

func NewPlayer(name string) Player {
	return &player{Name: name}
}

type player struct {
	Position int
	Moves    int
	Name     string
}

func (p *player) GetName() string {
	return p.Name
}

func (p *player) CurrentPosition() int {
	return p.Position
}

func (p *player) MovesPlayed() int {
	return p.Moves
}

func (p *player) NewPosition(position int) {
	p.Position = position
	p.Moves++
}
