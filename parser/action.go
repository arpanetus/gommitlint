package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func parse(data []byte) (*Commit, error) {
	var (
		fl, body, rest string
		ct             CommitType
	)
	fl, body = sepFirstLineAndBody(data)

	if len(fl) > 72 {
		return nil, ErrLongCommit
	}

	ct, rest = stripCommitType(fl)

	if !IsValidCommitType(ct, CommitTypes) {
		return nil, fmt.Errorf("commit type has to be one of: %v", CommitTypes)
	}

	scope, subject, err := ScopeNSubject(rest)
	if err != nil {
		return nil, err
	}

	subject, err = CleanedSubject(subject)
	if err != nil {
		return nil, err
	}

	return &Commit{
		Ct:      ct,
		Scope:   scope,
		Subject: subject,
		Body:    body,
	}, nil
}

func Parse(r io.Reader) {
	commit, err := ioutil.ReadAll(r)
	if err != nil {
		ReportFatal("can't read given data: %v", err)
	}

	_, err = parse(commit)
	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
	}
}
