package AWSAT002_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/tools/providerlint/passes/AWSAT002"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAWSAT002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AWSAT002.Analyzer, "a")
}
