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
		Messages: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMessagesUnion{
			OfMapOfAnyMap: []map[string]any{{
				"content": "bar",
				"role":    "bar",
			}},
		},
		Model: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsModelUnion{
			OfModelID: githubcomdedaluslabsdedalussdkgo.String("openai/gpt-4"),
		},
		AgentAttributes: map[string]float64{
			"accuracy":   0.9,
			"complexity": 0.8,
			"efficiency": 0.7,
		},
		Audio: map[string]any{
			"format": "bar",
			"voice":  "bar",
		},
		AutoExecuteTools:                githubcomdedaluslabsdedalussdkgo.Bool(true),
		DisableAutomaticFunctionCalling: githubcomdedaluslabsdedalussdkgo.Bool(true),
		FrequencyPenalty:                githubcomdedaluslabsdedalussdkgo.Float(-0.5),
		FunctionCall: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsFunctionCallUnion{
			OfString: githubcomdedaluslabsdedalussdkgo.String("string"),
		},
		Functions: []map[string]any{{
			"foo": "bar",
		}},
		GenerationConfig: map[string]any{
			"candidateCount":   "bar",
			"responseMimeType": "bar",
		},
		Guardrails: []map[string]any{{
			"foo": "bar",
		}},
		HandoffConfig: map[string]any{
			"foo": "bar",
		},
		Input: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsInputUnion{
			OfString: githubcomdedaluslabsdedalussdkgo.String("Translate this paragraph into French."),
		},
		Instructions: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsInstructionsUnion{
			OfString: githubcomdedaluslabsdedalussdkgo.String("You are a concise assistant."),
		},
		LogitBias: map[string]int64{
			"50256": -100,
		},
		Logprobs:            githubcomdedaluslabsdedalussdkgo.Bool(true),
		MaxCompletionTokens: githubcomdedaluslabsdedalussdkgo.Int(1000),
		MaxTokens:           githubcomdedaluslabsdedalussdkgo.Int(100),
		MaxTurns:            githubcomdedaluslabsdedalussdkgo.Int(5),
		MCPServers: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMCPServersUnion{
			OfMCPServers: []githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMCPServersMCPServerUnion{{
				OfMCPServerSlug: githubcomdedaluslabsdedalussdkgo.String("dedalus-labs/brave-search"),
			}, {
				OfMCPServerSlug: githubcomdedaluslabsdedalussdkgo.String("dedalus-labs/github-api"),
			}},
		},
		Metadata: map[string]string{
			"session": "abc",
			"user_id": "123",
		},
		Modalities: []string{"text"},
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
		N:                 githubcomdedaluslabsdedalussdkgo.Int(1),
		ParallelToolCalls: githubcomdedaluslabsdedalussdkgo.Bool(true),
		Prediction: map[string]any{
			"foo": "bar",
		},
		PresencePenalty: githubcomdedaluslabsdedalussdkgo.Float(-0.5),
		PromptCacheKey:  githubcomdedaluslabsdedalussdkgo.String("prompt_cache_key"),
		ReasoningEffort: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsReasoningEffortMedium,
		ResponseFormat: map[string]any{
			"type": "bar",
		},
		SafetyIdentifier: githubcomdedaluslabsdedalussdkgo.String("safety_identifier"),
		SafetySettings: []map[string]any{{
			"category":  "bar",
			"threshold": "bar",
		}},
		Seed:        githubcomdedaluslabsdedalussdkgo.Int(42),
		ServiceTier: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsServiceTierAuto,
		Stop:        []string{"\n", "END"},
		Store:       githubcomdedaluslabsdedalussdkgo.Bool(true),
		StreamOptions: map[string]any{
			"include_usage": "bar",
		},
		System: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsSystemUnion{
			OfString: githubcomdedaluslabsdedalussdkgo.String("You are a helpful assistant."),
		},
		Temperature: githubcomdedaluslabsdedalussdkgo.Float(0),
		Thinking: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsThinkingUnion{
			OfEnabled: &githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsThinkingEnabled{
				BudgetTokens: 2048,
			},
		},
		ToolChoice: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsToolChoiceUnion{
			OfString: githubcomdedaluslabsdedalussdkgo.String("auto"),
		},
		ToolConfig: map[string]any{
			"function_calling_config": "bar",
		},
		Tools: []map[string]any{{
			"function": "bar",
			"type":     "bar",
		}},
		TopK:        githubcomdedaluslabsdedalussdkgo.Int(40),
		TopLogprobs: githubcomdedaluslabsdedalussdkgo.Int(5),
		TopP:        githubcomdedaluslabsdedalussdkgo.Float(0.1),
		User:        githubcomdedaluslabsdedalussdkgo.String("user-123"),
		Verbosity:   githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsVerbosityLow,
		WebSearchOptions: map[string]any{
			"foo": "bar",
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
