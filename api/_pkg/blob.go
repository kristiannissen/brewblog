package pkg

import "time"

type Blob struct {
	URL        string    `json:"url"`
	PathName   string    `json:"pathname"`
	Size       uint64    `json:"size"`
	UploadedAt time.Time `json:"uploadedAt"`
}
