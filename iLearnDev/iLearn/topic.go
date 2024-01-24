package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

//topic code will come here

func topicSetup(logger logrus.FieldLogger, ms string, topic string) error {
	skill := os.Getenv("ILEARNSKILL")
	if skill == "" {
		errStr := "env ILEARNSKILL not set"
		logger.Error(errStr)
		return errors.New(errStr)
	}
	logger.Debug("Setting up topic : ", ms, topic, " skill: ", skill)
	if !IsMileStoneValid(ms) {
		return errors.New("Milestone " + ms + " is not valid milestone")
	}

	//copy list of directories from skillDir/.
	skillDir, err := getSkillDir(logger, skill)
	if err != nil {
		return err
	}
	srcDir := ".MyGuru/" + ms + "/" + topic
	srcDir = filepath.Join(skillDir, srcDir)

	tgtDir := "myRepo/" + ms + "/" + topic
	tgtDir = filepath.Join(skillDir, tgtDir)

	logger.Info("topic source Dir: ", srcDir, " target Dir: ", tgtDir)

	err = DuplicateFiles(logger, srcDir, tgtDir)
	return err

}

func topicUpdate(logger logrus.FieldLogger, ms string, topic string) error {
	return topicSetup(logger, ms, topic)
}
