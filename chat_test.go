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
				"foo": 0.000000,
			}),
			FrequencyPenalty: githubcomdedaluslabsdedalussdkgo.F(-2.000000),
			Input:            githubcomdedaluslabsdedalussdkgo.F([]interface{}{map[string]interface{}{}}),
			LogitBias: githubcomdedaluslabsdedalussdkgo.F(map[string]int64{
				"foo": int64(0),
			}),
			MaxTokens:  githubcomdedaluslabsdedalussdkgo.F(int64(1)),
			MaxTurns:   githubcomdedaluslabsdedalussdkgo.F(int64(1)),
			McpServers: githubcomdedaluslabsdedalussdkgo.F([]string{"string"}),
			Model:      githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.CompletionRequestModelUnionParam](shared.UnionString("string")),
			ModelAttributes: githubcomdedaluslabsdedalussdkgo.F(map[string]map[string]float64{
				"foo": {
					"foo": 0.000000,
				},
			}),
			N:               githubcomdedaluslabsdedalussdkgo.F(int64(1)),
			PresencePenalty: githubcomdedaluslabsdedalussdkgo.F(-2.000000),
			Stop:            githubcomdedaluslabsdedalussdkgo.F([]string{"string"}),
			Stream:          githubcomdedaluslabsdedalussdkgo.F(true),
			Temperature:     githubcomdedaluslabsdedalussdkgo.F(0.000000),
			ToolChoice:      githubcomdedaluslabsdedalussdkgo.F[any](map[string]interface{}{}),
			Tools:           githubcomdedaluslabsdedalussdkgo.F([]interface{}{map[string]interface{}{}}),
			TopP:            githubcomdedaluslabsdedalussdkgo.F(0.000000),
			User:            githubcomdedaluslabsdedalussdkgo.F("user"),
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
