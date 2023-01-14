package rss

import (
	"html/template"
	"path/filepath"

	library "github.com/themooer1/audiobook-library"
)

type RssTemplateArgs struct {
	Book     *library.AudioBook
	BasePath string
}

var templateString = `
<?xml version="1.0" ?><rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:media="http://search.yahoo.com/mrss/" xmlns:creativeCommons="http://backend.userland.com/creativeCommonsRssModule" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0">
<channel>
  <title><![CDATA[{{ .Book.Title }}]]></title>
  <!--<link><![CDATA[.Book.Link]]></link>-->
  <!--<atom:link rel="self" href="https://librivox.org/rss/12129" /> -->
  <description><![CDATA[{{ .Book.Description }}]]></description>
  <!--<genre>project element=Genre</genre>-->
  <!--<language>project element=lang.code</language>-->
  <itunes:type>serial</itunes:type>
  <itunes:author><!CDATA[{{ .Book.Author }}]]></itunes:author>
  <itunes:summary><![CDATA[{{ .Book.Description }}]]></itunes:summary>
  <!--
  <itunes:owner>
    <itunes:name>LibriVox</itunes:name>
    <itunes:email>info@librivox.org</itunes:email>
  </itunes:owner>
  -->
  <itunes:category text="Arts">
    <itunes:category text="Literature" />
  </itunes:category>
  <!-- file loop -->
  {{ range $index, $chapter := .Book.Chapters }}
    <item>
    <title><![CDATA[{{ $chapter.Title }}]]></title>
    <itunes:episode><![CDATA[{{ $index }}]]></itunes:episode>
    <!--<reader>file element=reader</reader> -->
    <!--<link><![CDATA[https://librivox.org/crime-and-punishment-version-3-by-fyodor-dostoyevsky/]]></link>-->
            <enclosure url="{{ join $.BasePath $chapter.Url }}" length="0" type="audio/mpeg" />
	<itunes:explicit>No</itunes:explicit>
	<itunes:block>No</itunes:block>
	<itunes:duration><![CDATA[]]></itunes:duration>
	<!--<pubDate>file element=rss.pubDate</pubDate>-->
	<media:content url="{{ join $.BasePath $chapter.Url }}" type="audio/mpeg" />
	</item>
 {{ end }}
    <!-- end file loop -->
</channel>
</rss>
`

func RssTemplate() *template.Template {
	rssTemplate := template.New("rss-feed")

	rssTemplate = rssTemplate.Funcs(
		template.FuncMap{
			"join": filepath.Join,
		},
	)

	return template.Must(rssTemplate.Parse(templateString))
}
