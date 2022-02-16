package parser

import (
	"testing"
)

type returns struct {
	commit *Commit
	err    error
}

var tests = []struct {
	name    string
	args    []byte
	returns returns
}{{"ok", []byte("feat(scope): something"), returns{&Commit{
	Ct:      Feat,
	Scope:   "scope",
	Subject: "something",
	Body:    "",
}, nil}}}

func TestParse(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if c, err := parse(tt.args); err != tt.returns.err && c != tt.returns.commit {
				t.Errorf(
					"expected {commit:%v, err:%v}, but got {commit:%v, err:%v}",
					tt.returns.commit, tt.returns.err, c, err,
				)
			}
		})
	}
}
