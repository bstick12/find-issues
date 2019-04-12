package issues_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestIssues(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("issues_junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Issues Suite", []Reporter{junitReporter})
}
