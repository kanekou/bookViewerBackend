package bookservice

import (
	chapter "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Chapter"
	page "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Page"
	question "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Question"
	sentence "github.com/KazuwoKiwame12/bookViewerBackend/DB/Model/Sentence"
)

//BookData ...本のデータ
type BookData struct {
	Book []chapterData
}

type chapterData struct {
	ChapterID int
	Pages     []pageData
}

type pageData struct {
	PageID    int
	Sentences []sentenceData
}

type sentenceData struct {
	SentenceID  int
	Content     string
	HasQuestion bool
}

//GetContentForClient ...本のデータをフロント向けに整形する
func GetContentForClient(bookID int) BookData {
	response := BookData{}
	chapterDataList := []chapterData{}

	for _, chapter := range chapter.GetListByBook(bookID) {
		pageDataList := []pageData{}
		for _, page := range page.GetListByChapter(chapter.ID) {
			sentenceDataList := []sentenceData{}
			for _, sentence := range sentence.GetListByPage(page.ID) {
				data := sentenceData{SentenceID: sentence.ID, Content: sentence.Content, HasQuestion: checkHasQuestion(sentence.ID)}
				sentenceDataList = append(sentenceDataList, data)
			}
			pageDataList = append(pageDataList, pageData{PageID: page.ID, Sentences: sentenceDataList})
		}
		chapterDataList = append(chapterDataList, chapterData{ChapterID: chapter.ID, Pages: pageDataList})
	}

	response.Book = chapterDataList
	return response
}

func checkHasQuestion(sentenceID int) bool {
	questionList := question.GetListBySentence(sentenceID)
	if len(questionList) == 0 {
		return false
	}
	return true
}
