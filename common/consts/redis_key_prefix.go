package consts

const (
	ProjectNamePrefix = "douyin:"

	TokenPrefix = ProjectNamePrefix + "token:"
	TokenTTL    = -1

	LastMessagePrefix  = ProjectNamePrefix + "lastMessage:"
	VideoCommentPrefix = ProjectNamePrefix + "videoComment:"

	VideoRankKey        = ProjectNamePrefix + "videoRank"
	UserVideoRankPrefix = ProjectNamePrefix + "userVideoRank:"

	VideoFavoriteCountPrefix = ProjectNamePrefix + "videoFavoriteCount:"
	UserFavoriteIdPrefix     = ProjectNamePrefix + "userFavoriteId:"
	VideoFavoritedIdPrefix   = ProjectNamePrefix + "videoFavoritedId:"
	UserFavoritedCountPrefix = ProjectNamePrefix + "userFavoritedCount:"
)
