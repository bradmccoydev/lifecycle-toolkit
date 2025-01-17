// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"sync"
)

// EventObjectMock is a mock implementation of interfaces.EventObject.
//
// 	func TestSomethingThatUsesEventObject(t *testing.T) {
//
// 		// make and configure a mocked interfaces.EventObject
// 		mockedEventObject := &EventObjectMock{
// 			GetEventAnnotationsFunc: func() map[string]string {
// 				panic("mock out the GetEventAnnotations method")
// 			},
// 		}
//
// 		// use mockedEventObject in code that requires interfaces.EventObject
// 		// and then make assertions.
//
// 	}
type EventObjectMock struct {
	// GetEventAnnotationsFunc mocks the GetEventAnnotations method.
	GetEventAnnotationsFunc func() map[string]string

	// calls tracks calls to the methods.
	calls struct {
		// GetEventAnnotations holds details about calls to the GetEventAnnotations method.
		GetEventAnnotations []struct {
		}
	}
	lockGetEventAnnotations sync.RWMutex
}

// GetEventAnnotations calls GetEventAnnotationsFunc.
func (mock *EventObjectMock) GetEventAnnotations() map[string]string {
	if mock.GetEventAnnotationsFunc == nil {
		panic("EventObjectMock.GetEventAnnotationsFunc: method is nil but EventObject.GetEventAnnotations was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetEventAnnotations.Lock()
	mock.calls.GetEventAnnotations = append(mock.calls.GetEventAnnotations, callInfo)
	mock.lockGetEventAnnotations.Unlock()
	return mock.GetEventAnnotationsFunc()
}

// GetEventAnnotationsCalls gets all the calls that were made to GetEventAnnotations.
// Check the length with:
//     len(mockedEventObject.GetEventAnnotationsCalls())
func (mock *EventObjectMock) GetEventAnnotationsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetEventAnnotations.RLock()
	calls = mock.calls.GetEventAnnotations
	mock.lockGetEventAnnotations.RUnlock()
	return calls
}
