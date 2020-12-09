package template_method

import (
	"testing"
)

func TestHTTPDownloader(t *testing.T) {
	var downloader Downloader = NewHTTPDownloader()

	downloader.Download("http://example.com/abc.zip")
}

func TestFTPDownloader(t *testing.T) {
	var downloader Downloader = NewFTPDownloader()

	downloader.Download("ftp://example.com/abc.zip")
}