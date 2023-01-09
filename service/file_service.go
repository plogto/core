package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	guuid "github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

func (s *Service) SingleUploadFile(ctx context.Context, file graphql.Upload) (*db.File, error) {
	if util.IsImage(file.ContentType) {
		name := guuid.New().String()
		ext := strings.ToLower(filepath.Ext(file.Filename))
		fileName := fmt.Sprintf("%v%v", name, ext)
		thumbnailWidth := uint(240)
		thumbnailName := fmt.Sprintf("%v-thumbnail-%v%v", name, thumbnailWidth, ext)

		// read file
		util.ReadFile(file.File, fileName)

		// calculate hash
		hash := util.CalculateHash(file.File, fileName)
		isFile, _ := s.Files.GetFileByHash(ctx, hash)
		if isFile != nil && len(isFile.ID) > 0 {
			_ = os.Remove(fileName)

			return isFile, nil
		}

		// upload original image
		img, _ := util.ConvertFileToImage(fileName, file.Size)
		originalSize, _, _ := util.CreateImage(img, fileName, false, nil)
		s.UploadImage(fileName, originalSize)

		// upload thumbnail
		thumbnailSize, _, _ := util.CreateImage(img, thumbnailName, true, &thumbnailWidth)
		s.UploadImage(thumbnailName, thumbnailSize)

		// save image data
		width := int32(img.Bounds().Size().X)
		height := int32(img.Bounds().Size().Y)
		newFile := db.CreateFileParams{
			Name:   fileName,
			Hash:   hash,
			Width:  width,
			Height: height,
		}

		os.Remove(fileName)
		os.Remove(thumbnailName)

		// save image
		return s.Files.CreateFile(ctx, newFile)
	}

	return nil, nil
}

func (s *Service) UploadFiles(ctx context.Context, files []*graphql.Upload) ([]*db.File, error) {
	var uploadedFiles []*db.File
	for _, file := range files {
		uploadedFile, _ := s.SingleUploadFile(ctx, *file)
		uploadedFiles = append(uploadedFiles, uploadedFile)
	}

	return uploadedFiles, nil
}

func (s *Service) UploadImage(fileName string, fileSize int64) {
	fileBytes := util.ReadImage(fileName, fileSize)

	SpaceRegion := os.Getenv("SPACE_REGION")
	SpaceName := os.Getenv("SPACE_NAME")
	accessKey := os.Getenv("SPACE_ACCESS_KEY")
	secretKey := os.Getenv("SPACE_SECRET_KEY")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(os.Getenv("SPACE_ENDPOINT")),
		Region:      aws.String(SpaceRegion),
	}
	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	object := s3.PutObjectInput{
		Bucket: aws.String(SpaceName),
		Key:    aws.String(fileName),
		Body:   fileBytes,
		ACL:    aws.String("public-read"),
	}

	if _, uploadErr := s3Client.PutObject(&object); uploadErr != nil {
		fmt.Errorf("error uploading file: %v", uploadErr)
	}
}

func (s *Service) GetFileByFileId(ctx context.Context, fileID uuid.UUID) (*db.File, error) {
	return s.Files.GetFileByID(ctx, fileID)
}
