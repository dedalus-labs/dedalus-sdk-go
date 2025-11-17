// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/param"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
)

// AudioSpeechService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioSpeechService] method instead.
type AudioSpeechService struct {
	Options []option.RequestOption
}

// NewAudioSpeechService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioSpeechService(opts ...option.RequestOption) (r *AudioSpeechService) {
	r = &AudioSpeechService{}
	r.Options = opts
	return
}

// Generate speech audio from text.
//
// Generates audio from the input text using text-to-speech models. Supports
// multiple voices and output formats including mp3, opus, aac, flac, wav, and pcm.
//
// Returns streaming audio data that can be saved to a file or streamed directly to
// users.
func (r *AudioSpeechService) New(ctx context.Context, body AudioSpeechNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "v1/audio/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AudioSpeechNewParams struct {
	// The text to generate audio for. The maximum length is 4096 characters.
	Input param.Field[string] `json:"input,required"`
	// One of the available [TTS models](https://platform.openai.com/docs/models#tts):
	// `openai/tts-1`, `openai/tts-1-hd` or `openai/gpt-4o-mini-tts`.
	Model param.Field[string] `json:"model,required"`
	// The voice to use when generating the audio. Supported voices are `alloy`, `ash`,
	// `ballad`, `coral`, `echo`, `fable`, `onyx`, `nova`, `sage`, `shimmer`, and
	// `verse`. Previews of the voices are available in the
	// [Text to speech guide](https://platform.openai.com/docs/guides/text-to-speech#voice-options).
	Voice param.Field[AudioSpeechNewParamsVoice] `json:"voice,required"`
	// Control the voice of your generated audio with additional instructions. Does not
	// work with `tts-1` or `tts-1-hd`.
	Instructions param.Field[string] `json:"instructions"`
	// The format to audio in. Supported formats are `mp3`, `opus`, `aac`, `flac`,
	// `wav`, and `pcm`.
	ResponseFormat param.Field[AudioSpeechNewParamsResponseFormat] `json:"response_format"`
	// The speed of the generated audio. Select a value from `0.25` to `4.0`. `1.0` is
	// the default.
	Speed param.Field[float64] `json:"speed"`
	// The format to stream the audio in. Supported formats are `sse` and `audio`.
	// `sse` is not supported for `tts-1` or `tts-1-hd`.
	StreamFormat param.Field[AudioSpeechNewParamsStreamFormat] `json:"stream_format"`
}

func (r AudioSpeechNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The voice to use when generating the audio. Supported voices are `alloy`, `ash`,
// `ballad`, `coral`, `echo`, `fable`, `onyx`, `nova`, `sage`, `shimmer`, and
// `verse`. Previews of the voices are available in the
// [Text to speech guide](https://platform.openai.com/docs/guides/text-to-speech#voice-options).
type AudioSpeechNewParamsVoice string

const (
	AudioSpeechNewParamsVoiceAlloy   AudioSpeechNewParamsVoice = "alloy"
	AudioSpeechNewParamsVoiceAsh     AudioSpeechNewParamsVoice = "ash"
	AudioSpeechNewParamsVoiceBallad  AudioSpeechNewParamsVoice = "ballad"
	AudioSpeechNewParamsVoiceCoral   AudioSpeechNewParamsVoice = "coral"
	AudioSpeechNewParamsVoiceEcho    AudioSpeechNewParamsVoice = "echo"
	AudioSpeechNewParamsVoiceFable   AudioSpeechNewParamsVoice = "fable"
	AudioSpeechNewParamsVoiceOnyx    AudioSpeechNewParamsVoice = "onyx"
	AudioSpeechNewParamsVoiceNova    AudioSpeechNewParamsVoice = "nova"
	AudioSpeechNewParamsVoiceSage    AudioSpeechNewParamsVoice = "sage"
	AudioSpeechNewParamsVoiceShimmer AudioSpeechNewParamsVoice = "shimmer"
	AudioSpeechNewParamsVoiceVerse   AudioSpeechNewParamsVoice = "verse"
)

func (r AudioSpeechNewParamsVoice) IsKnown() bool {
	switch r {
	case AudioSpeechNewParamsVoiceAlloy, AudioSpeechNewParamsVoiceAsh, AudioSpeechNewParamsVoiceBallad, AudioSpeechNewParamsVoiceCoral, AudioSpeechNewParamsVoiceEcho, AudioSpeechNewParamsVoiceFable, AudioSpeechNewParamsVoiceOnyx, AudioSpeechNewParamsVoiceNova, AudioSpeechNewParamsVoiceSage, AudioSpeechNewParamsVoiceShimmer, AudioSpeechNewParamsVoiceVerse:
		return true
	}
	return false
}

// The format to audio in. Supported formats are `mp3`, `opus`, `aac`, `flac`,
// `wav`, and `pcm`.
type AudioSpeechNewParamsResponseFormat string

const (
	AudioSpeechNewParamsResponseFormatMP3  AudioSpeechNewParamsResponseFormat = "mp3"
	AudioSpeechNewParamsResponseFormatOpus AudioSpeechNewParamsResponseFormat = "opus"
	AudioSpeechNewParamsResponseFormatAac  AudioSpeechNewParamsResponseFormat = "aac"
	AudioSpeechNewParamsResponseFormatFlac AudioSpeechNewParamsResponseFormat = "flac"
	AudioSpeechNewParamsResponseFormatWav  AudioSpeechNewParamsResponseFormat = "wav"
	AudioSpeechNewParamsResponseFormatPcm  AudioSpeechNewParamsResponseFormat = "pcm"
)

func (r AudioSpeechNewParamsResponseFormat) IsKnown() bool {
	switch r {
	case AudioSpeechNewParamsResponseFormatMP3, AudioSpeechNewParamsResponseFormatOpus, AudioSpeechNewParamsResponseFormatAac, AudioSpeechNewParamsResponseFormatFlac, AudioSpeechNewParamsResponseFormatWav, AudioSpeechNewParamsResponseFormatPcm:
		return true
	}
	return false
}

// The format to stream the audio in. Supported formats are `sse` and `audio`.
// `sse` is not supported for `tts-1` or `tts-1-hd`.
type AudioSpeechNewParamsStreamFormat string

const (
	AudioSpeechNewParamsStreamFormatSse   AudioSpeechNewParamsStreamFormat = "sse"
	AudioSpeechNewParamsStreamFormatAudio AudioSpeechNewParamsStreamFormat = "audio"
)

func (r AudioSpeechNewParamsStreamFormat) IsKnown() bool {
	switch r {
	case AudioSpeechNewParamsStreamFormatSse, AudioSpeechNewParamsStreamFormatAudio:
		return true
	}
	return false
}
