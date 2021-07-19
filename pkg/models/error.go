package models

type GoweeklyNameNotFoundError struct{}

func (k *GoweeklyNameNotFoundError) Error() string {
	return "Goweekly name not found"
}
