// File generated from our OpenAPI spec by Stainless.

package challenges

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WidgetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWidgetService] method instead.
type WidgetService struct {
	Options []option.RequestOption
}

// NewWidgetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWidgetService(opts ...option.RequestOption) (r *WidgetService) {
	r = &WidgetService{}
	r.Options = opts
	return
}
