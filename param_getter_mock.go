// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 23:23 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

type flagParamGetterMock struct {
}

func newFlagParamGetterMock() paramGetter {
	return flagParamGetterMock{}
}

func (f flagParamGetterMock) String(name string, value string, usage string) *string {
	return &value
}

func (f flagParamGetterMock) Bool(name string, value bool, usage string) *bool {
	return &value
}

func (f flagParamGetterMock) List(name string, value []string, usage string) []string {
	return value
}

func (f flagParamGetterMock) Parse() {
	return
}
