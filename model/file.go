package model

import "time"

// File struct.
type File struct {
	Name      string    `json:"name" description:"File name"`
	Extension string    `json:"extension" description:"File name"`
	Path      string    `json:"path"  description:"File stored path(url)"`
	Size      int64     `json:"size"  description:"File size(byte)"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"  description:"File created time"`
}
