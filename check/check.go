package check

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/capnm/sysinfo"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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
			Value:  "24h",
		},
		cli.StringFlag{
			Name:   "file",
			Usage:  "Path of the file to create.",
			EnvVar: "UC_FILE_PATH",
			Value:  "/var/run/reboot-required",
		},
		cli.StringFlag{
			Name:   "interval",
			Usage:  "Interval between 2 checks. In time.Duration format.",
			EnvVar: "UC_INTERVAL",
			Value:  "5m",
		},
		cli.StringFlag{
			Name:   "window-start-time",
			Usage:  "Window start hour and minute.",
			EnvVar: "UC_WINDOW_START",
			Value:  "00:00",
		},
		cli.StringFlag{
			Name:   "window-end-time",
			Usage:  "Window end hour and minute.",
			EnvVar: "UC_WINDOW_END",
			Value:  "23:59",
		},
	},
}

func check(c *cli.Context) {

	log := confLogging(c.GlobalString("l"), c.GlobalString("lf"))

	intervalDuration, err := time.ParseDuration(c.String("interval"))
	if err != nil {
		log.Error("Error parsing interval.", err)
	}

	for {
		newLayout := "15:04"
		start, _ := time.Parse(newLayout, c.String("window-start-time"))
		end, _ := time.Parse(newLayout, c.String("window-end-time"))
		currentHour, currentMinute, _ := time.Now().Clock()
		currentTS := strconv.Itoa(currentHour) + ":" + strconv.Itoa(currentMinute)
		log.Debug("current timestamp: ", currentTS)
		check, _ := time.Parse(newLayout, currentTS)

		if inTimeWindow(start, end, check) {
			log.Info("Current time is included in the time window, checking uptime...")
			uptimeDuration := sysinfo.Get().Uptime

			uptimeLimitDuration, err := time.ParseDuration(c.String("uptime-limit"))
			if err != nil {
				log.Error("Error parsing uptime-limit.", err)
			}

			if uptimeDuration > uptimeLimitDuration {
				log.Info("Uptime is above the limit.")
				touchSentinelFile(c.String("file"), log)
			}
		} else {
			log.Info("Current time is NOT included in the time window, doing nothing...")
		}

		time.Sleep(intervalDuration)
	}

}

func touchSentinelFile(filepath string, log *log.Logger) {
	err := ioutil.WriteFile(filepath, []byte(""), 0600)
	if err != nil {
		log.Error("Error when writing to ", filepath)
	} else {
		log.Info("file created: ", filepath)
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

func inTimeWindow(start, end, check time.Time) bool {
	_end := end
	_check := check
	log.Debug("start: ", start)
	log.Debug("end: ", end)
	log.Debug("check: ", check)
	if end.Before(start) {
		_end = end.Add(24 * time.Hour)
		if check.Before(start) {
			_check = check.Add(24 * time.Hour)
		}
	}
	return _check.After(start) && _check.Before(_end)
}
