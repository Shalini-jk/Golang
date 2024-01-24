package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

//Duplicate all the directories from the srcDir into tgtDir. It is not done recursively
func DuplicateDirs(logger logrus.FieldLogger, srcDir, tgtDir string) error {
	srcInfo, err := os.Stat(srcDir)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if !srcInfo.IsDir() {
		errStr := srcDir + "is not a directory"
		logger.Error(errStr)
		return errors.New(errStr)
	}

	err = os.MkdirAll(tgtDir, srcInfo.Mode())
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			tgtDirPath := filepath.Join(tgtDir, file.Name())
			err = os.MkdirAll(tgtDirPath, srcInfo.Mode())
			if err != nil {
				logger.Error(err.Error())
				return err
			}
		}
	}

	return nil
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

//Duplicate all the files and directories from the srcDir into tgtDir. It is not done recursively
func DuplicateFiles(logger logrus.FieldLogger, srcDir, tgtDir string) error {
	srcInfo, err := os.Stat(srcDir)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if !srcInfo.IsDir() {
		errStr := srcDir + "is not a directory"
		logger.Error(errStr)
		return errors.New(errStr)
	}

	err = os.MkdirAll(tgtDir, srcInfo.Mode())
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			tgtDirPath := filepath.Join(tgtDir, file.Name())
			err = os.MkdirAll(tgtDirPath, srcInfo.Mode())
			if err != nil {
				logger.Error(err.Error())
				return err
			}
		} else {
			srcFile := filepath.Join(srcDir, file.Name())
			tgtFile := filepath.Join(tgtDir, file.Name())
			err = CopyFile(srcFile, tgtFile)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
		}
	}

	return nil
}
