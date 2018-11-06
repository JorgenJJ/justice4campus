package storage

// Public interface handlers for this package
var Idea IdeaStorage
var Room RoomStorage

// IdeaStorage interface options
type IdeaStorage interface {
	Init() error
}

// RoomStorage interface options
type RoomStorage interface {
	Init() error
	Add(room RoomStruct) (RoomStruct, error)
	FindWithTitle(title string) (RoomStruct, error)
	FindAll() ([]RoomStruct, error)
	AddMember(member UserStruct, roomTitle, roomPassword string) error
	DeleteWithTitle(title string) error
}
