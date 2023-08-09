package consts

const (
	UserId = "user_id"

	MsgTypeRecv = 0
	MsgTypeSend = 1

	CommentAdd = 1
	CommentDel = 2

	FavoriteAdd = 1
	FavoriteDel = 2
)

const (
	APIsMachineId = iota
	UserMachineId
	VideoMachineId
	ChatMachineId
	CommentMachineId
)

const (
	DefaultAvatar          = "https://avatars.githubusercontent.com/u/1967156"
	DefaultBackGroundImage = "https://kitkot.oss-cn-shanghai.aliyuncs.com/default_background_image.png"
	DefaultSignature       = "这个人很懒，什么都没有留下"
)

const (
	FilePath    = "D:/Desktop/douyin-file"
	FileTmpPath = FilePath + "/tmp"
	ImagePath   = FilePath + "/images"
	VideoPath   = FilePath + "/videos"
)
