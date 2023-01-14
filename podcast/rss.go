package podcast

import (
	"io"

	library "github.com/themooer1/audiobook-library"
	"github.com/themooer1/podgen/podcast/rss"
)

func Create(book *library.AudioBook, basePath string, out io.Writer) error {
	args := rss.RssTemplateArgs{
		Book:     book,
		BasePath: basePath,
	}

	rssTemplate := rss.RssTemplate()

	return rssTemplate.Execute(out, args)
}
