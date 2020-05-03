package forms

import "mime/multipart"

type SpiderForm struct {
	SpiderUploadForm

	Name string `form:"name" json:"name" binding:"required,min=1,max=32"`
}

type SpiderUploadForm struct {
	BaseForm

	File *multipart.FileHeader `form:"file" json:"-" binding:"required"`
}