package common

var (
	states *AllStates
)

type AllStates struct {
	Application *ApplicationStates
	Task        *TaskStates
	Scheduler   *SchedulerStates
}

type SchedulerStates struct {
	New        string
	Registered string
	Syncing    string
	Running    string
	Draining   string
	Stopping   string
	Stopped    string
}

type ApplicationStates struct {
	New       string
	Submitted string
	Accepted  string
	Running   string
	Rejected  string
	Completed string
	Killing   string
	Killed    string
	Failed    string
}

type TaskStates struct {
	Pending    string
	Scheduling string
	Allocated  string
	Rejected   string
	Bound      string
	Killing    string
	Killed     string
	Failed     string
	Completed  string
}

func States() *AllStates {
	if states == nil {
		states = &AllStates{
			Application: &ApplicationStates{
				New:       "New",
				Submitted: "Submitted",
				Accepted:  "Accepted",
				Running:   "Running",
				Rejected:  "Rejected",
				Completed: "Completed",
				Killing:   "Killing",
				Killed:    "Killed",
				Failed:    "Failed",
			},
			Task: &TaskStates{
				Pending:    "Pending",
				Scheduling: "Scheduling",
				Allocated:  "Allocated",
				Rejected:   "Rejected",
				Bound:      "Bound",
				Killing:    "Killing",
				Killed:     "Killed",
				Failed:     "Failed",
				Completed:  "Completed",
			},
			Scheduler: &SchedulerStates{
				New:        "New",
				Registered: "Registered",
				Syncing:    "Syncing",
				Running:    "Running",
				Draining:   "Draining",
				Stopping:   "Stopping",
				Stopped:    "Stopped",
			},
		}
	}
	return states
}
