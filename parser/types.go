package parser

type CommitType string

const (
	Feat     CommitType = "feat"
	Fix      CommitType = "fix"
	Refactor CommitType = "refactor"
)

var CommitTypes = []CommitType{Feat, Fix, Refactor}

type Commit struct {
	Ct                   CommitType
	Scope, Subject, Body string
}
