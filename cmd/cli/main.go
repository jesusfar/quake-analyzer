package main

import (
	"flag"
	"github.com/jesusfar/quake-analyzer/internal/parser"
	"github.com/jesusfar/quake-analyzer/internal/reader"
	"github.com/jesusfar/quake-analyzer/internal/report"
	"github.com/jesusfar/quake-analyzer/pkg/util"
	"log"
	"os"
)

func main() {
	// Configure log
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set the log output to the log file
	log.SetOutput(logFile)

	// Read input flag
	reportType := flag.String("report", "grouped-matches", "Specify the report type. [grouped-matches, by-death-cause]")
	flag.Parse()

	// Read the file
	fileReader := reader.NewFileReader()
	lineCh, errCh := fileReader.Read(util.GetFilePath("./assets/qgames.log"))

	// Process log lines
	logParser := parser.NewQuakeLogParser()
	matches, err := logParser.ProcessSync(lineCh, errCh)
	if err != nil {
		return
	}

	// Report out
	reportService := report.NewService()
	switch *reportType {
	case "by-death-cause":
		util.PrettyPrint(reportService.GroupedMatchesByDeathCause(matches))
	default:
		util.PrettyPrint(reportService.GroupedMatches(matches))
	}
}
