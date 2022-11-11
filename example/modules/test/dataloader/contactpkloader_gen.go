// Code generated by github.com/shiqiyue/dataloaden, DO NOT EDIT.

package dataloader

import (
	"context"
	"sync"
	"time"

	model "github.com/shiqiyue/go-admin-gen/example"
)

// ContactPkLoaderConfig captures the config to create a new ContactPkLoader
type ContactPkLoaderConfig struct {
	// Fetch is a method that provides the data for the loader
	Fetch func(ctx context.Context, keys []int64) ([]*model.Contact, []error)

	// Wait is how long wait before sending a batch
	Wait time.Duration

	// MaxBatch will limit the maximum number of keys to send in one batch, 0 = not limit
	MaxBatch int
}

// NewContactPkLoader creates a new ContactPkLoader given a fetch, wait, and maxBatch
func NewContactPkLoader(config ContactPkLoaderConfig) *ContactPkLoader {
	return &ContactPkLoader{
		fetch:    config.Fetch,
		wait:     config.Wait,
		maxBatch: config.MaxBatch,
	}
}

// ContactPkLoader batches
type ContactPkLoader struct {
	// this method provides the data for the loader
	fetch func(ctx context.Context, keys []int64) ([]*model.Contact, []error)

	// how long to done before sending a batch
	wait time.Duration

	// this will limit the maximum number of keys to send in one batch, 0 = no limit
	maxBatch int

	// INTERNAL

	// the current batch. keys will continue to be collected until timeout is hit,
	// then everything will be sent to the fetch method and out to the listeners
	batch *contactPkLoaderBatch

	// mutex to prevent races
	mu sync.Mutex
}

type contactPkLoaderBatch struct {
	keys    []int64
	data    []*model.Contact
	error   []error
	closing bool
	done    chan struct{}
}

// Load a Contact by key, batching and caching will be applied automatically
func (l *ContactPkLoader) Load(ctx context.Context, key int64) (*model.Contact, error) {
	return l.LoadThunk(ctx, key)()
}

// LoadThunk returns a function that when called will block waiting for a Contact.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *ContactPkLoader) LoadThunk(ctx context.Context, key int64) func() (*model.Contact, error) {
	l.mu.Lock()
	if l.batch == nil {
		l.batch = &contactPkLoaderBatch{done: make(chan struct{})}
	}
	batch := l.batch
	pos := batch.keyIndex(ctx, l, key)
	l.mu.Unlock()

	return func() (*model.Contact, error) {
		<-batch.done

		var data *model.Contact
		if pos < len(batch.data) {
			data = batch.data[pos]
		}

		var err error
		// its convenient to be able to return a single error for everything
		if len(batch.error) == 1 {
			err = batch.error[0]
		} else if batch.error != nil {
			err = batch.error[pos]
		}

		return data, err
	}
}

// LoadAll fetches many keys at once. It will be broken into appropriate sized
// sub batches depending on how the loader is configured
func (l *ContactPkLoader) LoadAll(ctx context.Context, keys []int64) ([]*model.Contact, []error) {
	results := make([]func() (*model.Contact, error), len(keys))

	for i, key := range keys {
		results[i] = l.LoadThunk(ctx, key)
	}

	contacts := make([]*model.Contact, len(keys))
	errors := make([]error, len(keys))
	for i, thunk := range results {
		contacts[i], errors[i] = thunk()
	}
	return contacts, errors
}

// LoadAllThunk returns a function that when called will block waiting for a Contacts.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *ContactPkLoader) LoadAllThunk(ctx context.Context, keys []int64) func() ([]*model.Contact, []error) {
	results := make([]func() (*model.Contact, error), len(keys))
	for i, key := range keys {
		results[i] = l.LoadThunk(ctx, key)
	}
	return func() ([]*model.Contact, []error) {
		contacts := make([]*model.Contact, len(keys))
		errors := make([]error, len(keys))
		for i, thunk := range results {
			contacts[i], errors[i] = thunk()
		}
		return contacts, errors
	}
}

// keyIndex will return the location of the key in the batch, if its not found
// it will add the key to the batch
func (b *contactPkLoaderBatch) keyIndex(ctx context.Context, l *ContactPkLoader, key int64) int {
	for i, existingKey := range b.keys {
		if key == existingKey {
			return i
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(ctx, l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(ctx, l)
		}
	}

	return pos
}

func (b *contactPkLoaderBatch) startTimer(ctx context.Context, l *ContactPkLoader) {
	time.Sleep(l.wait)
	l.mu.Lock()

	// we must have hit a batch limit and are already finalizing this batch
	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(ctx, l)
}

func (b *contactPkLoaderBatch) end(ctx context.Context, l *ContactPkLoader) {
	b.data, b.error = l.fetch(ctx, b.keys)
	close(b.done)
}
