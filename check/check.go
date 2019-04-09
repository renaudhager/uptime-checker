package check

import (
	"io/ioutil"
	"time"

	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
	"github.com/capnm/sysinfo"
)

// CheckCmd start an infinite loop and check uptime every interval.
var CheckCmd = cli.Command{
	Name:   "check",
	Usage:  "Check if uptime is above a limit.",
	Action: check,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "uptime-limit",
			Usage:  "Limit above a file will be created. In time.Duration format.",
			EnvVar: "UC_UPTIME_LIMIT",
			Value: "24h",
		},
		cli.StringFlag{
			Name:   "file",
			Usage:  "Path of the file to create.",
			EnvVar: "UC_FILE_PATH",
			Value: "/var/run/reboot-required",
		},
		cli.StringFlag{
			Name:   "interval",
			Usage:  "Interval between 2 checks. In time.Duration format.",
			EnvVar: "UC_INTERVAL",
			Value: "5m",
		},
	},
}

func check(c *cli.Context) {

	log := confLogging(c.GlobalString("l"),c.GlobalString("lf"))

	intervalDuration, err := time.ParseDuration(c.String("interval"))
	if err != nil {
		log.Error("Error parsing interval.", err)
	}

	for {
		uptimeDuration := sysinfo.Get().Uptime

		uptimeLimitDuration, err := time.ParseDuration(c.String("uptime-limit"))
		if err != nil {
			log.Error("Error parsing uptime-limit.", err)
		}

		if uptimeDuration > uptimeLimitDuration {
			log.Info("Uptime is above the limit.")
			touchSentinelFile(c.String("file"), log)
		}

		time.Sleep(intervalDuration)
	}

}

func touchSentinelFile(filepath string, log *log.Logger) {

	err := ioutil.WriteFile(filepath, []byte(""), 0600)
	if err != nil {
		log.Error("Error when writing to ", filepath)
	} else {
    log.Info("file created: ",filepath)
  }

}

func confLogging(level string, format string) *log.Logger {

	var logger = log.New()

	switch format {
	case "json":
		logger.SetFormatter(&log.JSONFormatter{})
	case "text":
		logger.SetFormatter(&log.TextFormatter{})
	}

	switch level {
	case "INFO":
		logger.SetLevel(log.InfoLevel)
		logger.Info("Log level is ", level)
	case "DEBUG":
		logger.SetLevel(log.DebugLevel)
		logger.Debug("Log level is ", level)
	case "WARN":
		logger.SetLevel(log.WarnLevel)
		logger.Warn("Log level is ", level)
	case "ERROR":
		logger.SetLevel(log.ErrorLevel)
		logger.Error("Log level is ", level)
	}

	return logger
}
