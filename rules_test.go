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

	t.Run("tags", tags)
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
