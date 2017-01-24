package gitlab_webhook_parse

import (
	"time"
	"strings"
)

type customTime struct {
	time.Time
}

func (t customTime) UnmarshalJSON(b []byte) (err error) {
	layout := []string{
		"2006-01-02 15:04:05 MST",
		time.RFC3339,
	}
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	for _, l := range (layout) {
		t.Time, err = time.Parse(l, s)
		if err == nil {
			break
		}
	}
	return
}


type IssueEvent struct {
	ObjectKind       string `json:"object_kind"`
	User             User `json:"user"`
	Project          Project `json:"project"`
	Repository       Repository `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Assignee         Assignee `json:"assignee"`
}

type MergeRequestEvent struct {
	ObjectKind       string `json:"object_kind"`
	User             User `json:"user"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

type PushEvent struct {
	ObjectKind        string `json:"object_kind"`
	Before            string `json:"before"`
	After             string `json:"after"`
	Ref               string `json:"ref"`
	CheckoutSha       string `json:"checkout_sha"`
	UserID            int `json:"user_id"`
	UserName          string `json:"user_name"`
	UserEmail         string `json:"user_email"`
	UserAvatar        string `json:"user_avatar"`
	ProjectID         int `json:"project_id"`
	Project           Project `json:"Project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}

type TagEvent struct {
	ObjectKind        string `json:"object_kind"`
	Before            string `json:"before"`
	After             string `json:"after"`
	Ref               string `json:"ref"`
	CheckoutSha       string `json:"checkout_sha"`
	UserID            int `json:"user_id"`
	UserName          string `json:"user_name"`
	UserAvatar        string `json:"user_avatar"`
	ProjectID         int `json:"project_id"`
	Project           Project `json:"Project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}

type WikiPageEvent struct {
	ObjectKind       string `json:"object_kind"`
	User             User `json:"user"`
	Project          Project `json:"project"`
	Wiki             Wiki `json:"wiki"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

type PipelineEvent struct {
	ObjectKind       string `json:"object_kind"`
	User             User `json:"user"`
	Project          Project `json:"project"`
	Commit           Commit `json:"commit"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Builds           []Build `json:"builds"`
}

type CommentEvent struct {
	ObjectKind       string `json:"object_kind"`
	User             User `json:"user"`
	ProjectID        int `json:"project_id"`
	Project          Project `json:"project"`
	Repository       Repository `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Commit           Commit `json:"commit"`
	Issue            Issue `json:"issue"`
	Snippet          Snippet `json:"snippet"`
}

type Issue struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	AssigneeID  int `json:"assignee_id"`
	AuthorID    int `json:"author_id"`
	ProjectID   int `json:"project_id"`
	CreatedAt   customTime `json:"created_at"`
	UpdatedAt   customTime `json:"updated_at"`
	Position    int `json:"position"`
	BranchName  string `json:"branch_name"`
	Description string `json:"description"`
	MilestoneID int `json:"milestone_id"`
	State       string `json:"state"`
	IID         int `json:"iid"`
}

type Build struct {
	ID            int `json:"id"`
	Stage         string `json:"stage"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	CreatedAt     customTime `json:"created_at"`
	StartedAt     customTime `json:"started_at"`
	FinishedAt    customTime `json:"finished_at"`
	When          string `json:"when"`
	Manual        bool `json:"manual"`
	User          User `json:"user"`
	Runner        string `json:"runner"`
	ArtifactsFile Artifactsfile `json:"artifactsfile"`
}

type Artifactsfile struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
}

type Wiki struct {
	WebURL            string "web_url"
	GitSshURL         string "git_ssh_url"
	GitHttpURL        string "git_http_url"
	PathWithNamespace string "path_with_namespace"
	DefaultBranch     string "default_branch"
}

type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp customTime `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
	Added     []string `json:"added"`
	Modified  []string `json:"modified"`
	Removed   []string `json:"removed"`
}

type Snippet struct {
	ID              int `json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	AuthorID        int `json:"author_id"`
	ProjectID       int `json:"project_id"`
	CreatedAt       customTime `json:"created_at"`
	UpdatedAt       customTime `json:"updated_at"`
	FileName        string `json:"file_name"`
	ExpiresAt       customTime `json:"expires_at"`
	Type            string `json:"type"`
	VisibilityLevel int "visibility_level"
}

type User struct {
	Name      string `json:"name"`
	UserName  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type Project struct {
	Name              string `json"name"`
	Description       string "description"
	WebURL            string "web_url"
	AvatarURL         string "avatar_url"
	GitSshURL         string "git_ssh_url"
	GitHttpURL        string "git_http_url"
	Namespace         string "namespace"
	VisibilityLevel   int "visibility_level"
	PathWithNamespace string "path_with_namespace"
	DefaultBranch     string "default_branch"
	Homepage          string "homepage"
	URL               string "url"
	SshURL            string "ssh_url"
	HttpURL           string "http_url"
}

type Repository struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

type ObjectAttributes struct {
	ID              int `json:"id"`
	Title           string `json:"title"`
	AssigneeID      int `json:"assignee_id"`
	AuthorID        int `json:"author_id"`
	ProjectID       int `json:"project_id"`
	CreatedAt       customTime `json:"created_at"`
	UpdatedAt       customTime `json:"updated_at"`
	Position        int `json:"position"`
	BranchName      string `json:"branch_name"`
	Description     string `json:"description"`
	MilestoneID     int `json:"milestone_id"`
	State           string `json:"state"`
	IID             int `json:"iid"`
	URL             string `json:"url"`
	Action          string `json:"action"`
	TargetBranch    string `json:"target_branch"`
	SourceBranch    string `json:"source_branch"`
	SourceProjectId int `json:"source_project_id"`
	TargetProjectID int `json:"target_project_id"`
	StCommits       string `json:"st_commits"`
	MergeStatus     string `json:"merge_status"`
	Content         string `json:"content"`
	Format          string `json:"format"`
	Message         string `json:"message"`
	Slug            string `json:"slug"`
	Ref             string `json:"ref"`
	Tag             bool `json:"tag"`
	Sha             string `json:"sha"`
	BeforeSha       string `json:"before_sha"`
	Status          string `json:"status"`
	Stages          []string `json:"stages"`
	Duration        int `json:"duration"`
	Note            string `json:"note"`
	NotebookType    string `json:"noteable_type"`
	At              customTime `json:"attachment"`
	LineCode        string `json:"line_code"`
	CommitID        string `json:"commit_id"`
	NoteableID      int `json:"noteable_id"`
	System          bool `json:"system"`
	WorkInProgress  bool `json:"work_in_progress"`
	StDiffs         []StDiff `json:"st_diffs"`
	Source          Source `json:"source"`
	Target          Target `json:"target"`
	LastCommit      LastCommit `json:"last_commit"`
	Assignee        Assignee `json:"assignee"`
}

type MergeRequest struct {
	ID              int `json:"id"`
	TargetBranch    string `json:"target_branch"`
	SourceBranch    string `json:"source_branch"`
	SourceProjectId string `json:"source_project_id"`
	AssigneeID      int `json:"assignee_id"`
	AuthorID        int `json:"author_id"`
	Title           string `json:"title"`
	CreatedAt       customTime `json:"created_at"`
	UpdatedAt       customTime `json:"updated_at"`
	MilestoneID     int `json:"milestone_id"`
	State           string `json:"state"`
	MergeStatus     string `json:"merge_status"`
	TargetProjectID int `json:"target_project_id"`
	IID             int `json:"iid"`
	Description     string `json:"description"`
	Position        int `json:"position"`
	LockedAt        customTime `"locked_at"`
	Source          Source `json:"source"`
	Target          Target `json:"target"`
	LastCommit      LastCommit `json:"last_commit"`
	WorkInProgress  bool `json:"work_in_progress"`
	Assignee        Assignee `json:"assignee"`
}

type Assignee struct {
	Name      string `json"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type StDiff struct {
	Diff        string `"diff"`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool `json:"new_file"`
	RenamedFile bool `json:"renamed_file"`
	DeletedFile bool `json:"deleted_file"`
}

type Source struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSshURL         string `json:"git_ssh_url"`
	GitHttpURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int `json"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SshURL            string `json:"ssh_url"`
	HttpURL           string `json:"http_url"`
}

type Target struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSshURL         string `json:"git_ssh_url"`
	GitHttpURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int `json"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SshURL            string `json:"ssh_url"`
	HttpURL           string `json:"http_url"`
}

type LastCommit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp customTime `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
