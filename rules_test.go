package gorules

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/quasilyte/go-ruleguard/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRulesTags(t *testing.T) {
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

func TestRulesFactory(t *testing.T) {
	testdata := analysistest.TestData()
	if err := analyzer.Analyzer.Flags.Set("rules", "factory/rules.go"); err != nil {
		t.Fatalf("set rules flag: %v", err)
	}

	analysistest.Run(
		t,
		RootDir(),
		analyzer.Analyzer,
		filepath.Join(testdata, "src", "factory"),
	)
}

func TestRulesWrite(t *testing.T) {
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
