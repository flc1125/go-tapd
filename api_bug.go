package tapd

// BugService todo: add more methods
type BugService struct {
	client *Client
}

func NewBugService(client *Client) *BugService {
	return &BugService{
		client: client,
	}
}
