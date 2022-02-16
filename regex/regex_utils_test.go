package regexUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDefineRegexPatternsMessages aims to validate regexUtils.DefineRegexPatternsMessages return
func TestDefineRegexPatternsMessages(t *testing.T) {
	expected := Patterns{
		Pattern{Regex: "[A-Z]",
			Message: "password must have at least one upper case character",
			Found:   true,
		},
		Pattern{Regex: "[a-z]",
			Message: "password must have at least one lower case character",
			Found:   true,
		},
	}
	patterns := []string{"[A-Z]:password must have at least one upper case character:t",
		"[a-z]:password must have at least one lower case character:t",
	}

	assert.EqualValues(t, expected, DefineRegexPatternsAndMessages(patterns))

}

// TestAssertRegexPattern aims to validate regexUtils.AssertRegexPattern return
func TestAssertRegexPattern(t *testing.T) {
	found, nFound := AssertRegexPattern("TestPassword", "[A-Z]")
	assert.EqualValues(t, true, found)
	assert.EqualValues(t, nFound, 2)
}
