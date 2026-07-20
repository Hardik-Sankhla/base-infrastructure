package models

// Policy represents the system configuration policies that influence the planner
type Policy struct {
	AllowBeta          bool   `json:"allow_beta"`
	AllowRoot          bool   `json:"allow_root"`
	InternetRequired   bool   `json:"internet_required"`
	PreferredContainer string `json:"preferred_container"`
	MaxParallelTasks   int    `json:"max_parallel_tasks"`
	OfflineMode        bool   `json:"offline_mode"`
}
