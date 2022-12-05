package main

import "testing"

func TestDefaultConstruct(t *testing.T) {
	answer := Answer{}

	t.Run("test default construct 1", func(t *testing.T) {
		commit := `feat(cmt): update(v1.4.0-BF-6666)`
		defaultConstruct(&answer, commit)
		t.Log(answer)
	})

	t.Run("test default construct 2", func(t *testing.T) {
		commit := `feat(cmt): update`
		defaultConstruct(&answer, commit)
		t.Log(answer)
	})
}
