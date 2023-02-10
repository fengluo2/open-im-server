package cmsstruct

import (
	"Open_IM/pkg/api_struct"
	sdkws "Open_IM/pkg/proto/sdkws"
)

type AdminLoginRequest struct {
	AdminName   string `json:"adminID" binding:"required"`
	Secret      string `json:"secret" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}

type AdminLoginResponse struct {
	Token    string `json:"token"`
	UserName string `json:"userName"`
	FaceURL  string `json:"faceURL"`
}

type GetUserTokenRequest struct {
	UserID      string `json:"userID" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
	PlatFormID  int32  `json:"platformID" binding:"required"`
}

type GetUserTokenResponse struct {
	Token   string `json:"token"`
	ExpTime int64  `json:"expTime"`
}

type AddUserRegisterAddFriendIDListRequest struct {
	OperationID string   `json:"operationID" binding:"required"`
	UserIDList  []string `json:"userIDList" binding:"required"`
}

type AddUserRegisterAddFriendIDListResponse struct {
}

type ReduceUserRegisterAddFriendIDListRequest struct {
	OperationID string   `json:"operationID" binding:"required"`
	UserIDList  []string `json:"userIDList" binding:"required"`
	Operation   int32    `json:"operation"`
}

type ReduceUserRegisterAddFriendIDListResponse struct {
}

type GetUserRegisterAddFriendIDListRequest struct {
	OperationID string `json:"operationID" binding:"required"`
	apistruct.RequestPagination
}

type GetUserRegisterAddFriendIDListResponse struct {
	Users []*sdkws.UserInfo `json:"users"`
	apistruct.ResponsePagination
}