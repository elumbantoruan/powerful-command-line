package cmd

import (
	"powerful-command-line/pomo/pomodoro"
	"powerful-command-line/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
