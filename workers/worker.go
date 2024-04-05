// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// WorkerService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWorkerService] method instead.
type WorkerService struct {
	Options         []option.RequestOption
	AI              *AIService
	Scripts         *ScriptService
	Filters         *FilterService
	Routes          *RouteService
	AccountSettings *AccountSettingService
	Deployments     *DeploymentService
	Domains         *DomainService
	Subdomains      *SubdomainService
	Services        *ServiceService
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r *WorkerService) {
	r = &WorkerService{}
	r.Options = opts
	r.AI = NewAIService(opts...)
	r.Scripts = NewScriptService(opts...)
	r.Filters = NewFilterService(opts...)
	r.Routes = NewRouteService(opts...)
	r.AccountSettings = NewAccountSettingService(opts...)
	r.Deployments = NewDeploymentService(opts...)
	r.Domains = NewDomainService(opts...)
	r.Subdomains = NewSubdomainService(opts...)
	r.Services = NewServiceService(opts...)
	return
}

// A binding to allow the Worker to communicate with resources
type Binding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type BindingType `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Name of Worker to bind to
	Service string `json:"service"`
	// The exported class name of the Durable Object
	ClassName string `json:"class_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string `json:"script_name"`
	// R2 bucket to bind to
	BucketName string `json:"bucket_name"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding"`
	// ID of the D1 database to bind to
	ID string `json:"id"`
	// Namespace to bind to
	Namespace string      `json:"namespace"`
	Outbound  interface{} `json:"outbound,required"`
	// ID of the certificate to bind to
	CertificateID string      `json:"certificate_id"`
	Certificate   interface{} `json:"certificate,required"`
	JSON          bindingJSON `json:"-"`
	union         BindingUnion
}

// bindingJSON contains the JSON metadata for the struct [Binding]
type bindingJSON struct {
	Name          apijson.Field
	NamespaceID   apijson.Field
	Type          apijson.Field
	Environment   apijson.Field
	Service       apijson.Field
	ClassName     apijson.Field
	ScriptName    apijson.Field
	BucketName    apijson.Field
	QueueName     apijson.Field
	Binding       apijson.Field
	ID            apijson.Field
	Namespace     apijson.Field
	Outbound      apijson.Field
	CertificateID apijson.Field
	Certificate   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r bindingJSON) RawJSON() string {
	return r.raw
}

func (r *Binding) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Binding) AsUnion() BindingUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by [workers.BindingWorkersKVNamespaceBinding],
// [workers.BindingWorkersServiceBinding], [workers.BindingWorkersDoBinding],
// [workers.BindingWorkersR2Binding], [workers.BindingWorkersQueueBinding],
// [workers.BindingWorkersD1Binding], [workers.DispatchNamespaceBinding] or
// [workers.MTLSCERTBinding].
type BindingUnion interface {
	implementsWorkersBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BindingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MTLSCERTBinding{}),
		},
	)
}

type BindingWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersKVNamespaceBindingType `json:"type,required"`
	JSON bindingWorkersKVNamespaceBindingJSON `json:"-"`
}

// bindingWorkersKVNamespaceBindingJSON contains the JSON metadata for the struct
// [BindingWorkersKVNamespaceBinding]
type bindingWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersKVNamespaceBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersKVNamespaceBindingType string

const (
	BindingWorkersKVNamespaceBindingTypeKVNamespace BindingWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r BindingWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case BindingWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type BindingWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersServiceBindingType `json:"type,required"`
	JSON bindingWorkersServiceBindingJSON `json:"-"`
}

// bindingWorkersServiceBindingJSON contains the JSON metadata for the struct
// [BindingWorkersServiceBinding]
type bindingWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersServiceBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersServiceBindingType string

const (
	BindingWorkersServiceBindingTypeService BindingWorkersServiceBindingType = "service"
)

func (r BindingWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case BindingWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type BindingWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                      `json:"script_name"`
	JSON       bindingWorkersDoBindingJSON `json:"-"`
}

// bindingWorkersDoBindingJSON contains the JSON metadata for the struct
// [BindingWorkersDoBinding]
type bindingWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersDoBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersDoBindingType string

const (
	BindingWorkersDoBindingTypeDurableObjectNamespace BindingWorkersDoBindingType = "durable_object_namespace"
)

func (r BindingWorkersDoBindingType) IsKnown() bool {
	switch r {
	case BindingWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type BindingWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersR2BindingType `json:"type,required"`
	JSON bindingWorkersR2BindingJSON `json:"-"`
}

// bindingWorkersR2BindingJSON contains the JSON metadata for the struct
// [BindingWorkersR2Binding]
type bindingWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersR2Binding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersR2BindingType string

const (
	BindingWorkersR2BindingTypeR2Bucket BindingWorkersR2BindingType = "r2_bucket"
)

func (r BindingWorkersR2BindingType) IsKnown() bool {
	switch r {
	case BindingWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type BindingWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersQueueBindingType `json:"type,required"`
	JSON bindingWorkersQueueBindingJSON `json:"-"`
}

// bindingWorkersQueueBindingJSON contains the JSON metadata for the struct
// [BindingWorkersQueueBinding]
type bindingWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersQueueBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersQueueBindingType string

const (
	BindingWorkersQueueBindingTypeQueue BindingWorkersQueueBindingType = "queue"
)

func (r BindingWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case BindingWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type BindingWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersD1BindingType `json:"type,required"`
	JSON bindingWorkersD1BindingJSON `json:"-"`
}

// bindingWorkersD1BindingJSON contains the JSON metadata for the struct
// [BindingWorkersD1Binding]
type bindingWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersD1Binding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersD1BindingType string

const (
	BindingWorkersD1BindingTypeD1 BindingWorkersD1BindingType = "d1"
)

func (r BindingWorkersD1BindingType) IsKnown() bool {
	switch r {
	case BindingWorkersD1BindingTypeD1:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type BindingType string

const (
	BindingTypeKVNamespace            BindingType = "kv_namespace"
	BindingTypeService                BindingType = "service"
	BindingTypeDurableObjectNamespace BindingType = "durable_object_namespace"
	BindingTypeR2Bucket               BindingType = "r2_bucket"
	BindingTypeQueue                  BindingType = "queue"
	BindingTypeD1                     BindingType = "d1"
	BindingTypeDispatchNamespace      BindingType = "dispatch_namespace"
	BindingTypeMTLSCertificate        BindingType = "mtls_certificate"
)

func (r BindingType) IsKnown() bool {
	switch r {
	case BindingTypeKVNamespace, BindingTypeService, BindingTypeDurableObjectNamespace, BindingTypeR2Bucket, BindingTypeQueue, BindingTypeD1, BindingTypeDispatchNamespace, BindingTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceBindingJSON contains the JSON metadata for the struct
// [DispatchNamespaceBinding]
type dispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceBinding) implementsWorkersBinding() {}

func (r DispatchNamespaceBinding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type DispatchNamespaceBindingType string

const (
	DispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceBindingOutboundJSON contains the JSON metadata for the struct
// [DispatchNamespaceBindingOutbound]
type dispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                     `json:"service"`
	JSON    dispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceBindingOutboundWorkerJSON contains the JSON metadata for the
// struct [DispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceBindingParam struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[DispatchNamespaceBindingOutboundParam] `json:"outbound"`
}

func (r DispatchNamespaceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r DispatchNamespaceBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r DispatchNamespaceBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r DispatchNamespaceBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r DispatchNamespaceBindingParam) implementsWorkersBindingItemUnionParam() {}

// Outbound worker
type DispatchNamespaceBindingOutboundParam struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[DispatchNamespaceBindingOutboundWorkerParam] `json:"worker"`
}

func (r DispatchNamespaceBindingOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type DispatchNamespaceBindingOutboundWorkerParam struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceBindingOutboundWorkerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MigrationStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []MigrationStepRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []MigrationStepTransferredClass `json:"transferred_classes"`
	JSON               migrationStepJSON               `json:"-"`
}

// migrationStepJSON contains the JSON metadata for the struct [MigrationStep]
type migrationStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MigrationStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepJSON) RawJSON() string {
	return r.raw
}

type MigrationStepRenamedClass struct {
	From string                        `json:"from"`
	To   string                        `json:"to"`
	JSON migrationStepRenamedClassJSON `json:"-"`
}

// migrationStepRenamedClassJSON contains the JSON metadata for the struct
// [MigrationStepRenamedClass]
type migrationStepRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MigrationStepRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepRenamedClassJSON) RawJSON() string {
	return r.raw
}

type MigrationStepTransferredClass struct {
	From       string                            `json:"from"`
	FromScript string                            `json:"from_script"`
	To         string                            `json:"to"`
	JSON       migrationStepTransferredClassJSON `json:"-"`
}

// migrationStepTransferredClassJSON contains the JSON metadata for the struct
// [MigrationStepTransferredClass]
type migrationStepTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MigrationStepTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepTransferredClassJSON) RawJSON() string {
	return r.raw
}

type MigrationStepParam struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]MigrationStepRenamedClassParam] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]MigrationStepTransferredClassParam] `json:"transferred_classes"`
}

func (r MigrationStepParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MigrationStepRenamedClassParam struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r MigrationStepRenamedClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MigrationStepTransferredClassParam struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r MigrationStepTransferredClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type MTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string              `json:"certificate_id"`
	JSON          mtlscertBindingJSON `json:"-"`
}

// mtlscertBindingJSON contains the JSON metadata for the struct [MTLSCERTBinding]
type mtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r MTLSCERTBinding) implementsWorkersBinding() {}

func (r MTLSCERTBinding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type MTLSCERTBindingType string

const (
	MTLSCERTBindingTypeMTLSCertificate MTLSCERTBindingType = "mtls_certificate"
)

func (r MTLSCERTBindingType) IsKnown() bool {
	switch r {
	case MTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

type MTLSCERTBindingParam struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
	// The class of resource that the binding provides.
	Type param.Field[MTLSCERTBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r MTLSCERTBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MTLSCERTBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r MTLSCERTBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r MTLSCERTBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r MTLSCERTBindingParam) implementsWorkersBindingItemUnionParam() {}

func (r MTLSCERTBindingParam) implementsWorkersBindingItemUnionParam() {}

type PlacementConfiguration struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode PlacementConfigurationMode `json:"mode"`
	JSON placementConfigurationJSON `json:"-"`
}

// placementConfigurationJSON contains the JSON metadata for the struct
// [PlacementConfiguration]
type placementConfigurationJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementConfigurationJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type PlacementConfigurationMode string

const (
	PlacementConfigurationModeSmart PlacementConfigurationMode = "smart"
)

func (r PlacementConfigurationMode) IsKnown() bool {
	switch r {
	case PlacementConfigurationModeSmart:
		return true
	}
	return false
}

type PlacementConfigurationParam struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[PlacementConfigurationMode] `json:"mode"`
}

func (r PlacementConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single set of migrations to apply.
type SingleStepMigration struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []SingleStepMigrationRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []SingleStepMigrationTransferredClass `json:"transferred_classes"`
	JSON               singleStepMigrationJSON               `json:"-"`
}

// singleStepMigrationJSON contains the JSON metadata for the struct
// [SingleStepMigration]
type singleStepMigrationJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SingleStepMigration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleStepMigrationJSON) RawJSON() string {
	return r.raw
}

func (r SingleStepMigration) implementsWorkersSettingsItemMigrations() {}

type SingleStepMigrationRenamedClass struct {
	From string                              `json:"from"`
	To   string                              `json:"to"`
	JSON singleStepMigrationRenamedClassJSON `json:"-"`
}

// singleStepMigrationRenamedClassJSON contains the JSON metadata for the struct
// [SingleStepMigrationRenamedClass]
type singleStepMigrationRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleStepMigrationRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleStepMigrationRenamedClassJSON) RawJSON() string {
	return r.raw
}

type SingleStepMigrationTransferredClass struct {
	From       string                                  `json:"from"`
	FromScript string                                  `json:"from_script"`
	To         string                                  `json:"to"`
	JSON       singleStepMigrationTransferredClassJSON `json:"-"`
}

// singleStepMigrationTransferredClassJSON contains the JSON metadata for the
// struct [SingleStepMigrationTransferredClass]
type singleStepMigrationTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleStepMigrationTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleStepMigrationTransferredClassJSON) RawJSON() string {
	return r.raw
}

// A single set of migrations to apply.
type SingleStepMigrationParam struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]SingleStepMigrationRenamedClassParam] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]SingleStepMigrationTransferredClassParam] `json:"transferred_classes"`
}

func (r SingleStepMigrationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SingleStepMigrationParam) implementsWorkersScriptUpdateParamsVariant0MetadataMigrationsUnion() {
}

func (r SingleStepMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SingleStepMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SingleStepMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SingleStepMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SingleStepMigrationParam) implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsVariant0MetadataMigrationsUnion() {
}

type SingleStepMigrationRenamedClassParam struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r SingleStepMigrationRenamedClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleStepMigrationTransferredClassParam struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r SingleStepMigrationTransferredClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SteppedMigration struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []MigrationStep      `json:"steps"`
	JSON  steppedMigrationJSON `json:"-"`
}

// steppedMigrationJSON contains the JSON metadata for the struct
// [SteppedMigration]
type steppedMigrationJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SteppedMigration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r steppedMigrationJSON) RawJSON() string {
	return r.raw
}

func (r SteppedMigration) implementsWorkersSettingsItemMigrations() {}

type SteppedMigrationParam struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r SteppedMigrationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SteppedMigrationParam) implementsWorkersScriptUpdateParamsVariant0MetadataMigrationsUnion() {}

func (r SteppedMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SteppedMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SteppedMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SteppedMigrationParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

func (r SteppedMigrationParam) implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsVariant0MetadataMigrationsUnion() {
}
