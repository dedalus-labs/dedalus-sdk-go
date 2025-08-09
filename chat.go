// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"net/http"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/param"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/ssestream"
)

// ChatService contains methods and other services that help with interacting with
// the Dedalus API.
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
func NewChatService(opts ...option.RequestOption) (r *ChatService) {
	r = &ChatService{}
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
func (r *ChatService) NewStreaming(ctx context.Context, body ChatNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[StreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	path := "v1/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[StreamChunk](ssestream.NewDecoder(raw), err)
}

type Completion struct {
	ID                string                 `json:"id,required"`
	Choices           []CompletionChoice     `json:"choices,required"`
	Created           int64                  `json:"created,required"`
	Model             string                 `json:"model,required"`
	Object            CompletionObject       `json:"object,required"`
	ServiceTier       CompletionServiceTier  `json:"service_tier,nullable"`
	SystemFingerprint string                 `json:"system_fingerprint,nullable"`
	Usage             CompletionUsage        `json:"usage,nullable"`
	ExtraFields       map[string]interface{} `json:"-,extras"`
	JSON              completionJSON         `json:"-"`
}

// completionJSON contains the JSON metadata for the struct [Completion]
type completionJSON struct {
	ID                apijson.Field
	Choices           apijson.Field
	Created           apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	ServiceTier       apijson.Field
	SystemFingerprint apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Completion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionJSON) RawJSON() string {
	return r.raw
}

type CompletionChoice struct {
	FinishReason CompletionChoicesFinishReason `json:"finish_reason,required"`
	Index        int64                         `json:"index,required"`
	Message      CompletionChoicesMessage      `json:"message,required"`
	Logprobs     CompletionChoicesLogprobs     `json:"logprobs,nullable"`
	ExtraFields  map[string]interface{}        `json:"-,extras"`
	JSON         completionChoiceJSON          `json:"-"`
}

// completionChoiceJSON contains the JSON metadata for the struct
// [CompletionChoice]
type completionChoiceJSON struct {
	FinishReason apijson.Field
	Index        apijson.Field
	Message      apijson.Field
	Logprobs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoiceJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesFinishReason string

const (
	CompletionChoicesFinishReasonStop          CompletionChoicesFinishReason = "stop"
	CompletionChoicesFinishReasonLength        CompletionChoicesFinishReason = "length"
	CompletionChoicesFinishReasonToolCalls     CompletionChoicesFinishReason = "tool_calls"
	CompletionChoicesFinishReasonContentFilter CompletionChoicesFinishReason = "content_filter"
	CompletionChoicesFinishReasonFunctionCall  CompletionChoicesFinishReason = "function_call"
)

func (r CompletionChoicesFinishReason) IsKnown() bool {
	switch r {
	case CompletionChoicesFinishReasonStop, CompletionChoicesFinishReasonLength, CompletionChoicesFinishReasonToolCalls, CompletionChoicesFinishReasonContentFilter, CompletionChoicesFinishReasonFunctionCall:
		return true
	}
	return false
}

type CompletionChoicesMessage struct {
	Role         CompletionChoicesMessageRole         `json:"role,required"`
	Annotations  []CompletionChoicesMessageAnnotation `json:"annotations,nullable"`
	Audio        CompletionChoicesMessageAudio        `json:"audio,nullable"`
	Content      string                               `json:"content,nullable"`
	FunctionCall CompletionChoicesMessageFunctionCall `json:"function_call,nullable"`
	Refusal      string                               `json:"refusal,nullable"`
	ToolCalls    []CompletionChoicesMessageToolCall   `json:"tool_calls,nullable"`
	ExtraFields  map[string]interface{}               `json:"-,extras"`
	JSON         completionChoicesMessageJSON         `json:"-"`
}

// completionChoicesMessageJSON contains the JSON metadata for the struct
// [CompletionChoicesMessage]
type completionChoicesMessageJSON struct {
	Role         apijson.Field
	Annotations  apijson.Field
	Audio        apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Refusal      apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionChoicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageRole string

const (
	CompletionChoicesMessageRoleAssistant CompletionChoicesMessageRole = "assistant"
)

func (r CompletionChoicesMessageRole) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageRoleAssistant:
		return true
	}
	return false
}

type CompletionChoicesMessageAnnotation struct {
	Type        CompletionChoicesMessageAnnotationsType        `json:"type,required"`
	URLCitation CompletionChoicesMessageAnnotationsURLCitation `json:"url_citation,required"`
	ExtraFields map[string]interface{}                         `json:"-,extras"`
	JSON        completionChoicesMessageAnnotationJSON         `json:"-"`
}

// completionChoicesMessageAnnotationJSON contains the JSON metadata for the struct
// [CompletionChoicesMessageAnnotation]
type completionChoicesMessageAnnotationJSON struct {
	Type        apijson.Field
	URLCitation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageAnnotationJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageAnnotationsType string

const (
	CompletionChoicesMessageAnnotationsTypeURLCitation CompletionChoicesMessageAnnotationsType = "url_citation"
)

func (r CompletionChoicesMessageAnnotationsType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageAnnotationsTypeURLCitation:
		return true
	}
	return false
}

type CompletionChoicesMessageAnnotationsURLCitation struct {
	EndIndex    int64                                              `json:"end_index,required"`
	StartIndex  int64                                              `json:"start_index,required"`
	Title       string                                             `json:"title,required"`
	URL         string                                             `json:"url,required"`
	ExtraFields map[string]interface{}                             `json:"-,extras"`
	JSON        completionChoicesMessageAnnotationsURLCitationJSON `json:"-"`
}

// completionChoicesMessageAnnotationsURLCitationJSON contains the JSON metadata
// for the struct [CompletionChoicesMessageAnnotationsURLCitation]
type completionChoicesMessageAnnotationsURLCitationJSON struct {
	EndIndex    apijson.Field
	StartIndex  apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageAnnotationsURLCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageAnnotationsURLCitationJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageAudio struct {
	ID          string                            `json:"id,required"`
	Data        string                            `json:"data,required"`
	ExpiresAt   int64                             `json:"expires_at,required"`
	Transcript  string                            `json:"transcript,required"`
	ExtraFields map[string]interface{}            `json:"-,extras"`
	JSON        completionChoicesMessageAudioJSON `json:"-"`
}

// completionChoicesMessageAudioJSON contains the JSON metadata for the struct
// [CompletionChoicesMessageAudio]
type completionChoicesMessageAudioJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	ExpiresAt   apijson.Field
	Transcript  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageAudioJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageFunctionCall struct {
	Arguments   string                                   `json:"arguments,required"`
	Name        string                                   `json:"name,required"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
	JSON        completionChoicesMessageFunctionCallJSON `json:"-"`
}

// completionChoicesMessageFunctionCallJSON contains the JSON metadata for the
// struct [CompletionChoicesMessageFunctionCall]
type completionChoicesMessageFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageFunctionCallJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageToolCall struct {
	ID          string                                    `json:"id,required"`
	Function    CompletionChoicesMessageToolCallsFunction `json:"function,required"`
	Type        CompletionChoicesMessageToolCallsType     `json:"type,required"`
	ExtraFields map[string]interface{}                    `json:"-,extras"`
	JSON        completionChoicesMessageToolCallJSON      `json:"-"`
}

// completionChoicesMessageToolCallJSON contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCall]
type completionChoicesMessageToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageToolCallsFunction struct {
	Arguments   string                                        `json:"arguments,required"`
	Name        string                                        `json:"name,required"`
	ExtraFields map[string]interface{}                        `json:"-,extras"`
	JSON        completionChoicesMessageToolCallsFunctionJSON `json:"-"`
}

// completionChoicesMessageToolCallsFunctionJSON contains the JSON metadata for the
// struct [CompletionChoicesMessageToolCallsFunction]
type completionChoicesMessageToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageToolCallsType string

const (
	CompletionChoicesMessageToolCallsTypeFunction CompletionChoicesMessageToolCallsType = "function"
)

func (r CompletionChoicesMessageToolCallsType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageToolCallsTypeFunction:
		return true
	}
	return false
}

type CompletionChoicesLogprobs struct {
	Content     []CompletionChoicesLogprobsContent `json:"content,nullable"`
	Refusal     []CompletionChoicesLogprobsRefusal `json:"refusal,nullable"`
	ExtraFields map[string]interface{}             `json:"-,extras"`
	JSON        completionChoicesLogprobsJSON      `json:"-"`
}

// completionChoicesLogprobsJSON contains the JSON metadata for the struct
// [CompletionChoicesLogprobs]
type completionChoicesLogprobsJSON struct {
	Content     apijson.Field
	Refusal     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesLogprobsContent struct {
	Token       string                                       `json:"token,required"`
	Logprob     float64                                      `json:"logprob,required"`
	TopLogprobs []CompletionChoicesLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                      `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
	JSON        completionChoicesLogprobsContentJSON         `json:"-"`
}

// completionChoicesLogprobsContentJSON contains the JSON metadata for the struct
// [CompletionChoicesLogprobsContent]
type completionChoicesLogprobsContentJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsContentJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesLogprobsContentTopLogprob struct {
	Token       string                                         `json:"token,required"`
	Logprob     float64                                        `json:"logprob,required"`
	Bytes       []int64                                        `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                         `json:"-,extras"`
	JSON        completionChoicesLogprobsContentTopLogprobJSON `json:"-"`
}

// completionChoicesLogprobsContentTopLogprobJSON contains the JSON metadata for
// the struct [CompletionChoicesLogprobsContentTopLogprob]
type completionChoicesLogprobsContentTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsContentTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsContentTopLogprobJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesLogprobsRefusal struct {
	Token       string                                       `json:"token,required"`
	Logprob     float64                                      `json:"logprob,required"`
	TopLogprobs []CompletionChoicesLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                      `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
	JSON        completionChoicesLogprobsRefusalJSON         `json:"-"`
}

// completionChoicesLogprobsRefusalJSON contains the JSON metadata for the struct
// [CompletionChoicesLogprobsRefusal]
type completionChoicesLogprobsRefusalJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsRefusal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsRefusalJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesLogprobsRefusalTopLogprob struct {
	Token       string                                         `json:"token,required"`
	Logprob     float64                                        `json:"logprob,required"`
	Bytes       []int64                                        `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                         `json:"-,extras"`
	JSON        completionChoicesLogprobsRefusalTopLogprobJSON `json:"-"`
}

// completionChoicesLogprobsRefusalTopLogprobJSON contains the JSON metadata for
// the struct [CompletionChoicesLogprobsRefusalTopLogprob]
type completionChoicesLogprobsRefusalTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsRefusalTopLogprobJSON) RawJSON() string {
	return r.raw
}

type CompletionObject string

const (
	CompletionObjectChatCompletion CompletionObject = "chat.completion"
)

func (r CompletionObject) IsKnown() bool {
	switch r {
	case CompletionObjectChatCompletion:
		return true
	}
	return false
}

type CompletionServiceTier string

const (
	CompletionServiceTierAuto     CompletionServiceTier = "auto"
	CompletionServiceTierDefault  CompletionServiceTier = "default"
	CompletionServiceTierFlex     CompletionServiceTier = "flex"
	CompletionServiceTierScale    CompletionServiceTier = "scale"
	CompletionServiceTierPriority CompletionServiceTier = "priority"
)

func (r CompletionServiceTier) IsKnown() bool {
	switch r {
	case CompletionServiceTierAuto, CompletionServiceTierDefault, CompletionServiceTierFlex, CompletionServiceTierScale, CompletionServiceTierPriority:
		return true
	}
	return false
}

type CompletionUsage struct {
	CompletionTokens        int64                                  `json:"completion_tokens,required"`
	PromptTokens            int64                                  `json:"prompt_tokens,required"`
	TotalTokens             int64                                  `json:"total_tokens,required"`
	CompletionTokensDetails CompletionUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	PromptTokensDetails     CompletionUsagePromptTokensDetails     `json:"prompt_tokens_details,nullable"`
	ExtraFields             map[string]interface{}                 `json:"-,extras"`
	JSON                    completionUsageJSON                    `json:"-"`
}

// completionUsageJSON contains the JSON metadata for the struct [CompletionUsage]
type completionUsageJSON struct {
	CompletionTokens        apijson.Field
	PromptTokens            apijson.Field
	TotalTokens             apijson.Field
	CompletionTokensDetails apijson.Field
	PromptTokensDetails     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CompletionUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionUsageJSON) RawJSON() string {
	return r.raw
}

type CompletionUsageCompletionTokensDetails struct {
	AcceptedPredictionTokens int64                                      `json:"accepted_prediction_tokens,nullable"`
	AudioTokens              int64                                      `json:"audio_tokens,nullable"`
	ReasoningTokens          int64                                      `json:"reasoning_tokens,nullable"`
	RejectedPredictionTokens int64                                      `json:"rejected_prediction_tokens,nullable"`
	ExtraFields              map[string]interface{}                     `json:"-,extras"`
	JSON                     completionUsageCompletionTokensDetailsJSON `json:"-"`
}

// completionUsageCompletionTokensDetailsJSON contains the JSON metadata for the
// struct [CompletionUsageCompletionTokensDetails]
type completionUsageCompletionTokensDetailsJSON struct {
	AcceptedPredictionTokens apijson.Field
	AudioTokens              apijson.Field
	ReasoningTokens          apijson.Field
	RejectedPredictionTokens apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CompletionUsageCompletionTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionUsageCompletionTokensDetailsJSON) RawJSON() string {
	return r.raw
}

type CompletionUsagePromptTokensDetails struct {
	AudioTokens  int64                                  `json:"audio_tokens,nullable"`
	CachedTokens int64                                  `json:"cached_tokens,nullable"`
	ExtraFields  map[string]interface{}                 `json:"-,extras"`
	JSON         completionUsagePromptTokensDetailsJSON `json:"-"`
}

// completionUsagePromptTokensDetailsJSON contains the JSON metadata for the struct
// [CompletionUsagePromptTokensDetails]
type completionUsagePromptTokensDetailsJSON struct {
	AudioTokens  apijson.Field
	CachedTokens apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionUsagePromptTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionUsagePromptTokensDetailsJSON) RawJSON() string {
	return r.raw
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
	// Attributes for the agent itself, influencing behavior and model selection.
	// Format: {'attribute': value}, where values are 0.0-1.0. Common attributes:
	// 'complexity', 'accuracy', 'efficiency', 'creativity', 'friendliness'. Higher
	// values indicate stronger preference for that characteristic.
	AgentAttributes param.Field[map[string]float64] `json:"agent_attributes"`
	// Frequency penalty (-2 to 2). Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing likelihood of repeated
	// phrases.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Guardrails to apply to the agent for input/output validation and safety checks.
	// Reserved for future use - guardrails configuration format not yet finalized.
	Guardrails param.Field[[]map[string]interface{}] `json:"guardrails"`
	// Configuration for multi-model handoffs and agent orchestration. Reserved for
	// future use - handoff configuration format not yet finalized.
	HandoffConfig param.Field[map[string]interface{}] `json:"handoff_config"`
	// Input to the model - can be messages, images, or other modalities. Supports
	// OpenAI chat format with role/content structure. For multimodal inputs, content
	// can include text, images, or other media types.
	Input param.Field[[]map[string]interface{}] `json:"input"`
	// Modify likelihood of specified tokens appearing in the completion. Maps token
	// IDs (as strings) to bias values (-100 to 100). -100 = completely ban token, +100
	// = strongly favor token.
	LogitBias param.Field[map[string]int64] `json:"logit_bias"`
	// Maximum number of tokens to generate in the completion. Does not include tokens
	// in the input messages.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Maximum number of turns for agent execution before terminating (default: 10).
	// Each turn represents one model inference cycle. Higher values allow more complex
	// reasoning but increase cost and latency.
	MaxTurns param.Field[int64] `json:"max_turns"`
	// MCP (Model Context Protocol) server addresses to make available for server-side
	// tool execution. Can be URLs (e.g., 'https://mcp.example.com') or slugs (e.g.,
	// 'dedalus-labs/brave-search'). MCP tools are executed server-side and billed
	// separately.
	McpServers param.Field[[]string] `json:"mcp_servers"`
	// Model(s) to use for completion. Can be a single model ID, a ModelConfig object
	// (optionally with per-model 'settings'), or a list for multi-model routing.
	// Single model: 'gpt-4', 'claude-3-5-sonnet-20241022', 'gpt-4o-mini', or a Model
	// instance. Multi-model routing: ['gpt-4o-mini', 'gpt-4', 'claude-3-5-sonnet'] or
	// list of ModelConfig objects - agent will choose optimal model based on task
	// complexity.
	Model param.Field[CompletionRequestModelUnionParam] `json:"model"`
	// Attributes for individual models used in routing decisions during multi-model
	// execution. Format: {'model_name': {'attribute': value}}, where values are
	// 0.0-1.0. Common attributes: 'intelligence', 'speed', 'cost', 'creativity',
	// 'accuracy'. Used by agent to select optimal model based on task requirements.
	ModelAttributes param.Field[map[string]map[string]float64] `json:"model_attributes"`
	// Number of completions to generate. Note: only n=1 is currently supported.
	N param.Field[int64] `json:"n"`
	// Presence penalty (-2 to 2). Positive values penalize new tokens based on whether
	// they appear in the text so far, encouraging the model to talk about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Up to 4 sequences where the API will stop generating further tokens. The model
	// will stop as soon as it encounters any of these sequences.
	Stop param.Field[[]string] `json:"stop"`
	// Whether to stream back partial message deltas as Server-Sent Events. When true,
	// partial message deltas will be sent as chunks in OpenAI format.
	Stream param.Field[bool] `json:"stream"`
	// Sampling temperature (0 to 2). Higher values make output more random, lower
	// values make it more focused and deterministic. 0 = deterministic, 1 = balanced,
	// 2 = very creative.
	Temperature param.Field[float64] `json:"temperature"`
	// Controls which tool is called by the model. Options: 'auto' (default), 'none',
	// 'required', or specific tool name. Can also be a dict specifying a particular
	// tool.
	ToolChoice param.Field[CompletionRequestToolChoiceUnionParam] `json:"tool_choice"`
	// List of tools available to the model in OpenAI function calling format. Tools
	// are executed client-side and returned as JSON for the application to handle. Use
	// 'mcp_servers' for server-side tool execution.
	Tools param.Field[[]map[string]interface{}] `json:"tools"`
	// Nucleus sampling parameter (0 to 1). Alternative to temperature. 0.1 = only top
	// 10% probability mass, 1.0 = consider all tokens.
	TopP param.Field[float64] `json:"top_p"`
	// Unique identifier representing your end-user. Used for monitoring and abuse
	// detection. Should be consistent across requests from the same user.
	User param.Field[string] `json:"user"`
}

func (r CompletionRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Model(s) to use for completion. Can be a single model ID, a ModelConfig object
// (optionally with per-model 'settings'), or a list for multi-model routing.
// Single model: 'gpt-4', 'claude-3-5-sonnet-20241022', 'gpt-4o-mini', or a Model
// instance. Multi-model routing: ['gpt-4o-mini', 'gpt-4', 'claude-3-5-sonnet'] or
// list of ModelConfig objects - agent will choose optimal model based on task
// complexity.
//
// Satisfied by [shared.UnionString], [CompletionRequestModelModelIDListParam],
// [CompletionRequestModelModelConfigParam],
// [CompletionRequestModelModelConfigListParam].
type CompletionRequestModelUnionParam interface {
	ImplementsCompletionRequestModelUnionParam()
}

type CompletionRequestModelModelIDListParam []string

func (r CompletionRequestModelModelIDListParam) ImplementsCompletionRequestModelUnionParam() {}

// A model with optional routing attributes and per-model settings.
type CompletionRequestModelModelConfigParam struct {
	// Model identifier, e.g. 'gpt-4' or 'claude-3-5-sonnet-20241022'.
	Name param.Field[string] `json:"name,required"`
	// Numeric attributes used by routing (0.0–1.0), e.g. intelligence, speed, cost.
	Attributes param.Field[map[string]float64] `json:"attributes"`
	// Per-model generation settings like temperature, max_tokens, etc.
	Settings param.Field[map[string]interface{}] `json:"settings"`
}

func (r CompletionRequestModelModelConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CompletionRequestModelModelConfigParam) ImplementsCompletionRequestModelUnionParam() {}

type CompletionRequestModelModelConfigListParam []CompletionRequestModelModelConfigListItemParam

func (r CompletionRequestModelModelConfigListParam) ImplementsCompletionRequestModelUnionParam() {}

// A model with optional routing attributes and per-model settings.
type CompletionRequestModelModelConfigListItemParam struct {
	// Model identifier, e.g. 'gpt-4' or 'claude-3-5-sonnet-20241022'.
	Name param.Field[string] `json:"name,required"`
	// Numeric attributes used by routing (0.0–1.0), e.g. intelligence, speed, cost.
	Attributes param.Field[map[string]float64] `json:"attributes"`
	// Per-model generation settings like temperature, max_tokens, etc.
	Settings param.Field[map[string]interface{}] `json:"settings"`
}

func (r CompletionRequestModelModelConfigListItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls which tool is called by the model. Options: 'auto' (default), 'none',
// 'required', or specific tool name. Can also be a dict specifying a particular
// tool.
//
// Satisfied by [shared.UnionString], [CompletionRequestToolChoiceMapParam].
type CompletionRequestToolChoiceUnionParam interface {
	ImplementsCompletionRequestToolChoiceUnionParam()
}

type CompletionRequestToolChoiceMapParam map[string]interface{}

func (r CompletionRequestToolChoiceMapParam) ImplementsCompletionRequestToolChoiceUnionParam() {}

type StreamChunk struct {
	ID                string                 `json:"id,required"`
	Choices           []StreamChunkChoice    `json:"choices,required"`
	Created           int64                  `json:"created,required"`
	Model             string                 `json:"model,required"`
	Object            StreamChunkObject      `json:"object,required"`
	ServiceTier       StreamChunkServiceTier `json:"service_tier,nullable"`
	SystemFingerprint string                 `json:"system_fingerprint,nullable"`
	Usage             StreamChunkUsage       `json:"usage,nullable"`
	ExtraFields       map[string]interface{} `json:"-,extras"`
	JSON              streamChunkJSON        `json:"-"`
}

// streamChunkJSON contains the JSON metadata for the struct [StreamChunk]
type streamChunkJSON struct {
	ID                apijson.Field
	Choices           apijson.Field
	Created           apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	ServiceTier       apijson.Field
	SystemFingerprint apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoice struct {
	Delta        StreamChunkChoicesDelta        `json:"delta,required"`
	Index        int64                          `json:"index,required"`
	FinishReason StreamChunkChoicesFinishReason `json:"finish_reason,nullable"`
	Logprobs     StreamChunkChoicesLogprobs     `json:"logprobs,nullable"`
	ExtraFields  map[string]interface{}         `json:"-,extras"`
	JSON         streamChunkChoiceJSON          `json:"-"`
}

// streamChunkChoiceJSON contains the JSON metadata for the struct
// [StreamChunkChoice]
type streamChunkChoiceJSON struct {
	Delta        apijson.Field
	Index        apijson.Field
	FinishReason apijson.Field
	Logprobs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StreamChunkChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoiceJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDelta struct {
	Content      string                              `json:"content,nullable"`
	FunctionCall StreamChunkChoicesDeltaFunctionCall `json:"function_call,nullable"`
	Refusal      string                              `json:"refusal,nullable"`
	Role         StreamChunkChoicesDeltaRole         `json:"role,nullable"`
	ToolCalls    []StreamChunkChoicesDeltaToolCall   `json:"tool_calls,nullable"`
	ExtraFields  map[string]interface{}              `json:"-,extras"`
	JSON         streamChunkChoicesDeltaJSON         `json:"-"`
}

// streamChunkChoicesDeltaJSON contains the JSON metadata for the struct
// [StreamChunkChoicesDelta]
type streamChunkChoicesDeltaJSON struct {
	Content      apijson.Field
	FunctionCall apijson.Field
	Refusal      apijson.Field
	Role         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StreamChunkChoicesDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaFunctionCall struct {
	Arguments   string                                  `json:"arguments,nullable"`
	Name        string                                  `json:"name,nullable"`
	ExtraFields map[string]interface{}                  `json:"-,extras"`
	JSON        streamChunkChoicesDeltaFunctionCallJSON `json:"-"`
}

// streamChunkChoicesDeltaFunctionCallJSON contains the JSON metadata for the
// struct [StreamChunkChoicesDeltaFunctionCall]
type streamChunkChoicesDeltaFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesDeltaFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaFunctionCallJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaRole string

const (
	StreamChunkChoicesDeltaRoleDeveloper StreamChunkChoicesDeltaRole = "developer"
	StreamChunkChoicesDeltaRoleSystem    StreamChunkChoicesDeltaRole = "system"
	StreamChunkChoicesDeltaRoleUser      StreamChunkChoicesDeltaRole = "user"
	StreamChunkChoicesDeltaRoleAssistant StreamChunkChoicesDeltaRole = "assistant"
	StreamChunkChoicesDeltaRoleTool      StreamChunkChoicesDeltaRole = "tool"
)

func (r StreamChunkChoicesDeltaRole) IsKnown() bool {
	switch r {
	case StreamChunkChoicesDeltaRoleDeveloper, StreamChunkChoicesDeltaRoleSystem, StreamChunkChoicesDeltaRoleUser, StreamChunkChoicesDeltaRoleAssistant, StreamChunkChoicesDeltaRoleTool:
		return true
	}
	return false
}

type StreamChunkChoicesDeltaToolCall struct {
	Index       int64                                    `json:"index,required"`
	ID          string                                   `json:"id,nullable"`
	Function    StreamChunkChoicesDeltaToolCallsFunction `json:"function,nullable"`
	Type        StreamChunkChoicesDeltaToolCallsType     `json:"type,nullable"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
	JSON        streamChunkChoicesDeltaToolCallJSON      `json:"-"`
}

// streamChunkChoicesDeltaToolCallJSON contains the JSON metadata for the struct
// [StreamChunkChoicesDeltaToolCall]
type streamChunkChoicesDeltaToolCallJSON struct {
	Index       apijson.Field
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesDeltaToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaToolCallJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaToolCallsFunction struct {
	Arguments   string                                       `json:"arguments,nullable"`
	Name        string                                       `json:"name,nullable"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
	JSON        streamChunkChoicesDeltaToolCallsFunctionJSON `json:"-"`
}

// streamChunkChoicesDeltaToolCallsFunctionJSON contains the JSON metadata for the
// struct [StreamChunkChoicesDeltaToolCallsFunction]
type streamChunkChoicesDeltaToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesDeltaToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaToolCallsType string

const (
	StreamChunkChoicesDeltaToolCallsTypeFunction StreamChunkChoicesDeltaToolCallsType = "function"
)

func (r StreamChunkChoicesDeltaToolCallsType) IsKnown() bool {
	switch r {
	case StreamChunkChoicesDeltaToolCallsTypeFunction:
		return true
	}
	return false
}

type StreamChunkChoicesFinishReason string

const (
	StreamChunkChoicesFinishReasonStop          StreamChunkChoicesFinishReason = "stop"
	StreamChunkChoicesFinishReasonLength        StreamChunkChoicesFinishReason = "length"
	StreamChunkChoicesFinishReasonToolCalls     StreamChunkChoicesFinishReason = "tool_calls"
	StreamChunkChoicesFinishReasonContentFilter StreamChunkChoicesFinishReason = "content_filter"
	StreamChunkChoicesFinishReasonFunctionCall  StreamChunkChoicesFinishReason = "function_call"
)

func (r StreamChunkChoicesFinishReason) IsKnown() bool {
	switch r {
	case StreamChunkChoicesFinishReasonStop, StreamChunkChoicesFinishReasonLength, StreamChunkChoicesFinishReasonToolCalls, StreamChunkChoicesFinishReasonContentFilter, StreamChunkChoicesFinishReasonFunctionCall:
		return true
	}
	return false
}

type StreamChunkChoicesLogprobs struct {
	Content     []StreamChunkChoicesLogprobsContent `json:"content,nullable"`
	Refusal     []StreamChunkChoicesLogprobsRefusal `json:"refusal,nullable"`
	ExtraFields map[string]interface{}              `json:"-,extras"`
	JSON        streamChunkChoicesLogprobsJSON      `json:"-"`
}

// streamChunkChoicesLogprobsJSON contains the JSON metadata for the struct
// [StreamChunkChoicesLogprobs]
type streamChunkChoicesLogprobsJSON struct {
	Content     apijson.Field
	Refusal     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesLogprobsContent struct {
	Token       string                                        `json:"token,required"`
	Logprob     float64                                       `json:"logprob,required"`
	TopLogprobs []StreamChunkChoicesLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                       `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                        `json:"-,extras"`
	JSON        streamChunkChoicesLogprobsContentJSON         `json:"-"`
}

// streamChunkChoicesLogprobsContentJSON contains the JSON metadata for the struct
// [StreamChunkChoicesLogprobsContent]
type streamChunkChoicesLogprobsContentJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsContentJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesLogprobsContentTopLogprob struct {
	Token       string                                          `json:"token,required"`
	Logprob     float64                                         `json:"logprob,required"`
	Bytes       []int64                                         `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                          `json:"-,extras"`
	JSON        streamChunkChoicesLogprobsContentTopLogprobJSON `json:"-"`
}

// streamChunkChoicesLogprobsContentTopLogprobJSON contains the JSON metadata for
// the struct [StreamChunkChoicesLogprobsContentTopLogprob]
type streamChunkChoicesLogprobsContentTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsContentTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsContentTopLogprobJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesLogprobsRefusal struct {
	Token       string                                        `json:"token,required"`
	Logprob     float64                                       `json:"logprob,required"`
	TopLogprobs []StreamChunkChoicesLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                       `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                        `json:"-,extras"`
	JSON        streamChunkChoicesLogprobsRefusalJSON         `json:"-"`
}

// streamChunkChoicesLogprobsRefusalJSON contains the JSON metadata for the struct
// [StreamChunkChoicesLogprobsRefusal]
type streamChunkChoicesLogprobsRefusalJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsRefusal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsRefusalJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesLogprobsRefusalTopLogprob struct {
	Token       string                                          `json:"token,required"`
	Logprob     float64                                         `json:"logprob,required"`
	Bytes       []int64                                         `json:"bytes,nullable"`
	ExtraFields map[string]interface{}                          `json:"-,extras"`
	JSON        streamChunkChoicesLogprobsRefusalTopLogprobJSON `json:"-"`
}

// streamChunkChoicesLogprobsRefusalTopLogprobJSON contains the JSON metadata for
// the struct [StreamChunkChoicesLogprobsRefusalTopLogprob]
type streamChunkChoicesLogprobsRefusalTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsRefusalTopLogprobJSON) RawJSON() string {
	return r.raw
}

type StreamChunkObject string

const (
	StreamChunkObjectChatCompletionChunk StreamChunkObject = "chat.completion.chunk"
)

func (r StreamChunkObject) IsKnown() bool {
	switch r {
	case StreamChunkObjectChatCompletionChunk:
		return true
	}
	return false
}

type StreamChunkServiceTier string

const (
	StreamChunkServiceTierAuto     StreamChunkServiceTier = "auto"
	StreamChunkServiceTierDefault  StreamChunkServiceTier = "default"
	StreamChunkServiceTierFlex     StreamChunkServiceTier = "flex"
	StreamChunkServiceTierScale    StreamChunkServiceTier = "scale"
	StreamChunkServiceTierPriority StreamChunkServiceTier = "priority"
)

func (r StreamChunkServiceTier) IsKnown() bool {
	switch r {
	case StreamChunkServiceTierAuto, StreamChunkServiceTierDefault, StreamChunkServiceTierFlex, StreamChunkServiceTierScale, StreamChunkServiceTierPriority:
		return true
	}
	return false
}

type StreamChunkUsage struct {
	CompletionTokens        int64                                   `json:"completion_tokens,required"`
	PromptTokens            int64                                   `json:"prompt_tokens,required"`
	TotalTokens             int64                                   `json:"total_tokens,required"`
	CompletionTokensDetails StreamChunkUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	PromptTokensDetails     StreamChunkUsagePromptTokensDetails     `json:"prompt_tokens_details,nullable"`
	ExtraFields             map[string]interface{}                  `json:"-,extras"`
	JSON                    streamChunkUsageJSON                    `json:"-"`
}

// streamChunkUsageJSON contains the JSON metadata for the struct
// [StreamChunkUsage]
type streamChunkUsageJSON struct {
	CompletionTokens        apijson.Field
	PromptTokens            apijson.Field
	TotalTokens             apijson.Field
	CompletionTokensDetails apijson.Field
	PromptTokensDetails     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *StreamChunkUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkUsageJSON) RawJSON() string {
	return r.raw
}

type StreamChunkUsageCompletionTokensDetails struct {
	AcceptedPredictionTokens int64                                       `json:"accepted_prediction_tokens,nullable"`
	AudioTokens              int64                                       `json:"audio_tokens,nullable"`
	ReasoningTokens          int64                                       `json:"reasoning_tokens,nullable"`
	RejectedPredictionTokens int64                                       `json:"rejected_prediction_tokens,nullable"`
	ExtraFields              map[string]interface{}                      `json:"-,extras"`
	JSON                     streamChunkUsageCompletionTokensDetailsJSON `json:"-"`
}

// streamChunkUsageCompletionTokensDetailsJSON contains the JSON metadata for the
// struct [StreamChunkUsageCompletionTokensDetails]
type streamChunkUsageCompletionTokensDetailsJSON struct {
	AcceptedPredictionTokens apijson.Field
	AudioTokens              apijson.Field
	ReasoningTokens          apijson.Field
	RejectedPredictionTokens apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *StreamChunkUsageCompletionTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkUsageCompletionTokensDetailsJSON) RawJSON() string {
	return r.raw
}

type StreamChunkUsagePromptTokensDetails struct {
	AudioTokens  int64                                   `json:"audio_tokens,nullable"`
	CachedTokens int64                                   `json:"cached_tokens,nullable"`
	ExtraFields  map[string]interface{}                  `json:"-,extras"`
	JSON         streamChunkUsagePromptTokensDetailsJSON `json:"-"`
}

// streamChunkUsagePromptTokensDetailsJSON contains the JSON metadata for the
// struct [StreamChunkUsagePromptTokensDetails]
type streamChunkUsagePromptTokensDetailsJSON struct {
	AudioTokens  apijson.Field
	CachedTokens apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StreamChunkUsagePromptTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkUsagePromptTokensDetailsJSON) RawJSON() string {
	return r.raw
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
	CompletionRequest CompletionRequestParam `json:"completion_request,required"`
}

func (r ChatNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CompletionRequest)
}
