package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func validate(config Config) error {
	var errs []error

	if isEmptyString(config.GitAuthorEmail) {
		errs = append(errs, fmt.Errorf("No git author email"))
	}

	if isEmptyString(config.GitAuthorName) {
		errs = append(errs, fmt.Errorf("No git author name"))
	}

	if err := validateDirectory(config.RepoPath, "repo path"); err != nil {
		errs = append(errs, err)
	}

	if err := validateDirectory(config.SSHKeyPath, "SSH key path"); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func isEmptyString(param string) bool {
	return strings.TrimSpace(param) == ""
}

func fileOrDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func validateDirectory(path string, keyword string) error {
	if isEmptyString(path) {
		return fmt.Errorf("No %s", keyword)
	}

	if exists, err := fileOrDirectoryExists(path); err != nil || exists == false {
		if err != nil {
			return err
		} else {
			return fmt.Errorf("Not exist: %s", keyword)
		}
	}

	return nil
}
