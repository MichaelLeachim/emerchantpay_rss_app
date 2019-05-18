// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 00:20 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"encoding/json"
)

func storeOnDisk(da dataAccessor, st storage, fpath string, links []string) error {
	b, err := json.Marshal(da.Parse(links))
	if err != nil {
		return err
	}
	return st.Put(fpath, string(b))
}
