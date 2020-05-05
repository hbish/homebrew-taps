/*
Copyright © 2020 Ben Shi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package xml

import (
	"bytes"
	"encoding/xml"
	"regexp"
)

// Sitemap XML Protocol Implementation
// - info: https://www.sitemaps.org/protocol.html

// URLSet
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNs   string   `xml:"xmlns,attr"`
	URL     []URL    `xml:"url"`
}

// URL is for every single location url
type URL struct {
	Loc        string  `xml:"loc" json:"loc"`
	LastMod    string  `xml:"lastmod,omitempty" json:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty" json:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty" json:"priority,omitempty"`
	Image      []Image `xml:"image,omitempty" json:"image,omitempty"`
}

type Image struct {
	Loc         string `xml:"loc" json:"loc"`
	Title       string `xml:"title,omitempty" json:"title,omitempty"`
	Caption     string `xml:"caption,omitempty" json:"caption,omitempty"`
	GeoLocation string `xml:"geo_location,omitempty" json:"geo_location,omitempty"`
	License     string `xml:"license,omitempty" json:"license,omitempty"`
}

func unmarshalXML(rawXml []byte) (*URLSet, error) {
	urlSet := URLSet{}

	// validate xml without storing
	if err := xml.Unmarshal(rawXml, new(interface{})); err != nil {
		return nil, err
	}

	// decode xml and trim white spaces
	reader := bytes.NewReader(rawXml)
	d := xml.NewDecoder(reader)
	td := xml.NewTokenDecoder(TrimmingTokenReader{d})
	err := td.Decode(&urlSet)

	if err != nil {
		return nil, err
	}

	return &urlSet, nil
}

// unmarshal xml and filter by pattern
func UnmarshalXMLP(rawXml []byte, pattern string) (*URLSet, error) {
	urlSet, err := unmarshalXML(rawXml)
	if err != nil {
		return nil, err
	}

	if pattern != "" {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}

		var filteredUrls []URL
		for _, url := range urlSet.URL {
			matched := regex.MatchString(url.Loc)
			if matched {
				filteredUrls = append(filteredUrls, url)
			}
		}

		urlSet.URL = filteredUrls
	}

	return urlSet, nil
}

// Trimming TokenReader
type TrimmingTokenReader struct {
	dec *xml.Decoder
}

// Trimming token
func (tr TrimmingTokenReader) Token() (xml.Token, error) {
	t, err := tr.dec.Token()
	if cd, ok := t.(xml.CharData); ok {
		t = xml.CharData(bytes.TrimSpace(cd))
	}
	return t, err
}
