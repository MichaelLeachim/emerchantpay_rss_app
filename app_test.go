// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 22:41 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadingPersistingAndLoading(t *testing.T) {

	// store some data on disk
	storager := newDataOnDiskMockPersister()
	err := storeOnDisk(newMockDataAccessor(), storager, "/home/mik/blab/blip/blop.json", []string{})
	assert.Equal(t, err, nil)

	// load some data from disk
	mockPrinter := newMockPrinter()
	printOnTheScreen(storager, mockPrinter, "/home/mik/blab/blip/blop.json")
	assert.Equal(t, len(mockPrinter.ReadPrinted()), 6)
	assert.Equal(t, mockPrinter.ReadPrinted()[0].Title, "Google suspends some business with Huawei")
}

func TestApp(t *testing.T) {
	assert.Equal(t, 1, 1)
}
