package vuls

import "github.com/inconshreveable/log15"

// TaskSchema : TaskSchema
var TaskSchema = `
type Task {
	serverName: String!
	organizationID: String!
	comments: String
}

input TaskInput {
	serverName: String!
	organizationID: String!
	comments: String
}
`

func init() {
}

var tasks = make(map[string][]*task)

type task struct {
	ServerName     string
	OrganizationID string
	Comments       *string
}

type taskInput struct {
	ServerName     string
	OrganizationID string
	Comments       *string
}

// TaskResolver : TaskResolver
type TaskResolver struct {
	t *task
}

// ServerName : ServerName
func (r *TaskResolver) ServerName() string {
	return r.t.ServerName
}

// OrganizationID : OrganizationID
func (r *TaskResolver) OrganizationID() string {
	return r.t.OrganizationID
}

// Comments : Comments
func (r *TaskResolver) Comments() *string {
	return r.t.Comments
}

// Tasks : tasks
func (r *Resolver) Tasks(args *struct{ ServerName string }) []*TaskResolver {
	var t []*TaskResolver
	for _, task := range tasks[args.ServerName] {
		t = append(t, &TaskResolver{task})
	}
	return t
}

// CreateTask : create task
func (r *Resolver) CreateTask(args *struct {
	CveID  string
	UserID string
	Task   *taskInput
}) *TaskResolver {
	task := &task{
		ServerName:     args.Task.ServerName,
		OrganizationID: args.Task.OrganizationID,
		Comments:       args.Task.Comments,
	}
	log15.Info("Call CreateTask", "method", "vuls.CreateTask", "args", args)
	tasks[args.CveID] = append(tasks[args.CveID], task)

	log15.Info("Stored Data", "method", "vuls.CreateTask", "tasks", tasks)
	return &TaskResolver{task}
}
