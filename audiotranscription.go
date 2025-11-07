// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apiform"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
	"github.com/dedalus-labs/dedalus-sdk-go/shared/constant"
)

// AudioTranscriptionService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioTranscriptionService] method instead.
type AudioTranscriptionService struct {
	Options []option.RequestOption
}

// NewAudioTranscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAudioTranscriptionService(opts ...option.RequestOption) (r AudioTranscriptionService) {
	r = AudioTranscriptionService{}
	r.Options = opts
	return
}

// Transcribe audio into text.
//
// Transcribes audio files using OpenAI's Whisper model. Supports multiple audio
// formats including mp3, mp4, mpeg, mpga, m4a, wav, and webm. Maximum file size is
// 25 MB.
//
// Args: file: Audio file to transcribe (required) model: Model ID to use (e.g.,
// "openai/whisper-1") language: ISO-639-1 language code (e.g., "en", "es") -
// improves accuracy prompt: Optional text to guide the model's style
// response_format: Format of the output (json, text, srt, verbose_json, vtt)
// temperature: Sampling temperature between 0 and 1
//
// Returns: Transcription object with the transcribed text
func (r *AudioTranscriptionService) New(ctx context.Context, body AudioTranscriptionNewParams, opts ...option.RequestOption) (res *AudioTranscriptionNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/audio/transcriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// AudioTranscriptionNewResponseUnion contains all possible properties and values
// from [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson],
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJson].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AudioTranscriptionNewResponseUnion struct {
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson].
	Duration float64 `json:"duration"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson].
	Language string `json:"language"`
	Text     string `json:"text"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson].
	Segments []AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonSegment `json:"segments"`
	// This field is a union of
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonUsage],
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion]
	Usage AudioTranscriptionNewResponseUnionUsage `json:"usage"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson].
	Words []AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonWord `json:"words"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJson].
	Logprobs []AudioTranscriptionNewResponseCreateTranscriptionResponseJsonLogprob `json:"logprobs"`
	JSON     struct {
		Duration respjson.Field
		Language respjson.Field
		Text     respjson.Field
		Segments respjson.Field
		Usage    respjson.Field
		Words    respjson.Field
		Logprobs respjson.Field
		raw      string
	} `json:"-"`
}

func (u AudioTranscriptionNewResponseUnion) AsCreateTranscriptionResponseVerboseJson() (v AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AudioTranscriptionNewResponseUnion) AsCreateTranscriptionResponseJson() (v AudioTranscriptionNewResponseCreateTranscriptionResponseJson) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AudioTranscriptionNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AudioTranscriptionNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AudioTranscriptionNewResponseUnionUsage is an implicit subunion of
// [AudioTranscriptionNewResponseUnion]. AudioTranscriptionNewResponseUnionUsage
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AudioTranscriptionNewResponseUnion].
type AudioTranscriptionNewResponseUnionUsage struct {
	Seconds float64 `json:"seconds"`
	Type    string  `json:"type"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion].
	InputTokens int64 `json:"input_tokens"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion].
	OutputTokens int64 `json:"output_tokens"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion].
	TotalTokens int64 `json:"total_tokens"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion].
	InputTokenDetails AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokensInputTokenDetails `json:"input_token_details"`
	JSON              struct {
		Seconds           respjson.Field
		Type              respjson.Field
		InputTokens       respjson.Field
		OutputTokens      respjson.Field
		TotalTokens       respjson.Field
		InputTokenDetails respjson.Field
		raw               string
	} `json:"-"`
}

func (r *AudioTranscriptionNewResponseUnionUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a verbose json transcription response returned by model, based on the
// provided input.
//
// Fields:
//
// - language (required): str
// - duration (required): float
// - text (required): str
// - words (optional): list[TranscriptionWord]
// - segments (optional): list[TranscriptionSegment]
// - usage (optional): TranscriptTextUsageDuration
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson struct {
	// The duration of the input audio.
	Duration float64 `json:"duration,required"`
	// The language of the input audio.
	Language string `json:"language,required"`
	// The transcribed text.
	Text string `json:"text,required"`
	// Segments of the transcribed text and their corresponding details.
	Segments []AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonSegment `json:"segments"`
	// Usage statistics for models billed by audio input duration.
	Usage AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonUsage `json:"usage"`
	// Extracted words and their corresponding timestamps.
	Words []AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonWord `json:"words"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Language    respjson.Field
		Text        respjson.Field
		Segments    respjson.Field
		Usage       respjson.Field
		Words       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields:
//
// - id (required): int
// - seek (required): int
// - start (required): float
// - end (required): float
// - text (required): str
// - tokens (required): list[int]
// - temperature (required): float
// - avg_logprob (required): float
// - compression_ratio (required): float
// - no_speech_prob (required): float
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonSegment struct {
	// Unique identifier of the segment.
	ID int64 `json:"id,required"`
	// Average logprob of the segment. If the value is lower than -1, consider the
	// logprobs failed.
	AvgLogprob float64 `json:"avg_logprob,required"`
	// Compression ratio of the segment. If the value is greater than 2.4, consider the
	// compression failed.
	CompressionRatio float64 `json:"compression_ratio,required"`
	// End time of the segment in seconds.
	End float64 `json:"end,required"`
	// Probability of no speech in the segment. If the value is higher than 1.0 and the
	// `avg_logprob` is below -1, consider this segment silent.
	NoSpeechProb float64 `json:"no_speech_prob,required"`
	// Seek offset of the segment.
	Seek int64 `json:"seek,required"`
	// Start time of the segment in seconds.
	Start float64 `json:"start,required"`
	// Temperature parameter used for generating the segment.
	Temperature float64 `json:"temperature,required"`
	// Text content of the segment.
	Text string `json:"text,required"`
	// Array of token IDs for the text content.
	Tokens []int64 `json:"tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvgLogprob       respjson.Field
		CompressionRatio respjson.Field
		End              respjson.Field
		NoSpeechProb     respjson.Field
		Seek             respjson.Field
		Start            respjson.Field
		Temperature      respjson.Field
		Text             respjson.Field
		Tokens           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage statistics for models billed by audio input duration.
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonUsage struct {
	// Duration of the input audio in seconds.
	Seconds float64 `json:"seconds,required"`
	// The type of the usage object. Always `duration` for this variant.
	Type constant.Duration `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Seconds     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonUsage) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields:
//
// - word (required): str
// - start (required): float
// - end (required): float
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonWord struct {
	// End time of the word in seconds.
	End float64 `json:"end,required"`
	// Start time of the word in seconds.
	Start float64 `json:"start,required"`
	// The text content of the word.
	Word string `json:"word,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		Word        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonWord) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJsonWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a transcription response returned by model, based on the provided
// input.
//
// Fields:
//
// - text (required): str
// - logprobs (optional): list[LogprobsItem]
// - usage (optional): Usage
type AudioTranscriptionNewResponseCreateTranscriptionResponseJson struct {
	// The transcribed text.
	Text string `json:"text,required"`
	// The log probabilities of the tokens in the transcription. Only returned with the
	// models `gpt-4o-transcribe` and `gpt-4o-mini-transcribe` if `logprobs` is added
	// to the `include` array.
	Logprobs []AudioTranscriptionNewResponseCreateTranscriptionResponseJsonLogprob `json:"logprobs"`
	// Token usage statistics for the request.
	Usage AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Logprobs    respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJson) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields:
//
// - token (optional): str
// - logprob (optional): float
// - bytes (optional): list[float]
type AudioTranscriptionNewResponseCreateTranscriptionResponseJsonLogprob struct {
	// The token in the transcription.
	Token string `json:"token"`
	// The bytes of the token.
	Bytes []float64 `json:"bytes"`
	// The log probability of the token.
	Logprob float64 `json:"logprob"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Bytes       respjson.Field
		Logprob     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJsonLogprob) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJsonLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion contains
// all possible properties and values from
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens],
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration].
//
// Use the
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion struct {
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens].
	InputTokens int64 `json:"input_tokens"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens].
	OutputTokens int64 `json:"output_tokens"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens].
	TotalTokens int64 `json:"total_tokens"`
	// Any of "tokens", "duration".
	Type string `json:"type"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens].
	InputTokenDetails AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokensInputTokenDetails `json:"input_token_details"`
	// This field is from variant
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration].
	Seconds float64 `json:"seconds"`
	JSON    struct {
		InputTokens       respjson.Field
		OutputTokens      respjson.Field
		TotalTokens       respjson.Field
		Type              respjson.Field
		InputTokenDetails respjson.Field
		Seconds           respjson.Field
		raw               string
	} `json:"-"`
}

// anyAudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsage is
// implemented by each variant of
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion] to add
// type safety for the return type of
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion.AsAny]
type anyAudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsage interface {
	implAudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion()
}

func (AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens) implAudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion() {
}
func (AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration) implAudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion.AsAny().(type) {
//	case githubcomdedaluslabsdedalussdkgo.AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens:
//	case githubcomdedaluslabsdedalussdkgo.AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion) AsAny() anyAudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsage {
	switch u.Type {
	case "tokens":
		return u.AsTokens()
	case "duration":
		return u.AsDuration()
	}
	return nil
}

func (u AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion) AsTokens() (v AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion) AsDuration() (v AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage statistics for models billed by token usage.
//
// Fields:
//
// - type (required): Literal['tokens']
// - input_tokens (required): int
// - input_token_details (optional): InputTokenDetails
// - output_tokens (required): int
// - total_tokens (required): int
type AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens struct {
	// Number of input tokens billed for this request.
	InputTokens int64 `json:"input_tokens,required"`
	// Number of output tokens generated.
	OutputTokens int64 `json:"output_tokens,required"`
	// Total number of tokens used (input + output).
	TotalTokens int64 `json:"total_tokens,required"`
	// The type of the usage object. Always `tokens` for this variant.
	Type constant.Tokens `json:"type,required"`
	// Details about the input tokens billed for this request.
	InputTokenDetails AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokensInputTokenDetails `json:"input_token_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens       respjson.Field
		OutputTokens      respjson.Field
		TotalTokens       respjson.Field
		Type              respjson.Field
		InputTokenDetails respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokens) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the input tokens billed for this request.
type AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokensInputTokenDetails struct {
	// Number of audio tokens billed for this request.
	AudioTokens int64 `json:"audio_tokens"`
	// Number of text tokens billed for this request.
	TextTokens int64 `json:"text_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioTokens respjson.Field
		TextTokens  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokensInputTokenDetails) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageTokensInputTokenDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage statistics for models billed by audio input duration.
//
// Fields:
//
// - type (required): Literal['duration']
// - seconds (required): float
type AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration struct {
	// Duration of the input audio in seconds.
	Seconds float64 `json:"seconds,required"`
	// The type of the usage object. Always `duration` for this variant.
	Type constant.Duration `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Seconds     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJsonUsageDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranscriptionNewParams struct {
	File           io.Reader          `json:"file,omitzero,required" format:"binary"`
	Model          string             `json:"model,required"`
	Language       param.Opt[string]  `json:"language,omitzero"`
	Prompt         param.Opt[string]  `json:"prompt,omitzero"`
	ResponseFormat param.Opt[string]  `json:"response_format,omitzero"`
	Temperature    param.Opt[float64] `json:"temperature,omitzero"`
	paramObj
}

func (r AudioTranscriptionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
