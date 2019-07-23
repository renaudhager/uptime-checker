# Uptime-checker

[![Docker Pulls](https://img.shields.io/docker/pulls/renaudhager/uptime-checker.svg)](https://hub.docker.com/r/renaudhager/uptime-checker)[![Go Report Card](https://goreportcard.com/badge/github.com/renaudhager/uptime-checker)](https://goreportcard.com/report/github.com/renaudhager/uptime-checker) 

## Description
Small tool, that creates a file if uptime exceed a limit.

## Usage
```
~$ uptime-checker -h
NAME:
   uptime-checker - Check uptime, if its above the limit, creates a file.

USAGE:
   CLI tool to check if uptime is above a specific threshold and touch a file if uptime is above the threshold. [global options] command [command options] [arguments...]

VERSION:
   alpha

COMMANDS:
     check    Check if uptime is above a limit.
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -l value, --log-level value     Setting log level. (default: "INFO") [$UC_LOGLEVEL]
   --lf value, --log-format value  Setting log format. Value [text|json]. (default: "json") [$UC_LOGLEVEL]
   --help, -h                      show help
   --version, -v                   print the version
```

### Check command
```
~$ uptime-checker check -h
NAME:
   CLI tool to check if uptime is above a specific threshold and touch a file if uptime is above the threshold. check - Check if uptime is above a limit.

USAGE:
   CLI tool to check if uptime is above a specific threshold and touch a file if uptime is above the threshold. check [command options] [arguments...]

OPTIONS:
   --uptime-limit value  Limit above a file will be created. In time.Duration format. (default: "24h") [$UC_UPTIME_LIMIT]
   --file value          Path of the file to create. (default: "/var/run/reboot-required") [$UC_FILE_PATH]
   --interval value      Interval between 2 checks. In time.Duration format. (default: "5m") [$UC_INTERVAL]
```
