package handler

import (
	"CloudStorage/core/helper"
	"CloudStorage/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"CloudStorage/core/internal/logic"
	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		// Determine whether the file exists or not
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		get, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if get {
			httpx.OkJson(w, &types.FileUploadResponse{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}
		// Upload file to COS
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
