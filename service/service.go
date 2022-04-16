package service

import (
	"context"
	_ "image/png"
	"sync"

	"github.com/plogto/core/database"
	"github.com/plogto/core/graph/model"
)

type Service struct {
	User             database.User
	Password         database.Password
	Post             database.Post
	File             database.File
	Connection       database.Connection
	Tag              database.Tag
	PostAttachment   database.PostAttachment
	PostTag          database.PostTag
	PostLike         database.PostLike
	PostSave         database.PostSave
	OnlineUser       database.OnlineUser
	Notification     database.Notification
	NotificationType database.NotificationType
	Notifications    map[string]chan *model.Notification
	mu               sync.Mutex
}

func New(service Service) *Service {
	return &Service{
		User:             service.User,
		Password:         service.Password,
		Post:             service.Post,
		File:             service.File,
		Connection:       service.Connection,
		Tag:              service.Tag,
		PostTag:          service.PostTag,
		PostAttachment:   service.PostAttachment,
		PostLike:         service.PostLike,
		PostSave:         service.PostSave,
		OnlineUser:       service.OnlineUser,
		Notification:     service.Notification,
		NotificationType: service.NotificationType,
		Notifications:    map[string]chan *model.Notification{},
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	// file, err := os.Open("1.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// defer file.Close()

	// var files []model.File

	// s.File.DB.Model(&files).Select()

	// for _, v := range files {
	// 	url := "https://plogspace.fra1.digitaloceanspaces.com/" + v.Name
	// 	fmt.Println(url)
	// 	// don't worry about errors
	// 	response, e := http.Get(url)
	// 	if e != nil {
	// 		log.Fatal(e)
	// 	}
	// 	defer response.Body.Close()

	// 	img, _, err := image.Decode(response.Body)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil, err
	// 	}

	// 	file := model.File{
	// 		Width:  int32(img.Bounds().Size().X),
	// 		Height: int32(img.Bounds().Size().Y),
	// 	}
	// 	fmt.Println(file)

	// 	s.File.DB.Model(&file).Where("id = ?", v.ID).Set("width = ?width").Set("height = ?height").Returning("*").Update()
	// 	fmt.Println(img.Bounds().Size())
	// }

	// blurred := song2.GaussianBlur(img, 75.0)
	// newImage := resize.Resize(150, 0, blurred, resize.Lanczos3)

	// fmt.Println(newImage.Bounds().Size())

	// out, err := os.Create("./output.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// defer out.Close()

	// if err := png.Encode(out, newImage); err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	return nil, nil
}
