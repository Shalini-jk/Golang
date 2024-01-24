package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	

	"github.com/sirupsen/logrus"
)

//get the list of skill currently supported
func skillList() []string {
	skills := []string{"GoLang", "Python"}
	return skills
}

//check that given skill is a valid skill
func IsSkillValid(skill string) bool {
	skills := skillList()
	for _, s := range skills {
		if s == skill {
			return true
		}
	}
	return false
}

//returns the map of iLearn Git Repos, it should be coming from the server
func MyGuruILearnGitRepos() map[string]string {
	gitRepos := map[string]string{
		"GoLang": "https://github.com/anurag-kqi/iLearn.git",
		"Python": "https://github.com/anurag-kqi/iLearnPython.git",
	}

	return gitRepos
}

//this is a server side function
func MyGuruUserSetup(skill string, userEmail string, userGitRepo string) (string, error) {
	gitRepos := MyGuruILearnGitRepos()

	return gitRepos[skill], nil
}

func getSkillDir(logger logrus.FieldLogger, skill string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errStr := "failed to get home directory"
		logger.Error(errStr)
		return "", errors.New(errStr)
	}
	dirName := "iLearn/" + skill
	dirName = filepath.Join(homeDir, dirName)
	return dirName, nil
}

//this function should validate that user is a valid in the iLearn portal

func skillSetup(logger logrus.FieldLogger, skill string, userEmail string, gitRepo string) error {

	logger.Info("setup skill: ", skill)
	if !IsSkillValid(skill) {
		errStr := "Skill " + skill + " is not valid skill"
		logger.Error(errStr)
		return errors.New(errStr)
	}
	logger.Info("setup userEmail: ", userEmail)
	if !IsEmailValid(userEmail) {
		errStr := "Email " + userEmail + " is not valid email"
		logger.Error(errStr)
		return errors.New(errStr)
	}

	logger.Info("setup git Repo: ", gitRepo)
	if !IsURLValid(gitRepo) {
		errStr := "girRepo " + gitRepo + " is not valid URL"
		logger.Error(errStr)
		return errors.New(errStr)
	}

	// Validate the user is valid, and has been configured for the given skill
	// ideally, user should also pass a token given to the user from MyGuru portal, AA: TBD
	// This call should return the iLearn git repo for the skill
	iLearnGitRepo, err := MyGuruUserSetup(skill, userEmail, gitRepo)

	if err != nil {
		errStr := "MyGuruUserSet failed: " + err.Error()
		logger.Error(errStr)
		return errors.New(errStr)
	} else {

		logger.Info("iLearn Git Repo is: ", iLearnGitRepo)

		//create the following directory structure and checkout the git repo
		// iLearn
		// 	- SkillName
		// 		- .MyGuru
		// 		 	-iLearnGitRepo
		// 		- gitRepo

		dirName, err := getSkillDir(logger, skill)
		if err != nil {
			return err
		}
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			errStr := "failed to create directory: " + dirName + err.Error()
			logger.Error(errStr)
			return errors.New(errStr)
		}

		//cloning iLearn Git repo as myGuru repo in the hidden directory
		cmd := exec.Command("git", "clone", iLearnGitRepo, ".MyGuru")
		cmd.Dir = dirName

		output, err := cmd.CombinedOutput()
		if err != nil {
			logger.Error("git clone of repo: ", iLearnGitRepo, "Failed, err: ", err)
			logger.Error("cmd output: ", string(output))
			return errors.New("Failed to clone " + iLearnGitRepo)
		}
		logger.Info("git clone ", gitRepo, "successful ", string(output))

		cmd = exec.Command("git", "clone", gitRepo, "myRepo")
		cmd.Dir = dirName

		output, err = cmd.CombinedOutput()
		if err != nil {
			logger.Error("git clone of repo: ", gitRepo, "Failed, err: ", err)
			logger.Error("cmd output: ", string(output))
			return errors.New("Failed to clone " + gitRepo)
		}
		logger.Info("git clone ", gitRepo, "successful ", string(output))
	}

	return nil
}

// IsEmailValid checks if the given string is a valid email address.
// AA: Move this to helper library
func IsEmailValid(email string) bool {
	// Regular expression for a basic email validation
	// Note: This is a simplified version and may not cover all valid email formats
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		// Handle regex compilation error
		fmt.Println("Error:", err)
		return false
	}

	return match
}

func IsURLValid(inputURL string) bool {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		// Parsing error indicates that the URL is not valid
		return false
	}

	// Check if the parsed URL has a scheme (http, https, etc.) and a host
	return parsedURL.Scheme != "" && parsedURL.Host != ""
}
