package anchore

import (
	"errors"

	"k8s.io/klog"
)

var ErrImageDoesNotExist = errors.New("image does not exist")

// ImageProvider is a consumer-focused abstraction that describes the needs of the admission controller with respect
// to Anchore image-related API operations.
type ImageProvider interface {
	Get(imageReference string) (Image, error)
	Analyze(imageReference string) error
	DoesPolicyCheckPass(imageDigest, imageTag, policyBundleID string) (bool, error)
}

// GetImageProvider returns an abstract ImageProvider with an underlying implementation for fulfilling Anchore image
// -related API requests.
func GetImageProvider(authConfig AuthConfiguration, user, endpoint string) ImageProvider {
	for _, entry := range authConfig.Users {
		if entry.Username == user {
			klog.Info("Found selector match for user ", "Username=", entry.Username)
			return newAuthenticatedImageProvider(entry.Username, entry.Password, endpoint)
		}
	}

	return nil
}
