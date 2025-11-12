// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Structured model selection entry used in request payloads.
//
// Supports OpenAI-style semantics (string model id) while enabling optional
// per-model default settings for Dedalus multi-model routing.
//
// The property Model is required.
type DedalusModelParam struct {
	// Model identifier with provider prefix (e.g., 'openai/gpt-5',
	// 'anthropic/claude-3-5-sonnet').
	Model string `json:"model,required"`
	// Optional default generation settings (e.g., temperature, max_tokens) applied
	// when this model is selected.
	Settings DedalusModelSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r DedalusModelParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DedalusModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional default generation settings (e.g., temperature, max_tokens) applied
// when this model is selected.
type DedalusModelSettingsParam struct {
	Deferred                        param.Opt[bool]                    `json:"deferred,omitzero"`
	FrequencyPenalty                param.Opt[float64]                 `json:"frequency_penalty,omitzero"`
	IncludeUsage                    param.Opt[bool]                    `json:"include_usage,omitzero"`
	InputAudioFormat                param.Opt[string]                  `json:"input_audio_format,omitzero"`
	Logprobs                        param.Opt[bool]                    `json:"logprobs,omitzero"`
	MaxCompletionTokens             param.Opt[int64]                   `json:"max_completion_tokens,omitzero"`
	MaxTokens                       param.Opt[int64]                   `json:"max_tokens,omitzero"`
	N                               param.Opt[int64]                   `json:"n,omitzero"`
	OutputAudioFormat               param.Opt[string]                  `json:"output_audio_format,omitzero"`
	ParallelToolCalls               param.Opt[bool]                    `json:"parallel_tool_calls,omitzero"`
	PresencePenalty                 param.Opt[float64]                 `json:"presence_penalty,omitzero"`
	PromptCacheKey                  param.Opt[string]                  `json:"prompt_cache_key,omitzero"`
	ReasoningEffort                 param.Opt[string]                  `json:"reasoning_effort,omitzero"`
	SafetyIdentifier                param.Opt[string]                  `json:"safety_identifier,omitzero"`
	Seed                            param.Opt[int64]                   `json:"seed,omitzero"`
	ServiceTier                     param.Opt[string]                  `json:"service_tier,omitzero"`
	Store                           param.Opt[bool]                    `json:"store,omitzero"`
	Stream                          param.Opt[bool]                    `json:"stream,omitzero"`
	Temperature                     param.Opt[float64]                 `json:"temperature,omitzero"`
	Timeout                         param.Opt[float64]                 `json:"timeout,omitzero"`
	TopK                            param.Opt[int64]                   `json:"top_k,omitzero"`
	TopLogprobs                     param.Opt[int64]                   `json:"top_logprobs,omitzero"`
	TopP                            param.Opt[float64]                 `json:"top_p,omitzero"`
	User                            param.Opt[string]                  `json:"user,omitzero"`
	Verbosity                       param.Opt[string]                  `json:"verbosity,omitzero"`
	Voice                           param.Opt[string]                  `json:"voice,omitzero"`
	DisableAutomaticFunctionCalling param.Opt[bool]                    `json:"disable_automatic_function_calling,omitzero"`
	UseResponses                    param.Opt[bool]                    `json:"use_responses,omitzero"`
	Audio                           map[string]any                     `json:"audio,omitzero"`
	ExtraArgs                       map[string]any                     `json:"extra_args,omitzero"`
	ExtraHeaders                    map[string]string                  `json:"extra_headers,omitzero"`
	ExtraQuery                      map[string]any                     `json:"extra_query,omitzero"`
	GenerationConfig                map[string]any                     `json:"generation_config,omitzero"`
	InputAudioTranscription         map[string]any                     `json:"input_audio_transcription,omitzero"`
	LogitBias                       map[string]int64                   `json:"logit_bias,omitzero"`
	Metadata                        map[string]string                  `json:"metadata,omitzero"`
	Modalities                      []string                           `json:"modalities,omitzero"`
	Prediction                      map[string]any                     `json:"prediction,omitzero"`
	Reasoning                       DedalusModelSettingsReasoningParam `json:"reasoning,omitzero"`
	ResponseFormat                  map[string]any                     `json:"response_format,omitzero"`
	// Any of "file_search_call.results", "web_search_call.results",
	// "web_search_call.action.sources", "message.input_image.image_url",
	// "computer_call_output.output.image_url", "code_interpreter_call.outputs",
	// "reasoning.encrypted_content", "message.output_text.logprobs".
	ResponseInclude   []string                                 `json:"response_include,omitzero"`
	SafetySettings    []map[string]any                         `json:"safety_settings,omitzero"`
	SearchParameters  map[string]any                           `json:"search_parameters,omitzero"`
	Stop              DedalusModelSettingsStopUnionParam       `json:"stop,omitzero"`
	StreamOptions     map[string]any                           `json:"stream_options,omitzero"`
	SystemInstruction map[string]any                           `json:"system_instruction,omitzero"`
	Thinking          map[string]any                           `json:"thinking,omitzero"`
	ToolChoice        DedalusModelSettingsToolChoiceUnionParam `json:"tool_choice,omitzero"`
	ToolConfig        map[string]any                           `json:"tool_config,omitzero"`
	// Any of "auto", "disabled".
	Truncation       string         `json:"truncation,omitzero"`
	TurnDetection    map[string]any `json:"turn_detection,omitzero"`
	WebSearchOptions map[string]any `json:"web_search_options,omitzero"`
	Attributes       map[string]any `json:"attributes,omitzero"`
	StructuredOutput any            `json:"structured_output,omitzero"`
	paramObj
}

func (r DedalusModelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DedalusModelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DedalusModelSettingsParam](
		"truncation", "auto", "disabled",
	)
}

type DedalusModelSettingsReasoningParam struct {
	// Any of "minimal", "low", "medium", "high".
	Effort string `json:"effort,omitzero"`
	// Any of "auto", "concise", "detailed".
	GenerateSummary string `json:"generate_summary,omitzero"`
	// Any of "auto", "concise", "detailed".
	Summary     string         `json:"summary,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r DedalusModelSettingsReasoningParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelSettingsReasoningParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *DedalusModelSettingsReasoningParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DedalusModelSettingsReasoningParam](
		"effort", "minimal", "low", "medium", "high",
	)
	apijson.RegisterFieldValidator[DedalusModelSettingsReasoningParam](
		"generate_summary", "auto", "concise", "detailed",
	)
	apijson.RegisterFieldValidator[DedalusModelSettingsReasoningParam](
		"summary", "auto", "concise", "detailed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DedalusModelSettingsStopUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelSettingsStopUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *DedalusModelSettingsStopUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelSettingsStopUnionParam) asAny() any {
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
type DedalusModelSettingsToolChoiceUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfDedalusModelSettingsToolChoiceString)
	OfDedalusModelSettingsToolChoiceString param.Opt[string]                                 `json:",omitzero,inline"`
	OfString                               param.Opt[string]                                 `json:",omitzero,inline"`
	OfAnyMap                               map[string]any                                    `json:",omitzero,inline"`
	OfMCPToolChoice                        *DedalusModelSettingsToolChoiceMCPToolChoiceParam `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelSettingsToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDedalusModelSettingsToolChoiceString, u.OfString, u.OfAnyMap, u.OfMCPToolChoice)
}
func (u *DedalusModelSettingsToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelSettingsToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDedalusModelSettingsToolChoiceString) {
		return &u.OfDedalusModelSettingsToolChoiceString
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfMCPToolChoice) {
		return u.OfMCPToolChoice
	}
	return nil
}

type DedalusModelSettingsToolChoiceString string

const (
	DedalusModelSettingsToolChoiceStringAuto     DedalusModelSettingsToolChoiceString = "auto"
	DedalusModelSettingsToolChoiceStringRequired DedalusModelSettingsToolChoiceString = "required"
	DedalusModelSettingsToolChoiceStringNone     DedalusModelSettingsToolChoiceString = "none"
)

// The properties Name, ServerLabel are required.
type DedalusModelSettingsToolChoiceMCPToolChoiceParam struct {
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	paramObj
}

func (r DedalusModelSettingsToolChoiceMCPToolChoiceParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelSettingsToolChoiceMCPToolChoiceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DedalusModelSettingsToolChoiceMCPToolChoiceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func DedalusModelChoiceParamOfDedalusModel(model string) DedalusModelChoiceUnionParam {
	var variant DedalusModelParam
	variant.Model = model
	return DedalusModelChoiceUnionParam{OfDedalusModel: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DedalusModelChoiceUnionParam struct {
	OfModelID      param.Opt[githubcomdedaluslabsdedalussdkgo.ModelID] `json:",omitzero,inline"`
	OfDedalusModel *DedalusModelParam                                  `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModelID, u.OfDedalusModel)
}
func (u *DedalusModelChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfModelID) {
		return &u.OfModelID.Value
	} else if !param.IsOmitted(u.OfDedalusModel) {
		return u.OfDedalusModel
	}
	return nil
}

func NewResponseFormatJSONObjectParam() ResponseFormatJSONObjectParam {
	return ResponseFormatJSONObjectParam{
		Type: "json_object",
	}
}

// JSON object response format. An older method of generating JSON responses.
//
// Using `json_schema` is recommended for models that support it. Note that the
// model will not generate JSON without a system or user message instructing it to
// do so.
//
// Fields:
//
// - type (required): Literal['json_object']
//
// This struct has a constant value, construct it with
// [NewResponseFormatJSONObjectParam].
type ResponseFormatJSONObjectParam struct {
	// The type of response format being defined. Always `json_object`.
	Type constant.JSONObject `json:"type,required"`
	paramObj
}

func (r ResponseFormatJSONObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatJSONObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatJSONObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON Schema response format. Used to generate structured JSON responses.
//
// Learn more about
// [Structured Outputs](https://platform.openai.com/docs/guides/structured-outputs).
//
// Fields:
//
// - type (required): Literal['json_schema']
// - json_schema (required): JSONSchema
//
// The properties JSONSchema, Type are required.
type ResponseFormatJSONSchemaParam struct {
	// Structured Outputs configuration options, including a JSON Schema.
	JSONSchema ResponseFormatJSONSchemaJSONSchemaParam `json:"json_schema,omitzero,required"`
	// The type of response format being defined. Always `json_schema`.
	//
	// This field can be elided, and will marshal its zero value as "json_schema".
	Type constant.JSONSchema `json:"type,required"`
	paramObj
}

func (r ResponseFormatJSONSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatJSONSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatJSONSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured Outputs configuration options, including a JSON Schema.
//
// The property Name is required.
type ResponseFormatJSONSchemaJSONSchemaParam struct {
	// The name of the response format. Must be a-z, A-Z, 0-9, or contain underscores
	// and dashes, with a maximum length of 64.
	Name string `json:"name,required"`
	// Whether to enable strict schema adherence when generating the output. If set to
	// true, the model will always follow the exact schema defined in the `schema`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`. To
	// learn more, read the
	// [Structured Outputs guide](https://platform.openai.com/docs/guides/structured-outputs).
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// A description of what the response format is for, used by the model to determine
	// how to respond in the format.
	Description param.Opt[string] `json:"description,omitzero"`
	// The schema for the response format, described as a JSON Schema object. Learn how
	// to build JSON schemas [here](https://json-schema.org/).
	Schema map[string]any `json:"schema,omitzero"`
	paramObj
}

func (r ResponseFormatJSONSchemaJSONSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatJSONSchemaJSONSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatJSONSchemaJSONSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewResponseFormatTextParam() ResponseFormatTextParam {
	return ResponseFormatTextParam{
		Type: "text",
	}
}

// Default response format. Used to generate text responses.
//
// Fields:
//
// - type (required): Literal['text']
//
// This struct has a constant value, construct it with
// [NewResponseFormatTextParam].
type ResponseFormatTextParam struct {
	// The type of response format being defined. Always `text`.
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r ResponseFormatTextParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
