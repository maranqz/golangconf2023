package gorules

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/quasilyte/go-ruleguard/analyzer"
	_ "github.com/quasilyte/go-ruleguard/dsl"
	"golang.org/x/tools/go/analysis/analysistest"
)

// Using common test because of global analyzer.Analyzer
func Test(t *testing.T) {
	analyzer.ForceNewEngine = true

	t.Run("write", write)
	t.Run("tags", tags)
	t.Run("factory", factory)
}

func tags(t *testing.T) {
	testdata := analysistest.TestData()

	if err := analyzer.Analyzer.Flags.Set("rules", "tags/rules.go"); err != nil {
		t.Fatalf("set rules flag: %v", err)
	}

	analysistest.Run(
		t,
		TestdataDir(),
		analyzer.Analyzer,
		filepath.Join(testdata, "src", "tags"),
	)
}

func factory(t *testing.T) {
	testdata := analysistest.TestData()
	if err := analyzer.Analyzer.Flags.Set("rules", "factory/rules.go"); err != nil {
		t.Fatalf("set rules flag: %v", err)
	}

	factoryTestDir := filepath.Join(testdata, "src", "factory")

	tests := map[string]struct{ dir string }{
		"root":    {dir: "..."},
		"nested":  {dir: "nested/..."},
		"nested2": {dir: "nested/nested2/..."},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			analysistest.Run(
				t,
				RootDir(),
				analyzer.Analyzer,
				filepath.Join(factoryTestDir, tt.dir),
			)
		})
	}
}

func write(t *testing.T) {
	testdata := analysistest.TestData()
	if err := analyzer.Analyzer.Flags.Set("rules", "write/rules.go"); err != nil {
		t.Fatalf("set rules flag: %v", err)
	}

	analysistest.Run(
		t,
		RootDir(),
		analyzer.Analyzer,
		filepath.Join(testdata, "src", "write"),
	)
}

func TestdataDir() string {
	return filepath.Join(RootDir(), "testdata")
}

func RootDir() string {
	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get current test filename")
	}

	return filepath.Dir(testFilename)
}
