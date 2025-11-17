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
		Model: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsModelUnion](shared.UnionString("openai/gpt-4")),
		AgentAttributes: githubcomdedaluslabsdedalussdkgo.F(map[string]float64{
			"accuracy":   0.900000,
			"complexity": 0.800000,
			"efficiency": 0.700000,
		}),
		Audio: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"format": "bar",
			"voice":  "bar",
		}),
		AutoExecuteTools:                githubcomdedaluslabsdedalussdkgo.F(true),
		Deferred:                        githubcomdedaluslabsdedalussdkgo.F(true),
		DisableAutomaticFunctionCalling: githubcomdedaluslabsdedalussdkgo.F(true),
		FrequencyPenalty:                githubcomdedaluslabsdedalussdkgo.F(-0.500000),
		FunctionCall:                    githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsFunctionCallUnion](shared.UnionString("string")),
		Functions: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
			"foo": "bar",
		}}),
		GenerationConfig: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"candidateCount":   "bar",
			"responseMimeType": "bar",
		}),
		Guardrails: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
			"foo": "bar",
		}}),
		HandoffConfig: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Input:        githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsInputUnion](shared.UnionString("Translate this paragraph into French.")),
		Instructions: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsInstructionsUnion](shared.UnionString("You are a concise assistant.")),
		LogitBias: githubcomdedaluslabsdedalussdkgo.F(map[string]int64{
			"50256": int64(-100),
		}),
		Logprobs:            githubcomdedaluslabsdedalussdkgo.F(true),
		MaxCompletionTokens: githubcomdedaluslabsdedalussdkgo.F(int64(1000)),
		MaxTokens:           githubcomdedaluslabsdedalussdkgo.F(int64(100)),
		MaxTurns:            githubcomdedaluslabsdedalussdkgo.F(int64(5)),
		MCPServers:          githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMCPServersUnion](githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMCPServersArray([]string{"dedalus-labs/brave-search", "dedalus-labs/github-api"})),
		Messages: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMessagesUnion](githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMessagesArray([]map[string]interface{}{{
			"content": "bar",
			"role":    "bar",
		}})),
		Metadata: githubcomdedaluslabsdedalussdkgo.F(map[string]string{
			"session": "abc",
			"user_id": "123",
		}),
		Modalities: githubcomdedaluslabsdedalussdkgo.F([]string{"text"}),
		ModelAttributes: githubcomdedaluslabsdedalussdkgo.F(map[string]map[string]float64{
			"anthropic/claude-3-5-sonnet": {
				"cost":         0.700000,
				"creativity":   0.800000,
				"intelligence": 0.950000,
			},
			"openai/gpt-4": {
				"cost":         0.800000,
				"intelligence": 0.900000,
				"speed":        0.600000,
			},
			"openai/gpt-4o-mini": {
				"cost":         0.200000,
				"intelligence": 0.700000,
				"speed":        0.900000,
			},
		}),
		N:                 githubcomdedaluslabsdedalussdkgo.F(int64(1)),
		ParallelToolCalls: githubcomdedaluslabsdedalussdkgo.F(true),
		Prediction: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		PresencePenalty: githubcomdedaluslabsdedalussdkgo.F(-0.500000),
		PromptCacheKey:  githubcomdedaluslabsdedalussdkgo.F("prompt_cache_key"),
		ReasoningEffort: githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsReasoningEffortMedium),
		ResponseFormat: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsResponseFormatUnion](shared.ResponseFormatTextParam{
			Type: githubcomdedaluslabsdedalussdkgo.F(shared.ResponseFormatTextTypeText),
		}),
		SafetyIdentifier: githubcomdedaluslabsdedalussdkgo.F("safety_identifier"),
		SafetySettings: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
			"category":  "bar",
			"threshold": "bar",
		}}),
		SearchParameters: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Seed:        githubcomdedaluslabsdedalussdkgo.F(int64(42)),
		ServiceTier: githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsServiceTierAuto),
		Stop:        githubcomdedaluslabsdedalussdkgo.F([]string{"\n", "END"}),
		Store:       githubcomdedaluslabsdedalussdkgo.F(true),
		StreamOptions: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"include_usage": "bar",
		}),
		System:      githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsSystemUnion](shared.UnionString("You are a helpful assistant.")),
		Temperature: githubcomdedaluslabsdedalussdkgo.F(0.000000),
		Thinking: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsThinkingUnion](githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsThinkingThinkingConfigEnabled{
			BudgetTokens: githubcomdedaluslabsdedalussdkgo.F(int64(2048)),
			Type:         githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsThinkingThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsToolChoiceUnion](shared.UnionString("auto")),
		ToolConfig: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"function_calling_config": "bar",
		}),
		Tools: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
			"function": "bar",
			"type":     "bar",
		}}),
		TopK:        githubcomdedaluslabsdedalussdkgo.F(int64(40)),
		TopLogprobs: githubcomdedaluslabsdedalussdkgo.F(int64(5)),
		TopP:        githubcomdedaluslabsdedalussdkgo.F(0.100000),
		User:        githubcomdedaluslabsdedalussdkgo.F("user-123"),
		Verbosity:   githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsVerbosityLow),
		WebSearchOptions: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
	})
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
