// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 00:25 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"bytes"
	rss_reader "github.com/MichaelLeachim/emerchantpay_rss_reader"
	"html/template"
	"log"
)

type printer interface {
	Print([]rss_reader.RssItem) error
}

type consolePrinter struct {
}

func (c consolePrinter) Print(items []rss_reader.RssItem) error {
	for _, item := range items {
		log.Println("[", item.PubishDate.String(), "]", "============================================")
		log.Println(item.Title)
		log.Println(item.Description)
		log.Println(item.Link)
		log.Println("============================================")
	}
	return nil
}

func newConsolePrinter() printer {
	return consolePrinter{}
}

type htmlPrinter struct {
	storage storager
	tmpname string
}

func newHtmlPrinter(storage storager, tmpname string) printer {
	return htmlPrinter{storage: storage, tmpname: tmpname}
}

func (c htmlPrinter) Print(items []rss_reader.RssItem) error {

	tpl, err := template.New("webpage").Parse(`<!DOCTYPE html>
    <html>
    	<head>
    		<meta charset="UTF-8">
    		<title>{{.Title}}</title>
    	</head>
    	<body>
        <table>
          <thead>
            <tr><th>Title</th><th>Description</th><th>Link</th></tr>
          </thead>
          <tbody>
            {{range .Items}}
              <tr>
                <td>{{ .Title }}</td>
                <td>{{ .Description }}</td>
                <td>{{ .Link    }}</td>
              </tr>
            {{else}}
              <tr><td><div><strong>no feeds</strong></td></tr>
            {{end}}
          </tbody>
        </table>
    	</body>
    </html>`)

	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)

	err = tpl.Execute(buf, struct {
		Title string
		Items []rss_reader.RssItem
	}{
		Title: "Emerchant pay printout",
		Items: items,
	})
	if err != nil {
		return err
	}
	return c.storage.Put(c.tmpname, buf.String())

}
