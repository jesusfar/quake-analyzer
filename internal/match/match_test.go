package match

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatch(t *testing.T) {
	newMatch := NewMatch(1)
	assert.Equal(t, "Game_1", newMatch.ID)
	assert.Equal(t, 0, newMatch.TotalKills)
	assert.Equal(t, make(map[string]bool), newMatch.Players)
	assert.Equal(t, make(map[string]int), newMatch.Kills)
	assert.Equal(t, make(map[string]int), newMatch.KillsByMeans)
}

func TestMatch_AddPlayer(t *testing.T) {
	tests := []struct {
		name          string
		player        string
		expectPlayers map[string]bool
	}{
		{
			name:          "test case 1: add new player should add the record in the set",
			expectPlayers: map[string]bool{"Zen": true},
			player:        "Zen",
		},
		{
			name:          "test case 2: add player <world> should not be consider",
			expectPlayers: map[string]bool{},
			player:        "<world>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Match{
				Players: tt.expectPlayers,
			}
			m.AddPlayer(tt.player)
		})
	}
}

func TestMatch_CountDeathCause(t *testing.T) {
	newMatch := NewMatch(1)
	for i := 0; i < 10; i++ {
		newMatch.CountDeathCause("MOD_RAILGUN")
	}
	assert.Equal(t, 10, newMatch.KillsByMeans["MOD_RAILGUN"])
}

func TestMatch_ScoreKill(t *testing.T) {
	tests := []struct {
		name        string
		kill        Kill
		expectKills map[string]int
	}{
		{
			name: "test case 1: score kill normal player",
			kill: Kill{
				Killer:     "A",
				Victim:     "B",
				DeathCause: "MOD_RAILGUN",
			},
			expectKills: map[string]int{"A": 1},
		},
		{
			name: "test case 2: when player <world> is killer should decrease score from victim player",
			kill: Kill{
				Killer:     "<world>",
				Victim:     "B",
				DeathCause: "MOD_RAILGUN",
			},
			expectKills: map[string]int{"B": -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Match{
				Kills: tt.expectKills,
			}
			m.ScoreKill(tt.kill)
		})
	}
}
