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

// AudioTranslationService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioTranslationService] method instead.
type AudioTranslationService struct {
	Options []option.RequestOption
}

// NewAudioTranslationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAudioTranslationService(opts ...option.RequestOption) (r *AudioTranslationService) {
	r = &AudioTranslationService{}
	r.Options = opts
	return
}

// Translate audio into English.
//
// Translates audio files in any supported language to English text using OpenAI's
// Whisper model. Supports the same audio formats as transcription. Maximum file
// size is 25 MB.
//
// Args: file: Audio file to translate (required) model: Model ID to use (e.g.,
// "openai/whisper-1") prompt: Optional text to guide the model's style
// response_format: Format of the output (json, text, srt, verbose_json, vtt)
// temperature: Sampling temperature between 0 and 1
//
// Returns: Translation object with the English translation
func (r *AudioTranslationService) New(ctx context.Context, body AudioTranslationNewParams, opts ...option.RequestOption) (res *AudioTranslationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/audio/translations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fields:
//
// - language (required): str
// - duration (required): float
// - text (required): str
// - segments (optional): list[TranscriptionSegment]
type AudioTranslationNewResponse struct {
	// The translated text.
	Text string `json:"text,required"`
	// The duration of the input audio.
	Duration float64 `json:"duration"`
	// The language of the output translation (always `english`).
	Language string `json:"language"`
	// This field can have the runtime type of
	// [[]AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment].
	Segments interface{}                     `json:"segments"`
	JSON     audioTranslationNewResponseJSON `json:"-"`
	union    AudioTranslationNewResponseUnion
}

// audioTranslationNewResponseJSON contains the JSON metadata for the struct
// [AudioTranslationNewResponse]
type audioTranslationNewResponseJSON struct {
	Text        apijson.Field
	Duration    apijson.Field
	Language    apijson.Field
	Segments    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r audioTranslationNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AudioTranslationNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AudioTranslationNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AudioTranslationNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AudioTranslationNewResponseCreateTranslationResponseVerboseJSON],
// [AudioTranslationNewResponseCreateTranslationResponseJSON].
func (r AudioTranslationNewResponse) AsUnion() AudioTranslationNewResponseUnion {
	return r.union
}

// Fields:
//
// - language (required): str
// - duration (required): float
// - text (required): str
// - segments (optional): list[TranscriptionSegment]
//
// Union satisfied by
// [AudioTranslationNewResponseCreateTranslationResponseVerboseJSON] or
// [AudioTranslationNewResponseCreateTranslationResponseJSON].
type AudioTranslationNewResponseUnion interface {
	implementsAudioTranslationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AudioTranslationNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranslationNewResponseCreateTranslationResponseVerboseJSON{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AudioTranslationNewResponseCreateTranslationResponseJSON{}),
		},
	)
}

// Fields:
//
// - language (required): str
// - duration (required): float
// - text (required): str
// - segments (optional): list[TranscriptionSegment]
type AudioTranslationNewResponseCreateTranslationResponseVerboseJSON struct {
	// The duration of the input audio.
	Duration float64 `json:"duration,required"`
	// The language of the output translation (always `english`).
	Language string `json:"language,required"`
	// The translated text.
	Text string `json:"text,required"`
	// Segments of the translated text and their corresponding details.
	Segments []AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment `json:"segments"`
	JSON     audioTranslationNewResponseCreateTranslationResponseVerboseJSONJSON      `json:"-"`
}

// audioTranslationNewResponseCreateTranslationResponseVerboseJSONJSON contains the
// JSON metadata for the struct
// [AudioTranslationNewResponseCreateTranslationResponseVerboseJSON]
type audioTranslationNewResponseCreateTranslationResponseVerboseJSONJSON struct {
	Duration    apijson.Field
	Language    apijson.Field
	Text        apijson.Field
	Segments    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranslationNewResponseCreateTranslationResponseVerboseJSON) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranslationNewResponseCreateTranslationResponseVerboseJSONJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranslationNewResponseCreateTranslationResponseVerboseJSON) implementsAudioTranslationNewResponse() {
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
type AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment struct {
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
	Tokens []int64                                                                    `json:"tokens,required"`
	JSON   audioTranslationNewResponseCreateTranslationResponseVerboseJSONSegmentJSON `json:"-"`
}

// audioTranslationNewResponseCreateTranslationResponseVerboseJSONSegmentJSON
// contains the JSON metadata for the struct
// [AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment]
type audioTranslationNewResponseCreateTranslationResponseVerboseJSONSegmentJSON struct {
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

func (r *AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranslationNewResponseCreateTranslationResponseVerboseJSONSegmentJSON) RawJSON() string {
	return r.raw
}

// Fields:
//
// - text (required): str
type AudioTranslationNewResponseCreateTranslationResponseJSON struct {
	Text string                                                       `json:"text,required"`
	JSON audioTranslationNewResponseCreateTranslationResponseJSONJSON `json:"-"`
}

// audioTranslationNewResponseCreateTranslationResponseJSONJSON contains the JSON
// metadata for the struct
// [AudioTranslationNewResponseCreateTranslationResponseJSON]
type audioTranslationNewResponseCreateTranslationResponseJSONJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AudioTranslationNewResponseCreateTranslationResponseJSON) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r audioTranslationNewResponseCreateTranslationResponseJSONJSON) RawJSON() string {
	return r.raw
}

func (r AudioTranslationNewResponseCreateTranslationResponseJSON) implementsAudioTranslationNewResponse() {
}

type AudioTranslationNewParams struct {
	File           param.Field[io.Reader] `json:"file,required" format:"binary"`
	Model          param.Field[string]    `json:"model,required"`
	Prompt         param.Field[string]    `json:"prompt"`
	ResponseFormat param.Field[string]    `json:"response_format"`
	Temperature    param.Field[float64]   `json:"temperature"`
}

func (r AudioTranslationNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
