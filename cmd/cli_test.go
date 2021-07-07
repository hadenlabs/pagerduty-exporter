package cmd_test

import (

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCLI struct {
	Args             []string
	ExpectedLogLevel string
	ShouldErr        bool
}

var testsCLI []testCLI = []testCLI{
	{
		Args:             []string{"./pagerduty-exporter"},
		ExpectedLogLevel: "info",
		ShouldErr:        false,
	},
	{
		Args:             []string{"./pagerduty-exporter", "--logLevel", "debug"},
		ExpectedLogLevel: "debug",
		ShouldErr:        false,
	},
	{
		Args:             []string{"./pagerduty-exporter", "--logLevel", "42"},
		ShouldErr:        true,
		ExpectedLogLevel: "info",
	},
}

func TestCLI(t *testing.T) {
	for i, test := range testsCLI {
		cli := cmd.NewCLI()
		err := cli.Run(test.Args)
		if test.ShouldErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.ExpectedLogLevel, log.GetLevel().String())
		t.Logf("Test at index %d passed\n", i)
		// Reset logLevel to default
		log.SetLevel(log.InfoLevel)
	}
}
