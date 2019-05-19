// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-18 23:56 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

func app(da dataAccessor, dp dataPersister, pg paramGetter, log logger, print printer) {

	urlset := pg.List("urlset", []string{}, "Urls, that contain the data to parse")
	outfile := pg.String("out", "tempdata/output.json", "Path to the output file")
	infile := pg.String("in", "tempdata/input.json", "Path to the input file")
	isDisplay := pg.Bool("print", false, "Shold print the data from the <in> file")
	isSave := pg.Bool("save", false, "Should save the data into the <out> file")
	pg.Parse()

	if isDisplay != nil && *isDisplay {
		if infile == nil {
			log.Error("<in> flag should not be empty")
			return
		}
		err := printOnTheScreen(dp, print, *infile)
		if err != nil {
			log.Error(err)
		}
		return
	}
	if isSave != nil && *isSave {
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
	app(newUrlDataAccessor(),
		newDataOnDiskPersister(),
		newFlagParamGetter(),
		newLogrusLogger(),
		newConsolePrinter())
}
