package storage

// Public interface handlers for this package
var Idea IdeaStorage
var Room RoomStorage

// IdeaStorage interface options
type IdeaStorage interface {
	Init() error
	Add(idea IdeaStruct) (IdeaStruct, error)
	Vote(ideaId string, vote int) error
	Comment(ideaId string, comment CommentStruct) error
}

// RoomStorage interface options
type RoomStorage interface {
	Init() error
	Add(room RoomStruct) (RoomStruct, error)
	Find(id string) (RoomStruct, error)
	FindAll() ([]RoomStruct, error)
	FindWithID(ideaID string) (RoomStruct, error)
	AddMember(member UserStruct, roomTitle, roomPassword string) error
	DeleteWithTitle(title string) error
	AddIdeaID(roomID, ideaID string) error
	GetIdeaIDs(roomID string) ([]string, error)
}
