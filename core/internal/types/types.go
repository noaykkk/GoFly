// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Identity     string `json:"identity"`
}

type UserDetailRequest struct {
	Identity string `json:"userIdentity"`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailVerificationCodeRequest struct {
	Email string `json:"email"`
}

type MailVerificationCodeResponse struct {
	Code string `json:"code"`
}

type UserRegisterRequest struct {
	Name     string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterResponse struct {
}

type FileUploadRequest struct {
	Hash string `json:"hash, optional"`
	Name string `json:"name, optional"`
	Ext  string `json:"ext, optional"`
	Size int64  `json:"size, optional"`
	Path string `json:"path, optional"`
}

type FileUploadResponse struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveResponse struct {
	Identity string `json:"identity"`
}

type UserFileListRequest struct {
	Id   int `json:"id,optional"`
	Page int `json:"page,optional"`
	Size int `json:"size,optional"`
}

type UserFileListResponse struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserFileRenameRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileRenameResponse struct {
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFolderCreateResponse struct {
	Identity string `json:"identity"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteResponse struct {
}

type UserFileMoveRequest struct {
	Idnetity       string `json:"identity"`
	ParentIdnetity string `json:"parent_identity"`
}

type UserFileMoveResponse struct {
}

type UserFileShareRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpiredTime            int    `json:"expired_time"`
}

type UserFileShareResponse struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity,optional"`
}

type ShareBasicDetailResponse struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId           int64  `json:"parent_id"`
}

type ShareBasicSaveResponse struct {
	Identity string `json:"identity"`
}

type RefreshAuthorizationRequest struct {
}

type RefreshAuthorizationResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
