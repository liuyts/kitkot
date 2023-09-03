package consts

const (
	ProjectNamePrefix = "douyin:"

	TokenPrefix = ProjectNamePrefix + "token:"
	TokenTTL    = 18000

	LastMessagePrefix       = ProjectNamePrefix + "lastMessage:"
	VideoCommentCountPrefix = ProjectNamePrefix + "videoCommentCount:"

	VideoRankKey        = ProjectNamePrefix + "videoRank"
	UserVideoRankPrefix = ProjectNamePrefix + "userVideoRank:"

	UserFavoriteIdPrefix     = ProjectNamePrefix + "userFavoriteId:"
	VideoFavoritedIdPrefix   = ProjectNamePrefix + "videoFavoritedId:"
	UserFavoritedCountPrefix = ProjectNamePrefix + "userFavoritedCount:"

	UserFollowPrefix   = ProjectNamePrefix + "userFollow:"
	UserFollowerPrefix = ProjectNamePrefix + "userFollower:"
)
