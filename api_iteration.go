package tapd

// IterationService todo: add more methods
type IterationService struct {
	client *Client
}

func NewIterationService(client *Client) *IterationService {
	return &IterationService{
		client: client,
	}
}
