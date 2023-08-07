package report

import (
	"github.com/jesusfar/quake-analyzer/internal/match"
	"reflect"
	"testing"
)

func TestService_GroupedMatches(t *testing.T) {
	tests := []struct {
		name    string
		matches []match.Match
		want    map[string]GroupedMatch
	}{
		{
			name: "test case 1: should return a group of matches successfully.",
			matches: []match.Match{{
				ID:           "Game_1",
				TotalKills:   6,
				Players:      map[string]bool{"A": true},
				Kills:        map[string]int{"A": 1},
				KillsByMeans: map[string]int{"MOD_RAILGUN": 10},
			}},
			want: map[string]GroupedMatch{
				"Game_1": {
					TotalKills: 6,
					Players:    []string{"A"},
					Kills:      map[string]int{"A": 1},
				},
			},
		},
		{
			name:    "test case 1: should return a empty group of matches successfully.",
			matches: []match.Match{},
			want:    map[string]GroupedMatch{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Service{}
			if got := r.GroupedMatches(tt.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupedMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GroupedMatchesByDeathCause(t *testing.T) {

	tests := []struct {
		name    string
		matches []match.Match
		want    map[string]GroupByDeath
	}{
		{
			name: "test case 1: should return group matches by death cause",
			matches: []match.Match{{
				ID:           "Game_1",
				TotalKills:   6,
				Players:      map[string]bool{"A": true},
				Kills:        map[string]int{"A": 1},
				KillsByMeans: map[string]int{"MOD_RAILGUN": 10},
			}},
			want: map[string]GroupByDeath{
				"Game_1": {
					KillByMeans: map[string]int{"MOD_RAILGUN": 10},
				},
			},
		},
		{
			name:    "test case 1: should return a empty group by death successfully.",
			matches: []match.Match{},
			want:    map[string]GroupByDeath{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Service{}
			if got := r.GroupedMatchesByDeathCause(tt.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupedMatchesByDeathCause() = %v, want %v", got, tt.want)
			}
		})
	}
}
