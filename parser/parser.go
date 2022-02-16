package parser

import "strings"

const (
	openingCol       = '('
	closingParColSpc = "): "
	newline          = '\n'
	newlinePair      = "\n\n"
	empty            = ""
)

func IsValidCommitType(toMatch CommitType, matches []CommitType) bool {
	var notMatches uint8

	for i := 0; i < len(matches); i++ {
		if toMatch == matches[i] {
			notMatches |= 1
		}
	}

	if notMatches == 1 {
		return true
	}

	return false
}

func sepFirstLineAndBody(commit []byte) (first, rest string) {
	msg := string(commit)
	splitBy := strings.Index(msg, newlinePair)
	if splitBy == -1 {
		return msg, empty
	}

	return msg[:splitBy], msg[splitBy+len(newlinePair):]
}

func stripCommitType(fl string) (ct CommitType, rest string) {
	splitBy := strings.IndexRune(fl, openingCol)

	if splitBy == -1 {
		return CommitType(fl), ""
	}

	return CommitType(fl[:splitBy]), fl[splitBy:]
}

func ScopeNSubject(rest string) (scope, subject string, err error) {
	fp := strings.IndexRune(rest, openingCol)
	if fp != 0 {
		return empty, empty, ErrEmptyScope
	}

	rest = rest[0+1:]

	sp := strings.Index(rest, closingParColSpc)
	if sp == -1 || sp == 0 {
		return "", "", ErrEmptyScope
	}

	return rest[:sp], rest[sp+len(closingParColSpc):], nil
}

const spacesSet = "      　     "

func CleanedSubject(maybeSubject string) (string, error) {
	maybeSubject = strings.Trim(maybeSubject, spacesSet)
	if len(maybeSubject) == 0 {
		return empty, ErrEmptySubject
	}

	return maybeSubject, nil
}
