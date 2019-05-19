// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 21:59 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"flag"
)

type paramGetter interface {
	String(name string, value string, usage string) *string
	Bool(name string, value bool, usage string) *bool
	List(name string, value []string, usage string) []string
	Parse()
}

type listOfUrls []string

func (i *listOfUrls) String() string {
	return "my string representation"
}

func (i *listOfUrls) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type flagParamGetter struct {
}

func newFlagParamGetter() paramGetter {
	return flagParamGetter{}
}

func (f flagParamGetter) String(name string, value string, usage string) *string {
	return flag.String(name, value, usage)
}

func (f flagParamGetter) Bool(name string, value bool, usage string) *bool {
	return flag.Bool(name, value, usage)
}

func (f flagParamGetter) List(name string, value []string, usage string) []string {

	urlset := listOfUrls(value)
	flag.Var(&urlset, "urlset", "Urls, that contain the data to parse")
	return urlset
}

func (f flagParamGetter) Parse() {
	flag.Parse()

}
