// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalus_test

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
	client := dedalus.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chat.New(context.TODO(), dedalus.ChatNewParams{
		CompletionRequest: dedalus.CompletionRequestParam{
			AgentAttributes: map[string]float64{
				"accuracy":   0.9,
				"complexity": 0.8,
				"efficiency": 0.7,
			},
			FrequencyPenalty: dedalus.Float(-0.5),
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
			MaxTokens:  dedalus.Int(100),
			MaxTurns:   dedalus.Int(5),
			McpServers: []string{"dedalus-labs/brave-search", "dedalus-labs/github-api"},
			Messages: []map[string]any{{
				"foo": "bar",
			}},
			Model: dedalus.CompletionRequestModelUnionParam{
				OfString: dedalus.String("gpt-4"),
			},
			ModelAttributes: map[string]map[string]float64{
				"claude-3-5-sonnet": {
					"cost":         0.7,
					"creativity":   0.8,
					"intelligence": 0.95,
				},
				"gpt-4": {
					"cost":         0.8,
					"intelligence": 0.9,
					"speed":        0.6,
				},
				"gpt-4o-mini": {
					"cost":         0.2,
					"intelligence": 0.7,
					"speed":        0.9,
				},
			},
			N:               dedalus.Int(1),
			PresencePenalty: dedalus.Float(-0.5),
			Stop:            []string{"\\n", "END"},
			Stream:          dedalus.Bool(true),
			Temperature:     dedalus.Float(0),
			ToolChoice: dedalus.CompletionRequestToolChoiceUnionParam{
				OfString: dedalus.String("auto"),
			},
			Tools: []map[string]any{{
				"function": "bar",
				"type":     "bar",
			}},
			TopP: dedalus.Float(0.1),
			User: dedalus.String("user-123"),
		},
	})
	if err != nil {
		var apierr *dedalus.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
