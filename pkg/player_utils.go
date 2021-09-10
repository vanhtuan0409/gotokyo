package pkg

func FindPlayer(id uint, players []*PlayerInfo) *PlayerInfo {
	for _, p := range players {
		if p.Id == id {
			return p
		}
	}
	return nil
}
