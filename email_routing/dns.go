// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options []option.RequestOption
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *DNSService) New(ctx context.Context, body DNSNewParams, opts ...option.RequestOption) (res *Settings, err error) {
	var env DNSNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *DNSService) Delete(ctx context.Context, body DNSDeleteParams, opts ...option.RequestOption) (res *Settings, err error) {
	var env DNSDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Unlock MX Records previously locked by Email Routing.
func (r *DNSService) Edit(ctx context.Context, body DNSEditParams, opts ...option.RequestOption) (res *Settings, err error) {
	var env DNSEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *DNSService) Get(ctx context.Context, query DNSGetParams, opts ...option.RequestOption) (res *[]DNSRecord, err error) {
	var env DNSGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of records needed to enable an Email Routing zone.
type DNSRecord struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL DNSRecordTTL `json:"ttl"`
	// DNS record type.
	Type DNSRecordType `json:"type"`
	JSON dnsRecordJSON `json:"-"`
}

// dnsRecordJSON contains the JSON metadata for the struct [DNSRecord]
type dnsRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordJSON) RawJSON() string {
	return r.raw
}

type DNSRecordTTL float64

const (
	DNSRecordTTL1 DNSRecordTTL = 1
)

func (r DNSRecordTTL) IsKnown() bool {
	switch r {
	case DNSRecordTTL1:
		return true
	}
	return false
}

// DNS record type.
type DNSRecordType string

const (
	DNSRecordTypeA      DNSRecordType = "A"
	DNSRecordTypeAAAA   DNSRecordType = "AAAA"
	DNSRecordTypeCNAME  DNSRecordType = "CNAME"
	DNSRecordTypeHTTPS  DNSRecordType = "HTTPS"
	DNSRecordTypeTXT    DNSRecordType = "TXT"
	DNSRecordTypeSRV    DNSRecordType = "SRV"
	DNSRecordTypeLOC    DNSRecordType = "LOC"
	DNSRecordTypeMX     DNSRecordType = "MX"
	DNSRecordTypeNS     DNSRecordType = "NS"
	DNSRecordTypeCERT   DNSRecordType = "CERT"
	DNSRecordTypeDNSKEY DNSRecordType = "DNSKEY"
	DNSRecordTypeDS     DNSRecordType = "DS"
	DNSRecordTypeNAPTR  DNSRecordType = "NAPTR"
	DNSRecordTypeSMIMEA DNSRecordType = "SMIMEA"
	DNSRecordTypeSSHFP  DNSRecordType = "SSHFP"
	DNSRecordTypeSVCB   DNSRecordType = "SVCB"
	DNSRecordTypeTLSA   DNSRecordType = "TLSA"
	DNSRecordTypeURI    DNSRecordType = "URI"
)

func (r DNSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordTypeA, DNSRecordTypeAAAA, DNSRecordTypeCNAME, DNSRecordTypeHTTPS, DNSRecordTypeTXT, DNSRecordTypeSRV, DNSRecordTypeLOC, DNSRecordTypeMX, DNSRecordTypeNS, DNSRecordTypeCERT, DNSRecordTypeDNSKEY, DNSRecordTypeDS, DNSRecordTypeNAPTR, DNSRecordTypeSMIMEA, DNSRecordTypeSSHFP, DNSRecordTypeSVCB, DNSRecordTypeTLSA, DNSRecordTypeURI:
		return true
	}
	return false
}

type DNSNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Settings                      `json:"result"`
	JSON    dnsNewResponseEnvelopeJSON    `json:"-"`
}

// dnsNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSNewResponseEnvelope]
type dnsNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSNewResponseEnvelopeSuccess bool

const (
	DNSNewResponseEnvelopeSuccessTrue DNSNewResponseEnvelopeSuccess = true
)

func (r DNSNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  Settings                         `json:"result"`
	JSON    dnsDeleteResponseEnvelopeJSON    `json:"-"`
}

// dnsDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSDeleteResponseEnvelope]
type dnsDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSDeleteResponseEnvelopeSuccess bool

const (
	DNSDeleteResponseEnvelopeSuccessTrue DNSDeleteResponseEnvelopeSuccess = true
)

func (r DNSDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSEditResponseEnvelopeSuccess `json:"success,required"`
	Result  Settings                       `json:"result"`
	JSON    dnsEditResponseEnvelopeJSON    `json:"-"`
}

// dnsEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSEditResponseEnvelope]
type dnsEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSEditResponseEnvelopeSuccess bool

const (
	DNSEditResponseEnvelopeSuccessTrue DNSEditResponseEnvelopeSuccess = true
)

func (r DNSEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    DNSGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     []DNSRecord                      `json:"result"`
	ResultInfo DNSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dnsGetResponseEnvelopeJSON       `json:"-"`
}

// dnsGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelope]
type dnsGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSGetResponseEnvelopeSuccess bool

const (
	DNSGetResponseEnvelopeSuccessTrue DNSGetResponseEnvelopeSuccess = true
)

func (r DNSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       dnsGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dnsGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeResultInfo]
type dnsGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
