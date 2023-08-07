package match

import "fmt"

// GameMatch defines the basic operations fpr a match.
type GameMatch interface {

	// AddPlayer adds a new player to the match
	AddPlayer(player string)

	// CountDeathCause counts the death cause
	CountDeathCause(deathCause string)

	// ScoreKill scores the kill count for the killer player
	ScoreKill(kill Kill)
}

// Match defines the attributes for the match.
type Match struct {
	ID           string
	TotalKills   int
	Players      map[string]bool
	Kills        map[string]int
	KillsByMeans map[string]int
}

// Kill defines the attributes for kill action.
type Kill struct {
	Killer     string
	Victim     string
	DeathCause string
}

// NewMatch creates a new instance of Match struct with its identifier.
func NewMatch(matchNumber int) *Match {
	return &Match{
		ID:           fmt.Sprintf("Game_%d", matchNumber),
		Players:      make(map[string]bool),
		Kills:        make(map[string]int),
		KillsByMeans: make(map[string]int),
	}
}

// AddPlayer adds a new player to the match
func (m *Match) AddPlayer(player string) {
	if isWorldPlayer(player) {
		return
	}
	m.Players[player] = true
}

// CountDeathCause counts the death cause
func (m *Match) CountDeathCause(deathCause string) {
	m.KillsByMeans[deathCause]++
}

// ScoreKill scores the kill count for the killer player
func (m *Match) ScoreKill(kill Kill) {
	if isWorldPlayer(kill.Killer) {
		m.Kills[kill.Victim]--
		return
	}

	m.Kills[kill.Killer]++
}

func isWorldPlayer(player string) bool {
	return player == "<world>"
}
