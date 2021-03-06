/*
* MIT License

* Copyright (c) 2021 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/


package api

import (
	"time"
)

// Error class.             
type ApiError struct {

	// Gets or sets api error code.
	Code string `json:"Code,omitempty"`

	// Gets or sets error message.
	Message string `json:"Message,omitempty"`

	// Gets or sets error description.
	Description string `json:"Description,omitempty"`

	// Gets or sets server datetime.
	DateTime time.Time `json:"DateTime,omitempty"`

	// Gets or sets inner error.
	InnerError *ApiError `json:"InnerError,omitempty"`
}
type IApiError interface {
	IsApiError() bool
}
func (ApiError) IsApiError() bool {
	return true;
}

