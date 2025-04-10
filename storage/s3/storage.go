package s3

import (
	"bytes"
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/minamijoyo/tfmigrate/storage"
)

// Storage is a storage.Storage implementation for AWS S3.
type Storage struct {
	// config is a storage config for s3.
	config *Config
	// client is an instance of S3Client interface to call API.
	// It is intended to be replaced with a mock for testing.
	client Client
}

var _ storage.Storage = (*Storage)(nil)

// NewStorage returns a new instance of Storage.
func NewStorage(config *Config, client Client) (*Storage, error) {
	if client == nil {
		var err error
		client, err = newClient(config)
		if err != nil {
			return nil, err
		}
	}

	s := &Storage{
		config: config,
		client: client,
	}

	return s, nil
}

// Write writes migration history data to storage.
func (s *Storage) Write(ctx context.Context, b []byte) error {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s.config.Key),
		Body:   bytes.NewReader(b),
	}
	if s.config.KmsKeyID != "" {
		input.SSEKMSKeyId = &s.config.KmsKeyID
		input.ServerSideEncryption = types.ServerSideEncryptionAwsKms
	}

	_, err := s.client.PutObject(ctx, input)

	return err
}

// Read reads migration history data from storage.
// If the key does not exist, it is assumed to be uninitialized and returns
// an empty array instead of an error.
func (s *Storage) Read(ctx context.Context) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s.config.Key),
	}

	output, err := s.client.GetObject(ctx, input)
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			// If the key does not exist
			return []byte{}, nil
		}
		// unexpected error
		return nil, err
	}

	defer output.Body.Close()

	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(output.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
