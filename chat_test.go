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
			AgentAttributes: map[string]any{
				"accuracy":   "bar",
				"complexity": "bar",
				"efficiency": "bar",
			},
			FrequencyPenalty: githubcomdedaluslabsdedalussdkgo.Float(-0.5),
			Guardrails: []map[string]any{{
				"foo": "bar",
			}},
			HandoffConfig: map[string]any{
				"foo": "bar",
			},
			Input: []map[string]any{{
				"content": "bar",
				"role":    "bar",
			}},
			LogitBias: map[string]int64{
				"50256": -100,
			},
			MaxTokens:  githubcomdedaluslabsdedalussdkgo.Int(100),
			MaxTurns:   githubcomdedaluslabsdedalussdkgo.Int(5),
			McpServers: []string{"dedalus-labs/brave-search", "dedalus-labs/github-api"},
			Model: githubcomdedaluslabsdedalussdkgo.CompletionRequestModelUnionParam{
				OfString: githubcomdedaluslabsdedalussdkgo.String("gpt-4"),
			},
			ModelAttributes: map[string]map[string]any{
				"gpt-4": {
					"context_window": "bar",
					"description":    "bar",
					"multimodal":     "bar",
					"strengths":      "bar",
				},
				"gpt-4o-mini": {
					"cost_tier":   "bar",
					"description": "bar",
					"strengths":   "bar",
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
