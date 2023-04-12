package enum

type Status string

const (
	Pending Status = "pending"
	Paid    Status = "paid"
	Failed  Status = "failed"
)
