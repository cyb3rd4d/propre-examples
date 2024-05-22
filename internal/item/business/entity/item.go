package entity

var (
	ItemEventCreated = ItemEvent{name: "item_created"}
	ItemEventUpdated = ItemEvent{name: "name_updated"}
)

type ItemEvent struct {
	name string
}

type Item struct {
	id     int
	name   string
	events []ItemEvent
}

type ItemOption func(*Item)

func ItemWithID(id int) ItemOption {
	return func(item *Item) {
		item.id = id
	}
}

func ItemWithName(name string) ItemOption {
	return func(item *Item) {
		item.name = name
	}
}

func NewItem(opts ...ItemOption) Item {
	item := Item{}

	for _, option := range opts {
		option(&item)
	}

	if item.id == 0 {
		item.events = []ItemEvent{ItemEventCreated}
	}

	return item
}

func (item Item) ID() int {
	return item.id
}

func (item Item) Name() string {
	return item.name
}

func (item Item) Events() []ItemEvent {
	return item.events
}

func (item *Item) SetName(name string) {
	item.events = append(item.events, ItemEventUpdated)
	item.name = name
}
