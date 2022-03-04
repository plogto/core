package service

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	guuid "github.com/google/uuid"
	"github.com/plogto/core/graph/model"
)

func (s *Service) SingleUploadFile(ctx context.Context, file graphql.Upload) (*model.File, error) {
	SpaceRegion := os.Getenv("DO_SPACE_REGION")
	SpaceName := os.Getenv("DO_SPACE_NAME")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(os.Getenv("SPACE_ENDPOINT")),
		Region:      aws.String(SpaceRegion),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	fileName := fmt.Sprintf("%v%v", guuid.New().String(), strings.ToLower(filepath.Ext(file.Filename)))

	stream, readErr := ioutil.ReadAll(file.File)
	if readErr != nil {
		fmt.Printf("error from file %v", readErr)
	}

	fileErr := ioutil.WriteFile(fileName, stream, 0644)
	if fileErr != nil {
		fmt.Printf("file err %v", fileErr)
	}

	// calculate hash file
	tempFile, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Printf("Error opening file: %v", openErr)
	}

	encryption := sha256.New()
	if _, err := io.Copy(encryption, tempFile); err != nil {
		log.Fatal(err)
	}
	hash := fmt.Sprintf("%x\n", encryption.Sum(nil))

	defer tempFile.Close()

	isFile, _ := s.File.GetFileByHash(hash)

	if isFile != nil && len(isFile.ID) > 0 {
		_ = os.Remove(fileName)

		return isFile, nil
	}

	// store file
	f, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Printf("Error opening file: %v", openErr)
	}

	defer f.Close()

	buffer := make([]byte, file.Size)

	_, _ = f.Read(buffer)

	fileBytes := bytes.NewReader(buffer)

	object := s3.PutObjectInput{
		Bucket: aws.String(SpaceName),
		Key:    aws.String(fileName),
		Body:   fileBytes,
		ACL:    aws.String("public-read"),
	}

	if _, uploadErr := s3Client.PutObject(&object); uploadErr != nil {
		return nil, fmt.Errorf("error uploading file: %v", uploadErr)
	}

	newFile := &model.File{
		Name: fileName,
		Hash: hash,
	}
	s.File.CreateFile(newFile)

	_ = os.Remove(fileName)

	return newFile, nil
}
