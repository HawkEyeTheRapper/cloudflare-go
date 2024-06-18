// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// AIService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options []option.RequestOption
	Models  *AIModelService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r *AIService) {
	r = &AIService{}
	r.Options = opts
	r.Models = NewAIModelService(opts...)
	return
}

// This endpoint provides users with the capability to run specific AI models
// on-demand.
//
// By submitting the required input data, users can receive real-time predictions
// or results generated by the chosen AI model. The endpoint supports various AI
// model types, ensuring flexibility and adaptability for diverse use cases.
//
// Model specific inputs available in
// [Cloudflare Docs](https://developers.cloudflare.com/workers-ai/models/).
func (r *AIService) Run(ctx context.Context, modelName string, params AIRunParams, opts ...option.RequestOption) (res *AIRunResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AIRunResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if modelName == "" {
		err = errors.New("missing required model_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/%s", params.AccountID, modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [workers.AIRunResponseTextClassification],
// [shared.UnionString], [workers.AIRunResponseSentenceSimilarity],
// [workers.AIRunResponseTextEmbeddings], [workers.AIRunResponseSpeechRecognition],
// [workers.AIRunResponseImageClassification],
// [workers.AIRunResponseObjectDetection], [workers.AIRunResponseObject],
// [shared.UnionString], [workers.AIRunResponseTranslation],
// [workers.AIRunResponseSummarization] or [workers.AIRunResponseImageToText].
type AIRunResponseUnion interface {
	ImplementsWorkersAIRunResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTextClassification{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseSentenceSimilarity{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTextEmbeddings{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseSpeechRecognition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseImageClassification{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseObjectDetection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseTranslation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseSummarization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AIRunResponseImageToText{}),
		},
	)
}

type AIRunResponseTextClassification []AIRunResponseTextClassification

func (r AIRunResponseTextClassification) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseSentenceSimilarity []float64

func (r AIRunResponseSentenceSimilarity) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseTextEmbeddings struct {
	Data  [][]float64                     `json:"data"`
	Shape []float64                       `json:"shape"`
	JSON  aiRunResponseTextEmbeddingsJSON `json:"-"`
}

// aiRunResponseTextEmbeddingsJSON contains the JSON metadata for the struct
// [AIRunResponseTextEmbeddings]
type aiRunResponseTextEmbeddingsJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseTextEmbeddings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTextEmbeddingsJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseTextEmbeddings) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseSpeechRecognition struct {
	Text      string                               `json:"text,required"`
	Vtt       string                               `json:"vtt"`
	WordCount float64                              `json:"word_count"`
	Words     []AIRunResponseSpeechRecognitionWord `json:"words"`
	JSON      aiRunResponseSpeechRecognitionJSON   `json:"-"`
}

// aiRunResponseSpeechRecognitionJSON contains the JSON metadata for the struct
// [AIRunResponseSpeechRecognition]
type aiRunResponseSpeechRecognitionJSON struct {
	Text        apijson.Field
	Vtt         apijson.Field
	WordCount   apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseSpeechRecognition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseSpeechRecognitionJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseSpeechRecognition) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseSpeechRecognitionWord struct {
	End   float64                                `json:"end"`
	Start float64                                `json:"start"`
	Word  string                                 `json:"word"`
	JSON  aiRunResponseSpeechRecognitionWordJSON `json:"-"`
}

// aiRunResponseSpeechRecognitionWordJSON contains the JSON metadata for the struct
// [AIRunResponseSpeechRecognitionWord]
type aiRunResponseSpeechRecognitionWordJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseSpeechRecognitionWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseSpeechRecognitionWordJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseImageClassification []AIRunResponseImageClassification

func (r AIRunResponseImageClassification) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseObjectDetection []AIRunResponseObjectDetection

func (r AIRunResponseObjectDetection) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseObject struct {
	Response  string                        `json:"response"`
	ToolCalls []AIRunResponseObjectToolCall `json:"tool_calls"`
	JSON      aiRunResponseObjectJSON       `json:"-"`
}

// aiRunResponseObjectJSON contains the JSON metadata for the struct
// [AIRunResponseObject]
type aiRunResponseObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseObject) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseObjectToolCall struct {
	Arguments interface{}                     `json:"arguments"`
	Name      string                          `json:"name"`
	JSON      aiRunResponseObjectToolCallJSON `json:"-"`
}

// aiRunResponseObjectToolCallJSON contains the JSON metadata for the struct
// [AIRunResponseObjectToolCall]
type aiRunResponseObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseObjectToolCallJSON) RawJSON() string {
	return r.raw
}

type AIRunResponseTranslation struct {
	TranslatedText string                       `json:"translated_text"`
	JSON           aiRunResponseTranslationJSON `json:"-"`
}

// aiRunResponseTranslationJSON contains the JSON metadata for the struct
// [AIRunResponseTranslation]
type aiRunResponseTranslationJSON struct {
	TranslatedText apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIRunResponseTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseTranslationJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseTranslation) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseSummarization struct {
	Summary string                         `json:"summary"`
	JSON    aiRunResponseSummarizationJSON `json:"-"`
}

// aiRunResponseSummarizationJSON contains the JSON metadata for the struct
// [AIRunResponseSummarization]
type aiRunResponseSummarizationJSON struct {
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseSummarization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseSummarizationJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseSummarization) ImplementsWorkersAIRunResponseUnion() {}

type AIRunResponseImageToText struct {
	Description string                       `json:"description"`
	JSON        aiRunResponseImageToTextJSON `json:"-"`
}

// aiRunResponseImageToTextJSON contains the JSON metadata for the struct
// [AIRunResponseImageToText]
type aiRunResponseImageToTextJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseImageToText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseImageToTextJSON) RawJSON() string {
	return r.raw
}

func (r AIRunResponseImageToText) ImplementsWorkersAIRunResponseUnion() {}

type AIRunParams struct {
	AccountID param.Field[string]  `path:"account_id,required"`
	Body      AIRunParamsBodyUnion `json:"body,required"`
}

func (r AIRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AIRunParamsBody struct {
	Text        param.Field[interface{}] `json:"text,required"`
	Guidance    param.Field[float64]     `json:"guidance"`
	Image       param.Field[interface{}] `json:"image,required"`
	Mask        param.Field[interface{}] `json:"mask,required"`
	NumSteps    param.Field[int64]       `json:"num_steps"`
	Prompt      param.Field[string]      `json:"prompt"`
	Strength    param.Field[float64]     `json:"strength"`
	Sentences   param.Field[interface{}] `json:"sentences,required"`
	Source      param.Field[string]      `json:"source"`
	Audio       param.Field[interface{}] `json:"audio,required"`
	Lora        param.Field[string]      `json:"lora"`
	MaxTokens   param.Field[int64]       `json:"max_tokens"`
	Raw         param.Field[bool]        `json:"raw"`
	Stream      param.Field[bool]        `json:"stream"`
	Messages    param.Field[interface{}] `json:"messages,required"`
	SourceLang  param.Field[string]      `json:"source_lang"`
	TargetLang  param.Field[string]      `json:"target_lang"`
	InputText   param.Field[string]      `json:"input_text"`
	MaxLength   param.Field[int64]       `json:"max_length"`
	Temperature param.Field[float64]     `json:"temperature"`
}

func (r AIRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBody) implementsWorkersAIRunParamsBodyUnion() {}

// Satisfied by [workers.AIRunParamsBodyTextClassification],
// [workers.AIRunParamsBodyTextToImage],
// [workers.AIRunParamsBodySentenceSimilarity],
// [workers.AIRunParamsBodyTextEmbeddings],
// [workers.AIRunParamsBodySpeechRecognition],
// [workers.AIRunParamsBodyImageClassification],
// [workers.AIRunParamsBodyObjectDetection],
// [workers.AIRunParamsBodyTextGeneration], [workers.AIRunParamsBodyTranslation],
// [workers.AIRunParamsBodySummarization], [workers.AIRunParamsBodyImageToText],
// [AIRunParamsBody].
type AIRunParamsBodyUnion interface {
	implementsWorkersAIRunParamsBodyUnion()
}

type AIRunParamsBodyTextClassification struct {
	Text param.Field[string] `json:"text,required"`
}

func (r AIRunParamsBodyTextClassification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextClassification) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextToImage struct {
	Prompt   param.Field[string]    `json:"prompt,required"`
	Guidance param.Field[float64]   `json:"guidance"`
	Image    param.Field[[]float64] `json:"image"`
	Mask     param.Field[[]float64] `json:"mask"`
	NumSteps param.Field[int64]     `json:"num_steps"`
	Strength param.Field[float64]   `json:"strength"`
}

func (r AIRunParamsBodyTextToImage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextToImage) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodySentenceSimilarity struct {
	Sentences param.Field[[]string] `json:"sentences,required"`
	Source    param.Field[string]   `json:"source,required"`
}

func (r AIRunParamsBodySentenceSimilarity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodySentenceSimilarity) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextEmbeddings struct {
	Text param.Field[AIRunParamsBodyTextEmbeddingsTextUnion] `json:"text,required"`
}

func (r AIRunParamsBodyTextEmbeddings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextEmbeddings) implementsWorkersAIRunParamsBodyUnion() {}

// Satisfied by [shared.UnionString],
// [workers.AIRunParamsBodyTextEmbeddingsTextArray].
type AIRunParamsBodyTextEmbeddingsTextUnion interface {
	ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion()
}

type AIRunParamsBodyTextEmbeddingsTextArray []string

func (r AIRunParamsBodyTextEmbeddingsTextArray) ImplementsWorkersAIRunParamsBodyTextEmbeddingsTextUnion() {
}

type AIRunParamsBodySpeechRecognition struct {
	Audio param.Field[[]float64] `json:"audio,required"`
}

func (r AIRunParamsBodySpeechRecognition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodySpeechRecognition) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageClassification struct {
	Image param.Field[[]float64] `json:"image,required"`
}

func (r AIRunParamsBodyImageClassification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyImageClassification) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyObjectDetection struct {
	Image param.Field[[]float64] `json:"image"`
}

func (r AIRunParamsBodyObjectDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyObjectDetection) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextGeneration struct {
	Lora      param.Field[string]                                 `json:"lora"`
	MaxTokens param.Field[int64]                                  `json:"max_tokens"`
	Messages  param.Field[[]AIRunParamsBodyTextGenerationMessage] `json:"messages"`
	Prompt    param.Field[string]                                 `json:"prompt"`
	Raw       param.Field[bool]                                   `json:"raw"`
	Stream    param.Field[bool]                                   `json:"stream"`
}

func (r AIRunParamsBodyTextGeneration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTextGeneration) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyTextGenerationMessage struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIRunParamsBodyTextGenerationMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIRunParamsBodyTranslation struct {
	TargetLang param.Field[string] `json:"target_lang,required"`
	Text       param.Field[string] `json:"text,required"`
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AIRunParamsBodyTranslation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyTranslation) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodySummarization struct {
	InputText param.Field[string] `json:"input_text,required"`
	MaxLength param.Field[int64]  `json:"max_length"`
}

func (r AIRunParamsBodySummarization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodySummarization) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageToText struct {
	Image       param.Field[[]float64]                           `json:"image,required"`
	MaxTokens   param.Field[int64]                               `json:"max_tokens"`
	Messages    param.Field[[]AIRunParamsBodyImageToTextMessage] `json:"messages"`
	Prompt      param.Field[string]                              `json:"prompt"`
	Raw         param.Field[bool]                                `json:"raw"`
	Temperature param.Field[float64]                             `json:"temperature"`
}

func (r AIRunParamsBodyImageToText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AIRunParamsBodyImageToText) implementsWorkersAIRunParamsBodyUnion() {}

type AIRunParamsBodyImageToTextMessage struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIRunParamsBodyImageToTextMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIRunResponseEnvelope struct {
	Result AIRunResponseUnion        `json:"result" format:"binary"`
	JSON   aiRunResponseEnvelopeJSON `json:"-"`
}

// aiRunResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIRunResponseEnvelope]
type aiRunResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIRunResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiRunResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
