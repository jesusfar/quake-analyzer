package parser

import (
	"github.com/jesusfar/quake-analyzer/internal/match"
	"log"
	"regexp"
)

type quakeLogLine struct {
	TimeStamp  string
	LogType    string
	LogMessage string
}

// QuakeParser interface defines the operation for process log lines.
type QuakeParser interface {
	// ProcessSync receives the log lines from line channel, parses each line and returns a slice of matches with the analyzed matches.
	ProcessSync(lineCh <-chan string, errCh <-chan error) ([]match.Match, error)
}

// QuakeLogParser implements the QuakeParser.
type QuakeLogParser struct {
}

// NewQuakeLogParser returns a instance of QuakeLogParser.
func NewQuakeLogParser() *QuakeLogParser {
	return &QuakeLogParser{}
}

// ProcessSync receives the log lines from line channel, parses each line and returns a slice of matches with the analyzed matches.
func (q QuakeLogParser) ProcessSync(lineCh <-chan string, errCh <-chan error) ([]match.Match, error) {
	var newMatch *match.Match
	var matches = make([]match.Match, 0)
	var matchCount = 0
	for {
		select {
		case line, ok := <-lineCh:
			if !ok {
				log.Println("file reader channel closed, finishing parser process")
				return matches, nil
			}
			quakeLogLine := parseLine(line)
			if quakeLogLine.LogType == "InitGame" {
				if matchCount > 0 {
					matches = append(matches, *newMatch)
				}
				matchCount++
				newMatch = match.NewMatch(matchCount)
			}
			if quakeLogLine.LogType == "Kill" {
				kill := parseKillLog(quakeLogLine.LogMessage)
				processKill(kill, newMatch)
			}
		case err := <-errCh:
			if err != nil {
				log.Println("error received from file reader: ", err)
				return matches, err
			}
		}
	}

}

func processKill(kill match.Kill, newMatch *match.Match) {
	newMatch.TotalKills++
	newMatch.AddPlayer(kill.Killer)
	newMatch.AddPlayer(kill.Victim)
	newMatch.CountDeathCause(kill.DeathCause)
	newMatch.ScoreKill(kill)
}

func parseKillLog(text string) match.Kill {
	re := regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s+\d+:\s+(.+?)\s+killed\s+(.+?)\s+by\s+(\w+)$`)
	if matches := re.FindStringSubmatch(text); matches != nil {
		return match.Kill{
			Killer:     matches[3],
			Victim:     matches[4],
			DeathCause: matches[5],
		}
	}
	return match.Kill{}
}

func parseLine(line string) quakeLogLine {
	reg := regexp.MustCompile(`^\s+(\S+)\s+(\S+):+(.*)$`)
	if matches := reg.FindStringSubmatch(line); matches != nil {
		return quakeLogLine{
			TimeStamp:  matches[1],
			LogType:    matches[2],
			LogMessage: matches[3],
		}
	}

	return quakeLogLine{}
}
