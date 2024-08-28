package tapd

import (
	"context"
	"errors"
	"io"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// WebhookDispatcher is a dispatcher for webhook events.
type WebhookDispatcher struct {
	storyCreateListeners []StoryCreateListener
	storyUpdateListeners []StoryUpdateListener
	bugCreateListeners   []BugCreateListener
}

type WebhookDispatcherOption func(*WebhookDispatcher)

func WithWebhookDispatcherRegister(listeners ...any) WebhookDispatcherOption {
	return func(d *WebhookDispatcher) {
		d.Register(listeners...)
	}
}

// NewWebhookDispatcher returns a new WebhookDispatcher instance.
func NewWebhookDispatcher(opts ...WebhookDispatcherOption) *WebhookDispatcher {
	dispatcher := &WebhookDispatcher{}
	for _, opt := range opts {
		opt(dispatcher)
	}
	return dispatcher
}

func (d *WebhookDispatcher) Register(listeners ...any) {
	for _, listener := range listeners {
		if l, ok := listener.(StoryUpdateListener); ok {
			d.RegisterStoryUpdateListener(l)
		}
	}
}

func (d *WebhookDispatcher) Dispatch(ctx context.Context, event any) error {
	switch e := event.(type) {
	case *BugCreateEvent:
		return d.processBugCreate(ctx, e)
	case *StoryUpdateEvent:
		return d.processStoryUpdate(ctx, e)
	case *StoryCreateEvent:
		return d.processStoryCreate(ctx, e)
	default:
		return errors.New("tapd: webhook dispatcher unsupported event")
	}
}

func (d *WebhookDispatcher) DispatchPayload(ctx context.Context, payload []byte) error {
	_, event, err := ParseWebhookEvent(payload)
	if err != nil {
		return err
	}
	return d.Dispatch(ctx, event)
}

func (d *WebhookDispatcher) DispatchRequest(req *http.Request) error {
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return d.DispatchPayload(req.Context(), payload)
}

func (d *WebhookDispatcher) RegisterStoryCreateListener(listeners ...StoryCreateListener) {
	d.storyCreateListeners = append(d.storyCreateListeners, listeners...)
}

func (d *WebhookDispatcher) RegisterStoryUpdateListener(listeners ...StoryUpdateListener) {
	d.storyUpdateListeners = append(d.storyUpdateListeners, listeners...)
}

func (d *WebhookDispatcher) RegisterBugCreateListener(listeners ...BugCreateListener) {
	d.bugCreateListeners = append(d.bugCreateListeners, listeners...)
}

func (d *WebhookDispatcher) processStoryCreate(ctx context.Context, event *StoryCreateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyCreateListeners {
		listener := listener
		eg.Go(func() error {
			return listener.OnStoryCreate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *WebhookDispatcher) processStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyUpdateListeners {
		listener := listener
		eg.Go(func() error {
			return listener.OnStoryUpdate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *WebhookDispatcher) processBugCreate(ctx context.Context, event *BugCreateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.bugCreateListeners {
		listener := listener
		eg.Go(func() error {
			return listener.OnBugCreate(ctx, event)
		})
	}
	return eg.Wait()
}
