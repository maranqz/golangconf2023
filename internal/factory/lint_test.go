package factory_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/maranqz/go-factory-lint"
)

func TestLinterSuite(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()

	tests := map[string]struct {
		pkgs    []string
		prepare func(t *testing.T, a *analysis.Analyzer)
	}{
		"simple": {pkgs: []string{"simple/..."}},
		"blockedPkgs": {
			pkgs: []string{"blockedPkgs/..."},
			prepare: func(t *testing.T, a *analysis.Analyzer) {
				if err := a.Flags.Set("blockedPkgs", "factory/blockedPkgs/blocked"); err != nil {
					t.Fatal(err)
				}
			},
		},
		"onlyBlockedPkgs": {
			pkgs: []string{"onlyBlockedPkgs/main/..."},
			prepare: func(t *testing.T, a *analysis.Analyzer) {
				if err := a.Flags.Set("b", "factory/onlyBlockedPkgs/blocked"); err != nil {
					t.Fatal(err)
				}

				if err := a.Flags.Set("ob", "true"); err != nil {
					t.Fatal(err)
				}

				if err := a.Flags.Set("onlyBlockedPkgs", "true"); err != nil {
					t.Fatal(err)
				}
			},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dirs := make([]string, 0, len(tt.pkgs))

			for _, pkg := range tt.pkgs {
				dirs = append(dirs, filepath.Join(testdata, "src", "factory", pkg))
			}

			analyzer := factory.NewAnalyzer()

			if tt.prepare != nil {
				tt.prepare(t, analyzer)
			}

			analysistest.Run(t, TestdataDir(),
				analyzer, dirs...)
		})
	}
}

func TestdataDir() string {
	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get current test filename")
	}

	return filepath.Join(filepath.Dir(testFilename), "testdata")
}
