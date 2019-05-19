// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-18 23:56 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"flag"
	rss_reader "github.com/MichaelLeachim/emerchantpay_rss_reader"
	
)
type listOfUrls []string

func (i *listOfUrls) String() string {
	return "my string representation"
}

func (i *listOfUrls) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func app(da dataAccessor, dp dataPersister, log logger, print printer) {
	urlset := listOfUrls{}
	flag.Var(&urlset, "urlset", "Urls, that contain the data to parse")
	outfile := flag.String("out", "tempdata/output.json", "Path to the output file")
	infile := flag.String("in", "tempdata/input.json", "Path to the input file")
	isDisplay := flag.Bool("print", false, "Shold print the data from the <in> file")
	isSave := flag.Bool("save", false, "Should save the data into the <out> file")
	flag.Parse()

	if isDisplay {
		if infile == nil {
			log.Error("<in> flag should not be empty")
			return
		}
		res, err := dp.Get(*infile)
		if err != nil {
			log.Error(err)
			return
		}
		err = printOnTheScreen(dp, print, *infile)
		if err != nil {
			log.Error(err)
		}
		return
	}
	if isSave {
		if outfile == nil {
			log.Error("<out> flag should not be empty")
			return
		}
		err := storeOnDisk(da, dp, *outfile, urlset)
		if err != nil {
			log.Error(err)
		}
		return
	}
	log.Warn("Use either: <save> or <print> flags")
	return
}

func main() {
	app(newUrlDataAccessor(), newDataOnDiskPersister(),
		newLogrusLogger(),
		log logger, print printer)

}
