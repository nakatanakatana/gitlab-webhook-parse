package gitlab_webhook_parse

import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"testing"
	"fmt"
)

func TestIssueEvent(t *testing.T){
	filename, _ := filepath.Abs("sampledata/issue_event.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d IssueEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}

func TestMergeRequestEvent(t *testing.T){
	filename, _ := filepath.Abs("sampledata/merge_request.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d MergeRequestEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}

func TestPipelineEvent(t *testing.T){
	filename, _ := filepath.Abs("sampledata/pipeline_event.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d PipelineEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}

func TestPushEvent(t *testing.T){
	filename, _ := filepath.Abs("sampledata/push_event.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d PushEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}

func TestPTagEvent(t *testing.T){
	filename, _ := filepath.Abs("sampledata/tag_event.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d TagEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}

func TestWikiPageEvent(t *testing.T){
	filename, _ := filepath.Abs("sampledata/wiki_page_event.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d WikiPageEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}

func TestCommentEventSnippet(t *testing.T){
	filename, _ := filepath.Abs("sampledata/comment_event_code_snippet.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d CommentEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}

func TestCommentEventCommit(t *testing.T){
	filename, _ := filepath.Abs("sampledata/comment_event_commit.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d CommentEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}
func TestCommentEventIssue(t *testing.T){
	filename, _ := filepath.Abs("sampledata/comment_event_issue.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d CommentEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}
func TestCommentEventMergeRequest(t *testing.T){
	filename, _ := filepath.Abs("sampledata/comment_event_merge_request.json")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("file not found, %s", filename)
	}
	var d CommentEvent
	err = json.Unmarshal(data, &d)
	if err != nil {
		t.Errorf("cannot Unmarshal, %v", err)
	}
	fmt.Println(d)
}