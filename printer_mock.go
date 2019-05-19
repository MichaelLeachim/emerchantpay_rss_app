// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 22:53 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	rss_reader "github.com/MichaelLeachim/emerchantpay_rss_reader"
)

type mockPrinter struct {
	data []rss_reader.RssItem
}

func newMockPrinter() *mockPrinter {
	return &mockPrinter{data: []rss_reader.RssItem{}}
}

func (m *mockPrinter) ReadPrinted() []rss_reader.RssItem {
	return m.data
}

func (m *mockPrinter) Print(items []rss_reader.RssItem) error {
	for _, item := range items {
		m.data = append(m.data, item)
	}
	return nil
}
