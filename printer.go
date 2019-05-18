// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 00:25 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	rss_reader "github.com/MichaelLeachim/emerchantpay_rss_reader"
)

type printer interface {
	Print([]rss_reader.RssItem) error
}
