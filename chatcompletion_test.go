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
)

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Chat.Completions.New(context.TODO(), githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParams{
		CompletionRequest: githubcomdedaluslabsdedalussdkgo.CompletionRequestParam{
			AgentAttributes: map[string]float64{
				"accuracy":   0.9,
				"complexity": 0.8,
				"efficiency": 0.7,
			},
			FrequencyPenalty: githubcomdedaluslabsdedalussdkgo.Float(-0.5),
			Guardrails: []map[string]any{{
				"foo": "bar",
			}},
			HandoffConfig: map[string]any{
				"foo": "bar",
			},
			LogitBias: map[string]int64{
				"50256": -100,
			},
			MaxTokens:  githubcomdedaluslabsdedalussdkgo.Int(100),
			MaxTurns:   githubcomdedaluslabsdedalussdkgo.Int(5),
			MCPServers: []string{"dedalus-labs/brave-search", "dedalus-labs/github-api"},
			Messages: []map[string]any{{
				"content": "bar",
				"role":    "bar",
			}},
			Model: githubcomdedaluslabsdedalussdkgo.CompletionRequestModelUnionParam{
				OfModelID: githubcomdedaluslabsdedalussdkgo.String("openai/gpt-4"),
			},
			ModelAttributes: map[string]map[string]float64{
				"anthropic/claude-3-5-sonnet": {
					"cost":         0.7,
					"creativity":   0.8,
					"intelligence": 0.95,
				},
				"openai/gpt-4": {
					"cost":         0.8,
					"intelligence": 0.9,
					"speed":        0.6,
				},
				"openai/gpt-4o-mini": {
					"cost":         0.2,
					"intelligence": 0.7,
					"speed":        0.9,
				},
			},
			N:               githubcomdedaluslabsdedalussdkgo.Int(1),
			PresencePenalty: githubcomdedaluslabsdedalussdkgo.Float(-0.5),
			Stop:            []string{"\\n", "END"},
			Stream:          githubcomdedaluslabsdedalussdkgo.Bool(true),
			Temperature:     githubcomdedaluslabsdedalussdkgo.Float(0),
			ToolChoice: githubcomdedaluslabsdedalussdkgo.CompletionRequestToolChoiceUnionParam{
				OfString: githubcomdedaluslabsdedalussdkgo.String("auto"),
			},
			Tools: []map[string]any{{
				"function": "bar",
				"type":     "bar",
			}},
			TopP: githubcomdedaluslabsdedalussdkgo.Float(0.1),
			User: githubcomdedaluslabsdedalussdkgo.String("user-123"),
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
