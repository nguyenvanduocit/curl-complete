package main

import (
	"github.com/posener/complete"
)

func main() {
	run := complete.Command{
		Flags: complete.Flags{
			"--user-agent":  complete.PredictNothing,
			"--cookie":      complete.PredictNothing,
			"--cookie-jar":  complete.PredictFiles("*"),
			"--output":      complete.PredictFiles("*"),
			"--trace-ascii": complete.PredictFiles("*"),
			"--upload-file": complete.PredictFiles("*"),
			"--compressed":  complete.PredictNothing,
			"--insecure":    complete.PredictNothing,
			"--remote-name": complete.PredictNothing,
			"--head":        complete.PredictNothing,
			"--header":      complete.PredictNothing,
			"--location":    complete.PredictNothing,
			"--verbose":     complete.PredictNothing,
			"--write-out":   complete.PredictNothing,
			"--proxy":       complete.PredictNothing,
			"--user":        complete.PredictSet("user:password"),
			"--silent":      complete.PredictNothing,
			"--limit-rate":  complete.PredictSet("200K", "3m", "1G"),
			"--max-time":    complete.PredictSet("60", "120", "3600"),
			"--data":        complete.PredictFiles("*"),
			"--form":        complete.PredictFiles("*"),
		},
	}
	complete.New("curl", run).Run()
}
