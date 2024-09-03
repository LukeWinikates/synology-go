package docker

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateRepositoryName(repository string) error {
	if strings.Contains(repository, ".") {
		hint := ""
		hint = fmt.Sprintf(" - maybe you want '%s'?", RepositoryShortName(repository))
		return errors.New("invalid repository name: must use short form, rather than fully-qualified version" + hint)
	}
	return nil
}

func RepositoryShortName(repository string) string {
	if strings.Contains(repository, ".") {
		elems := strings.Split(repository, "/")[1:]
		return strings.Join(elems, "/")
	}
	return repository
}
