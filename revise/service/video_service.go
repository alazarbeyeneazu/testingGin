package service

import (
	"testingme/entity"
)

type VideoService interface {
	FindAll() []entity.Video
	Save(video entity.Video) entity.Video
}

type videoAdapter struct {
	Videos []entity.Video
}

func NewAdapter() *videoAdapter {
	return &videoAdapter{}
}
func (videos *videoAdapter) FindAll() []entity.Video {
	return videos.Videos
}
func (videosrv *videoAdapter) Save(video entity.Video) entity.Video {
	videosrv.Videos = append(videosrv.Videos, video)
	return video
}
