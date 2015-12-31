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
	output := blackfriday.MarkdownCommon([]byte(input))
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
