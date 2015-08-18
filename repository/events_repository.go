package repository

import (
	"github.com/satori/go.uuid"
	"strings"
	"github.com/ricardolonga/openv2/entity"
)

type EventsRepository struct {
	events map[string]*entity.Event
}

func NewEventsRepository() *EventsRepository {
	return &EventsRepository{}
}

func (this *EventsRepository) Save(event *entity.Event) *entity.Event {
	if this.events == nil {
		this.events = make(map[string]*entity.Event, 0)
	}

	if event.Id == "" {
		event.Id = uuid.NewV4().String()
	}

	this.events[event.Id] = event

	return this.events[event.Id]
}

func (this *EventsRepository) GetAll() *[]entity.Event {
	if this.events == nil {
		this.events = make(map[string]*entity.Event, 0)
	}

	events := make([]entity.Event, 0)

	for _, event := range this.events {
		events = append(events, *event)
	}

	return &events
}

func (this *EventsRepository) GetByName(name string) *[]entity.Event {
	events := make([]entity.Event, 0)

	if this.events == nil {
		this.events = make(map[string]*entity.Event, 0)
		return &events
	}

	for _, event := range this.events {
		if strings.Contains(event.Name, name) {
			events = append(events, *event)
		}
	}

	return &events
}

func (this *EventsRepository) Get(id string) *entity.Event {
	if this.events == nil {
		this.events = make(map[string]*entity.Event, 0)
		return nil
	}

	for _, event := range this.events {
		if strings.EqualFold(event.Id, id) {
			return event
		}
	}

	return nil
}
