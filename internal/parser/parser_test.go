package parser

import (
	"github.com/jesusfar/quake-analyzer/internal/match"
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name      string
		inputLine string
		want      quakeLogLine
	}{
		{
			name:      "Test case 1: parse a kill log",
			inputLine: " 21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			want: quakeLogLine{
				TimeStamp:  "21:42",
				LogType:    "Kill",
				LogMessage: " 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			},
		},
		{
			name:      "Test case 2: parse a Item log",
			inputLine: " 981:35 Item: 4 weapon_rocketlauncher",
			want: quakeLogLine{
				TimeStamp:  "981:35",
				LogType:    "Item",
				LogMessage: " 4 weapon_rocketlauncher",
			},
		},
		{
			name:      "Test case 3: parse a Init Game",
			inputLine: ` 0:00 InitGame: \capturelimit\8\g_maxGameClients\0\timelimit\15\fraglimit\20\dmflags\0\bot_minplayers\0\sv_allowDownload\0\sv_maxclients\16\sv_privateClients\2\g_gametype\4\sv_hostname\Code Miner Server\sv_minRate\0\sv_maxRate\10000\sv_minPing\0\sv_maxPing\0\sv_floodProtect\1\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`,
			want: quakeLogLine{
				TimeStamp:  "0:00",
				LogType:    "InitGame",
				LogMessage: ` \capturelimit\8\g_maxGameClients\0\timelimit\15\fraglimit\20\dmflags\0\bot_minplayers\0\sv_allowDownload\0\sv_maxclients\16\sv_privateClients\2\g_gametype\4\sv_hostname\Code Miner Server\sv_minRate\0\sv_maxRate\10000\sv_minPing\0\sv_maxPing\0\sv_floodProtect\1\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`,
			},
		},
		{
			name:      "Test case 4: parse a ShutdownGame",
			inputLine: ` 14:11 ShutdownGame:`,
			want: quakeLogLine{
				TimeStamp:  "14:11",
				LogType:    "ShutdownGame",
				LogMessage: "",
			},
		},
		{
			name:      "Test case 5: parse separator should return empty log line",
			inputLine: " 54:21 ------------------------------------------------------------",
			want:      quakeLogLine{},
		},
		{
			name:      "Test case 6: parse empty line should return empty log line",
			inputLine: " 54:21 ------------------------------------------------------------",
			want:      quakeLogLine{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLine(tt.inputLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseKillLog(t *testing.T) {
	tests := []struct {
		name string
		text string
		want match.Kill
	}{
		{
			name: "test case1: parse world killer",
			text: " 1022 2 19: <world> killed Dono da Bola by MOD_FALLING",
			want: match.Kill{Killer: "<world>", Victim: "Dono da Bola", DeathCause: "MOD_FALLING"},
		},
		{
			name: "test case2: parse Dono da Bolas as killer",
			text: " 2 4 6: Dono da Bola killed Zeh by MOD_ROCKET",
			want: match.Kill{Killer: "Dono da Bola", Victim: "Zeh", DeathCause: "MOD_ROCKET"},
		},
		{
			name: "test case3: parse Isgalamido as killer",
			text: " 3 2 6: Isgalamido killed Dono da Bola by MOD_ROCKET",
			want: match.Kill{Killer: "Isgalamido", Victim: "Dono da Bola", DeathCause: "MOD_ROCKET"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseKillLog(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKillLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
