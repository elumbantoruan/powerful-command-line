package pomodoro_test

import (
	"powerful-command-line/pomo/pomodoro"
	"powerful-command-line/pomo/pomodoro/repository"
	"testing"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	return repository.NewInMemoryRepo(), func() {}
}
