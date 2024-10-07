package main

import (
	"cloud_tinamic/kitex_gen/base"
	"cloud_tinamic/kitex_gen/data/source"
	"context"
	"testing"

	"github.com/zeebo/assert"
)

func TestSourceServiceImpl_AddItem(t *testing.T) {
	// Create a mock SourceRepo
	mockRepo := &mockSourceRepo{}

	// Create a SourceServiceImpl with the mock repo
	sourceService := &SourceServiceImpl{SourceRepo: mockRepo}

	// Create a mock request
	req := &source.AddItemRequest{
		SourceCategory: source.SourceCategory_VECTOR,
		CurrentFolder:  "folder123",
		Item: &source.Item{
			Name: "test_item",
			Key:  "item123",
		},
	}

	// Set up expectations for the mock
	mockRepo.On("AddItem", req.SourceCategory, req.CurrentFolder, req.Item).Return(true, nil)

	// Call the AddItem function
	resp, err := sourceService.AddItem(context.Background(), req)

	// Assert the results
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, base.Code_SUCCESS, resp.Base.Code)
	assert.Equal(t, "添加成功", resp.Base.Msg)

	// Verify that the mock method was called
	mockRepo.AssertExpectations(t)
}
