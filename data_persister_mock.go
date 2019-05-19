// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 22:48 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"errors"
)

type dataOnDiskMockPersister struct {
	fileList map[string]string
}

func newDataOnDiskMockPersister() dataPersister {
	return dataOnDiskMockPersister{fileList: map[string]string{}}
}

func (d dataOnDiskMockPersister) Get(fpath string) (string, error) {
	res, ok := d.fileList[fpath]
	if !ok {
		return "", errors.New("No such file error")
	}
	return res, nil
}

func (d dataOnDiskMockPersister) Put(fpath string, data string) error {
	d.fileList[fpath] = data
	return nil
}
