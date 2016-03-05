package md

import (
	//"bytes"
	"blogo/reader/meta"
	"github.com/russross/blackfriday"
	//"io/ioutil"
	//"os"
)

type MdArticle struct {
	meta        map[string]string
	ContentHtml string
	ContentMd   string
}

func md2html(input string) string {
	renderFlag := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	extensions := 0 |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

	htmlRenderer := blackfriday.HtmlRenderer(renderFlag, "", "")
	output := blackfriday.Markdown([]byte(input), htmlRenderer, extensions)
	return string(output)
}

func (article *MdArticle) Parse(input string) error {
	article.meta = make(map[string]string)
	article.ContentMd = input
	metaStr, md := meta.Split(article.ContentMd)
	err := meta.Prase(metaStr, article.meta)
	if err != nil {
		return err
	}
	article.ContentHtml = md2html(string(md))
	//ioutil.WriteFile("out.html", []byte(output), os.ModePerm)
	return nil

}

func (article *MdArticle) Meta() map[string]string {
	return article.meta
}

func (article *MdArticle) Html() string {
	return article.ContentHtml
}
