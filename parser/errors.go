package parser

import "errors"

var (
	ErrInvalidCommitType = errors.New("invalid commit type")
	ErrLongCommit        = errors.New("commit head is longer than 72 symbols allowed")
	ErrEmptyScope        = errors.New("scope is empty")
	ErrEmptySubject      = errors.New("subject is empty")
)
