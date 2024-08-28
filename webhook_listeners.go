package tapd

import "context"

type StoryUpdateListener interface {
	OnStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error
}
