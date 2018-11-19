package storage

import (
	//"github.com/JorgenJJ/justice4campus/internal/storage"
	"testing"
)


func TestAdd(t *testing.T) {
	testTitle := "testTitle"
	testDescription := "testDescription"

	testIdea := IdeaStruct{Title: testTitle, Description: testDescription}

	resultIdea, err := 	Idea.Add(testIdea)

	if err != nil {
		t.Errorf("Got error %d", err)
	} else if resultIdea.Title != testIdea.Title && resultIdea.Description != testIdea.Description  {
		t.Errorf("Expected Title %v and Description %v, Got Title %v and Description %v",
			testIdea.Title, testIdea.Description, resultIdea.Title, resultIdea.Description)
	}

	/*if testIdea.Title != "testTitle" {
		t.Errorf("Expected %v", testIdea.Title)
	}*/
}