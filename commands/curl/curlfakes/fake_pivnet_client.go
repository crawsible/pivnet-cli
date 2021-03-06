// This file was generated by counterfeiter
package curlfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands/curl"
)

type FakePivnetClient struct {
	MakeRequestStub        func(method string, url string, expectedResponseCode int, body io.Reader, data interface{}) (*http.Response, []byte, error)
	makeRequestMutex       sync.RWMutex
	makeRequestArgsForCall []struct {
		method               string
		url                  string
		expectedResponseCode int
		body                 io.Reader
		data                 interface{}
	}
	makeRequestReturns struct {
		result1 *http.Response
		result2 []byte
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) MakeRequest(method string, url string, expectedResponseCode int, body io.Reader, data interface{}) (*http.Response, []byte, error) {
	fake.makeRequestMutex.Lock()
	fake.makeRequestArgsForCall = append(fake.makeRequestArgsForCall, struct {
		method               string
		url                  string
		expectedResponseCode int
		body                 io.Reader
		data                 interface{}
	}{method, url, expectedResponseCode, body, data})
	fake.recordInvocation("MakeRequest", []interface{}{method, url, expectedResponseCode, body, data})
	fake.makeRequestMutex.Unlock()
	if fake.MakeRequestStub != nil {
		return fake.MakeRequestStub(method, url, expectedResponseCode, body, data)
	} else {
		return fake.makeRequestReturns.result1, fake.makeRequestReturns.result2, fake.makeRequestReturns.result3
	}
}

func (fake *FakePivnetClient) MakeRequestCallCount() int {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return len(fake.makeRequestArgsForCall)
}

func (fake *FakePivnetClient) MakeRequestArgsForCall(i int) (string, string, int, io.Reader, interface{}) {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return fake.makeRequestArgsForCall[i].method, fake.makeRequestArgsForCall[i].url, fake.makeRequestArgsForCall[i].expectedResponseCode, fake.makeRequestArgsForCall[i].body, fake.makeRequestArgsForCall[i].data
}

func (fake *FakePivnetClient) MakeRequestReturns(result1 *http.Response, result2 []byte, result3 error) {
	fake.MakeRequestStub = nil
	fake.makeRequestReturns = struct {
		result1 *http.Response
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ curl.PivnetClient = new(FakePivnetClient)
