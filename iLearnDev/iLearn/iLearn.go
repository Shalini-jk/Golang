package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func iLearnUsage() {
	fmt.Println("USage: iLearn <command> <arguments> \n commands are: \n skill list \n skill setup -s skill -u useremail@abc.com -l https://github.com/userid/golang topic")

}

func main() {

	logger := logrus.New()
	// Uncomment following if want to print filename, line number in the log
	// logger.SetReportCaller(true)

	flag.NewFlagSet("iLearn", flag.ExitOnError)

	// Parse top-level command
	if len(os.Args) < 2 {
		iLearnUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "skill":
		if len(os.Args) < 3 {
			iLearnUsage()
			os.Exit(1)
		}

		subCommand := os.Args[2]
		switch subCommand {
		case "list":
			for _, skill := range skillList() {
				fmt.Println(skill)
			}

		case "setup":
			setupCmd := flag.NewFlagSet("skill setup", flag.ExitOnError)
			skill := setupCmd.String("s", "GoLang", "Name of skill")
			userEmail := setupCmd.String("u", "anurag@b.com", "User Email")
			gitRepo := setupCmd.String("l", "https://github.com/anurag-kqi/repro1", "User Git Repo")

			setupCmd.Parse(os.Args[3:])

			err := skillSetup(logger, *skill, *userEmail, *gitRepo)
			if err != nil {
				logger.Error(err.Error())
			}
		}

	case "milestone":
		if len(os.Args) < 3 {
			iLearnUsage()
			os.Exit(1)
		}

		subCommand := os.Args[2]
		switch subCommand {
		case "list":
			for _, ms := range mileStoneList() {
				fmt.Println(ms)
			}

		case "setup":
			setupCmd := flag.NewFlagSet("milestone setup", flag.ExitOnError)
			ms := setupCmd.String("m", "BasicTopics", "Name of milestone")

			setupCmd.Parse(os.Args[3:])

			err := mileStoneSetup(logger, *ms)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}
		case "update":
			setupCmd := flag.NewFlagSet("milestone update", flag.ExitOnError)
			up := setupCmd.String("m", "BasicTopics", "Name of milestone")

			setupCmd.Parse(os.Args[3:])

			err := mileStoneUpdate(logger, *up)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}

		}

	case "topic":

		if len(os.Args) < 3 {
			iLearnUsage()
			os.Exit(1)
		}

		subCommand := os.Args[2]
		switch subCommand {

		case "setup":
			setupCmd := flag.NewFlagSet("topic setup", flag.ExitOnError)
			ms := setupCmd.String("m", "BasicTopics", "Name of milestone")
			topic := setupCmd.String("t", "", "Name of topic")

			setupCmd.Parse(os.Args[3:])

			err := topicSetup(logger, *ms, *topic)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}

		case "update":
			setupCmd := flag.NewFlagSet("topic update", flag.ExitOnError)
			ms := setupCmd.String("m", "BasicTopics", "Name of milestone")
			topic := setupCmd.String("t", "", "Name of topic")

			setupCmd.Parse(os.Args[3:])

			err := topicUpdate(logger, *ms, *topic)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}

		}
	}

}
