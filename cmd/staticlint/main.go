package main

import (
	"strings"

	"github.com/gostaticanalysis/sqlrows/passes/sqlrows"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/stringintconv"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"honnef.co/go/tools/staticcheck"

	libVarsAnalyzer "github.com/sashamelentyev/usestdlibvars/pkg/analyzer"

	"github.com/lastbyte32/go-metric/internal/analyzers/exit"
)

type checkerRules []*analysis.Analyzer

func (a *checkerRules) addAnalyzers() {
	*a = append(*a,
		structtag.Analyzer,
		unmarshal.Analyzer,
		printf.Analyzer,
		shadow.Analyzer,
		sqlrows.Analyzer,
		unreachable.Analyzer,
		stringintconv.Analyzer,
		exit.Analyzer,
		libVarsAnalyzer.New(),
	)
}

func (a *checkerRules) addAnalyzersClassOf(class string) {
	for _, v := range staticcheck.Analyzers {
		if strings.Contains(v.Analyzer.Name, class) {
			*a = append(*a, v.Analyzer)
		}
	}
}

func main() {
	var rules checkerRules
	rules.addAnalyzersClassOf("SA")
	rules.addAnalyzers()

	multichecker.Main(
		rules...,
	)
}
