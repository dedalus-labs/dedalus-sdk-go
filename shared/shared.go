// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
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
	Attributes                      param.Field[map[string]interface{}]                   `json:"attributes"`
	Audio                           param.Field[map[string]interface{}]                   `json:"audio"`
	Deferred                        param.Field[bool]                                     `json:"deferred"`
	DisableAutomaticFunctionCalling param.Field[bool]                                     `json:"disable_automatic_function_calling"`
	ExtraArgs                       param.Field[map[string]interface{}]                   `json:"extra_args"`
	ExtraHeaders                    param.Field[map[string]string]                        `json:"extra_headers"`
	ExtraQuery                      param.Field[map[string]interface{}]                   `json:"extra_query"`
	FrequencyPenalty                param.Field[float64]                                  `json:"frequency_penalty"`
	GenerationConfig                param.Field[map[string]interface{}]                   `json:"generation_config"`
	IncludeUsage                    param.Field[bool]                                     `json:"include_usage"`
	InputAudioFormat                param.Field[string]                                   `json:"input_audio_format"`
	InputAudioTranscription         param.Field[map[string]interface{}]                   `json:"input_audio_transcription"`
	LogitBias                       param.Field[map[string]int64]                         `json:"logit_bias"`
	Logprobs                        param.Field[bool]                                     `json:"logprobs"`
	MaxCompletionTokens             param.Field[int64]                                    `json:"max_completion_tokens"`
	MaxTokens                       param.Field[int64]                                    `json:"max_tokens"`
	Metadata                        param.Field[map[string]string]                        `json:"metadata"`
	Modalities                      param.Field[[]string]                                 `json:"modalities"`
	N                               param.Field[int64]                                    `json:"n"`
	OutputAudioFormat               param.Field[string]                                   `json:"output_audio_format"`
	ParallelToolCalls               param.Field[bool]                                     `json:"parallel_tool_calls"`
	Prediction                      param.Field[map[string]interface{}]                   `json:"prediction"`
	PresencePenalty                 param.Field[float64]                                  `json:"presence_penalty"`
	PromptCacheKey                  param.Field[string]                                   `json:"prompt_cache_key"`
	Reasoning                       param.Field[DedalusModelSettingsReasoningParam]       `json:"reasoning"`
	ReasoningEffort                 param.Field[string]                                   `json:"reasoning_effort"`
	ResponseFormat                  param.Field[map[string]interface{}]                   `json:"response_format"`
	ResponseInclude                 param.Field[[]DedalusModelSettingsResponseInclude]    `json:"response_include"`
	SafetyIdentifier                param.Field[string]                                   `json:"safety_identifier"`
	SafetySettings                  param.Field[[]map[string]interface{}]                 `json:"safety_settings"`
	SearchParameters                param.Field[map[string]interface{}]                   `json:"search_parameters"`
	Seed                            param.Field[int64]                                    `json:"seed"`
	ServiceTier                     param.Field[string]                                   `json:"service_tier"`
	Stop                            param.Field[DedalusModelSettingsStopUnionParam]       `json:"stop"`
	Store                           param.Field[bool]                                     `json:"store"`
	Stream                          param.Field[bool]                                     `json:"stream"`
	StreamOptions                   param.Field[map[string]interface{}]                   `json:"stream_options"`
	StructuredOutput                param.Field[interface{}]                              `json:"structured_output"`
	SystemInstruction               param.Field[map[string]interface{}]                   `json:"system_instruction"`
	Temperature                     param.Field[float64]                                  `json:"temperature"`
	Thinking                        param.Field[map[string]interface{}]                   `json:"thinking"`
	Timeout                         param.Field[float64]                                  `json:"timeout"`
	ToolChoice                      param.Field[DedalusModelSettingsToolChoiceUnionParam] `json:"tool_choice"`
	ToolConfig                      param.Field[map[string]interface{}]                   `json:"tool_config"`
	TopK                            param.Field[int64]                                    `json:"top_k"`
	TopLogprobs                     param.Field[int64]                                    `json:"top_logprobs"`
	TopP                            param.Field[float64]                                  `json:"top_p"`
	Truncation                      param.Field[DedalusModelSettingsTruncation]           `json:"truncation"`
	TurnDetection                   param.Field[map[string]interface{}]                   `json:"turn_detection"`
	UseResponses                    param.Field[bool]                                     `json:"use_responses"`
	User                            param.Field[string]                                   `json:"user"`
	Verbosity                       param.Field[string]                                   `json:"verbosity"`
	Voice                           param.Field[string]                                   `json:"voice"`
	WebSearchOptions                param.Field[map[string]interface{}]                   `json:"web_search_options"`
}

func (r DedalusModelSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DedalusModelSettingsReasoningParam struct {
	Effort          param.Field[DedalusModelSettingsReasoningEffort]          `json:"effort"`
	GenerateSummary param.Field[DedalusModelSettingsReasoningGenerateSummary] `json:"generate_summary"`
	Summary         param.Field[DedalusModelSettingsReasoningSummary]         `json:"summary"`
	ExtraFields     map[string]interface{}                                    `json:"-,extras"`
}

func (r DedalusModelSettingsReasoningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DedalusModelSettingsReasoningEffort string

const (
	DedalusModelSettingsReasoningEffortMinimal DedalusModelSettingsReasoningEffort = "minimal"
	DedalusModelSettingsReasoningEffortLow     DedalusModelSettingsReasoningEffort = "low"
	DedalusModelSettingsReasoningEffortMedium  DedalusModelSettingsReasoningEffort = "medium"
	DedalusModelSettingsReasoningEffortHigh    DedalusModelSettingsReasoningEffort = "high"
)

func (r DedalusModelSettingsReasoningEffort) IsKnown() bool {
	switch r {
	case DedalusModelSettingsReasoningEffortMinimal, DedalusModelSettingsReasoningEffortLow, DedalusModelSettingsReasoningEffortMedium, DedalusModelSettingsReasoningEffortHigh:
		return true
	}
	return false
}

type DedalusModelSettingsReasoningGenerateSummary string

const (
	DedalusModelSettingsReasoningGenerateSummaryAuto     DedalusModelSettingsReasoningGenerateSummary = "auto"
	DedalusModelSettingsReasoningGenerateSummaryConcise  DedalusModelSettingsReasoningGenerateSummary = "concise"
	DedalusModelSettingsReasoningGenerateSummaryDetailed DedalusModelSettingsReasoningGenerateSummary = "detailed"
)

func (r DedalusModelSettingsReasoningGenerateSummary) IsKnown() bool {
	switch r {
	case DedalusModelSettingsReasoningGenerateSummaryAuto, DedalusModelSettingsReasoningGenerateSummaryConcise, DedalusModelSettingsReasoningGenerateSummaryDetailed:
		return true
	}
	return false
}

type DedalusModelSettingsReasoningSummary string

const (
	DedalusModelSettingsReasoningSummaryAuto     DedalusModelSettingsReasoningSummary = "auto"
	DedalusModelSettingsReasoningSummaryConcise  DedalusModelSettingsReasoningSummary = "concise"
	DedalusModelSettingsReasoningSummaryDetailed DedalusModelSettingsReasoningSummary = "detailed"
)

func (r DedalusModelSettingsReasoningSummary) IsKnown() bool {
	switch r {
	case DedalusModelSettingsReasoningSummaryAuto, DedalusModelSettingsReasoningSummaryConcise, DedalusModelSettingsReasoningSummaryDetailed:
		return true
	}
	return false
}

type DedalusModelSettingsResponseInclude string

const (
	DedalusModelSettingsResponseIncludeFileSearchCallResults            DedalusModelSettingsResponseInclude = "file_search_call.results"
	DedalusModelSettingsResponseIncludeWebSearchCallResults             DedalusModelSettingsResponseInclude = "web_search_call.results"
	DedalusModelSettingsResponseIncludeWebSearchCallActionSources       DedalusModelSettingsResponseInclude = "web_search_call.action.sources"
	DedalusModelSettingsResponseIncludeMessageInputImageImageURL        DedalusModelSettingsResponseInclude = "message.input_image.image_url"
	DedalusModelSettingsResponseIncludeComputerCallOutputOutputImageURL DedalusModelSettingsResponseInclude = "computer_call_output.output.image_url"
	DedalusModelSettingsResponseIncludeCodeInterpreterCallOutputs       DedalusModelSettingsResponseInclude = "code_interpreter_call.outputs"
	DedalusModelSettingsResponseIncludeReasoningEncryptedContent        DedalusModelSettingsResponseInclude = "reasoning.encrypted_content"
	DedalusModelSettingsResponseIncludeMessageOutputTextLogprobs        DedalusModelSettingsResponseInclude = "message.output_text.logprobs"
)

func (r DedalusModelSettingsResponseInclude) IsKnown() bool {
	switch r {
	case DedalusModelSettingsResponseIncludeFileSearchCallResults, DedalusModelSettingsResponseIncludeWebSearchCallResults, DedalusModelSettingsResponseIncludeWebSearchCallActionSources, DedalusModelSettingsResponseIncludeMessageInputImageImageURL, DedalusModelSettingsResponseIncludeComputerCallOutputOutputImageURL, DedalusModelSettingsResponseIncludeCodeInterpreterCallOutputs, DedalusModelSettingsResponseIncludeReasoningEncryptedContent, DedalusModelSettingsResponseIncludeMessageOutputTextLogprobs:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.DedalusModelSettingsStopArrayParam].
type DedalusModelSettingsStopUnionParam interface {
	ImplementsDedalusModelSettingsStopUnionParam()
}

type DedalusModelSettingsStopArrayParam []string

func (r DedalusModelSettingsStopArrayParam) ImplementsDedalusModelSettingsStopUnionParam() {}

// Satisfied by [shared.DedalusModelSettingsToolChoiceString],
// [shared.UnionString], [shared.DedalusModelSettingsToolChoiceMapParam],
// [shared.DedalusModelSettingsToolChoiceMCPToolChoiceParam].
type DedalusModelSettingsToolChoiceUnionParam interface {
	ImplementsDedalusModelSettingsToolChoiceUnionParam()
}

type DedalusModelSettingsToolChoiceString string

const (
	DedalusModelSettingsToolChoiceStringAuto     DedalusModelSettingsToolChoiceString = "auto"
	DedalusModelSettingsToolChoiceStringRequired DedalusModelSettingsToolChoiceString = "required"
	DedalusModelSettingsToolChoiceStringNone     DedalusModelSettingsToolChoiceString = "none"
)

func (r DedalusModelSettingsToolChoiceString) IsKnown() bool {
	switch r {
	case DedalusModelSettingsToolChoiceStringAuto, DedalusModelSettingsToolChoiceStringRequired, DedalusModelSettingsToolChoiceStringNone:
		return true
	}
	return false
}

func (r DedalusModelSettingsToolChoiceString) ImplementsDedalusModelSettingsToolChoiceUnionParam() {}

type DedalusModelSettingsToolChoiceMapParam map[string]interface{}

func (r DedalusModelSettingsToolChoiceMapParam) ImplementsDedalusModelSettingsToolChoiceUnionParam() {
}

type DedalusModelSettingsToolChoiceMCPToolChoiceParam struct {
	Name        param.Field[string] `json:"name,required"`
	ServerLabel param.Field[string] `json:"server_label,required"`
}

func (r DedalusModelSettingsToolChoiceMCPToolChoiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DedalusModelSettingsToolChoiceMCPToolChoiceParam) ImplementsDedalusModelSettingsToolChoiceUnionParam() {
}

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

// JSON object response format. An older method of generating JSON responses.
//
// Using `json_schema` is recommended for models that support it. Note that the
// model will not generate JSON without a system or user message instructing it to
// do so.
//
// Fields:
//
// - type (required): Literal['json_object']
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

// JSON Schema response format. Used to generate structured JSON responses.
//
// Learn more about
// [Structured Outputs](https://platform.openai.com/docs/guides/structured-outputs).
//
// Fields:
//
// - type (required): Literal['json_schema']
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
// - type (required): Literal['text']
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
