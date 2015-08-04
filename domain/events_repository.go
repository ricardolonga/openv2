package domain

import (
	"github.com/satori/go.uuid"
	"strings"
)

type EventsRepository struct {
	events map[string]*Event
}

func (this *EventsRepository) Save(event *Event) *Event {
	if this.events == nil {
		this.events = make(map[string]*Event)
	}

	if event.Id == "" {
		event.Id = uuid.NewV4().String()
	}

	this.events[event.Id] = event

	return event
}

func (this *EventsRepository) GetAll() *Events {
	if this.events == nil {
		this.events = make(map[string]*Event)
	}

	events := make([]Event, len(this.events))

	for _, event := range this.events {
		events = append(events, *event)
	}

	return &Events{Events: events}
}

func (this *EventsRepository) Get(id string) *Event {
	if this.events == nil {
		this.events = make(map[string]*Event)
		return nil
	}

	for _, event := range this.events {
		if strings.EqualFold(event.Id, id) {
			return event
		}
	}

	return nil
}
