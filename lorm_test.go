package lorm

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "lorm.db")
	if err != nil {
		t.Fatal("Faild to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engin := OpenDB(t)
	defer engin.Close()
}
