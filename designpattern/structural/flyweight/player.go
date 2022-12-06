package flyweight

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getSingleInstance().getDressByType(dressType)

	return &Player{
		dress:      dress,
		playerType: playerType,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat, p.long = lat, long
}
