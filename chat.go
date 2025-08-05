// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/ssestream"
	"github.com/dedalus-labs/dedalus-sdk-go/shared/constant"
)

// ChatService contains methods and other services that help with interacting with
// the dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options []option.RequestOption
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	return
}

// Create a chat completion using the Agent framework.
//
// This endpoint provides a vendor-agnostic chat completion API that works with
// 100+ LLM providers via the Agent framework. It supports both single and
// multi-model routing, client-side and server-side tool execution, and integration
// with MCP (Model Context Protocol) servers.
//
// Features: - Cross-vendor compatibility (OpenAI, Anthropic, Cohere, etc.) -
// Multi-model routing with intelligent agentic handoffs - Client-side tool
// execution (tools returned as JSON) - Server-side MCP tool execution with
// automatic billing - Streaming and non-streaming responses - Advanced agent
// attributes for routing decisions - Automatic usage tracking and billing
//
// Args: request: Chat completion request with messages, model, and configuration
// http_request: FastAPI request object for accessing headers and state
// background_tasks: FastAPI background tasks for async billing operations user:
// Authenticated user with validated API key and sufficient balance
//
// Returns: ChatCompletion: OpenAI-compatible completion response with usage data
//
// Raises: HTTPException: - 401 if authentication fails or insufficient balance -
// 400 if request validation fails - 500 if internal processing error occurs
//
// Billing: - Token usage billed automatically based on model pricing - MCP tool
// calls billed separately using credits system - Streaming responses billed after
// completion via background task
//
// Example: Basic chat completion: ```python import dedalus_labs
//
//	client = dedalus_labs.Client(api_key="your-api-key")
//
//	completion = client.chat.create(
//	    model="gpt-4",
//	    input=[{"role": "user", "content": "Hello, how are you?"}],
//	)
//
//	print(completion.choices[0].message.content)
//	```
//
//	With tools and MCP servers:
//	```python
//	completion = client.chat.create(
//	    model="gpt-4",
//	    input=[{"role": "user", "content": "Search for recent AI news"}],
//	    tools=[
//	        {
//	            "type": "function",
//	            "function": {
//	                "name": "search_web",
//	                "description": "Search the web for information",
//	            },
//	        }
//	    ],
//	    mcp_servers=["dedalus-labs/brave-search"],
//	)
//	```
//
//	Multi-model routing:
//	```python
//	completion = client.chat.create(
//	    model=["gpt-4o-mini", "gpt-4", "claude-3-5-sonnet"],
//	    input=[{"role": "user", "content": "Analyze this complex data"}],
//	    agent_attributes={"complexity": 0.8, "accuracy": 0.9},
//	)
//	```
//
//	Streaming response:
//	```python
//	stream = client.chat.create(
//	    model="gpt-4",
//	    input=[{"role": "user", "content": "Tell me a story"}],
//	    stream=True,
//	)
//
//	for chunk in stream:
//	    if chunk.choices[0].delta.content:
//	        print(chunk.choices[0].delta.content, end="")
//	```
func (r *ChatService) NewStreaming(ctx context.Context, body ChatNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[Completion]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	path := "v1/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[Completion](ssestream.NewDecoder(raw), err)
}

type Completion struct {
	ID      string                  `json:"id,required"`
	Choices []CompletionChoice      `json:"choices,required"`
	Created int64                   `json:"created,required"`
	Model   string                  `json:"model,required"`
	Object  constant.ChatCompletion `json:"object,required"`
	// Any of "auto", "default", "flex", "scale", "priority".
	ServiceTier       CompletionServiceTier `json:"service_tier,nullable"`
	SystemFingerprint string                `json:"system_fingerprint,nullable"`
	Usage             CompletionUsage       `json:"usage,nullable"`
	ExtraFields       map[string]any        `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Choices           respjson.Field
		Created           respjson.Field
		Model             respjson.Field
		Object            respjson.Field
		ServiceTier       respjson.Field
		SystemFingerprint respjson.Field
		Usage             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Completion) RawJSON() string { return r.JSON.raw }
func (r *Completion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoice struct {
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string                   `json:"finish_reason,required"`
	Index        int64                    `json:"index,required"`
	Message      CompletionChoiceMessage  `json:"message,required"`
	Logprobs     CompletionChoiceLogprobs `json:"logprobs,nullable"`
	ExtraFields  map[string]any           `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Message      respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoice) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceMessage struct {
	Role         constant.Assistant                  `json:"role,required"`
	Annotations  []CompletionChoiceMessageAnnotation `json:"annotations,nullable"`
	Audio        CompletionChoiceMessageAudio        `json:"audio,nullable"`
	Content      string                              `json:"content,nullable"`
	FunctionCall CompletionChoiceMessageFunctionCall `json:"function_call,nullable"`
	Refusal      string                              `json:"refusal,nullable"`
	ToolCalls    []CompletionChoiceMessageToolCall   `json:"tool_calls,nullable"`
	ExtraFields  map[string]any                      `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role         respjson.Field
		Annotations  respjson.Field
		Audio        respjson.Field
		Content      respjson.Field
		FunctionCall respjson.Field
		Refusal      respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessage) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceMessageAnnotation struct {
	Type        constant.URLCitation                         `json:"type,required"`
	URLCitation CompletionChoiceMessageAnnotationURLCitation `json:"url_citation,required"`
	ExtraFields map[string]any                               `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URLCitation respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageAnnotation) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceMessageAnnotationURLCitation struct {
	EndIndex    int64          `json:"end_index,required"`
	StartIndex  int64          `json:"start_index,required"`
	Title       string         `json:"title,required"`
	URL         string         `json:"url,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageAnnotationURLCitation) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceMessageAudio struct {
	ID          string         `json:"id,required"`
	Data        string         `json:"data,required"`
	ExpiresAt   int64          `json:"expires_at,required"`
	Transcript  string         `json:"transcript,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		ExpiresAt   respjson.Field
		Transcript  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageAudio) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceMessageFunctionCall struct {
	Arguments   string         `json:"arguments,required"`
	Name        string         `json:"name,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceMessageToolCall struct {
	ID          string                                  `json:"id,required"`
	Function    CompletionChoiceMessageToolCallFunction `json:"function,required"`
	Type        constant.Function                       `json:"type,required"`
	ExtraFields map[string]any                          `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageToolCall) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceMessageToolCallFunction struct {
	Arguments   string         `json:"arguments,required"`
	Name        string         `json:"name,required"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceLogprobs struct {
	Content     []CompletionChoiceLogprobsContent `json:"content,nullable"`
	Refusal     []CompletionChoiceLogprobsRefusal `json:"refusal,nullable"`
	ExtraFields map[string]any                    `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceLogprobsContent struct {
	Token       string                                      `json:"token,required"`
	Logprob     float64                                     `json:"logprob,required"`
	TopLogprobs []CompletionChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                     `json:"bytes,nullable"`
	ExtraFields map[string]any                              `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceLogprobsContentTopLogprob struct {
	Token       string         `json:"token,required"`
	Logprob     float64        `json:"logprob,required"`
	Bytes       []int64        `json:"bytes,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceLogprobsRefusal struct {
	Token       string                                      `json:"token,required"`
	Logprob     float64                                     `json:"logprob,required"`
	TopLogprobs []CompletionChoiceLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                     `json:"bytes,nullable"`
	ExtraFields map[string]any                              `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceLogprobsRefusal) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceLogprobsRefusal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoiceLogprobsRefusalTopLogprob struct {
	Token       string         `json:"token,required"`
	Logprob     float64        `json:"logprob,required"`
	Bytes       []int64        `json:"bytes,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceLogprobsRefusalTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionServiceTier string

const (
	CompletionServiceTierAuto     CompletionServiceTier = "auto"
	CompletionServiceTierDefault  CompletionServiceTier = "default"
	CompletionServiceTierFlex     CompletionServiceTier = "flex"
	CompletionServiceTierScale    CompletionServiceTier = "scale"
	CompletionServiceTierPriority CompletionServiceTier = "priority"
)

type CompletionUsage struct {
	CompletionTokens        int64                                  `json:"completion_tokens,required"`
	PromptTokens            int64                                  `json:"prompt_tokens,required"`
	TotalTokens             int64                                  `json:"total_tokens,required"`
	CompletionTokensDetails CompletionUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	PromptTokensDetails     CompletionUsagePromptTokensDetails     `json:"prompt_tokens_details,nullable"`
	ExtraFields             map[string]any                         `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens        respjson.Field
		PromptTokens            respjson.Field
		TotalTokens             respjson.Field
		CompletionTokensDetails respjson.Field
		PromptTokensDetails     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionUsage) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionUsageCompletionTokensDetails struct {
	AcceptedPredictionTokens int64          `json:"accepted_prediction_tokens,nullable"`
	AudioTokens              int64          `json:"audio_tokens,nullable"`
	ReasoningTokens          int64          `json:"reasoning_tokens,nullable"`
	RejectedPredictionTokens int64          `json:"rejected_prediction_tokens,nullable"`
	ExtraFields              map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedPredictionTokens respjson.Field
		AudioTokens              respjson.Field
		ReasoningTokens          respjson.Field
		RejectedPredictionTokens respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionUsagePromptTokensDetails struct {
	AudioTokens  int64          `json:"audio_tokens,nullable"`
	CachedTokens int64          `json:"cached_tokens,nullable"`
	ExtraFields  map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioTokens  respjson.Field
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request model for chat completions.
//
// Validates incoming chat requests with support for multimodality, multi-model
// routing, and agent-enhanced features. Compatible with OpenAI API format while
// extending functionality for advanced use cases.
//
// This model supports both the OpenAI-standard 'messages' field and the
// Dedalus-specific 'input' field for maximum compatibility. The 'input' field can
// handle various modalities beyond text messages.
//
// Key Features: - Multi-model routing with intelligent handoffs - MCP (Model
// Context Protocol) server integration - Advanced agent attributes for routing
// decisions - Client-side and server-side tool execution - Streaming and
// non-streaming responses - Automatic usage tracking and billing
//
// Examples: Basic chat completion:
// `python request = ChatCompletionRequest( model="gpt-4", input=[ {"role": "user", "content": "Hello, how are you?"} ] ) `
//
//	Multi-model routing with attributes:
//	```python
//	request = ChatCompletionRequest(
//	    model=["gpt-4o-mini", "gpt-4", "claude-3-5-sonnet"],
//	    input=[
//	        {"role": "user", "content": "Analyze this complex problem"}
//	    ],
//	    agent_attributes={
//	        "complexity": 0.8,
//	        "accuracy": 0.9
//	    },
//	    model_attributes={
//	        "gpt-4": {"intelligence": 0.9, "cost": 0.8},
//	        "claude-3-5-sonnet": {"intelligence": 0.95, "cost": 0.7}
//	    }
//	)
//	```
//
//	With tools and MCP servers:
//	```python
//	request = ChatCompletionRequest(
//	    model="gpt-4",
//	    input=[
//	        {"role": "user", "content": "Search for AI news"}
//	    ],
//	    tools=[
//	        {
//	            "type": "function",
//	            "function": {
//	                "name": "search_web",
//	                "description": "Search the web"
//	            }
//	        }
//	    ],
//	    mcp_servers=["dedalus-labs/brave-search"],
//	    temperature=0.7,
//	    max_tokens=1000
//	)
//	```
type CompletionRequestParam struct {
	// Frequency penalty (-2 to 2). Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing likelihood of repeated
	// phrases.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// Maximum number of tokens to generate in the completion. Does not include tokens
	// in the input messages.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Maximum number of turns for agent execution before terminating (default: 10).
	// Each turn represents one model inference cycle. Higher values allow more complex
	// reasoning but increase cost and latency.
	MaxTurns param.Opt[int64] `json:"max_turns,omitzero"`
	// Number of completions to generate. Note: only n=1 is currently supported.
	N param.Opt[int64] `json:"n,omitzero"`
	// Presence penalty (-2 to 2). Positive values penalize new tokens based on whether
	// they appear in the text so far, encouraging the model to talk about new topics.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// Whether to stream back partial message deltas as Server-Sent Events. When true,
	// partial message deltas will be sent as chunks in OpenAI format.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Sampling temperature (0 to 2). Higher values make output more random, lower
	// values make it more focused and deterministic. 0 = deterministic, 1 = balanced,
	// 2 = very creative.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Nucleus sampling parameter (0 to 1). Alternative to temperature. 0.1 = only top
	// 10% probability mass, 1.0 = consider all tokens.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Unique identifier representing your end-user. Used for monitoring and abuse
	// detection. Should be consistent across requests from the same user.
	User param.Opt[string] `json:"user,omitzero"`
	// Attributes for the agent itself, influencing behavior and model selection.
	// Format: {'attribute': value}, where values are 0.0-1.0. Common attributes:
	// 'complexity', 'accuracy', 'efficiency', 'creativity', 'friendliness'. Higher
	// values indicate stronger preference for that characteristic.
	AgentAttributes map[string]float64 `json:"agent_attributes,omitzero"`
	// Guardrails to apply to the agent for input/output validation and safety checks.
	// Reserved for future use - guardrails configuration format not yet finalized.
	Guardrails []map[string]any `json:"guardrails,omitzero"`
	// Configuration for multi-model handoffs and agent orchestration. Reserved for
	// future use - handoff configuration format not yet finalized.
	HandoffConfig map[string]any `json:"handoff_config,omitzero"`
	// Input to the model - can be messages, images, or other modalities. Supports
	// OpenAI chat format with role/content structure. For multimodal inputs, content
	// can include text, images, or other media types.
	Input []map[string]any `json:"input,omitzero"`
	// Modify likelihood of specified tokens appearing in the completion. Maps token
	// IDs (as strings) to bias values (-100 to 100). -100 = completely ban token, +100
	// = strongly favor token.
	LogitBias map[string]int64 `json:"logit_bias,omitzero"`
	// MCP (Model Context Protocol) server addresses to make available for server-side
	// tool execution. Can be URLs (e.g., 'https://mcp.example.com') or slugs (e.g.,
	// 'dedalus-labs/brave-search'). MCP tools are executed server-side and billed
	// separately.
	McpServers []string `json:"mcp_servers,omitzero"`
	// Model(s) to use for completion. Can be a single model ID or a list for
	// multi-model routing. Single model: 'gpt-4', 'claude-3-5-sonnet-20241022',
	// 'gpt-4o-mini'. Multi-model routing: ['gpt-4o-mini', 'gpt-4',
	// 'claude-3-5-sonnet'] - agent will choose optimal model based on task complexity.
	Model CompletionRequestModelUnionParam `json:"model,omitzero"`
	// Attributes for individual models used in routing decisions during multi-model
	// execution. Format: {'model_name': {'attribute': value}}, where values are
	// 0.0-1.0. Common attributes: 'intelligence', 'speed', 'cost', 'creativity',
	// 'accuracy'. Used by agent to select optimal model based on task requirements.
	ModelAttributes map[string]map[string]float64 `json:"model_attributes,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The model
	// will stop as soon as it encounters any of these sequences.
	Stop []string `json:"stop,omitzero"`
	// Controls which tool is called by the model. Options: 'auto' (default), 'none',
	// 'required', or specific tool name. Can also be a dict specifying a particular
	// tool.
	ToolChoice CompletionRequestToolChoiceUnionParam `json:"tool_choice,omitzero"`
	// List of tools available to the model in OpenAI function calling format. Tools
	// are executed client-side and returned as JSON for the application to handle. Use
	// 'mcp_servers' for server-side tool execution.
	Tools []map[string]any `json:"tools,omitzero"`
	paramObj
}

func (r CompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CompletionRequestModelUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CompletionRequestModelUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CompletionRequestModelUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CompletionRequestModelUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CompletionRequestToolChoiceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u CompletionRequestToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *CompletionRequestToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CompletionRequestToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

type ChatNewParams struct {
	// Request model for chat completions.
	//
	// Validates incoming chat requests with support for multimodality, multi-model
	// routing, and agent-enhanced features. Compatible with OpenAI API format while
	// extending functionality for advanced use cases.
	//
	// This model supports both the OpenAI-standard 'messages' field and the
	// Dedalus-specific 'input' field for maximum compatibility. The 'input' field can
	// handle various modalities beyond text messages.
	//
	// Key Features: - Multi-model routing with intelligent handoffs - MCP (Model
	// Context Protocol) server integration - Advanced agent attributes for routing
	// decisions - Client-side and server-side tool execution - Streaming and
	// non-streaming responses - Automatic usage tracking and billing
	//
	// Examples: Basic chat completion:
	// `python request = ChatCompletionRequest( model="gpt-4", input=[ {"role": "user", "content": "Hello, how are you?"} ] ) `
	//
	//	Multi-model routing with attributes:
	//	```python
	//	request = ChatCompletionRequest(
	//	    model=["gpt-4o-mini", "gpt-4", "claude-3-5-sonnet"],
	//	    input=[
	//	        {"role": "user", "content": "Analyze this complex problem"}
	//	    ],
	//	    agent_attributes={
	//	        "complexity": 0.8,
	//	        "accuracy": 0.9
	//	    },
	//	    model_attributes={
	//	        "gpt-4": {"intelligence": 0.9, "cost": 0.8},
	//	        "claude-3-5-sonnet": {"intelligence": 0.95, "cost": 0.7}
	//	    }
	//	)
	//	```
	//
	//	With tools and MCP servers:
	//	```python
	//	request = ChatCompletionRequest(
	//	    model="gpt-4",
	//	    input=[
	//	        {"role": "user", "content": "Search for AI news"}
	//	    ],
	//	    tools=[
	//	        {
	//	            "type": "function",
	//	            "function": {
	//	                "name": "search_web",
	//	                "description": "Search the web"
	//	            }
	//	        }
	//	    ],
	//	    mcp_servers=["dedalus-labs/brave-search"],
	//	    temperature=0.7,
	//	    max_tokens=1000
	//	)
	//	```
	CompletionRequest CompletionRequestParam
	paramObj
}

func (r ChatNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.CompletionRequest)
}
func (r *ChatNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CompletionRequest)
}
