package alidrive

import (
	"FileTransfer/config"
	"FileTransfer/pkg/utils"
	"time"
)

type AliRespError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AliFiles struct {
	Items      []AliFile `json:"items"`
	NextMarker string    `json:"next_marker"`
}

type AliFile struct {
	DriveId       string     `json:"drive_id"`
	CreatedAt     *time.Time `json:"created_at"`
	FileExtension string     `json:"file_extension"`
	FileId        string     `json:"file_id"`
	Type          string     `json:"type"`
	Name          string     `json:"name"`
	Category      string     `json:"category"`
	ParentFileId  string     `json:"parent_file_id"`
	UpdatedAt     string     `json:"updated_at"`
	Size          int64      `json:"size"`
	Thumbnail     string     `json:"thumbnail"`
	Url           string     `json:"url"`
}

func (f AliFile) GetSize() uint64 {
	return uint64(f.Size)
}

func (f AliFile) GetName() string {
	return f.Name
}

func (f AliFile) GetType() int {
	if f.Type == "folder" {
		return config.FOLDER
	}
	if f.Category == "video" {
		return config.VIDEO
	}
	if f.Category == "image" {
		return config.IMAGE
	}
	return utils.GetFileType(f.FileExtension)
}
