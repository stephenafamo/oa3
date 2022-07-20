// Code generated by oa3 (https://github.com/aarondl/oa3). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package oa3gen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aarondl/chrono"
	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
)

// Authenticate post /auth
func (_c Client) Authenticate(_ctx context.Context) (AuthenticateResponse, *http.Response, error) {
	_urlStr := `/auth`
	_req, _err := http.NewRequestWithContext(_ctx, http.MethodPost, _urlStr, nil)
	if _err != nil {
		return nil, nil, _err
	}
	var _query url.Values
	if len(_query) > 0 {
		_req.URL.RawQuery = _query.Encode()
	}

	_httpResp, _err := _c.doRequest(_ctx, _req)
	if _err != nil {
		return nil, nil, _err
	}

	var _resp AuthenticateResponse
	switch _httpResp.Status {
	case `200`:
		_resp = HTTPStatusOk{}
	default:
		return nil, nil, fmt.Errorf("unknown response code")
	}

	return _resp, _httpResp, nil
}

// TestInlinePrimitiveBody get /test/inline
func (_c Client) TestInlinePrimitiveBody(_ctx context.Context, body string) (TestInlinePrimitiveBodyResponse, *http.Response, error) {
	_urlStr := `/test/inline`
	_req, _err := http.NewRequestWithContext(_ctx, http.MethodGet, _urlStr, nil)
	if _err != nil {
		return nil, nil, _err
	}
	_bodyBytes, _err := json.Marshal(body)
	if _err != nil {
		return nil, nil, _err
	}
	_req.Body = io.NopCloser(bytes.NewReader(_bodyBytes))
	var _query url.Values
	if len(_query) > 0 {
		_req.URL.RawQuery = _query.Encode()
	}

	_httpResp, _err := _c.doRequest(_ctx, _req)
	if _err != nil {
		return nil, nil, _err
	}

	var _resp TestInlinePrimitiveBodyResponse
	switch _httpResp.Status {
	case `200`:
		_resp = HTTPStatusOk{}
	default:
		return nil, nil, fmt.Errorf("unknown response code")
	}

	return _resp, _httpResp, nil
}

// TestInline post /test/inline
func (_c Client) TestInline(_ctx context.Context, body TestInlineInline) (TestInlineResponse, *http.Response, error) {
	_urlStr := `/test/inline`
	_req, _err := http.NewRequestWithContext(_ctx, http.MethodPost, _urlStr, nil)
	if _err != nil {
		return nil, nil, _err
	}
	_bodyBytes, _err := json.Marshal(body)
	if _err != nil {
		return nil, nil, _err
	}
	_req.Body = io.NopCloser(bytes.NewReader(_bodyBytes))
	var _query url.Values
	if len(_query) > 0 {
		_req.URL.RawQuery = _query.Encode()
	}

	_httpResp, _err := _c.doRequest(_ctx, _req)
	if _err != nil {
		return nil, nil, _err
	}

	var _resp TestInlineResponse
	switch _httpResp.Status {
	case `200`:
		var _respObject TestInline200Inline
		_b, _err := io.ReadAll(_httpResp.Body)
		if _err != nil {
			return nil, nil, _err
		}
		if _err = json.Unmarshal(_b, &_respObject); _err != nil {
			return nil, nil, _err
		}
		_resp = _respObject
	case `201`:
		var _respObject TestInline201Inline
		_b, _err := io.ReadAll(_httpResp.Body)
		if _err != nil {
			return nil, nil, _err
		}
		if _err = json.Unmarshal(_b, &_respObject); _err != nil {
			return nil, nil, _err
		}
		_resp = _respObject
	default:
		return nil, nil, fmt.Errorf("unknown response code")
	}

	return _resp, _httpResp, nil
}

// GetUser get /users/{id}
//
// Retrieves a user with a long description that spans multiple lines so
// that we can see that both wrapping and long-line support is not
// bleeding over the sacred 80 char limit.
func (_c Client) GetUser(_ctx context.Context, id string, validStr omitnull.Val[string], reqValidStr null.Val[string], validInt omit.Val[int], reqValidInt int, validNum omit.Val[float64], reqValidNum float64, validBool omit.Val[bool], reqValidBool bool, reqStrFormat string, dateTime chrono.DateTime, date chrono.Date, timeVal chrono.Time, durationVal time.Duration) (GetUserResponse, *http.Response, error) {
	_urlStr := `/users/{id}`
	_urlStr = strings.Replace(_urlStr, `{id}`, fmt.Sprintf("%v", id), 1)
	_req, _err := http.NewRequestWithContext(_ctx, http.MethodGet, _urlStr, nil)
	if _err != nil {
		return nil, nil, _err
	}
	if val, ok := validStr.Get(); ok {
		_req.Header.Set(`valid_str`, fmt.Sprintf("%v", val))
	}
	var _query url.Values
	if _query == nil {
		_query = make(url.Values)
	}
	if val, ok := reqValidStr.Get(); ok {
		_query.Set(`req_valid_str`, fmt.Sprintf("%v", val))
	}
	if val, ok := validInt.Get(); ok {
		_query.Set(`valid_int`, fmt.Sprintf("%v", val))
	}
	_query.Set(`req_valid_int`, fmt.Sprintf("%v", reqValidInt))
	if val, ok := validNum.Get(); ok {
		_query.Set(`valid_num`, fmt.Sprintf("%v", val))
	}
	_query.Set(`req_valid_num`, fmt.Sprintf("%v", reqValidNum))
	if val, ok := validBool.Get(); ok {
		_query.Set(`valid_bool`, fmt.Sprintf("%v", val))
	}
	_query.Set(`req_valid_bool`, fmt.Sprintf("%v", reqValidBool))
	_query.Set(`req_str_format`, fmt.Sprintf("%v", reqStrFormat))
	_query.Set(`date_time`, fmt.Sprintf("%v", dateTime))
	_query.Set(`date`, fmt.Sprintf("%v", date))
	_query.Set(`time_val`, fmt.Sprintf("%v", timeVal))
	_query.Set(`duration_val`, fmt.Sprintf("%v", durationVal))
	if len(_query) > 0 {
		_req.URL.RawQuery = _query.Encode()
	}

	_httpResp, _err := _c.doRequest(_ctx, _req)
	if _err != nil {
		return nil, nil, _err
	}

	var _resp GetUserResponse
	switch _httpResp.Status {
	case `304`:
		_resp = HTTPStatusNotModified{}
	default:
		return nil, nil, fmt.Errorf("unknown response code")
	}

	return _resp, _httpResp, nil
}

// SetUser post /users/{id}
//
// Sets a user
func (_c Client) SetUser(_ctx context.Context, body *Primitives) (SetUserResponse, *http.Response, error) {
	_urlStr := `/users/{id}`
	_req, _err := http.NewRequestWithContext(_ctx, http.MethodPost, _urlStr, nil)
	if _err != nil {
		return nil, nil, _err
	}
	_bodyBytes, _err := json.Marshal(body)
	if _err != nil {
		return nil, nil, _err
	}
	_req.Body = io.NopCloser(bytes.NewReader(_bodyBytes))
	var _query url.Values
	if len(_query) > 0 {
		_req.URL.RawQuery = _query.Encode()
	}

	_httpResp, _err := _c.doRequest(_ctx, _req)
	if _err != nil {
		return nil, nil, _err
	}

	var _resp SetUserResponse
	switch _httpResp.Status {
	case `200`:
		var _respObject SetUser200WrappedResponse
		_b, _err := io.ReadAll(_httpResp.Body)
		if _err != nil {
			return nil, nil, _err
		}
		if _err = json.Unmarshal(_b, &_respObject.Body); _err != nil {
			return nil, nil, _err
		}
		if hdr := _httpResp.Header.Get(`X-Response-Header`); len(hdr) != 0 {
			_respObject.HeaderXResponseHeader.Set(hdr)
		}
		_resp = _respObject
	default:
		var _respObject SetUserdefaultWrappedResponse
		_b, _err := io.ReadAll(_httpResp.Body)
		if _err != nil {
			return nil, nil, _err
		}
		if _err = json.Unmarshal(_b, &_respObject.Body); _err != nil {
			return nil, nil, _err
		}
		_resp = _respObject
	}

	return _resp, _httpResp, nil
}