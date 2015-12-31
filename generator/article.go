package generator

import (
	//"bytes"
	"blogo/reader/md"
	"blogo/util"
	//"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

type ArticleI interface {
	Parse(input string) (err error)
	Meta() map[string]string
	Html() string
}
type ArticleM struct {
	ID          string
	Title       string
	Time        time.Time
	ArticleUrl  string
	PostContent template.HTML
}

type articlesM []ArticleM

var (
	articles articlesM
)

func (list articlesM) Len() int {
	return len(list)
}
func (list articlesM) Less(i, j int) bool {
	return list[i].Time.Before(list[j].Time)
}
func (list articlesM) Swap(i, j int) {
	var temp ArticleM = list[j]
	list[j] = list[i]
	list[i] = temp
}

func loadArticle(filePath string, fileName string) {
	article := new(md.MdArticle)
	input, err := ioutil.ReadFile(filePath + fileName)
	util.CheckError(err)
	err = article.Parse(string(input))
	util.CheckError(err)

	meta := article.Meta()
	createTime, err := time.Parse("2006-1-2 15:04", meta["time"])
	util.CheckError(err)

	data := ArticleM{
		ID:          strings.TrimSuffix(fileName, ".md"),
		Title:       meta["title"],
		PostContent: template.HTML(article.Html()),
		Time:        createTime,
		ArticleUrl:  "/" + util.Config.ArticlePath + strings.TrimSuffix(fileName, ".md") + ".html",
	}
	articles = append(articles, data)

}

func loadArticles() {
	workDir, err := os.Getwd()
	util.CheckError(err)

	files, err := ioutil.ReadDir(workDir + "/articles")
	util.CheckError(err)

	for _, file := range files {
		if file.IsDir() == true {
			continue
		}
		loadArticle(workDir+"/articles/", file.Name())

	}
	sort.Sort(articles)

}

func generateArticles() {
	workDir, err := os.Getwd()
	util.CheckError(err)
	os.MkdirAll(workDir+"/build/"+util.Config.ArticlePath, os.ModePerm)
	//util.CheckError(err)

	t, err := template.ParseFiles(workDir + "/templates/article.html")
	for _, article := range articles {

		articleFile, err := os.Create(workDir + "/build/" + util.Config.ArticlePath + article.ID + ".html")
		util.CheckError(err)
		defer articleFile.Close()
		err = t.Execute(articleFile, article)
		util.CheckError(err)

	}

}
