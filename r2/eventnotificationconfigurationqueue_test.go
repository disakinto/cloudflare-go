// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/r2"
)

func TestEventNotificationConfigurationQueueUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.R2.EventNotifications.Configuration.Queues.Update(
		context.TODO(),
		"example-bucket",
		"queue_id",
		r2.EventNotificationConfigurationQueueUpdateParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Rules: cloudflare.F([]r2.EventNotificationConfigurationQueueUpdateParamsRule{{
				Actions: cloudflare.F([]r2.EventNotificationConfigurationQueueUpdateParamsRulesAction{r2.EventNotificationConfigurationQueueUpdateParamsRulesActionPutObject, r2.EventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  cloudflare.F("img/"),
				Suffix:  cloudflare.F(".jpeg"),
			}, {
				Actions: cloudflare.F([]r2.EventNotificationConfigurationQueueUpdateParamsRulesAction{r2.EventNotificationConfigurationQueueUpdateParamsRulesActionPutObject, r2.EventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  cloudflare.F("img/"),
				Suffix:  cloudflare.F(".jpeg"),
			}, {
				Actions: cloudflare.F([]r2.EventNotificationConfigurationQueueUpdateParamsRulesAction{r2.EventNotificationConfigurationQueueUpdateParamsRulesActionPutObject, r2.EventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  cloudflare.F("img/"),
				Suffix:  cloudflare.F(".jpeg"),
			}}),
			CfR2Jurisdiction: cloudflare.F(r2.EventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionDefault),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventNotificationConfigurationQueueDeleteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.R2.EventNotifications.Configuration.Queues.Delete(
		context.TODO(),
		"example-bucket",
		"queue_id",
		r2.EventNotificationConfigurationQueueDeleteParams{
			AccountID:        cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			CfR2Jurisdiction: cloudflare.F(r2.EventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionDefault),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
