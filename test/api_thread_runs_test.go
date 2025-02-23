/*
LangGraph Platform

Testing ThreadRunsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package langgraphgo_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/chelala/langgraphgo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tmaxmax/go-sse"
)

func Test_langgraphgo_ThreadRunsAPIService(t *testing.T) {

	configuration := langgraphgo.NewConfiguration()
	configuration.Servers = langgraphgo.ServerConfigurations{
		{
			URL:         "http://127.0.0.1:2024",
			Description: "local running server",
		},
	}
	apiClient := langgraphgo.NewAPIClient(configuration)

	t.Run("Test ThreadRunsAPIService CancelRunHttpThreadsThreadIdRunsRunIdCancelPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.ThreadRunsAPI.CancelRunHttpThreadsThreadIdRunsRunIdCancelPost(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreadRunsAPIService CreateRunThreadsThreadIdRunsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		threadRunRequest := apiClient.ThreadRunsAPI.CreateRunThreadsThreadIdRunsPost(context.Background(), threadId)
		threadRunReq := threadRunRequest.RunCreateStateful(langgraphgo.RunCreateStateful{})
		resp, httpRes, err := threadRunReq.Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreadRunsAPIService DeleteRunThreadsThreadIdRunsRunIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.ThreadRunsAPI.DeleteRunThreadsThreadIdRunsRunIdDelete(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreadRunsAPIService GetRunHttpThreadsThreadIdRunsRunIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.ThreadRunsAPI.GetRunHttpThreadsThreadIdRunsRunIdGet(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreadRunsAPIService JoinRunHttpThreadsThreadIdRunsRunIdJoinGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.ThreadRunsAPI.JoinRunHttpThreadsThreadIdRunsRunIdJoinGet(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreadRunsAPIService ListRunsHttpThreadsThreadIdRunsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.ThreadRunsAPI.ListRunsHttpThreadsThreadIdRunsGet(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreadRunsAPIService StreamRunHttpThreadsThreadIdRunsRunIdJoinGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string
		var runId string

		resp, httpRes, err := apiClient.ThreadRunsAPI.StreamRunHttpThreadsThreadIdRunsRunIdJoinGet(context.Background(), threadId, runId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThreadRunsAPIService StreamRunThreadsThreadIdRunsStreamPost", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		threadId := "86de033f-49a6-44ce-97fc-4d3cca3865c4"
		graphName := "agent"
		inputId := uuid.New()
		// inputtext := "Cual es mi nombre?"
		// inputtext := "Que hora es?"
		// inputtext := "Cuál es la temperatura en  NYC?"
		// inputtext := "Cuál es la temperatura en  San Francisco?"
		inputtext := "Requisitos para un prestamo Micro?"

		// We prepare a context with a cancel function to stop the stream
		myctx, cancelMyContext := context.WithCancel(context.Background())

		threadRunsStream := apiClient.ThreadRunsAPI.StreamRunThreadsThreadIdRunsStreamPost(myctx, threadId)

		input := map[string]interface{}{
			"messages": []interface{}{
				map[string]interface{}{
					"id":      inputId.String(),
					"content": inputtext,
					"type":    "human",
					"name":    "Harold", // this is a test to see if the name is used
				},
			},
		}
		// streamModeText := "updates"
		streamModeArray := []string{"messages", "updates", "values"}
		streamMode := langgraphgo.StreamMode{
			// String: &streamModeText,
			ArrayOfString: &streamModeArray,
		}
		threadRunsStreamReq := threadRunsStream.RunCreateStateful(langgraphgo.RunCreateStateful{
			AssistantId: langgraphgo.RunCreateStatefulAssistantId{
				String: &graphName,
			},
			Input: *langgraphgo.NewNullableInput1(&langgraphgo.Input1{
				MapmapOfStringAny: &input,
			}),
			StreamMode: &streamMode,
		})

		// resp, httpRes, err := threadRunsStreamReq.Execute()
		conn, err := threadRunsStreamReq.CreateSSEConnection()
		if err != nil {
			fmt.Printf("Error creating SSE Connection: %s", err)
		}

		require.Nil(t, err)
		require.NotNil(t, conn)
		// assert.Equal(t, 200, httpRes.StatusCode)

		eventMessagesMetadata := map[string]string{}
		// var unsubscribe sse.EventCallbackRemover
		unsubscribe := conn.SubscribeToAll(func(event sse.Event) {
			// fmt.Printf("******************* Received event - LastEventID: %s, Type: %s *******************", event.LastEventID, event.Type)

			var eventDataI interface{}
			err := apiClient.SSEDecode(&eventDataI, []byte(event.Data), "application/json")
			if err != nil {
				fmt.Printf("Error decoding event data: %s", err)
				t.Fail()
			}

			sseEvent, err := langgraphgo.ParseSSEEvent(event.Type, event.Data)
			if err != nil {
				fmt.Printf("Error parsing event data: %s\n", err)
				t.Fail()
			}

			if event.Type == "metadata" {
				fmt.Printf("Run ID: %v\n", sseEvent.Metadata.RunID)
				fmt.Printf("Attempt: %v\n", sseEvent.Metadata.Attempt)
			}

			if event.Type == "messages/metadata" {
				for key, value := range *sseEvent.MessagesMetadata {
					eventMessagesMetadata[key] = value.Metadata.LanggraphNode
					fmt.Printf("Metadata for node: %v %v\n", key, value.Metadata.LanggraphNode)
				}
			}

			if event.Type == "messages/partial" {
				for _, value := range *sseEvent.MessagesPartial {
					role := eventMessagesMetadata[*value.ID]
					switch role {
					case "generate", "creditofull", "agent":
						fmt.Printf("Partial content of %v: %v\n", role, value.Content)
					default:
						fmt.Print(".")
					}
				}
			}

			if event.Type == "updates" {
				fmt.Print("Update\n")
			}

			if event.Type == "values" {
				fmt.Print("Values:\n")
				inputIdFound := false

				var messageRagContexts = sseEvent.Values.RagContexts

				for _, message := range sseEvent.Values.Messages {

					inputIdFound = inputIdFound || *message.ID == inputId.String()
					if inputIdFound && *message.ID != inputId.String() {
						fmt.Printf("Message ID: %v\n", *message.ID)
						fmt.Printf("Message Type: %v\n", message.Type)
						if message.Name != nil {
							fmt.Printf("Message Name: %v\n", *message.Name)
						}
						fmt.Printf("Message Content: %v\n", message.Content)

						// if the message is a rag message and there are rag contexts, print them
						if message.Name != nil && *message.Name == "rag" && messageRagContexts != nil {
							if messageRagContext, exists := messageRagContexts[*message.ID]; exists {
								for _, document := range messageRagContext {

									fmt.Printf("Document ID: %v\n", document.ID)
									fmt.Printf("Document Container: %v\n", document.Metadata.Container)
									fmt.Printf("Document Filename: %v\n", document.Metadata.Filename)
									fmt.Printf("Document Page Number: %v\n", document.Metadata.PageNumber)
									fmt.Printf("Document Content Type: %v\n", document.Metadata.ContentType)
									if len(document.PageContent) > 50 {
										fmt.Printf("Document Page Content: %v[...]\n", document.PageContent[:50])
									} else {
										fmt.Printf("Document Page Content: %v\n", document.PageContent)
									}
								}
							} else {
								fmt.Printf("inputId %s not found in Rag contexts", inputId)
							}
						}

					}

				}
			}
		})
		_ = cancelMyContext
		_ = unsubscribe

		if err := conn.Connect(); err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("Connection closed with <EOF>: %s", err)
			} else {
				fmt.Printf("Connection closed with error: %s", err)
			}
		}

	})

	t.Run("Test ThreadRunsAPIService WaitRunThreadsThreadIdRunsWaitPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		threadId := "86de033f-49a6-44ce-97fc-4d3cca3865c4"
		graphName := "agent"
		inputId := uuid.New()
		// inputtext := "Que hora es?"
		// inputtext := "Cuál es la temperatura en  NYC?"
		// inputtext := "Cuál es la temperatura en  San Francisco?"
		inputtext := "Requisitos para un prestamo Micro?"

		threadRunsWaitAPI := apiClient.ThreadRunsAPI.WaitRunThreadsThreadIdRunsWaitPost(context.Background(), threadId)

		input := map[string]interface{}{
			"messages": []interface{}{
				map[string]interface{}{
					"id":      inputId.String(),
					"content": inputtext,
					"type":    "human",
				},
			},
		}
		threadRunsWaitAPIReq := threadRunsWaitAPI.RunCreateStateful(langgraphgo.RunCreateStateful{
			AssistantId: langgraphgo.RunCreateStatefulAssistantId{
				String: &graphName,
			},
			Input: *langgraphgo.NewNullableInput1(&langgraphgo.Input1{
				MapmapOfStringAny: &input,
			}),
		})

		resp, httpRes, err := threadRunsWaitAPIReq.Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

		responseMap, ok := resp.(map[string]interface{})
		require.True(t, ok, "response is not a map[string]interface{}")
		messages, ok := responseMap["messages"].([]interface{})
		require.True(t, ok, "messages is not a []interface{}")
		inputIdFound := false

		var messageRagContexts map[string]interface{} = nil
		if ragContextsItem, exists := responseMap["rag_contexts"]; exists {
			_messageRagContexts, ok := ragContextsItem.(map[string]interface{})
			require.True(t, ok, "rag_contexts is not a map[string]interface{}")
			messageRagContexts = _messageRagContexts

		}

		for _, messageItem := range messages {
			message, ok := messageItem.(map[string]interface{})
			require.True(t, ok, "message is not a map[string]interface{}")

			messageId, ok := message["id"].(string)
			require.True(t, ok, "messageId is not a string")
			messageType, ok := message["type"].(string)
			require.True(t, ok, "messageType is not a string")
			messageContent, ok := message["content"].(string)
			require.True(t, ok, "messageContent is not a string")

			// if the key "name" does not exist or the value under "name" is not a string, the type assertion will return
			// the zero value of the target type (an empty string in this case) without triggering a runtime error.
			// The error (the boolean value) is ignored because it’s assigned to the blank identifier (_).
			messageName, _ := message["name"].(string)

			inputIdFound = inputIdFound || messageId == inputId.String()
			if inputIdFound && messageId != inputId.String() {
				fmt.Printf("Message ID: %v\n", messageId)
				fmt.Printf("Message Type: %v\n", messageType)
				fmt.Printf("Message Name: %v\n", messageName)
				fmt.Printf("Message Content: %v\n", messageContent)

				// if the message is a rag message and there are rag contexts, print them
				if messageName == "rag" && messageRagContexts != nil {
					if _messageRagContext, exists := messageRagContexts[inputId.String()]; exists {
						fmt.Printf("Found inputId %s has rag contexts\n", inputId)
						messageRagContext, ok := _messageRagContext.([]interface{})
						require.True(t, ok, "messageRagContext is not a []interface{}")
						for _, documentItem := range messageRagContext {
							document := documentItem.(map[string]interface{})
							metadata, ok := document["metadata"].(map[string]interface{})
							require.True(t, ok, "metadata is not a map[string]interface{}")
							documentId, ok := metadata["id"].(string)
							require.True(t, ok, "documentId is not a string")
							documentContainer, ok := metadata["container"].(string)
							require.True(t, ok, "documentContainer is not a string")
							documentFilename, ok := metadata["filename"].(string)
							require.True(t, ok, "documentFilename is not a string")
							documentPageNumber, ok := metadata["page_number"].(float64)
							require.True(t, ok, "documentPageNumber is not a float64")
							documentContentType, ok := metadata["content_type"].(string)
							require.True(t, ok, "documentContentType is not a string")
							documentPageContent, ok := document["page_content"].(string)
							require.True(t, ok, "documentPageContent is not a string")

							fmt.Printf("Document ID: %v\n", documentId)
							fmt.Printf("Document Container: %v\n", documentContainer)
							fmt.Printf("Document Filename: %v\n", documentFilename)
							fmt.Printf("Document Page Number: %v\n", documentPageNumber)
							fmt.Printf("Document Content Type: %v\n", documentContentType)
							if len(documentPageContent) > 20 {
								fmt.Printf("Document Page Content: %v\n", documentPageContent[:20])
							} else {
								fmt.Printf("Document Page Content: %v\n", documentPageContent)
							}
						}
					} else {
						fmt.Printf("inputId %s not found in Rag contexts", inputId)
					}
				}

			}

		}

		fmt.Printf("Input ID: %v\n", inputId)
		// fmt.Printf("Response: %v", resp)

	})

}

func processEventMessages(t *testing.T, responseMap map[string]interface{}, inputId uuid.UUID) {
	messages, ok := responseMap["messages"].([]interface{})
	require.True(t, ok, "messages is not a []interface{}")
	inputIdFound := false

	var messageRagContexts map[string]interface{} = nil
	if ragContextsItem, exists := responseMap["rag_contexts"]; exists {
		_messageRagContexts, ok := ragContextsItem.(map[string]interface{})
		require.True(t, ok, "rag_contexts is not a map[string]interface{}")
		messageRagContexts = _messageRagContexts

	}

	for _, messageItem := range messages {
		message, ok := messageItem.(map[string]interface{})
		require.True(t, ok, "message is not a map[string]interface{}")

		messageId, ok := message["id"].(string)
		require.True(t, ok, "messageId is not a string")
		messageType, ok := message["type"].(string)
		require.True(t, ok, "messageType is not a string")
		messageContent, ok := message["content"].(string)
		require.True(t, ok, "messageContent is not a string")

		// if the key "name" does not exist or the value under "name" is not a string, the type assertion will return
		// the zero value of the target type (an empty string in this case) without triggering a runtime error.
		// The error (the boolean value) is ignored because it’s assigned to the blank identifier (_).
		messageName, _ := message["name"].(string)

		inputIdFound = inputIdFound || messageId == inputId.String()
		if inputIdFound && messageId != inputId.String() {
			fmt.Printf("Message ID: %v\n", messageId)
			fmt.Printf("Message Type: %v\n", messageType)
			fmt.Printf("Message Name: %v\n", messageName)
			fmt.Printf("Message Content: %v\n", messageContent)

			// if the message is a rag message and there are rag contexts, print them
			if messageName == "rag" && messageRagContexts != nil {
				if _messageRagContext, exists := messageRagContexts[inputId.String()]; exists {
					fmt.Printf("Found inputId %s has rag contexts\n", inputId)
					messageRagContext, ok := _messageRagContext.([]interface{})
					require.True(t, ok, "messageRagContext is not a []interface{}")
					for _, documentItem := range messageRagContext {
						document := documentItem.(map[string]interface{})
						metadata, ok := document["metadata"].(map[string]interface{})
						require.True(t, ok, "metadata is not a map[string]interface{}")
						documentId, ok := metadata["id"].(string)
						require.True(t, ok, "documentId is not a string")
						documentContainer, ok := metadata["container"].(string)
						require.True(t, ok, "documentContainer is not a string")
						documentFilename, ok := metadata["filename"].(string)
						require.True(t, ok, "documentFilename is not a string")
						documentPageNumber, ok := metadata["page_number"].(float64)
						require.True(t, ok, "documentPageNumber is not a float64")
						documentContentType, ok := metadata["content_type"].(string)
						require.True(t, ok, "documentContentType is not a string")
						documentPageContent, ok := document["page_content"].(string)
						require.True(t, ok, "documentPageContent is not a string")

						fmt.Printf("Document ID: %v\n", documentId)
						fmt.Printf("Document Container: %v\n", documentContainer)
						fmt.Printf("Document Filename: %v\n", documentFilename)
						fmt.Printf("Document Page Number: %v\n", documentPageNumber)
						fmt.Printf("Document Content Type: %v\n", documentContentType)
						if len(documentPageContent) > 20 {
							fmt.Printf("Document Page Content: %v\n", documentPageContent[:20])
						} else {
							fmt.Printf("Document Page Content: %v\n", documentPageContent)
						}
					}
				} else {
					fmt.Printf("inputId %s not found in Rag contexts", inputId)
				}
			}

		}

	}
}
