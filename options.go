// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Options holds the optional parameters for Diffbot client.
//
// See http://diffbot.com/products/automatic/
type Options struct {
	Fields                  string
	Timeout                 time.Duration
	Callback                string
	FrontpageAll            string
	Discussion              bool // discussion defaults to false
	ClassifierMode          string
	ClassifierStats         string
	BulkNotifyEmail         string
	BulkNotifyWebHook       string
	BulkRepeat              string
	BulkMaxRounds           string
	BulkPageProcessPattern  string
	CrawlMaxToCrawl         string
	CrawlMaxToProcess       string
	CrawlRestrictDomain     string
	CrawlNotifyEmail        string
	CrawlNotifyWebHook      string
	CrawlDelay              string
	CrawlRepeat             string
	CrawlOnlyProcessIfNew   string
	CrawlMaxRounds          string
	CrawlUrlPattern         string
	CrawlUrlRegexp          string
	CrawlUrlProcessPattern  string
	CrawlUrlProcessRegexp   string
	CrawlPageProcessPattern string
	CrawlMaxHops            string
	CrawlFormat             string
	CrawlType               string
	CrawlNumber             string
	BatchMethod             string
	BatchRelativeUrl        string
	CustomHeader            http.Header
}

// MethodParamString return string as the url params.
//
// If the Options is not empty, the return string begin with a '&'.
func (p *Options) MethodParamString(method string) string {
	if p == nil || method == "" {
		return ""
	}

	switch method {
	case "article", "image", "product":
		var s []byte
		if p.Fields != "" {
			s = append(s, ("&fields=" + p.Fields)...)
		}
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.Callback != "" {
			s = append(s, ("&callback=" + url.QueryEscape(p.Callback))...)
		}
		if !p.Discussion {
			s = append(s, ("&discussion=false")...)
		}
		return string(s)

	case "frontpage":
		var s []byte
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.FrontpageAll != "" {
			s = append(s, ("&all=" + p.FrontpageAll)...)
		}
		return string(s)

	case "analyze":
		var s []byte
		if p.ClassifierMode != "" {
			s = append(s, ("&mode=" + p.ClassifierMode)...)
		}
		if p.Fields != "" {
			s = append(s, ("&fields=" + p.Fields)...)
		}
		if p.ClassifierStats != "" {
			s = append(s, ("&stats=" + p.ClassifierStats)...)
		}
		if !p.Discussion {
			s = append(s, ("&discussion=false")...)
		}
		return string(s)

	case "bulk":
		var s []byte
		if p.BulkNotifyEmail != "" {
			s = append(s, ("&notifyEmail=" + p.BulkNotifyEmail)...)
		}
		if p.BulkNotifyWebHook != "" {
			s = append(s, ("&notifyWebHook=" + p.BulkNotifyWebHook)...)
		}
		if p.BulkRepeat != "" {
			s = append(s, ("&repeat=" + p.BulkRepeat)...)
		}
		if p.BulkMaxRounds != "" {
			s = append(s, ("&maxRounds=" + p.BulkMaxRounds)...)
		}
		if p.BulkPageProcessPattern != "" {
			s = append(s, ("&pageProcessPattern=" + p.BulkPageProcessPattern)...)
		}
		return string(s)

	case "crawl":
		var s []byte
		if p.CrawlMaxToCrawl != "" {
			s = append(s, ("&maxToCrawl=" + p.CrawlMaxToCrawl)...)
		}
		if p.CrawlMaxToProcess != "" {
			s = append(s, ("&maxToProcess=" + p.CrawlMaxToProcess)...)
		}
		if p.CrawlRestrictDomain != "" {
			s = append(s, ("&restrictDomain=" + p.CrawlRestrictDomain)...)
		}
		if p.CrawlNotifyEmail != "" {
			s = append(s, ("&notifyEmail=" + p.CrawlNotifyEmail)...)
		}
		if p.CrawlNotifyWebHook != "" {
			s = append(s, ("&notifyWebHook=" + p.CrawlNotifyWebHook)...)
		}
		if p.CrawlDelay != "" {
			s = append(s, ("&crawlDelay=" + p.CrawlDelay)...)
		}
		if p.CrawlRepeat != "" {
			s = append(s, ("&repeat=" + p.CrawlRepeat)...)
		}
		if p.CrawlOnlyProcessIfNew != "" {
			s = append(s, ("&onlyProcessIfNew=" + p.CrawlOnlyProcessIfNew)...)
		}
		if p.CrawlMaxRounds != "" {
			s = append(s, ("&maxRounds=" + p.CrawlMaxRounds)...)
		}
		if p.CrawlUrlPattern != "" {
			s = append(s, ("&urlCrawlPattern=" + p.CrawlUrlPattern)...)
		}
		if p.CrawlUrlRegexp != "" {
			s = append(s, ("&urlCrawlRegEx=" + p.CrawlUrlRegexp)...)
		}
		if p.CrawlUrlProcessPattern != "" {
			s = append(s, ("&urlProcessPattern=" + p.CrawlUrlProcessPattern)...)
		}
		if p.CrawlUrlProcessRegexp != "" {
			s = append(s, ("&urlProcessRegEx=" + p.CrawlUrlProcessRegexp)...)
		}
		if p.CrawlPageProcessPattern != "" {
			s = append(s, ("&pageProcessPattern=" + p.CrawlPageProcessPattern)...)
		}
		if p.CrawlMaxHops != "" {
			s = append(s, ("&maxHops=" + p.CrawlMaxHops)...)
		}
		return string(s)

	case "crawl/data":
		var s []byte
		if p.CrawlFormat != "" {
			s = append(s, ("&format=" + p.CrawlFormat)...)
		}
		if p.CrawlType != "" {
			s = append(s, ("&type=" + p.CrawlType)...)
		}
		if p.CrawlNumber != "" {
			s = append(s, ("&num=" + p.CrawlNumber)...)
		}
		return string(s)

	case "batch":
		var s []byte
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.BatchMethod != "" {
			s = append(s, ("&method=" + p.BatchMethod)...)
		}
		if p.BatchRelativeUrl != "" {
			s = append(s, ("&relative_url=" + url.QueryEscape(p.BatchRelativeUrl))...)
		}
		return string(s)

	default: // Custom APIs
		var s []byte
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.Callback != "" {
			s = append(s, ("&callback=" + url.QueryEscape(p.Callback))...)
		}
		return string(s)
	}

	return ""
}
