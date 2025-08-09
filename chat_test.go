// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
)

func TestChatNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomdedaluslabsdedalussdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chat.New(context.TODO(), githubcomdedaluslabsdedalussdkgo.ChatNewParams{
		CompletionRequest: githubcomdedaluslabsdedalussdkgo.CompletionRequestParam{
			AgentAttributes: githubcomdedaluslabsdedalussdkgo.F(map[string]float64{
				"accuracy":   0.900000,
				"complexity": 0.800000,
				"efficiency": 0.700000,
			}),
			FrequencyPenalty: githubcomdedaluslabsdedalussdkgo.F(-0.500000),
			Guardrails: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			HandoffConfig: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Input: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
				"content": "bar",
				"role":    "bar",
			}}),
			LogitBias: githubcomdedaluslabsdedalussdkgo.F(map[string]int64{
				"50256": int64(-100),
			}),
			MaxTokens:  githubcomdedaluslabsdedalussdkgo.F(int64(100)),
			MaxTurns:   githubcomdedaluslabsdedalussdkgo.F(int64(5)),
			McpServers: githubcomdedaluslabsdedalussdkgo.F([]string{"dedalus-labs/brave-search", "dedalus-labs/github-api"}),
			Model:      githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.CompletionRequestModelUnionParam](shared.UnionString("gpt-4")),
			ModelAttributes: githubcomdedaluslabsdedalussdkgo.F(map[string]map[string]float64{
				"claude-3-5-sonnet": {
					"cost":         0.700000,
					"creativity":   0.800000,
					"intelligence": 0.950000,
				},
				"gpt-4": {
					"cost":         0.800000,
					"intelligence": 0.900000,
					"speed":        0.600000,
				},
				"gpt-4o-mini": {
					"cost":         0.200000,
					"intelligence": 0.700000,
					"speed":        0.900000,
				},
			}),
			N:               githubcomdedaluslabsdedalussdkgo.F(int64(1)),
			PresencePenalty: githubcomdedaluslabsdedalussdkgo.F(-0.500000),
			Stop:            githubcomdedaluslabsdedalussdkgo.F([]string{"\n", "END"}),
			Stream:          githubcomdedaluslabsdedalussdkgo.F(true),
			Temperature:     githubcomdedaluslabsdedalussdkgo.F(0.000000),
			ToolChoice:      githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.CompletionRequestToolChoiceUnionParam](shared.UnionString("auto")),
			Tools: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
				"function": "bar",
				"type":     "bar",
			}}),
			TopP: githubcomdedaluslabsdedalussdkgo.F(0.100000),
			User: githubcomdedaluslabsdedalussdkgo.F("user-123"),
		},
	})
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
