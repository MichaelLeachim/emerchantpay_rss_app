// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 00:21 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"io/ioutil"
)

type dataPersister interface {
	Put(string, string) error
	Get(string) (string, error)
}

type dataOnDiskPersister struct {
}

func newDataOnDiskPersister() dataPersister {
	return dataOnDiskPersister{}
}

func (d dataOnDiskPersister) Get(fpath string) (string, error) {
	res, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", err
	}
	return string(res), err
}

func (d dataOnDiskPersister) Put(fpath string, data string) error {
	return ioutil.WriteFile(fpath, []byte(data), 0644)
}
