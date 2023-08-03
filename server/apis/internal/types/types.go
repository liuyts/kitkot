// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoRequest struct {
	UserId int64 `form:"user_id" vd:"$>0;msg:'user_id error'"`
}

type UserInfoResponse struct {
	BaseResponse
	User *User `json:"user"`
}

type User struct {
	Id              int64  `json:"id"`
	Username        string `json:"name"`
	Avatar          string `json:"avatar"`
	FollowCount     int64  `json:"follow_count"`
	TotalFavorited  int64  `json:"total_favorited"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	FollowerCount   int64  `json:"follower_count"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	IsFollow        bool   `json:"is_follow"`
}

type UserLoginRequest struct {
	Username string `form:"username" vd:"len($)>0&&len($)<32;msg:'用户名长度错误'"`
	Password string `form:"password" vd:"len($)>0&&len($)<32;msg:'密码长度错误'"`
}

type UserLoginResponse struct {
	BaseResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisteRequest struct {
	Username string `form:"username" vd:"len($)>0&&len($)<32;msg:'用户名长度错误'"`
	Password string `form:"password" vd:"len($)>0&&len($)<32;msg:'密码长度错误'"`
}

type UserRegisterResponse struct {
	BaseResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type MessageActionRequest struct {
	ToUserId   int64  `form:"to_user_id" vd:"$>0;msg:'to_user_id error'"`
	ActionType int64  `form:"action_type" vd:"$==1;msg:'action_type error'"`
	Content    string `form:"content" vd:"$!='';msg:'消息不能为空'"`
}

type MessageActionResponse struct {
	BaseResponse
}

type MessageChatRequest struct {
	ToUserId   int64 `form:"to_user_id" vd:"$>0;msg:'to_user_id error'"`
	PreMsgTime int64 `form:"pre_msg_time,optional"`
}

type MessageChatResponse struct {
	BaseResponse
	MessageList []*Message `json:"message_list"`
}

type Message struct {
	Id         int64  `json:"id"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type BaseResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg,omitempty"`
}
