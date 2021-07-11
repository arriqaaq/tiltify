package main

import (
	"os"
	"path/filepath"

	tiltcli "github.com/arriqaaq/tiltify/cmd/cli"
	"github.com/knadh/stuffbin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	// Version and date of the build. This is injected at build-time.
	buildVersion = "unknown"
	buildDate    = "unknown"
)

// initLogger initializes logger
func initLogger(verbose bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// Set logger level
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("verbose logging enabled")
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}

// initFileSystem initializes the stuffbin FileSystem to provide
// access to bunded static assets to the app.
func initFileSystem(binPath string) (stuffbin.FileSystem, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fs, err := stuffbin.UnStuff(filepath.Join(exPath, filepath.Base(os.Args[0])))
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func main() {
	// Intialize new CLI app
	app := cli.NewApp()
	app.Name = "tiltify"
	app.Version = buildVersion
	// Register command line args.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose logging",
		},
	}
	var (
		logger = initLogger(true)
	)
	// Initialize the static file system into which all
	// required static assets (.css, .js files etc.) are loaded.
	fs, err := initFileSystem(os.Args[0])
	if err != nil {
		logger.Errorf("error reading stuffed binary: %v", err)
		os.Exit(1)
	}

	tc := tiltcli.NewTiltify(logger, fs, buildVersion)

	// Register commands.
	app.Commands = []cli.Command{
		tc.Init(),
	}
	// Run the app.
	tc.Logger.Info("Starting tiltify...")
	err = app.Run(os.Args)
	if err != nil {
		logger.Errorf("Something terrbily went wrong: %s", err)
	}
}
