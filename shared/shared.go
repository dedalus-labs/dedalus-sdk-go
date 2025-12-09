// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/param"
)

// Structured model selection entry used in request payloads.
//
// Supports OpenAI-style semantics (string model id) while enabling optional
// per-model default settings for Dedalus multi-model routing.
type DedalusModelParam struct {
	// Model identifier with provider prefix (e.g., 'openai/gpt-5',
	// 'anthropic/claude-3-5-sonnet').
	Model param.Field[string] `json:"model,required"`
	// Optional default generation settings (e.g., temperature, max_tokens) applied
	// when this model is selected.
	Settings param.Field[DedalusModelSettingsParam] `json:"settings"`
}

func (r DedalusModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DedalusModelParam) ImplementsDedalusModelChoiceUnionParam() {}

func (r DedalusModelParam) ImplementsChatCompletionNewParamsModelUnion() {}

// Optional default generation settings (e.g., temperature, max_tokens) applied
// when this model is selected.
type DedalusModelSettingsParam struct {
	Attributes              param.Field[map[string]interface{}]                                `json:"attributes"`
	Audio                   param.Field[map[string]interface{}]                                `json:"audio"`
	Deferred                param.Field[bool]                                                  `json:"deferred"`
	ExtraArgs               param.Field[map[string]interface{}]                                `json:"extra_args"`
	ExtraHeaders            param.Field[map[string]string]                                     `json:"extra_headers"`
	ExtraQuery              param.Field[map[string]interface{}]                                `json:"extra_query"`
	FrequencyPenalty        param.Field[float64]                                               `json:"frequency_penalty"`
	GenerationConfig        param.Field[map[string]interface{}]                                `json:"generation_config"`
	IncludeUsage            param.Field[bool]                                                  `json:"include_usage"`
	InputAudioFormat        param.Field[string]                                                `json:"input_audio_format"`
	InputAudioTranscription param.Field[map[string]interface{}]                                `json:"input_audio_transcription"`
	LogitBias               param.Field[map[string]int64]                                      `json:"logit_bias"`
	Logprobs                param.Field[bool]                                                  `json:"logprobs"`
	MaxCompletionTokens     param.Field[int64]                                                 `json:"max_completion_tokens"`
	MaxTokens               param.Field[int64]                                                 `json:"max_tokens"`
	Metadata                param.Field[map[string]string]                                     `json:"metadata"`
	Modalities              param.Field[[]string]                                              `json:"modalities"`
	N                       param.Field[int64]                                                 `json:"n"`
	OutputAudioFormat       param.Field[string]                                                `json:"output_audio_format"`
	ParallelToolCalls       param.Field[bool]                                                  `json:"parallel_tool_calls"`
	Prediction              param.Field[map[string]interface{}]                                `json:"prediction"`
	PresencePenalty         param.Field[float64]                                               `json:"presence_penalty"`
	PromptCacheKey          param.Field[string]                                                `json:"prompt_cache_key"`
	Reasoning               param.Field[githubcomdedaluslabsdedalussdkgo.ReasoningParam]       `json:"reasoning"`
	ReasoningEffort         param.Field[string]                                                `json:"reasoning_effort"`
	ResponseFormat          param.Field[map[string]interface{}]                                `json:"response_format"`
	SafetyIdentifier        param.Field[string]                                                `json:"safety_identifier"`
	SafetySettings          param.Field[[]map[string]interface{}]                              `json:"safety_settings"`
	SearchParameters        param.Field[map[string]interface{}]                                `json:"search_parameters"`
	Seed                    param.Field[int64]                                                 `json:"seed"`
	ServiceTier             param.Field[string]                                                `json:"service_tier"`
	Stop                    param.Field[DedalusModelSettingsStopUnionParam]                    `json:"stop"`
	Store                   param.Field[bool]                                                  `json:"store"`
	Stream                  param.Field[bool]                                                  `json:"stream"`
	StreamOptions           param.Field[map[string]interface{}]                                `json:"stream_options"`
	StructuredOutput        param.Field[interface{}]                                           `json:"structured_output"`
	SystemInstruction       param.Field[map[string]interface{}]                                `json:"system_instruction"`
	Temperature             param.Field[float64]                                               `json:"temperature"`
	Thinking                param.Field[map[string]interface{}]                                `json:"thinking"`
	Timeout                 param.Field[float64]                                               `json:"timeout"`
	ToolChoice              param.Field[githubcomdedaluslabsdedalussdkgo.ToolChoiceUnionParam] `json:"tool_choice"`
	ToolConfig              param.Field[map[string]interface{}]                                `json:"tool_config"`
	TopK                    param.Field[int64]                                                 `json:"top_k"`
	TopLogprobs             param.Field[int64]                                                 `json:"top_logprobs"`
	TopP                    param.Field[float64]                                               `json:"top_p"`
	Truncation              param.Field[DedalusModelSettingsTruncation]                        `json:"truncation"`
	TurnDetection           param.Field[map[string]interface{}]                                `json:"turn_detection"`
	User                    param.Field[string]                                                `json:"user"`
	Verbosity               param.Field[string]                                                `json:"verbosity"`
	Voice                   param.Field[string]                                                `json:"voice"`
	WebSearchOptions        param.Field[map[string]interface{}]                                `json:"web_search_options"`
}

func (r DedalusModelSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.DedalusModelSettingsStopArrayParam].
type DedalusModelSettingsStopUnionParam interface {
	ImplementsDedalusModelSettingsStopUnionParam()
}

type DedalusModelSettingsStopArrayParam []string

func (r DedalusModelSettingsStopArrayParam) ImplementsDedalusModelSettingsStopUnionParam() {}

type DedalusModelSettingsTruncation string

const (
	DedalusModelSettingsTruncationAuto     DedalusModelSettingsTruncation = "auto"
	DedalusModelSettingsTruncationDisabled DedalusModelSettingsTruncation = "disabled"
)

func (r DedalusModelSettingsTruncation) IsKnown() bool {
	switch r {
	case DedalusModelSettingsTruncationAuto, DedalusModelSettingsTruncationDisabled:
		return true
	}
	return false
}

// Dedalus model choice - either a string ID or DedalusModel configuration object.
//
// Satisfied by [shared.UnionString], [shared.DedalusModelParam].
type DedalusModelChoiceUnionParam interface {
	ImplementsDedalusModelChoiceUnionParam()
}

// Schema for FunctionObject.
//
// Fields:
//
// - description (optional): str
// - name (required): str
// - parameters (optional): FunctionParameters
// - strict (optional): bool | None
type FunctionDefinitionParam struct {
	// The name of the function to be called. Must be a-z, A-Z, 0-9, or contain
	// underscores and dashes, with a maximum length of 64.
	Name param.Field[string] `json:"name,required"`
	// A description of what the function does, used by the model to choose when and
	// how to call the function.
	Description param.Field[string] `json:"description"`
	// The parameters the functions accepts, described as a JSON Schema object. See the
	// [guide](https://platform.openai.com/docs/guides/function-calling) for examples,
	// and the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for
	// documentation about the format.
	//
	// Omitting `parameters` defines a function with an empty parameter list.
	Parameters param.Field[FunctionParameters] `json:"parameters"`
	// Whether to enable strict schema adherence when generating the function call. If
	// set to true, the model will follow the exact schema defined in the `parameters`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`. Learn
	// more about Structured Outputs in the
	// [function calling guide](https://platform.openai.com/docs/guides/function-calling).
	Strict param.Field[bool] `json:"strict"`
}

func (r FunctionDefinitionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionParameters map[string]interface{}

// Single MCP server input: slug string or structured MCPServerSpec.
//
// Satisfied by [shared.UnionString], [shared.MCPServerSpecParam].
type MCPServerInputUnionParam interface {
	ImplementsMCPServerInputUnionParam()
}

// Structured MCP server specification.
//
// Slug-based: {"slug": "dedalus-labs/brave-search", "version": "v1.0.0"}
// URL-based: {"url": "https://mcp.dedaluslabs.ai/acme/my-server/mcp"}
type MCPServerSpecParam struct {
	// Connection name for credential matching. Must match a key in the client's
	// credentials list.
	Connection param.Field[string] `json:"connection"`
	// Schema declaring what credentials are needed. Maps field names to their bindings
	// (e.g., env var names).
	Credentials param.Field[map[string]MCPServerSpecCredentialsUnionParam] `json:"credentials"`
	// Client-encrypted credential values. Maps connection names to encrypted envelopes
	// (base64url JWE). SDK encrypts credentials client-side using the enclave's public
	// key from authorization server.
	EncryptedCredentials param.Field[map[string]string] `json:"encrypted_credentials"`
	// Marketplace slug.
	Slug param.Field[string] `json:"slug"`
	// Direct URL to MCP server endpoint.
	URL param.Field[string] `json:"url"`
	// Version constraint for slug-based servers.
	Version param.Field[string] `json:"version"`
}

func (r MCPServerSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MCPServerSpecParam) ImplementsMCPServerInputUnionParam() {}

func (r MCPServerSpecParam) ImplementsChatCompletionNewParamsMCPServersUnion() {}

// Detailed credential binding with options.
//
// Used when a binding needs default values, optional flags, or type casting.
//
// Satisfied by [shared.UnionString],
// [shared.MCPServerSpecCredentialsBindingSpecParam].
type MCPServerSpecCredentialsUnionParam interface {
	ImplementsMCPServerSpecCredentialsUnionParam()
}

// Detailed credential binding with options.
//
// Used when a binding needs default values, optional flags, or type casting.
type MCPServerSpecCredentialsBindingSpecParam struct {
	// Environment variable name or source identifier.
	Name param.Field[string] `json:"name,required"`
	// Type to cast value to (e.g., 'int', 'bool').
	Cast param.Field[string] `json:"cast"`
	// Default value if source not set.
	Default param.Field[interface{}] `json:"default"`
	// If true, missing value is allowed.
	Optional param.Field[bool] `json:"optional"`
}

func (r MCPServerSpecCredentialsBindingSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MCPServerSpecCredentialsBindingSpecParam) ImplementsMCPServerSpecCredentialsUnionParam() {}

type MCPServersParam []MCPServerInputUnionParam

func (r MCPServersParam) ImplementsChatCompletionNewParamsMCPServersUnion() {}

// JSON object response format. An older method of generating JSON responses. Using
// `json_schema` is recommended for models that support it. Note that the model
// will not generate JSON without a system or user message instructing it to do so.
//
// Fields:
//
// - type (required): Literal["json_object"]
type ResponseFormatJSONObjectParam struct {
	// The type of response format being defined. Always `json_object`.
	Type param.Field[ResponseFormatJSONObjectType] `json:"type,required"`
}

func (r ResponseFormatJSONObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseFormatJSONObjectParam) ImplementsChatCompletionNewParamsResponseFormatUnion() {}

// The type of response format being defined. Always `json_object`.
type ResponseFormatJSONObjectType string

const (
	ResponseFormatJSONObjectTypeJSONObject ResponseFormatJSONObjectType = "json_object"
)

func (r ResponseFormatJSONObjectType) IsKnown() bool {
	switch r {
	case ResponseFormatJSONObjectTypeJSONObject:
		return true
	}
	return false
}

// JSON Schema response format. Used to generate structured JSON responses. Learn
// more about
// [Structured Outputs](https://platform.openai.com/docs/guides/structured-outputs).
//
// Fields:
//
// - type (required): Literal["json_schema"]
// - json_schema (required): JSONSchema
type ResponseFormatJSONSchemaParam struct {
	// Structured Outputs configuration options, including a JSON Schema.
	JSONSchema param.Field[ResponseFormatJSONSchemaJSONSchemaParam] `json:"json_schema,required"`
	// The type of response format being defined. Always `json_schema`.
	Type param.Field[ResponseFormatJSONSchemaType] `json:"type,required"`
}

func (r ResponseFormatJSONSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseFormatJSONSchemaParam) ImplementsChatCompletionNewParamsResponseFormatUnion() {}

// Structured Outputs configuration options, including a JSON Schema.
type ResponseFormatJSONSchemaJSONSchemaParam struct {
	// The name of the response format. Must be a-z, A-Z, 0-9, or contain underscores
	// and dashes, with a maximum length of 64.
	Name param.Field[string] `json:"name,required"`
	// A description of what the response format is for, used by the model to determine
	// how to respond in the format.
	Description param.Field[string] `json:"description"`
	// The schema for the response format, described as a JSON Schema object. Learn how
	// to build JSON schemas [here](https://json-schema.org/).
	Schema param.Field[map[string]interface{}] `json:"schema"`
	// Whether to enable strict schema adherence when generating the output. If set to
	// true, the model will always follow the exact schema defined in the `schema`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`. To
	// learn more, read the
	// [Structured Outputs guide](https://platform.openai.com/docs/guides/structured-outputs).
	Strict param.Field[bool] `json:"strict"`
}

func (r ResponseFormatJSONSchemaJSONSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of response format being defined. Always `json_schema`.
type ResponseFormatJSONSchemaType string

const (
	ResponseFormatJSONSchemaTypeJSONSchema ResponseFormatJSONSchemaType = "json_schema"
)

func (r ResponseFormatJSONSchemaType) IsKnown() bool {
	switch r {
	case ResponseFormatJSONSchemaTypeJSONSchema:
		return true
	}
	return false
}

// Default response format. Used to generate text responses.
//
// Fields:
//
// - type (required): Literal["text"]
type ResponseFormatTextParam struct {
	// The type of response format being defined. Always `text`.
	Type param.Field[ResponseFormatTextType] `json:"type,required"`
}

func (r ResponseFormatTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseFormatTextParam) ImplementsChatCompletionNewParamsResponseFormatUnion() {}

// The type of response format being defined. Always `text`.
type ResponseFormatTextType string

const (
	ResponseFormatTextTypeText ResponseFormatTextType = "text"
)

func (r ResponseFormatTextType) IsKnown() bool {
	switch r {
	case ResponseFormatTextTypeText:
		return true
	}
	return false
}
