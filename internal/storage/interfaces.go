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
	FindAllPublic() ([]RoomStruct, error)
	AddMember(member MemberStruct, roomTitle string) error
	AddMemberWithPassword(member MemberStruct, roomTitle, roomPassword string) error
	DeleteWithTitle(title string) error
}
