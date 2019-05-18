// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 00:20 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"encoding/json"
	rss_reader "github.com/MichaelLeachim/emerchantpay_rss_reader"
)

func printOnTheScreen(st storage, p printer, fpath string) error {
	res, err := st.Get(fpath)
	if err != nil {
		return err
	}
	results := []rss_reader.RssItem{}

	err = json.Unmarshal([]byte(res), &results)
	if err != nil {
		return err
	}
	return p.Print(results)
}
