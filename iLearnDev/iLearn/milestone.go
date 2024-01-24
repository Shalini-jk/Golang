package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

//get the list of skill currently supported
func mileStoneList() []string {
	msList := []string{"BasicTopics", "CoreTopics", "AdvancedTopics"}
	return msList
}

//check that given milestone is a valid milestone
func IsMileStoneValid(mileStone string) bool {
	msList := mileStoneList()
	for _, ms := range msList {
		if ms == mileStone {
			return true
		}
	}
	return false
}

//Copy all the topics from the .MyGuru to myRepo for specified milestone
//don't copy all the kbits and problems for all the topics, that will be done duing setting up the topic

func mileStoneSetup(logger logrus.FieldLogger, mileStone string) error {
	//os.Setenv("ILEARNSKILL","GoLang")
	skill := os.Getenv("ILEARNSKILL")
	if skill == "" {
		errStr := "env ILEARNSKILL not set"
		logger.Error(errStr)
		return errors.New(errStr)
	}
	logger.Debug("Setting up milestone : ", mileStone, " skill: ", skill)
	if !IsMileStoneValid(mileStone) {
		return errors.New("Milestone " + mileStone + " is not valid milestone")
	}

	//copy list of directories from skillDir/.
	skillDir, err := getSkillDir(logger, skill)
	if err != nil {
		return err
	}
	srcDir := ".MyGuru/" + mileStone
	srcDir = filepath.Join(skillDir, srcDir)

	tgtDir := "myRepo/" + mileStone
	tgtDir = filepath.Join(skillDir, tgtDir)

	logger.Info("milestone source Dir: ", srcDir, " target Dir: ", tgtDir)

	err = os.MkdirAll(tgtDir, 0755)
	if err != nil {
		errStr := "failed to create directory: " + tgtDir + err.Error()
		logger.Error(errStr)
		return errors.New(errStr)
	}
	err = DuplicateDirs(logger, srcDir, tgtDir)
	return err
}

//For now, there is no difference in setting up and updating the milestone
//update creates all the directories again, that brings any newly created directory to the target
func mileStoneUpdate(logger logrus.FieldLogger, mileStone string) error {
	fmt.Println("Updating milestone : ", mileStone)
	if !IsMileStoneValid(mileStone) {
		return errors.New("Milestone " + mileStone + " is not valid milestone")
	}
	return mileStoneSetup(logger, mileStone)
}
