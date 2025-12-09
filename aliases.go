// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apierror"
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
)

type Error = apierror.Error

// Structured model selection entry used in request payloads.
//
// Supports OpenAI-style semantics (string model id) while enabling optional
// per-model default settings for Dedalus multi-model routing.
//
// This is an alias to an internal type.
type DedalusModelParam = shared.DedalusModelParam

// Optional default generation settings (e.g., temperature, max_tokens) applied
// when this model is selected.
//
// This is an alias to an internal type.
type DedalusModelSettingsParam = shared.DedalusModelSettingsParam

// This is an alias to an internal type.
type DedalusModelSettingsStopUnionParam = shared.DedalusModelSettingsStopUnionParam

// This is an alias to an internal type.
type DedalusModelSettingsStopArrayParam = shared.DedalusModelSettingsStopArrayParam

// This is an alias to an internal type.
type DedalusModelSettingsTruncation = shared.DedalusModelSettingsTruncation

// This is an alias to an internal value.
const DedalusModelSettingsTruncationAuto = shared.DedalusModelSettingsTruncationAuto

// This is an alias to an internal value.
const DedalusModelSettingsTruncationDisabled = shared.DedalusModelSettingsTruncationDisabled

// Dedalus model choice - either a string ID or DedalusModel configuration object.
//
// This is an alias to an internal type.
type DedalusModelChoiceUnionParam = shared.DedalusModelChoiceUnionParam

// Schema for FunctionObject.
//
// Fields:
//
// - description (optional): str
// - name (required): str
// - parameters (optional): FunctionParameters
// - strict (optional): bool | None
//
// This is an alias to an internal type.
type FunctionDefinitionParam = shared.FunctionDefinitionParam

// The parameters the functions accepts, described as a JSON Schema object. See the
// [guide](https://platform.openai.com/docs/guides/function-calling) for examples,
// and the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for
// documentation about the format.
//
// Omitting `parameters` defines a function with an empty parameter list.
//
// This is an alias to an internal type.
type FunctionParameters = shared.FunctionParameters

// Single MCP server input: slug string or structured MCPServerSpec.
//
// This is an alias to an internal type.
type MCPServerInputUnionParam = shared.MCPServerInputUnionParam

// Structured MCP server specification.
//
// Slug-based: {"slug": "dedalus-labs/brave-search", "version": "v1.0.0"}
// URL-based: {"url": "https://mcp.dedaluslabs.ai/acme/my-server/mcp"}
//
// This is an alias to an internal type.
type MCPServerSpecParam = shared.MCPServerSpecParam

// Detailed credential binding with options.
//
// Used when a binding needs default values, optional flags, or type casting.
//
// This is an alias to an internal type.
type MCPServerSpecCredentialsUnionParam = shared.MCPServerSpecCredentialsUnionParam

// Detailed credential binding with options.
//
// Used when a binding needs default values, optional flags, or type casting.
//
// This is an alias to an internal type.
type MCPServerSpecCredentialsBindingSpecParam = shared.MCPServerSpecCredentialsBindingSpecParam

// List of MCP server inputs.
//
// This is an alias to an internal type.
type MCPServersParam = shared.MCPServersParam

// JSON object response format. An older method of generating JSON responses. Using
// `json_schema` is recommended for models that support it. Note that the model
// will not generate JSON without a system or user message instructing it to do so.
//
// Fields:
//
// - type (required): Literal["json_object"]
//
// This is an alias to an internal type.
type ResponseFormatJSONObjectParam = shared.ResponseFormatJSONObjectParam

// The type of response format being defined. Always `json_object`.
//
// This is an alias to an internal type.
type ResponseFormatJSONObjectType = shared.ResponseFormatJSONObjectType

// This is an alias to an internal value.
const ResponseFormatJSONObjectTypeJSONObject = shared.ResponseFormatJSONObjectTypeJSONObject

// JSON Schema response format. Used to generate structured JSON responses. Learn
// more about
// [Structured Outputs](https://platform.openai.com/docs/guides/structured-outputs).
//
// Fields:
//
// - type (required): Literal["json_schema"]
// - json_schema (required): JSONSchema
//
// This is an alias to an internal type.
type ResponseFormatJSONSchemaParam = shared.ResponseFormatJSONSchemaParam

// Structured Outputs configuration options, including a JSON Schema.
//
// This is an alias to an internal type.
type ResponseFormatJSONSchemaJSONSchemaParam = shared.ResponseFormatJSONSchemaJSONSchemaParam

// The type of response format being defined. Always `json_schema`.
//
// This is an alias to an internal type.
type ResponseFormatJSONSchemaType = shared.ResponseFormatJSONSchemaType

// This is an alias to an internal value.
const ResponseFormatJSONSchemaTypeJSONSchema = shared.ResponseFormatJSONSchemaTypeJSONSchema

// Default response format. Used to generate text responses.
//
// Fields:
//
// - type (required): Literal["text"]
//
// This is an alias to an internal type.
type ResponseFormatTextParam = shared.ResponseFormatTextParam

// The type of response format being defined. Always `text`.
//
// This is an alias to an internal type.
type ResponseFormatTextType = shared.ResponseFormatTextType

// This is an alias to an internal value.
const ResponseFormatTextTypeText = shared.ResponseFormatTextTypeText
