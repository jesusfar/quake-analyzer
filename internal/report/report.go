package report

import (
	"github.com/jesusfar/quake-analyzer/internal/match"
)

// MatchesReport defines the operations for generating reports.
type MatchesReport interface {

	// GroupedMatches returns a group of matches.
	GroupedMatches(matches []match.Match) map[string]GroupedMatch

	// GroupedMatchesByDeathCause returns a group of matches by death cause.
	GroupedMatchesByDeathCause(matches []match.Match) map[string]GroupByDeath
}

// Service implement the MatchesReport.
type Service struct {
}

// NewService returns a new instance of Service.
func NewService() *Service {
	return &Service{}
}

// GroupedMatch defines the attributes for printing group of matches.
type GroupedMatch struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

// GroupByDeath defines the attributes for printing matches by death cause.
type GroupByDeath struct {
	KillByMeans map[string]int `json:"kill_by_means"`
}

// GroupedMatches returns a group of matches.
func (r *Service) GroupedMatches(matches []match.Match) map[string]GroupedMatch {
	out := make(map[string]GroupedMatch)
	for _, m := range matches {
		out[m.ID] = GroupedMatch{
			TotalKills: m.TotalKills,
			Players:    toSlicePlayers(m.Players),
			Kills:      m.Kills,
		}
	}
	return out
}

// GroupedMatchesByDeathCause returns a group of matches by death cause.
func (r *Service) GroupedMatchesByDeathCause(matches []match.Match) map[string]GroupByDeath {
	out := make(map[string]GroupByDeath)
	for _, m := range matches {
		out[m.ID] = GroupByDeath{
			KillByMeans: m.KillsByMeans,
		}
	}
	return out
}

func toSlicePlayers(players map[string]bool) []string {
	sPlayers := make([]string, 0, len(players))
	for key := range players {
		sPlayers = append(sPlayers, key)
	}
	return sPlayers
}
