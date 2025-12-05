// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apiform"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/param"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/tidwall/gjson"
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
func NewAudioTranscriptionService(opts ...option.RequestOption) (r *AudioTranscriptionService) {
	r = &AudioTranscriptionService{}
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
func (r *AudioTranscriptionService) New(ctx context.Context, body AudioTranscriptionNewParams, opts ...option.RequestOption) (res *AudioTranscriptionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/audio/transcriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
type AudioTranscriptionNewResponse struct {
	// The transcribed text.
	Text string `json:"text,required"`
	// The duration of the input audio.
	Duration float64 `json:"duration"`
	// The language of the input audio.
	Language string `json:"language"`
	// This field can have the runtime type of
	// [[]AudioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprob].
	Logprobs interface{} `json:"logprobs"`
	// This field can have the runtime type of
	// [[]AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegment].
	Segments interface{} `json:"segments"`
	// This field can have the runtime type of
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsage],
	// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage].
	Usage interface{} `json:"usage"`
	// This field can have the runtime type of
	// [[]AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWord].
	Words interface{}                       `json:"words"`
	JSON  audioTranscriptionNewResponseJSON `json:"-"`
	union AudioTranscriptionNewResponseUnion
}

// audioTranscriptionNewResponseJSON contains the JSON metadata for the struct
// [AudioTranscriptionNewResponse]
type audioTranscriptionNewResponseJSON struct {
	Text        apijson.Field
	Duration    apijson.Field
	Language    apijson.Field
	Logprobs    apijson.Field
	Segments    apijson.Field
	Usage       apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r audioTranscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AudioTranscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AudioTranscriptionNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AudioTranscriptionNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSON],
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSON].
func (r AudioTranscriptionNewResponse) AsUnion() AudioTranscriptionNewResponseUnion {
	return r.union
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
//
// Union satisfied by
// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSON] or
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSON].
type AudioTranscriptionNewResponseUnion interface {
	implementsAudioTranscriptionNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AudioTranscriptionNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSON{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranscriptionNewResponseCreateTranscriptionResponseJSON{}),
		},
	)
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
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSON struct {
	// The duration of the input audio.
	Duration float64 `json:"duration,required"`
	// The language of the input audio.
	Language string `json:"language,required"`
	// The transcribed text.
	Text string `json:"text,required"`
	// Segments of the transcribed text and their corresponding details.
	Segments []AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegment `json:"segments"`
	// Usage statistics for models billed by audio input duration.
	Usage AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsage `json:"usage"`
	// Extracted words and their corresponding timestamps.
	Words []AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWord `json:"words"`
	JSON  audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONJSON   `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONJSON contains
// the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSON]
type audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONJSON struct {
	Duration    apijson.Field
	Language    apijson.Field
	Text        apijson.Field
	Segments    apijson.Field
	Usage       apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSON) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSON) implementsAudioTranscriptionNewResponse() {
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
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegment struct {
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
	Tokens []int64                                                                        `json:"tokens,required"`
	JSON   audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegmentJSON `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegmentJSON
// contains the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegment]
type audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegmentJSON struct {
	ID               apijson.Field
	AvgLogprob       apijson.Field
	CompressionRatio apijson.Field
	End              apijson.Field
	NoSpeechProb     apijson.Field
	Seek             apijson.Field
	Start            apijson.Field
	Temperature      apijson.Field
	Text             apijson.Field
	Tokens           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONSegmentJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for models billed by audio input duration.
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsage struct {
	// Duration of the input audio in seconds.
	Seconds float64 `json:"seconds,required"`
	// The type of the usage object. Always `duration` for this variant.
	Type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageType `json:"type,required"`
	JSON audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageJSON `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageJSON
// contains the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsage]
type audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageJSON struct {
	Seconds     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageJSON) RawJSON() string {
	return r.raw
}

// The type of the usage object. Always `duration` for this variant.
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageType string

const (
	AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageTypeDuration AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageType = "duration"
)

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageType) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONUsageTypeDuration:
		return true
	}
	return false
}

// Fields:
//
// - word (required): str
// - start (required): float
// - end (required): float
type AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWord struct {
	// End time of the word in seconds.
	End float64 `json:"end,required"`
	// Start time of the word in seconds.
	Start float64 `json:"start,required"`
	// The text content of the word.
	Word string                                                                      `json:"word,required"`
	JSON audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWordJSON `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWordJSON
// contains the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWord]
type audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseVerboseJSONWordJSON) RawJSON() string {
	return r.raw
}

// Represents a transcription response returned by model, based on the provided
// input.
//
// Fields:
//
// - text (required): str
// - logprobs (optional): list[LogprobsItem]
// - usage (optional): Usage
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSON struct {
	// The transcribed text.
	Text string `json:"text,required"`
	// The log probabilities of the tokens in the transcription. Only returned with the
	// models `gpt-4o-transcribe` and `gpt-4o-mini-transcribe` if `logprobs` is added
	// to the `include` array.
	Logprobs []AudioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprob `json:"logprobs"`
	// Token usage statistics for the request.
	Usage AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage `json:"usage"`
	JSON  audioTranscriptionNewResponseCreateTranscriptionResponseJSONJSON  `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseJSONJSON contains the
// JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSON]
type audioTranscriptionNewResponseCreateTranscriptionResponseJSONJSON struct {
	Text        apijson.Field
	Logprobs    apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJSON) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseJSONJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJSON) implementsAudioTranscriptionNewResponse() {
}

// Fields:
//
// - token (optional): str
// - logprob (optional): float
// - bytes (optional): list[float]
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprob struct {
	// The token in the transcription.
	Token string `json:"token"`
	// The bytes of the token.
	Bytes []float64 `json:"bytes"`
	// The log probability of the token.
	Logprob float64                                                                 `json:"logprob"`
	JSON    audioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprobJSON `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprobJSON contains
// the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprob]
type audioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprobJSON struct {
	Token       apijson.Field
	Bytes       apijson.Field
	Logprob     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseJSONLogprobJSON) RawJSON() string {
	return r.raw
}

// Token usage statistics for the request.
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage struct {
	// The type of the usage object. Always `tokens` for this variant.
	Type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageType `json:"type,required"`
	// Details about the input tokens billed for this request.
	InputTokenDetails InputTokenDetails `json:"input_token_details"`
	// Number of input tokens billed for this request.
	InputTokens int64 `json:"input_tokens"`
	// Number of output tokens generated.
	OutputTokens int64 `json:"output_tokens"`
	// Duration of the input audio in seconds.
	Seconds float64 `json:"seconds"`
	// Total number of tokens used (input + output).
	TotalTokens int64                                                                 `json:"total_tokens"`
	JSON        audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageJSON `json:"-"`
	union       AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageUnion
}

// audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageJSON contains
// the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage]
type audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageJSON struct {
	Type              apijson.Field
	InputTokenDetails apijson.Field
	InputTokens       apijson.Field
	OutputTokens      apijson.Field
	Seconds           apijson.Field
	TotalTokens       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageJSON) RawJSON() string {
	return r.raw
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage) UnmarshalJSON(data []byte) (err error) {
	*r = AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokens],
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDuration].
func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage) AsUnion() AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageUnion {
	return r.union
}

// Token usage statistics for the request.
//
// Union satisfied by
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokens]
// or
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDuration].
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageUnion interface {
	implementsAudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokens{}),
			DiscriminatorValue: "tokens",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDuration{}),
			DiscriminatorValue: "duration",
		},
	)
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
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokens struct {
	// Number of input tokens billed for this request.
	InputTokens int64 `json:"input_tokens,required"`
	// Number of output tokens generated.
	OutputTokens int64 `json:"output_tokens,required"`
	// Total number of tokens used (input + output).
	TotalTokens int64 `json:"total_tokens,required"`
	// The type of the usage object. Always `tokens` for this variant.
	Type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensType `json:"type,required"`
	// Details about the input tokens billed for this request.
	InputTokenDetails InputTokenDetails                                                                              `json:"input_token_details"`
	JSON              audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensJSON `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensJSON
// contains the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokens]
type audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensJSON struct {
	InputTokens       apijson.Field
	OutputTokens      apijson.Field
	TotalTokens       apijson.Field
	Type              apijson.Field
	InputTokenDetails apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokens) implementsAudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage() {
}

// The type of the usage object. Always `tokens` for this variant.
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensType string

const (
	AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensTypeTokens AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensType = "tokens"
)

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensType) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageTokensTypeTokens:
		return true
	}
	return false
}

// Usage statistics for models billed by audio input duration.
//
// Fields:
//
// - type (required): Literal['duration']
// - seconds (required): float
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDuration struct {
	// Duration of the input audio in seconds.
	Seconds float64 `json:"seconds,required"`
	// The type of the usage object. Always `duration` for this variant.
	Type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationType `json:"type,required"`
	JSON audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationJSON `json:"-"`
}

// audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationJSON
// contains the JSON metadata for the struct
// [AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDuration]
type audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationJSON struct {
	Seconds     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDuration) implementsAudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsage() {
}

// The type of the usage object. Always `duration` for this variant.
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationType string

const (
	AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationTypeDuration AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationType = "duration"
)

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationType) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTranscriptTextUsageDurationTypeDuration:
		return true
	}
	return false
}

// The type of the usage object. Always `tokens` for this variant.
type AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageType string

const (
	AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTypeTokens   AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageType = "tokens"
	AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTypeDuration AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageType = "duration"
)

func (r AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageType) IsKnown() bool {
	switch r {
	case AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTypeTokens, AudioTranscriptionNewResponseCreateTranscriptionResponseJSONUsageTypeDuration:
		return true
	}
	return false
}

type AudioTranscriptionNewParams struct {
	File           param.Field[io.Reader] `json:"file,required" format:"binary"`
	Model          param.Field[string]    `json:"model,required"`
	Language       param.Field[string]    `json:"language"`
	Prompt         param.Field[string]    `json:"prompt"`
	ResponseFormat param.Field[string]    `json:"response_format"`
	Temperature    param.Field[float64]   `json:"temperature"`
}

func (r AudioTranscriptionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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
