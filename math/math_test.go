package math

import (
	"os"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	if Abs(-1) < 0 {
		t.Error("Negative value found in abs() with", -1)
	}
	if Abs(0) < 0 {
		t.Error("Negative value found in abs() with", 0)
	}
	if Abs(1) < 0 {
		t.Error("Negative value found in abs() with", 1)
	}
}

// testes com tables
func TestTableAbs(t *testing.T) {
	var tests = []struct {
		name     string
		input    int
		expected int
	}{
		{"-1 should be 1", -1, 1},
		{"1 should be 1", 1, 1},
		{"0 should be 0", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Abs(tt.input) != tt.expected {
				t.Errorf("Test Abs failed with input %v", tt.input)
			}
		})
	}
}

// escrevendo sub testes
func TestAbsSub(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		if Abs(1) < 0 {
			t.Error("Negative value found in abs()")
		}
	})
	t.Run("Negative", func(t *testing.T) {
		if Abs(-1) < 0 {
			t.Error("Negative value found in abs()")
		}
	})
	t.Run("Zero", func(t *testing.T) {
		if Abs(0) < 0 {
			t.Error("Negative value found in abs()")
		}
	})
}

// skip caso algo não ocorrer como esperado
func TestSkip(t *testing.T) {
	if len(os.Getenv("GOPATH")) == 0 {
		t.Skip("Skipping test because GOPATH isn't set")
	}
	// ...
	t.Log("Tested with GOPATH:", os.Getenv("GOPATH"))
}

// função de cleanup funciona basicamente como um defer, executa no final de tudo
func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("Cleanup")
	})

	t.Log("Running some test")
}

// rodar os testes paralelamente(usando concurrency)
func TestParallel1(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 3)
}

func TestParallel2(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 3)
}

func TestParallel3(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 3)
}
