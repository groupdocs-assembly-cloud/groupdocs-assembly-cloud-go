/*
* MIT License

* Copyright (c) 2020 Aspose Pty Ltd

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

// Assemble options data which is using for specifying assemble options, like template name, save format, report data and etc.
type AssembleOptionsData struct {

	// Gets or sets the template name which is located on storage.
	TemplateFileInfo *TemplateFileInfo `json:"TemplateFileInfo,omitempty"`

	// Gets or sets a save format for assembled document.
	SaveFormat string `json:"SaveFormat,omitempty"`

	// Gets or sets a data for report.
	ReportData string `json:"ReportData,omitempty"`

	// Gets or sets result path of a built document.
	OutputPath string `json:"OutputPath,omitempty"`
}
type IAssembleOptionsData interface {
	IsAssembleOptionsData() bool
}
func (AssembleOptionsData) IsAssembleOptionsData() bool {
	return true;
}

