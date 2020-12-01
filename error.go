package golinkedmin

import (
	"encoding/json"
	"errors"
)

const csrfFailed = "CSRF check failed."
const profileCantBeAccessed = "This profile can't be accessed"

// Error is error message from Linkedin
type Error struct {
	ExceptionClass string `json:"exceptionClass"`
	Message        string `json:"message"`
	Status         int    `json:"status"`
}

func (err Error) Error() string {
	return err.Message
}

// CSRFFailed return CSRF check error, usually this occur because you have invalid cookie
type CSRFFailed string

func (csrfErr CSRFFailed) Error() string {
	return csrfFailed
}

// ProfileCantBeAccessed occured when profile is not exist or have limited privacy
type ProfileCantBeAccessed string

func (profCantBeAcc ProfileCantBeAccessed) Error() string {
	return profileCantBeAccessed
}

// if error not parseable to Error nor CSRF error, it's probably invalid/expired cookie or Linkedin internal error
func parseErrMsg(msg string) error {
	err := new(Error)
	if er := json.Unmarshal([]byte(msg), err); er == nil {
		if err.Message == profileCantBeAccessed {
			profCantBeAcc := ProfileCantBeAccessed(err.Message)
			return profCantBeAcc
		}

		return *err
	}

	if msg == csrfFailed {
		var csrfErr CSRFFailed
		csrfErr = csrfFailed
		return csrfErr
	}

	return errors.New(msg)
}
