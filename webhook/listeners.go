package webhook

import "context"

type (
	StoryCreateListener interface {
		OnStoryCreate(ctx context.Context, event *StoryCreateEvent) error
	}

	StoryUpdateListener interface {
		OnStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error
	}

	BugCreateListener interface {
		OnBugCreate(ctx context.Context, event *BugCreateEvent) error
	}
)
