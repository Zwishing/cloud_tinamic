package handler

import (
	"cloud_tinamic/app/logger"
	"cloud_tinamic/kitex_gen/base/user/userservice"
	"cloud_tinamic/kitex_gen/data/source/sourceservice"
	"cloud_tinamic/kitex_gen/service/collection/servicecollection"
	"github.com/cloudwego/kitex/client"
)

var (
	userClient       userservice.Client
	collectionClient servicecollection.Client
	sourceClient     sourceservice.Client
)

func init() {
	var err error

	// Create Kitex clients
	clients := []struct {
		name   string
		port   string
		client interface{}
	}{
		{"base.user.userservice", "0.0.0.0:8810", &userClient},
		{"collectionservice", "0.0.0.0:8089", &collectionClient},
		{"data.source.sourceservice", "0.0.0.0:8813", &sourceClient},
	}

	for _, c := range clients {
		switch clientPtr := c.client.(type) {
		case *userservice.Client:
			*clientPtr, err = userservice.NewClient(c.name, client.WithHostPorts(c.port))
		case *servicecollection.Client:
			*clientPtr, err = servicecollection.NewClient(c.name, client.WithHostPorts(c.port))
		case *sourceservice.Client:
			*clientPtr, err = sourceservice.NewClient(c.name, client.WithHostPorts(c.port))
		}

		if err != nil {
			logger.Log.Fatal().Err(err).Msgf("Failed to create client for %s", c.name)
		}
	}
}
