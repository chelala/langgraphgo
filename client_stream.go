/*
LangGraph Platform

# API Client for Stream API

API version: 0.1.0
*/
package langgraphgo

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tmaxmax/go-sse"
)

// SSE Version for api_thread_runs.go L1016
func (r ApiStreamRunThreadsThreadIdRunsStreamPostRequest) CreateSSEConnection() (*sse.Connection, error) {
	return r.ApiService.StreamRunThreadsThreadIdRunsStreamPostSSEConnect(r)
}

// SSE Version for api_thread_runs.go L1039
// Execute executes the request
//
//	@return string
func (a *ThreadRunsAPIService) StreamRunThreadsThreadIdRunsStreamPostSSEConnect(r ApiStreamRunThreadsThreadIdRunsStreamPostRequest) (*sse.Connection, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadRunsAPIService.StreamRunThreadsThreadIdRunsStreamPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/threads/{thread_id}/runs/stream"
	localVarPath = strings.Replace(localVarPath, "{"+"thread_id"+"}", url.PathEscape(parameterValueToString(r.threadId, "threadId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.runCreateStateful == nil {
		return nil, reportError("runCreateStateful is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/event-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.runCreateStateful
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	client := sse.Client{
		Backoff: sse.Backoff{
			MaxRetries: -1,
		},
	}
	conn := client.NewConnection(req)

	return conn, nil

}

func (c *APIClient) SSEDecode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if filePtr, ok := v.(**os.File); ok {
		tmpFile, err := os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return err
		}
		*filePtr = tmpFile
		_, err = tmpFile.Write(b)
		if err != nil {
			return err
		}
		_, err = tmpFile.Seek(0, io.SeekStart)
		return err
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}
