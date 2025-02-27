/*
LangGraph Platform

# SSE response utils for API Client for Stream API

API version: 0.1.0
*/
package langgraphgo

import (
	"encoding/json"
	"fmt"
)

// Here are example structs that match the JSON shapes you'd expect for specific Types.
// For instance, "metadata" might look like this:
type Metadata struct {
	RunID   string `json:"run_id"`
	Attempt int    `json:"attempt"`
}

// Another possible shape: "values" (as an example).
type Message struct {
	Content          string                 `json:"content"`
	AdditionalKwargs map[string]interface{} `json:"additional_kwargs"`
	ResponseMetadata map[string]interface{} `json:"response_metadata"`
	Type             string                 `json:"type"`
	Name             *string                `json:"name"`
	ID               *string                `json:"id"`
	Example          bool                   `json:"example"`
}

// MessageMetadata represents the detailed metadata for messages/metadata event.
type MessageMetadata struct {
	CreatedBy              string                 `json:"created_by"`
	GraphID                string                 `json:"graph_id"`
	AssistantID            string                 `json:"assistant_id"`
	RunAttempt             int                    `json:"run_attempt"`
	LanggraphVersion       string                 `json:"langgraph_version"`
	LanggraphPlan          string                 `json:"langgraph_plan"`
	LanggraphHost          string                 `json:"langgraph_host"`
	LanggraphAuthUserID    string                 `json:"langgraph_auth_user_id"`
	ThreadID               string                 `json:"thread_id"`
	CheckpointID           string                 `json:"checkpoint_id"`
	CheckpointNs           string                 `json:"checkpoint_ns"`
	XAuthScheme            string                 `json:"x-auth-scheme"`
	RunID                  string                 `json:"run_id"`
	UserID                 string                 `json:"user_id"`
	LanggraphStep          int                    `json:"langgraph_step"`
	LanggraphNode          string                 `json:"langgraph_node"`
	LanggraphTriggers      []string               `json:"langgraph_triggers"`
	LanggraphPath          []string               `json:"langgraph_path"`
	LanggraphCheckpointNs  string                 `json:"langgraph_checkpoint_ns"`
	LSProvider             string                 `json:"ls_provider"`
	LSModelName            string                 `json:"ls_model_name"`
	LSModelType            string                 `json:"ls_model_type"`
	LSTemperature          float64                `json:"ls_temperature"`
	StructuredOutputFormat map[string]interface{} `json:"structured_output_format"`
}

// MessagesMetadata holds metadata keyed by dynamic run ids.
type MessagesMetadata map[string]struct {
	Metadata MessageMetadata `json:"metadata"`
}

type SSEValues struct {
	Next                string                `json:"next"`
	Messages            []Message             `json:"messages"`
	RagContainers       map[string]string     `json:"rag_containers"`
	RagQuestions        map[string]string     `json:"rag_questions"`
	RagContexts         map[string][]Document `json:"rag_contexts"`
	LastGenerationGrade string                `json:"last_generation_grade"`
}

type MessagesPartial struct {
	Content          string                 `json:"content"`
	AdditionalKwargs map[string]interface{} `json:"additional_kwargs"`
	ResponseMetadata map[string]interface{} `json:"response_metadata"`
	Type             string                 `json:"type"`
	Name             *string                `json:"name"`
	ID               *string                `json:"id"`
	Example          bool                   `json:"example"`
	ToolCalls        []interface{}          `json:"tool_calls"`
	InvalidToolCalls []interface{}          `json:"invalid_tool_calls"`
	UsageMetadata    interface{}            `json:"usage_metadata"`
}

type MessagesPartialList []MessagesPartial

type MessagesCompleteList []Message

type GradeGeneration struct {
	Messages            []Message `json:"messages"`
	LastGenerationGrade string    `json:"last_generation_grade"`
}

type GenerateUpdate struct {
	RagAnswers map[string]string `json:"rag_answers"`
}

type SupervisorUpdate struct {
	// Define fields as needed or leave it empty if not used.
}

type Captions struct {
	Text       string `json:"text"`
	Highlights string `json:"highlights"`
}

type DocumentMetadata struct {
	ID          string   `json:"id"`
	Filename    string   `json:"filename"`
	PageNumber  int      `json:"page_number"`
	Container   string   `json:"container"`
	Blob        string   `json:"blob"`
	ContentType string   `json:"content_type"`
	Captions    Captions `json:"captions"`
	Answers     string   `json:"answers"`
}

type Document struct {
	ID          interface{}       `json:"id"`
	Metadata    *DocumentMetadata `json:"metadata"`
	PageContent string            `json:"page_content"`
	Type        string            `json:"type"`
}

type GradeDocumentsUpdate struct {
	RagContexts map[string][]Document `json:"rag_contexts"`
}

type Updates struct {
	GradeGeneration       *GradeGeneration      `json:"grade_generation,omitempty"`
	SummarizeConversation interface{}           `json:"summarize_conversation,omitempty"`
	Supervisor            *SupervisorUpdate     `json:"supervisor,omitempty"`
	Generate              *GenerateUpdate       `json:"generate,omitempty"`
	GradeDocuments        *GradeDocumentsUpdate `json:"grade_documents,omitempty"`
}

// SSEEvent is a “wrapper” type that holds all possible typed fields in optional pointers.
// You can add more fields and structs to match other Types (e.g., "messages/partial", "updates", etc.).
type SSEEvent struct {
	Type             string                // echo the Type for reference
	Metadata         *Metadata             // populated if Type == "metadata"
	Values           *SSEValues            // populated if Type == "values"
	MessagesPartial  *MessagesPartialList  // populated if Type == "messages/partial"
	MessagesMetadata *MessagesMetadata     // populated if Type == "messages/metadata"
	Updates          *Updates              // populated if Type == "updates"
	Complete         *MessagesCompleteList // populated if Type == "messages/complete"
}

// ParseSSEEvent is the main function. It checks the Type, then JSON-unmarshals Data
// into the appropriate nested struct. The caller never has to do a type switch; they
// can just check e.Metadata or e.Values, etc., as long as they know the Type.
func ParseSSEEvent(eventType, data string) (*SSEEvent, error) {
	e := &SSEEvent{Type: eventType}

	switch eventType {
	case "metadata":
		var m Metadata
		if err := json.Unmarshal([]byte(data), &m); err != nil {
			return nil, fmt.Errorf("unmarshalling metadata failed: %w", err)
		}
		e.Metadata = &m

	case "values":
		var v SSEValues
		if err := json.Unmarshal([]byte(data), &v); err != nil {
			return nil, fmt.Errorf("unmarshalling values failed: %w", err)
		}
		e.Values = &v

	case "messages/metadata":
		var mm MessagesMetadata
		if err := json.Unmarshal([]byte(data), &mm); err != nil {
			return nil, fmt.Errorf("unmarshalling messages/metadata failed: %w", err)
		}
		e.MessagesMetadata = &mm

	case "messages/partial":
		var mp MessagesPartialList
		if err := json.Unmarshal([]byte(data), &mp); err != nil {
			return nil, fmt.Errorf("unmarshalling messages/partial failed: %w", err)
		}
		e.MessagesPartial = &mp

	case "updates":
		var upd Updates
		if err := json.Unmarshal([]byte(data), &upd); err != nil {
			return nil, fmt.Errorf("unmarshalling updates failed: %w", err)
		}
		e.Updates = &upd

	case "messages/complete":
		var mc MessagesCompleteList
		if err := json.Unmarshal([]byte(data), &mc); err != nil {
			return nil, fmt.Errorf("unmarshalling messages/complete failed: %w", err)
		}
		e.Complete = &mc

	default:
		// You could either treat unknown Types as an error or just store
		// the raw JSON for debugging.
		// return nil, errors.New("unknown SSE event type: " + eventType)
		return nil, nil
	}

	return e, nil
}
