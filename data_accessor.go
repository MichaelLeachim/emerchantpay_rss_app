// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 00:06 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	rss_reader "github.com/MichaelLeachim/emerchantpay_rss_reader"
)

type dataAccessor interface {
	Parse(links []string) []rss_reader.RssItem
}

type mockDataAccessor struct {
}

func (m *mockDataAccessor) Parse(links []string) []rss_reader.RssItem {
	return []rss_reader.RssItem{}
}
