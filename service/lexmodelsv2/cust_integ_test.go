// +build integration

package lexmodelsv2

import (
	"testing"

	"github.com/aws/aws-sdk-go/awstesting/integration"
)

func TestInteg_ListBots(t *testing.T) {
	sess := integration.SessionWithDefaultRegion("us-west-2")

	client := New(sess)

	_, err := client.ListBots(&ListBotsInput{})
	if err != nil {
		t.Fatalf("expect API call, got %v", err)
	}
}
