package storage

// Public interface handlers for this package
var Idea IdeaStorage
var Room RoomStorage
var User UserStorage

// IdeaStorage interface options
type IdeaStorage interface {
	Init() error
	Add(idea IdeaStruct, roomid string) (IdeaStruct, error)
	Vote(ideaId string, vote int) error
	Comment(ideaId string, comment CommentStruct) error
	Find(id string) (IdeaStruct, error)
	FindMany(ids []string) ([]IdeaStruct, error)
}

// RoomStorage interface options
type RoomStorage interface {
	Init() error
	Add(room RoomStruct) (RoomStruct, error)
	Find(id string) (RoomStruct, error)
	FindAll() ([]RoomStruct, error)
	FindWithID(ideaID string) (RoomStruct, error)
	AddMember(id string, roomTitle, roomPassword string) error
	DeleteWithTitle(title string) error
	IsUserInRoom(uid string, rid string) bool
	AddIdeaID(roomID, ideaID string) error
	GetIdeaIDs(roomID string) ([]string, error)
}

// UserStorage interface options
type UserStorage interface {
	Init() error
	Add(user UserStruct) (UserStruct, error)
	FindByName(username string) (UserStruct, error)
	FindByID(id string) (UserStruct, error)
	FindByCred(user UserStruct) (UserStruct, error)
}
