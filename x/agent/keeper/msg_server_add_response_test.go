package keeper_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParsing(t *testing.T) {
	k, _, _ := setupMsgServer(t)

	// default params
	testCases := []struct {
		name                  string
		newResponse           string
		existingResponses     []string
		existingFunctionCalls []string
		expResponses          []string
		expFunctionCalls      []string
	}{
		{
			name:                  "minimal",
			newResponse:           "new response",
			existingResponses:     []string{},
			existingFunctionCalls: []string{},
			expResponses:          []string{"new response"},
			expFunctionCalls:      []string{},
		},
		{
			name:                  "existing response",
			newResponse:           "new response",
			existingResponses:     []string{"old response"},
			existingFunctionCalls: []string{"fn1(\"asd\")"},
			expResponses:          []string{"old response", "new response"},
			expFunctionCalls:      []string{"fn1(\"asd\")"},
		},
		{
			name:                  "function call",
			newResponse:           "F_CALL: fn4(\"value\")",
			existingResponses:     []string{"old response"},
			existingFunctionCalls: []string{"fn1(\"asd\")"},
			expResponses:          []string{"old response"},
			expFunctionCalls:      []string{"fn1(\"asd\")", "fn4(\"value\")"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			responses, functionCalls := k.GetInputs(
				tc.newResponse,
				tc.existingResponses,
				tc.existingFunctionCalls,
			)
			require.Equal(t, tc.expResponses, responses)
			require.Equal(t, tc.expFunctionCalls, functionCalls)
		})
	}
}

func TestFunctions(t *testing.T) {
	k, _, _ := setupMsgServer(t)

	// default params
	testCases := []struct {
		name                    string
		existingFunctionResults []string
		functionCalls           []string
		newFunctionCalls        []string
		expectedFunctionResults []string
	}{
		{
			name:                    "minimal",
			existingFunctionResults: []string{},
			functionCalls:           []string{},
			newFunctionCalls:        []string{},
			expectedFunctionResults: []string{},
		},
		{
			name:                    "minimal",
			existingFunctionResults: []string{},
			functionCalls:           []string{},
			newFunctionCalls:        []string{"MultiplyBy8(\"8\")"},
			expectedFunctionResults: []string{"64"},
		},
		{
			name:                    "existing",
			existingFunctionResults: []string{"64"},
			functionCalls:           []string{"MultiplyBy8(\"8\")"},
			newFunctionCalls:        []string{"MultiplyBy8(\"8\")", "DivideBy2(\"8\")"},
			expectedFunctionResults: []string{"64", "4"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			functionResults := k.GetFunctionResults(
				tc.existingFunctionResults,
				tc.functionCalls,
				tc.newFunctionCalls,
			)
			require.Equal(t, tc.expectedFunctionResults, functionResults)
		})
	}
}
