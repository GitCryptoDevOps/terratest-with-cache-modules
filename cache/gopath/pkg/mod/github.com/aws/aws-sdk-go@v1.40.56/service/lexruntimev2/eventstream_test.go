// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

//go:build go1.16
// +build go1.16

package lexruntimev2

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/unit"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go/private/protocol/eventstream/eventstreamtest"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

var _ time.Time
var _ awserr.Error
var _ context.Context
var _ sync.WaitGroup
var _ strings.Reader

func TestStartConversation_Read(t *testing.T) {
	expectEvents, eventMsgs := mockStartConversationReadEvents()
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}
	defer resp.GetStream().Close()

	var i int
	for event := range resp.GetStream().Events() {
		if event == nil {
			t.Errorf("%d, expect event, got nil", i)
		}
		if e, a := expectEvents[i], event; !reflect.DeepEqual(e, a) {
			t.Errorf("%d, expect %T %v, got %T %v", i, e, e, a, a)
		}
		i++
	}

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func TestStartConversation_ReadClose(t *testing.T) {
	_, eventMsgs := mockStartConversationReadEvents()
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	// Assert calling Err before close does not close the stream.
	resp.GetStream().Err()
	select {
	case _, ok := <-resp.GetStream().Events():
		if !ok {
			t.Fatalf("expect stream not to be closed, but was")
		}
	default:
	}

	resp.GetStream().Close()
	<-resp.GetStream().Events()

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func TestStartConversation_ReadUnknownEvent(t *testing.T) {
	expectEvents, eventMsgs := mockStartConversationReadEvents()
	var eventOffset int

	unknownEvent := eventstream.Message{
		Headers: eventstream.Headers{
			eventstreamtest.EventMessageTypeHeader,
			{
				Name:  eventstreamapi.EventTypeHeader,
				Value: eventstream.StringValue("UnknownEventName"),
			},
		},
		Payload: []byte("some unknown event"),
	}

	eventMsgs = append(eventMsgs[:eventOffset],
		append([]eventstream.Message{unknownEvent}, eventMsgs[eventOffset:]...)...)

	expectEvents = append(expectEvents[:eventOffset],
		append([]StartConversationResponseEventStreamEvent{
			&StartConversationResponseEventStreamUnknownEvent{
				Type:    "UnknownEventName",
				Message: unknownEvent,
			},
		},
			expectEvents[eventOffset:]...)...)

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}
	defer resp.GetStream().Close()

	var i int
	for event := range resp.GetStream().Events() {
		if event == nil {
			t.Errorf("%d, expect event, got nil", i)
		}
		if e, a := expectEvents[i], event; !reflect.DeepEqual(e, a) {
			t.Errorf("%d, expect %T %v, got %T %v", i, e, e, a, a)
		}
		i++
	}

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func BenchmarkStartConversation_Read(b *testing.B) {
	_, eventMsgs := mockStartConversationReadEvents()
	var buf bytes.Buffer
	encoder := eventstream.NewEncoder(&buf)
	for _, msg := range eventMsgs {
		if err := encoder.Encode(msg); err != nil {
			b.Fatalf("failed to encode message, %v", err)
		}
	}
	stream := &loopReader{source: bytes.NewReader(buf.Bytes())}

	sess := unit.Session
	svc := New(sess, &aws.Config{
		Endpoint:               aws.String("https://example.com"),
		DisableParamValidation: aws.Bool(true),
	})
	svc.Handlers.Send.Swap(corehandlers.SendHandler.Name,
		request.NamedHandler{Name: "mockSend",
			Fn: func(r *request.Request) {
				r.HTTPResponse = &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
					Header:     http.Header{},
					Body:       ioutil.NopCloser(stream),
				}
			},
		},
	)

	resp, err := svc.StartConversation(nil)
	if err != nil {
		b.Fatalf("failed to create request, %v", err)
	}
	defer resp.GetStream().Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = resp.GetStream().Err(); err != nil {
			b.Fatalf("expect no error, got %v", err)
		}
		event := <-resp.GetStream().Events()
		if event == nil {
			b.Fatalf("expect event, got nil, %v, %d", resp.GetStream().Err(), i)
		}
	}
}

func mockStartConversationReadEvents() (
	[]StartConversationResponseEventStreamEvent,
	[]eventstream.Message,
) {
	expectEvents := []StartConversationResponseEventStreamEvent{
		&AudioResponseEvent{
			AudioChunk:  []byte("blob value goes here"),
			ContentType: aws.String("string value goes here"),
			EventId:     aws.String("string value goes here"),
		},
		&HeartbeatEvent{
			EventId: aws.String("string value goes here"),
		},
		&IntentResultEvent{
			EventId:   aws.String("string value goes here"),
			InputMode: aws.String("string value goes here"),
			Interpretations: []*Interpretation{
				{
					Intent: &Intent{
						ConfirmationState: aws.String("string value goes here"),
						Name:              aws.String("string value goes here"),
						Slots: map[string]*Slot{
							"a": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
							"b": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
							"c": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
						},
						State: aws.String("string value goes here"),
					},
					NluConfidence: &ConfidenceScore{
						Score: aws.Float64(123.45),
					},
					SentimentResponse: &SentimentResponse{
						Sentiment: aws.String("string value goes here"),
						SentimentScore: &SentimentScore{
							Mixed:    aws.Float64(123.45),
							Negative: aws.Float64(123.45),
							Neutral:  aws.Float64(123.45),
							Positive: aws.Float64(123.45),
						},
					},
				},
				{
					Intent: &Intent{
						ConfirmationState: aws.String("string value goes here"),
						Name:              aws.String("string value goes here"),
						Slots: map[string]*Slot{
							"a": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
							"b": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
							"c": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
						},
						State: aws.String("string value goes here"),
					},
					NluConfidence: &ConfidenceScore{
						Score: aws.Float64(123.45),
					},
					SentimentResponse: &SentimentResponse{
						Sentiment: aws.String("string value goes here"),
						SentimentScore: &SentimentScore{
							Mixed:    aws.Float64(123.45),
							Negative: aws.Float64(123.45),
							Neutral:  aws.Float64(123.45),
							Positive: aws.Float64(123.45),
						},
					},
				},
				{
					Intent: &Intent{
						ConfirmationState: aws.String("string value goes here"),
						Name:              aws.String("string value goes here"),
						Slots: map[string]*Slot{
							"a": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
							"b": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
							"c": {
								Shape: aws.String("string value goes here"),
								Value: &Value{
									InterpretedValue: aws.String("string value goes here"),
									OriginalValue:    aws.String("string value goes here"),
									ResolvedValues: []*string{
										aws.String("string value goes here"),
										aws.String("string value goes here"),
										aws.String("string value goes here"),
									},
								},
								Values: []*Slot{},
							},
						},
						State: aws.String("string value goes here"),
					},
					NluConfidence: &ConfidenceScore{
						Score: aws.Float64(123.45),
					},
					SentimentResponse: &SentimentResponse{
						Sentiment: aws.String("string value goes here"),
						SentimentScore: &SentimentScore{
							Mixed:    aws.Float64(123.45),
							Negative: aws.Float64(123.45),
							Neutral:  aws.Float64(123.45),
							Positive: aws.Float64(123.45),
						},
					},
				},
			},
			RequestAttributes: map[string]*string{
				"a": aws.String("string value goes here"),
				"b": aws.String("string value goes here"),
				"c": aws.String("string value goes here"),
			},
			SessionId: aws.String("string value goes here"),
			SessionState: &SessionState{
				ActiveContexts: []*ActiveContext{
					{
						ContextAttributes: map[string]*string{
							"a": aws.String("string value goes here"),
							"b": aws.String("string value goes here"),
							"c": aws.String("string value goes here"),
						},
						Name: aws.String("string value goes here"),
						TimeToLive: &ActiveContextTimeToLive{
							TimeToLiveInSeconds: aws.Int64(123),
							TurnsToLive:         aws.Int64(123),
						},
					},
					{
						ContextAttributes: map[string]*string{
							"a": aws.String("string value goes here"),
							"b": aws.String("string value goes here"),
							"c": aws.String("string value goes here"),
						},
						Name: aws.String("string value goes here"),
						TimeToLive: &ActiveContextTimeToLive{
							TimeToLiveInSeconds: aws.Int64(123),
							TurnsToLive:         aws.Int64(123),
						},
					},
					{
						ContextAttributes: map[string]*string{
							"a": aws.String("string value goes here"),
							"b": aws.String("string value goes here"),
							"c": aws.String("string value goes here"),
						},
						Name: aws.String("string value goes here"),
						TimeToLive: &ActiveContextTimeToLive{
							TimeToLiveInSeconds: aws.Int64(123),
							TurnsToLive:         aws.Int64(123),
						},
					},
				},
				DialogAction: &DialogAction{
					SlotToElicit: aws.String("string value goes here"),
					Type:         aws.String("string value goes here"),
				},
				Intent: &Intent{
					ConfirmationState: aws.String("string value goes here"),
					Name:              aws.String("string value goes here"),
					Slots: map[string]*Slot{
						"a": {
							Shape: aws.String("string value goes here"),
							Value: &Value{
								InterpretedValue: aws.String("string value goes here"),
								OriginalValue:    aws.String("string value goes here"),
								ResolvedValues: []*string{
									aws.String("string value goes here"),
									aws.String("string value goes here"),
									aws.String("string value goes here"),
								},
							},
							Values: []*Slot{},
						},
						"b": {
							Shape: aws.String("string value goes here"),
							Value: &Value{
								InterpretedValue: aws.String("string value goes here"),
								OriginalValue:    aws.String("string value goes here"),
								ResolvedValues: []*string{
									aws.String("string value goes here"),
									aws.String("string value goes here"),
									aws.String("string value goes here"),
								},
							},
							Values: []*Slot{},
						},
						"c": {
							Shape: aws.String("string value goes here"),
							Value: &Value{
								InterpretedValue: aws.String("string value goes here"),
								OriginalValue:    aws.String("string value goes here"),
								ResolvedValues: []*string{
									aws.String("string value goes here"),
									aws.String("string value goes here"),
									aws.String("string value goes here"),
								},
							},
							Values: []*Slot{},
						},
					},
					State: aws.String("string value goes here"),
				},
				OriginatingRequestId: aws.String("string value goes here"),
				SessionAttributes: map[string]*string{
					"a": aws.String("string value goes here"),
					"b": aws.String("string value goes here"),
					"c": aws.String("string value goes here"),
				},
			},
		},
		&PlaybackInterruptionEvent{
			CausedByEventId: aws.String("string value goes here"),
			EventId:         aws.String("string value goes here"),
			EventReason:     aws.String("string value goes here"),
		},
		&TextResponseEvent{
			EventId: aws.String("string value goes here"),
			Messages: []*Message{
				{
					Content:     aws.String("string value goes here"),
					ContentType: aws.String("string value goes here"),
					ImageResponseCard: &ImageResponseCard{
						Buttons: []*Button{
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
						},
						ImageUrl: aws.String("string value goes here"),
						Subtitle: aws.String("string value goes here"),
						Title:    aws.String("string value goes here"),
					},
				},
				{
					Content:     aws.String("string value goes here"),
					ContentType: aws.String("string value goes here"),
					ImageResponseCard: &ImageResponseCard{
						Buttons: []*Button{
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
						},
						ImageUrl: aws.String("string value goes here"),
						Subtitle: aws.String("string value goes here"),
						Title:    aws.String("string value goes here"),
					},
				},
				{
					Content:     aws.String("string value goes here"),
					ContentType: aws.String("string value goes here"),
					ImageResponseCard: &ImageResponseCard{
						Buttons: []*Button{
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
						},
						ImageUrl: aws.String("string value goes here"),
						Subtitle: aws.String("string value goes here"),
						Title:    aws.String("string value goes here"),
					},
				},
			},
		},
		&TranscriptEvent{
			EventId:    aws.String("string value goes here"),
			Transcript: aws.String("string value goes here"),
		},
	}

	var marshalers request.HandlerList
	marshalers.PushBackNamed(restjson.BuildHandler)
	payloadMarshaler := protocol.HandlerPayloadMarshal{
		Marshalers: marshalers,
	}
	_ = payloadMarshaler

	eventMsgs := []eventstream.Message{
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("AudioResponseEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[0]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("HeartbeatEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[1]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("IntentResultEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[2]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("PlaybackInterruptionEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[3]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("TextResponseEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[4]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("TranscriptEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[5]),
		},
	}

	return expectEvents, eventMsgs
}
func TestStartConversation_ReadException(t *testing.T) {
	expectEvents := []StartConversationResponseEventStreamEvent{
		&AccessDeniedException{
			RespMetadata: protocol.ResponseMetadata{
				StatusCode: 200,
			},
			Message_: aws.String("string value goes here"),
		},
	}

	var marshalers request.HandlerList
	marshalers.PushBackNamed(restjson.BuildHandler)
	payloadMarshaler := protocol.HandlerPayloadMarshal{
		Marshalers: marshalers,
	}

	eventMsgs := []eventstream.Message{
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventExceptionTypeHeader,
				{
					Name:  eventstreamapi.ExceptionTypeHeader,
					Value: eventstream.StringValue("AccessDeniedException"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, expectEvents[0]),
		},
	}

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:      t,
			Events: eventMsgs,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	defer resp.GetStream().Close()

	<-resp.GetStream().Events()

	err = resp.GetStream().Err()
	if err == nil {
		t.Fatalf("expect err, got none")
	}

	expectErr := &AccessDeniedException{
		RespMetadata: protocol.ResponseMetadata{
			StatusCode: 200,
		},
		Message_: aws.String("string value goes here"),
	}
	aerr, ok := err.(awserr.Error)
	if !ok {
		t.Errorf("expect exception, got %T, %#v", err, err)
	}
	if e, a := expectErr.Code(), aerr.Code(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectErr.Message(), aerr.Message(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	if e, a := expectErr, aerr; !reflect.DeepEqual(e, a) {
		t.Errorf("expect error %+#v, got %+#v", e, a)
	}
}

var _ awserr.Error = (*AccessDeniedException)(nil)
var _ awserr.Error = (*BadGatewayException)(nil)
var _ awserr.Error = (*ConflictException)(nil)
var _ awserr.Error = (*DependencyFailedException)(nil)
var _ awserr.Error = (*InternalServerException)(nil)
var _ awserr.Error = (*ResourceNotFoundException)(nil)
var _ awserr.Error = (*ThrottlingException)(nil)
var _ awserr.Error = (*ValidationException)(nil)

type loopReader struct {
	source *bytes.Reader
}

func (c *loopReader) Read(p []byte) (int, error) {
	if c.source.Len() == 0 {
		c.source.Seek(0, 0)
	}

	return c.source.Read(p)
}

func TestStartConversation_Write(t *testing.T) {
	clientEvents, expectedClientEvents := mockStartConversationWriteEvents()

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		&eventstreamtest.ServeEventStream{
			T:             t,
			ClientEvents:  expectedClientEvents,
			BiDirectional: true,
		},
		true)
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	stream := resp.GetStream()

	for _, event := range clientEvents {
		err = stream.Send(context.Background(), event)
		if err != nil {
			t.Fatalf("expect no error, got %v", err)
		}
	}

	if err := stream.Close(); err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestStartConversation_WriteClose(t *testing.T) {
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{T: t, BiDirectional: true},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	// Assert calling Err before close does not close the stream.
	resp.GetStream().Err()

	err = resp.GetStream().Send(context.Background(), &AudioInputEvent{})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	resp.GetStream().Close()

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func TestStartConversation_WriteError(t *testing.T) {
	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		eventstreamtest.ServeEventStream{
			T:               t,
			BiDirectional:   true,
			ForceCloseAfter: time.Millisecond * 500,
		},
		true,
	)
	if err != nil {
		t.Fatalf("expect no error, %v", err)
	}
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error got, %v", err)
	}

	defer resp.GetStream().Close()

	for {
		err = resp.GetStream().Send(context.Background(), &AudioInputEvent{})
		if err != nil {
			if strings.Contains("unable to send event", err.Error()) {
				t.Errorf("expected stream closed error, got %v", err)
			}
			break
		}
	}
}

func TestStartConversation_ReadWrite(t *testing.T) {
	expectedServiceEvents, serviceEvents := mockStartConversationReadEvents()
	clientEvents, expectedClientEvents := mockStartConversationWriteEvents()

	sess, cleanupFn, err := eventstreamtest.SetupEventStreamSession(t,
		&eventstreamtest.ServeEventStream{
			T:             t,
			ClientEvents:  expectedClientEvents,
			Events:        serviceEvents,
			BiDirectional: true,
		},
		true)
	defer cleanupFn()

	svc := New(sess)
	resp, err := svc.StartConversation(nil)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	stream := resp.GetStream()
	defer stream.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		var i int
		for event := range resp.GetStream().Events() {
			if event == nil {
				t.Errorf("%d, expect event, got nil", i)
			}
			if e, a := expectedServiceEvents[i], event; !reflect.DeepEqual(e, a) {
				t.Errorf("%d, expect %T %v, got %T %v", i, e, e, a, a)
			}
			i++
		}
	}()

	for _, event := range clientEvents {
		err = stream.Send(context.Background(), event)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	resp.GetStream().Close()

	wg.Wait()

	if err := resp.GetStream().Err(); err != nil {
		t.Errorf("expect no error, %v", err)
	}
}

func mockStartConversationWriteEvents() (
	[]StartConversationRequestEventStreamEvent,
	[]eventstream.Message,
) {
	inputEvents := []StartConversationRequestEventStreamEvent{
		&AudioInputEvent{
			AudioChunk:            []byte("blob value goes here"),
			ClientTimestampMillis: aws.Int64(1234),
			ContentType:           aws.String("string value goes here"),
			EventId:               aws.String("string value goes here"),
		},
		&ConfigurationEvent{
			ClientTimestampMillis: aws.Int64(1234),
			DisablePlayback:       aws.Bool(true),
			EventId:               aws.String("string value goes here"),
			RequestAttributes: map[string]*string{
				"a": aws.String("string value goes here"),
				"b": aws.String("string value goes here"),
				"c": aws.String("string value goes here"),
			},
			ResponseContentType: aws.String("string value goes here"),
			SessionState: &SessionState{
				ActiveContexts: []*ActiveContext{
					{
						ContextAttributes: map[string]*string{
							"a": aws.String("string value goes here"),
							"b": aws.String("string value goes here"),
							"c": aws.String("string value goes here"),
						},
						Name: aws.String("string value goes here"),
						TimeToLive: &ActiveContextTimeToLive{
							TimeToLiveInSeconds: aws.Int64(123),
							TurnsToLive:         aws.Int64(123),
						},
					},
					{
						ContextAttributes: map[string]*string{
							"a": aws.String("string value goes here"),
							"b": aws.String("string value goes here"),
							"c": aws.String("string value goes here"),
						},
						Name: aws.String("string value goes here"),
						TimeToLive: &ActiveContextTimeToLive{
							TimeToLiveInSeconds: aws.Int64(123),
							TurnsToLive:         aws.Int64(123),
						},
					},
					{
						ContextAttributes: map[string]*string{
							"a": aws.String("string value goes here"),
							"b": aws.String("string value goes here"),
							"c": aws.String("string value goes here"),
						},
						Name: aws.String("string value goes here"),
						TimeToLive: &ActiveContextTimeToLive{
							TimeToLiveInSeconds: aws.Int64(123),
							TurnsToLive:         aws.Int64(123),
						},
					},
				},
				DialogAction: &DialogAction{
					SlotToElicit: aws.String("string value goes here"),
					Type:         aws.String("string value goes here"),
				},
				Intent: &Intent{
					ConfirmationState: aws.String("string value goes here"),
					Name:              aws.String("string value goes here"),
					Slots: map[string]*Slot{
						"a": {
							Shape: aws.String("string value goes here"),
							Value: &Value{
								InterpretedValue: aws.String("string value goes here"),
								OriginalValue:    aws.String("string value goes here"),
								ResolvedValues: []*string{
									aws.String("string value goes here"),
									aws.String("string value goes here"),
									aws.String("string value goes here"),
								},
							},
							Values: []*Slot{},
						},
						"b": {
							Shape: aws.String("string value goes here"),
							Value: &Value{
								InterpretedValue: aws.String("string value goes here"),
								OriginalValue:    aws.String("string value goes here"),
								ResolvedValues: []*string{
									aws.String("string value goes here"),
									aws.String("string value goes here"),
									aws.String("string value goes here"),
								},
							},
							Values: []*Slot{},
						},
						"c": {
							Shape: aws.String("string value goes here"),
							Value: &Value{
								InterpretedValue: aws.String("string value goes here"),
								OriginalValue:    aws.String("string value goes here"),
								ResolvedValues: []*string{
									aws.String("string value goes here"),
									aws.String("string value goes here"),
									aws.String("string value goes here"),
								},
							},
							Values: []*Slot{},
						},
					},
					State: aws.String("string value goes here"),
				},
				OriginatingRequestId: aws.String("string value goes here"),
				SessionAttributes: map[string]*string{
					"a": aws.String("string value goes here"),
					"b": aws.String("string value goes here"),
					"c": aws.String("string value goes here"),
				},
			},
			WelcomeMessages: []*Message{
				{
					Content:     aws.String("string value goes here"),
					ContentType: aws.String("string value goes here"),
					ImageResponseCard: &ImageResponseCard{
						Buttons: []*Button{
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
						},
						ImageUrl: aws.String("string value goes here"),
						Subtitle: aws.String("string value goes here"),
						Title:    aws.String("string value goes here"),
					},
				},
				{
					Content:     aws.String("string value goes here"),
					ContentType: aws.String("string value goes here"),
					ImageResponseCard: &ImageResponseCard{
						Buttons: []*Button{
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
						},
						ImageUrl: aws.String("string value goes here"),
						Subtitle: aws.String("string value goes here"),
						Title:    aws.String("string value goes here"),
					},
				},
				{
					Content:     aws.String("string value goes here"),
					ContentType: aws.String("string value goes here"),
					ImageResponseCard: &ImageResponseCard{
						Buttons: []*Button{
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
							{
								Text:  aws.String("string value goes here"),
								Value: aws.String("string value goes here"),
							},
						},
						ImageUrl: aws.String("string value goes here"),
						Subtitle: aws.String("string value goes here"),
						Title:    aws.String("string value goes here"),
					},
				},
			},
		},
		&DTMFInputEvent{
			ClientTimestampMillis: aws.Int64(1234),
			EventId:               aws.String("string value goes here"),
			InputCharacter:        aws.String("string value goes here"),
		},
		&DisconnectionEvent{
			ClientTimestampMillis: aws.Int64(1234),
			EventId:               aws.String("string value goes here"),
		},
		&PlaybackCompletionEvent{
			ClientTimestampMillis: aws.Int64(1234),
			EventId:               aws.String("string value goes here"),
		},
		&TextInputEvent{
			ClientTimestampMillis: aws.Int64(1234),
			EventId:               aws.String("string value goes here"),
			Text:                  aws.String("string value goes here"),
		},
	}

	var marshalers request.HandlerList
	marshalers.PushBackNamed(restjson.BuildHandler)
	payloadMarshaler := protocol.HandlerPayloadMarshal{
		Marshalers: marshalers,
	}
	_ = payloadMarshaler

	eventMsgs := []eventstream.Message{
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("AudioInputEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, inputEvents[0]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("ConfigurationEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, inputEvents[1]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("DTMFInputEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, inputEvents[2]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("DisconnectionEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, inputEvents[3]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("PlaybackCompletionEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, inputEvents[4]),
		},
		{
			Headers: eventstream.Headers{
				eventstreamtest.EventMessageTypeHeader,
				{
					Name:  eventstreamapi.EventTypeHeader,
					Value: eventstream.StringValue("TextInputEvent"),
				},
			},
			Payload: eventstreamtest.MarshalEventPayload(payloadMarshaler, inputEvents[5]),
		},
	}

	return inputEvents, eventMsgs
}
