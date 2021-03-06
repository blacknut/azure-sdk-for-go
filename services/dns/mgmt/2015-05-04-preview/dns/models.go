package dns

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// RecordType enumerates the values for record type.
type RecordType string

const (
	// A ...
	A RecordType = "A"
	// AAAA ...
	AAAA RecordType = "AAAA"
	// CNAME ...
	CNAME RecordType = "CNAME"
	// MX ...
	MX RecordType = "MX"
	// NS ...
	NS RecordType = "NS"
	// PTR ...
	PTR RecordType = "PTR"
	// SOA ...
	SOA RecordType = "SOA"
	// SRV ...
	SRV RecordType = "SRV"
	// TXT ...
	TXT RecordType = "TXT"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// PossibleRecordTypeValues returns an array of possible values for the RecordType const type.
func PossibleRecordTypeValues() []RecordType {
	return []RecordType{A, AAAA, CNAME, MX, NS, PTR, SOA, SRV, TXT}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// AaaaRecord an AAAA record.
type AaaaRecord struct {
	// Ipv6Address - Gets or sets the IPv6 address of this AAAA record in string notation.
	Ipv6Address *string `json:"ipv6Address,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// ARecord an A record.
type ARecord struct {
	// Ipv4Address - Gets or sets the IPv4 address of this A record in string notation.
	Ipv4Address *string `json:"ipv4Address,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// CnameRecord a CNAME record.
type CnameRecord struct {
	// Cname - Gets or sets the canonical name for this record without a terminating dot.
	Cname *string `json:"cname,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// MxRecord an MX record.
type MxRecord struct {
	// Preference - Gets or sets the preference metric for this record.
	Preference *int32 `json:"preference,omitempty"`
	// Exchange - Gets or sets the domain name of the mail host, without a terminating dot.
	Exchange *string `json:"exchange,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// NsRecord an NS record.
type NsRecord struct {
	// Nsdname - Gets or sets the name server name for this record, without a terminating dot.
	Nsdname *string `json:"nsdname,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// PtrRecord a PTR record.
type PtrRecord struct {
	// Ptrdname - Gets or sets the PTR target domain name for this record without a terminating dot.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// RecordSet describes a DNS RecordSet (a set of DNS records with the same name and type).
type RecordSet struct {
	autorest.Response `json:"-"`
	// Etag - Gets or sets the ETag of the RecordSet.
	Etag *string `json:"etag,omitempty"`
	// Properties - Gets or sets the properties of the RecordSet.
	Properties *RecordSetProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// MarshalJSON is the custom marshaler for RecordSet.
func (rs RecordSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rs.Etag != nil {
		objectMap["etag"] = rs.Etag
	}
	if rs.Properties != nil {
		objectMap["properties"] = rs.Properties
	}
	if rs.ID != nil {
		objectMap["id"] = rs.ID
	}
	if rs.Name != nil {
		objectMap["name"] = rs.Name
	}
	if rs.Type != nil {
		objectMap["type"] = rs.Type
	}
	if rs.Location != nil {
		objectMap["location"] = rs.Location
	}
	if rs.Tags != nil {
		objectMap["tags"] = rs.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// RecordSetListResult the response to a RecordSet List operation.
type RecordSetListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets information about the RecordSets in the response.
	Value *[]RecordSet `json:"value,omitempty"`
	// NextLink - Gets or sets the continuation token for the next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// RecordSetListResultIterator provides access to a complete listing of RecordSet values.
type RecordSetListResultIterator struct {
	i    int
	page RecordSetListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RecordSetListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RecordSetListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Response returns the raw server response from the last page request.
func (iter RecordSetListResultIterator) Response() RecordSetListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RecordSetListResultIterator) Value() RecordSet {
	if !iter.page.NotDone() {
		return RecordSet{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// IsEmpty returns true if the ListResult contains no values.
func (rslr RecordSetListResult) IsEmpty() bool {
	return rslr.Value == nil || len(*rslr.Value) == 0
}

// recordSetListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rslr RecordSetListResult) recordSetListResultPreparer() (*http.Request, error) {
	if rslr.NextLink == nil || len(to.String(rslr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rslr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// RecordSetListResultPage contains a page of RecordSet values.
type RecordSetListResultPage struct {
	fn   func(RecordSetListResult) (RecordSetListResult, error)
	rslr RecordSetListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RecordSetListResultPage) Next() error {
	next, err := page.fn(page.rslr)
	if err != nil {
		return err
	}
	page.rslr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RecordSetListResultPage) NotDone() bool {
	return !page.rslr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Response returns the raw server response from the last page request.
func (page RecordSetListResultPage) Response() RecordSetListResult {
	return page.rslr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page RecordSetListResultPage) Values() []RecordSet {
	if page.rslr.IsEmpty() {
		return nil
	}
	return *page.rslr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// RecordSetProperties represents the properties of the records in the RecordSet.
type RecordSetProperties struct {
	// TTL - Gets or sets the TTL of the records in the RecordSet.
	TTL *int64 `json:"TTL,omitempty"`
	// ARecords - Gets or sets the list of A records in the RecordSet.
	ARecords *[]ARecord `json:"ARecords,omitempty"`
	// AAAARecords - Gets or sets the list of AAAA records in the RecordSet.
	AAAARecords *[]AaaaRecord `json:"AAAARecords,omitempty"`
	// MXRecords - Gets or sets the list of MX records in the RecordSet.
	MXRecords *[]MxRecord `json:"MXRecords,omitempty"`
	// NSRecords - Gets or sets the list of NS records in the RecordSet.
	NSRecords *[]NsRecord `json:"NSRecords,omitempty"`
	// PTRRecords - Gets or sets the list of PTR records in the RecordSet.
	PTRRecords *[]PtrRecord `json:"PTRRecords,omitempty"`
	// SRVRecords - Gets or sets the list of SRV records in the RecordSet.
	SRVRecords *[]SrvRecord `json:"SRVRecords,omitempty"`
	// TXTRecords - Gets or sets the list of TXT records in the RecordSet.
	TXTRecords *[]TxtRecord `json:"TXTRecords,omitempty"`
	// CNAMERecord - Gets or sets the CNAME record in the RecordSet.
	CNAMERecord *CnameRecord `json:"CNAMERecord,omitempty"`
	// SOARecord - Gets or sets the SOA record in the RecordSet.
	SOARecord *SoaRecord `json:"SOARecord,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Resource ...
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// SoaRecord an SOA record.
type SoaRecord struct {
	// Host - Gets or sets the domain name of the authoritative name server, without a temrinating dot.
	Host *string `json:"host,omitempty"`
	// Email - Gets or sets the email for this record.
	Email *string `json:"email,omitempty"`
	// SerialNumber - Gets or sets the serial number for this record.
	SerialNumber *int64 `json:"serialNumber,omitempty"`
	// RefreshTime - Gets or sets the refresh value for this record.
	RefreshTime *int64 `json:"refreshTime,omitempty"`
	// RetryTime - Gets or sets the retry time for this record.
	RetryTime *int64 `json:"retryTime,omitempty"`
	// ExpireTime - Gets or sets the expire time for this record.
	ExpireTime *int64 `json:"expireTime,omitempty"`
	// MinimumTTL - Gets or sets the minimum TTL value for this record.
	MinimumTTL *int64 `json:"minimumTTL,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// SrvRecord an SRV record.
type SrvRecord struct {
	// Priority - Gets or sets the priority metric for this record.
	Priority *int32 `json:"priority,omitempty"`
	// Weight - Gets or sets the weight metric for this this record.
	Weight *int32 `json:"weight,omitempty"`
	// Port - Gets or sets the port of the service for this record.
	Port *int32 `json:"port,omitempty"`
	// Target - Gets or sets the domain name of the target for this record, without a terminating dot.
	Target *string `json:"target,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// SubResource ...
type SubResource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// TxtRecord a TXT record.
type TxtRecord struct {
	// Value - Gets or sets the text value of this record.
	Value *[]string `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Zone describes a DNS zone.
type Zone struct {
	autorest.Response `json:"-"`
	// Etag - Gets or sets the ETag of the zone that is being updated, as received from a Get operation.
	Etag *string `json:"etag,omitempty"`
	// Properties - Gets or sets the properties of the zone.
	Properties *ZoneProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// MarshalJSON is the custom marshaler for Zone.
func (z Zone) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if z.Etag != nil {
		objectMap["etag"] = z.Etag
	}
	if z.Properties != nil {
		objectMap["properties"] = z.Properties
	}
	if z.ID != nil {
		objectMap["id"] = z.ID
	}
	if z.Name != nil {
		objectMap["name"] = z.Name
	}
	if z.Type != nil {
		objectMap["type"] = z.Type
	}
	if z.Location != nil {
		objectMap["location"] = z.Location
	}
	if z.Tags != nil {
		objectMap["tags"] = z.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// ZoneListResult the response to a Zone List or ListAll operation.
type ZoneListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets information about the zones in the response.
	Value *[]Zone `json:"value,omitempty"`
	// NextLink - Gets or sets the continuation token for the next page.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// ZoneListResultIterator provides access to a complete listing of Zone values.
type ZoneListResultIterator struct {
	i    int
	page ZoneListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ZoneListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ZoneListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Response returns the raw server response from the last page request.
func (iter ZoneListResultIterator) Response() ZoneListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ZoneListResultIterator) Value() Zone {
	if !iter.page.NotDone() {
		return Zone{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// IsEmpty returns true if the ListResult contains no values.
func (zlr ZoneListResult) IsEmpty() bool {
	return zlr.Value == nil || len(*zlr.Value) == 0
}

// zoneListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (zlr ZoneListResult) zoneListResultPreparer() (*http.Request, error) {
	if zlr.NextLink == nil || len(to.String(zlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(zlr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// ZoneListResultPage contains a page of Zone values.
type ZoneListResultPage struct {
	fn  func(ZoneListResult) (ZoneListResult, error)
	zlr ZoneListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ZoneListResultPage) Next() error {
	next, err := page.fn(page.zlr)
	if err != nil {
		return err
	}
	page.zlr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ZoneListResultPage) NotDone() bool {
	return !page.zlr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Response returns the raw server response from the last page request.
func (page ZoneListResultPage) Response() ZoneListResult {
	return page.zlr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page ZoneListResultPage) Values() []Zone {
	if page.zlr.IsEmpty() {
		return nil
	}
	return *page.zlr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/dns/mgmt/2015-05-04-preview/dns instead.
// ZoneProperties represents the properties of the zone.
type ZoneProperties struct {
	// MaxNumberOfRecordSets - Gets or sets the maximum number of record sets that can be created in this zone.
	MaxNumberOfRecordSets *int64 `json:"maxNumberOfRecordSets,omitempty"`
	// NumberOfRecordSets - Gets or sets the current number of record sets in this zone.
	NumberOfRecordSets *int64 `json:"numberOfRecordSets,omitempty"`
}
