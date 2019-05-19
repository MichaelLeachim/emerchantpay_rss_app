// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 22:33 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	rss_reader "github.com/MichaelLeachim/emerchantpay_rss_reader"
	"time"
)

type mockDataAccessor struct {
}

func newMockDataAccessor() mockDataAccessor {
	return mockDataAccessor{}
}

func (m mockDataAccessor) Parse(links []string) []rss_reader.RssItem {
	mockTimeNow := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	return []rss_reader.RssItem{
		rss_reader.RssItem{
			Title:       "Google suspends some business with Huawei",
			Source:      "",
			SourceURL:   "",
			Link:        "https://www.reuters.com/article/us-huawei-tech-alphabet-exclusive/exclusive-google-suspends-some-business-with-huawei-after-trump-blacklist-source-idUSKCN1SP0NB",
			Description: "<a href=\"https://news.ycombinator.com/item?id=19954673\">Comments</a>",
			PubishDate:  mockTimeNow,
		},
		rss_reader.RssItem{
			Title:       "Using an Atari ST as Unix/Linux Terminal",
			Source:      "",
			SourceURL:   "",
			Link:        "http://www.atari-wiki.com/index.php/Using_an_Atari_ST_as_Unix/Linux_Terminal",
			Description: "<a href=\"https://news.ycombinator.com/item?id=19954224\">Comments</a>",
			PubishDate:  mockTimeNow,
		},
		rss_reader.RssItem{
			Title:       "How modern SAT solvers work, heuristics and other tricks",
			Source:      "",
			SourceURL:   "",
			Link:        "https://codingnest.com/modern-sat-solvers-fast-neat-and-underused-part-3-of-n/",
			Description: "<a href=\"https://news.ycombinator.com/item?id=19953213\">Comments</a>",
			PubishDate:  mockTimeNow,
		},
		rss_reader.RssItem{
			Title:       "AMD Ryzen Mini-STX: ASRock’s DeskMini A300 – Finally",
			Source:      "",
			SourceURL:   "",
			Link:        "https://smallformfactor.net/reviews/systems/asrocks-deskmini-a300-finally",
			Description: "<a href=\"https://news.ycombinator.com/item?id=19954494\">Comments</a>",
			PubishDate:  mockTimeNow,
		},
		rss_reader.RssItem{
			Title:       "CodeGen: LLVM wrapper for just-in-time code generation and compilation",
			Source:      "",
			SourceURL:   "",
			Link:        "https://github.com/pdziepak/codegen",
			Description: "<a href=\"https://news.ycombinator.com/item?id=19954741\">Comments</a>",
			PubishDate:  mockTimeNow,
		},
		rss_reader.RssItem{
			Title:       "A Decade of Remote Work",
			Source:      "",
			SourceURL:   "",
			Link:        "https://blog.viktorpetersson.com/2019/05/18/a-decade-of-remote.html",
			Description: "<a href=\"https://news.ycombinator.com/item?id=19953854\">Comments</a>",
			PubishDate:  mockTimeNow,
		},
	}
}
