package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var googleDrive *GoogleDriveService

type GoogleDriveService struct {
	srv *drive.Service
}

func initGoogleDrive(ctx context.Context) error {

	srv, err := drive.NewService(ctx, option.WithCredentialsFile("project-eb52dd00-d7cd-4eb5-a75-e5e200351865.json"))
	if err != nil {
		logrus.Errorf("Unable to retrieve Drive client: %v", err)
		panic(err)
	}

	googleDrive = &GoogleDriveService{srv: srv}
	return nil
}
