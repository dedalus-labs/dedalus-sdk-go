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
		Model: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsModelUnion](shared.UnionString("openai/gpt-5")),
		AgentAttributes: githubcomdedaluslabsdedalussdkgo.F(map[string]float64{
			"accuracy":   0.900000,
			"complexity": 0.800000,
		}),
		Audio: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		AutomaticToolExecution: githubcomdedaluslabsdedalussdkgo.F(true),
		CachedContent:          githubcomdedaluslabsdedalussdkgo.F("cachedContent"),
		Deferred:               githubcomdedaluslabsdedalussdkgo.F(true),
		FrequencyPenalty:       githubcomdedaluslabsdedalussdkgo.F(-2.000000),
		FunctionCall:           githubcomdedaluslabsdedalussdkgo.F("function_call"),
		Functions: githubcomdedaluslabsdedalussdkgo.F([]githubcomdedaluslabsdedalussdkgo.ChatCompletionFunctionsParam{{
			Name:        githubcomdedaluslabsdedalussdkgo.F("name"),
			Description: githubcomdedaluslabsdedalussdkgo.F("description"),
			Parameters: githubcomdedaluslabsdedalussdkgo.F(shared.FunctionParameters{
				"foo": "bar",
			}),
		}}),
		GenerationConfig: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Guardrails: githubcomdedaluslabsdedalussdkgo.F([]map[string]interface{}{{
			"foo": "bar",
		}}),
		HandoffConfig: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		LogitBias: githubcomdedaluslabsdedalussdkgo.F(map[string]int64{
			"foo": int64(0),
		}),
		Logprobs:            githubcomdedaluslabsdedalussdkgo.F(true),
		MaxCompletionTokens: githubcomdedaluslabsdedalussdkgo.F(int64(0)),
		MaxTokens:           githubcomdedaluslabsdedalussdkgo.F(int64(1)),
		MaxTurns:            githubcomdedaluslabsdedalussdkgo.F(int64(5)),
		MCPServers:          githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMCPServersUnion](githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMCPServersArray([]string{"dedalus-labs/brave-search"})),
		Messages: githubcomdedaluslabsdedalussdkgo.F([]githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMessageUnion{githubcomdedaluslabsdedalussdkgo.ChatCompletionDeveloperMessageParam{
			Content: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionDeveloperMessageParamContentUnion](shared.UnionString("string")),
			Role:    githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionDeveloperMessageParamRoleDeveloper),
			Name:    githubcomdedaluslabsdedalussdkgo.F("name"),
		}}),
		Metadata: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Modalities: githubcomdedaluslabsdedalussdkgo.F([]string{"string"}),
		ModelAttributes: githubcomdedaluslabsdedalussdkgo.F(map[string]map[string]float64{
			"gpt-5": {
				"accuracy": 0.950000,
				"speed":    0.600000,
			},
		}),
		N:                 githubcomdedaluslabsdedalussdkgo.F(int64(1)),
		ParallelToolCalls: githubcomdedaluslabsdedalussdkgo.F(true),
		Prediction: githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.PredictionContentParam{
			Content: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.PredictionContentContentUnionParam](shared.UnionString("string")),
			Type:    githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.PredictionContentTypeContent),
		}),
		PresencePenalty:      githubcomdedaluslabsdedalussdkgo.F(-2.000000),
		PromptCacheKey:       githubcomdedaluslabsdedalussdkgo.F("prompt_cache_key"),
		PromptCacheRetention: githubcomdedaluslabsdedalussdkgo.F("prompt_cache_retention"),
		PromptMode: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		ReasoningEffort: githubcomdedaluslabsdedalussdkgo.F("reasoning_effort"),
		ResponseFormat: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsResponseFormatUnion](shared.ResponseFormatTextParam{
			Type: githubcomdedaluslabsdedalussdkgo.F(shared.ResponseFormatTextTypeText),
		}),
		SafePrompt:       githubcomdedaluslabsdedalussdkgo.F(true),
		SafetyIdentifier: githubcomdedaluslabsdedalussdkgo.F("safety_identifier"),
		SafetySettings: githubcomdedaluslabsdedalussdkgo.F([]githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsSafetySetting{{
			Category:  githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryUnspecified),
			Threshold: githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsSafetySettingsThresholdHarmBlockThresholdUnspecified),
		}}),
		SearchParameters: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Seed:          githubcomdedaluslabsdedalussdkgo.F(int64(0)),
		ServiceTier:   githubcomdedaluslabsdedalussdkgo.F("service_tier"),
		Stop:          githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsStopUnion](githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsStopArray([]string{"string"})),
		StopSequences: githubcomdedaluslabsdedalussdkgo.F([]string{"string"}),
		Store:         githubcomdedaluslabsdedalussdkgo.F(true),
		StreamOptions: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		SystemInstruction: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsSystemInstructionUnion](githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsSystemInstructionMap(map[string]interface{}{
			"foo": "bar",
		})),
		Temperature: githubcomdedaluslabsdedalussdkgo.F(0.000000),
		Thinking: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsThinkingUnion](githubcomdedaluslabsdedalussdkgo.ThinkingConfigEnabledParam{
			BudgetTokens: githubcomdedaluslabsdedalussdkgo.F(int64(1024)),
			Type:         githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsToolChoiceUnion](githubcomdedaluslabsdedalussdkgo.ToolChoiceAutoParam{
			Type:                   githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ToolChoiceAutoTypeAuto),
			DisableParallelToolUse: githubcomdedaluslabsdedalussdkgo.F(true),
		}),
		ToolConfig: githubcomdedaluslabsdedalussdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Tools: githubcomdedaluslabsdedalussdkgo.F([]githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsToolUnion{githubcomdedaluslabsdedalussdkgo.ChatCompletionToolParam{
			Function: githubcomdedaluslabsdedalussdkgo.F(shared.FunctionDefinitionParam{
				Name:        githubcomdedaluslabsdedalussdkgo.F("name"),
				Description: githubcomdedaluslabsdedalussdkgo.F("description"),
				Parameters: githubcomdedaluslabsdedalussdkgo.F(shared.FunctionParameters{
					"foo": "bar",
				}),
				Strict: githubcomdedaluslabsdedalussdkgo.F(true),
			}),
			Type: githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.ChatCompletionToolParamTypeFunction),
		}}),
		TopK:        githubcomdedaluslabsdedalussdkgo.F(int64(0)),
		TopLogprobs: githubcomdedaluslabsdedalussdkgo.F(int64(0)),
		TopP:        githubcomdedaluslabsdedalussdkgo.F(0.000000),
		User:        githubcomdedaluslabsdedalussdkgo.F("user"),
		Verbosity:   githubcomdedaluslabsdedalussdkgo.F("verbosity"),
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
