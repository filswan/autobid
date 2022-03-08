package models

import (
	"go-swan/database"
	"go-swan/logs"
)

type OfflineDeals struct {
	Id                   int     `json:"id"`
	DealCid              string  `json:"deal_cid"`
	FilePath             string  `json:"file_path"`
	FileName             string  `json:"file_name"`
	FileSourceUrl        *string `json:"file_source_url"`
	Md5Origin            string  `json:"md5_origin"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	Status               string  `json:"status"`
	MinerId              int     `json:"miner_id"`
	Md5Local             string  `json:"md5_local"`
	StartEpoch           *int    `json:"start_epoch"`
	FileDownloadedStatus string  `json:"file_downloaded_status"`
	UserId               int     `json:"user_id"`
	Note                 string  `json:"note"`
	TaskId               int     `json:"task_id"`
	IsPublic             int     `json:"is_public"`
	FileSize             *string `json:"file_size"`
	PayloadCid           *string `json:"payload_cid"`
	PieceCid             *string `json:"piece_cid"`
	DownloadedAt         string  `json:"downloaded_at"`
	FileSizeNumer        float64
}

func GetOfflineDealByTaskId(taskId int) ([]*OfflineDeals, error) {
	var offlineDeals []*OfflineDeals
	err := database.GetDB().Where("task_id=?", taskId).Find(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}
