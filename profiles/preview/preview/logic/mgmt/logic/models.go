// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/blacknut/azure-sdk-for-go/tools/profileBuilder

package logic

import original "github.com/blacknut/azure-sdk-for-go/services/preview/logic/mgmt/2015-08-01-preview/logic"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type IntegrationAccountAgreementsClient = original.IntegrationAccountAgreementsClient
type IntegrationAccountCertificatesClient = original.IntegrationAccountCertificatesClient
type IntegrationAccountMapsClient = original.IntegrationAccountMapsClient
type IntegrationAccountPartnersClient = original.IntegrationAccountPartnersClient
type IntegrationAccountsClient = original.IntegrationAccountsClient
type IntegrationAccountSchemasClient = original.IntegrationAccountSchemasClient
type AgreementType = original.AgreementType

const (
	AS2          AgreementType = original.AS2
	Edifact      AgreementType = original.Edifact
	NotSpecified AgreementType = original.NotSpecified
	X12          AgreementType = original.X12
)

type EdifactCharacterSet = original.EdifactCharacterSet

const (
	EdifactCharacterSetKECA         EdifactCharacterSet = original.EdifactCharacterSetKECA
	EdifactCharacterSetNotSpecified EdifactCharacterSet = original.EdifactCharacterSetNotSpecified
	EdifactCharacterSetUNOA         EdifactCharacterSet = original.EdifactCharacterSetUNOA
	EdifactCharacterSetUNOB         EdifactCharacterSet = original.EdifactCharacterSetUNOB
	EdifactCharacterSetUNOC         EdifactCharacterSet = original.EdifactCharacterSetUNOC
	EdifactCharacterSetUNOD         EdifactCharacterSet = original.EdifactCharacterSetUNOD
	EdifactCharacterSetUNOE         EdifactCharacterSet = original.EdifactCharacterSetUNOE
	EdifactCharacterSetUNOF         EdifactCharacterSet = original.EdifactCharacterSetUNOF
	EdifactCharacterSetUNOG         EdifactCharacterSet = original.EdifactCharacterSetUNOG
	EdifactCharacterSetUNOH         EdifactCharacterSet = original.EdifactCharacterSetUNOH
	EdifactCharacterSetUNOI         EdifactCharacterSet = original.EdifactCharacterSetUNOI
	EdifactCharacterSetUNOJ         EdifactCharacterSet = original.EdifactCharacterSetUNOJ
	EdifactCharacterSetUNOK         EdifactCharacterSet = original.EdifactCharacterSetUNOK
	EdifactCharacterSetUNOX         EdifactCharacterSet = original.EdifactCharacterSetUNOX
	EdifactCharacterSetUNOY         EdifactCharacterSet = original.EdifactCharacterSetUNOY
)

type EdifactDecimalIndicator = original.EdifactDecimalIndicator

const (
	EdifactDecimalIndicatorComma        EdifactDecimalIndicator = original.EdifactDecimalIndicatorComma
	EdifactDecimalIndicatorDecimal      EdifactDecimalIndicator = original.EdifactDecimalIndicatorDecimal
	EdifactDecimalIndicatorNotSpecified EdifactDecimalIndicator = original.EdifactDecimalIndicatorNotSpecified
)

type EncryptionAlgorithm = original.EncryptionAlgorithm

const (
	EncryptionAlgorithmAES128       EncryptionAlgorithm = original.EncryptionAlgorithmAES128
	EncryptionAlgorithmAES192       EncryptionAlgorithm = original.EncryptionAlgorithmAES192
	EncryptionAlgorithmAES256       EncryptionAlgorithm = original.EncryptionAlgorithmAES256
	EncryptionAlgorithmDES3         EncryptionAlgorithm = original.EncryptionAlgorithmDES3
	EncryptionAlgorithmNone         EncryptionAlgorithm = original.EncryptionAlgorithmNone
	EncryptionAlgorithmNotSpecified EncryptionAlgorithm = original.EncryptionAlgorithmNotSpecified
	EncryptionAlgorithmRC2          EncryptionAlgorithm = original.EncryptionAlgorithmRC2
)

type HashingAlgorithm = original.HashingAlgorithm

const (
	HashingAlgorithmNone         HashingAlgorithm = original.HashingAlgorithmNone
	HashingAlgorithmNotSpecified HashingAlgorithm = original.HashingAlgorithmNotSpecified
	HashingAlgorithmSHA2256      HashingAlgorithm = original.HashingAlgorithmSHA2256
	HashingAlgorithmSHA2384      HashingAlgorithm = original.HashingAlgorithmSHA2384
	HashingAlgorithmSHA2512      HashingAlgorithm = original.HashingAlgorithmSHA2512
)

type MapType = original.MapType

const (
	MapTypeNotSpecified MapType = original.MapTypeNotSpecified
	MapTypeXslt         MapType = original.MapTypeXslt
)

type MessageFilterType = original.MessageFilterType

const (
	MessageFilterTypeExclude      MessageFilterType = original.MessageFilterTypeExclude
	MessageFilterTypeInclude      MessageFilterType = original.MessageFilterTypeInclude
	MessageFilterTypeNotSpecified MessageFilterType = original.MessageFilterTypeNotSpecified
)

type PartnerType = original.PartnerType

const (
	PartnerTypeB2B          PartnerType = original.PartnerTypeB2B
	PartnerTypeNotSpecified PartnerType = original.PartnerTypeNotSpecified
)

type SchemaType = original.SchemaType

const (
	SchemaTypeNotSpecified SchemaType = original.SchemaTypeNotSpecified
	SchemaTypeXML          SchemaType = original.SchemaTypeXML
)

type SegmentTerminatorSuffix = original.SegmentTerminatorSuffix

const (
	SegmentTerminatorSuffixCR           SegmentTerminatorSuffix = original.SegmentTerminatorSuffixCR
	SegmentTerminatorSuffixCRLF         SegmentTerminatorSuffix = original.SegmentTerminatorSuffixCRLF
	SegmentTerminatorSuffixLF           SegmentTerminatorSuffix = original.SegmentTerminatorSuffixLF
	SegmentTerminatorSuffixNone         SegmentTerminatorSuffix = original.SegmentTerminatorSuffixNone
	SegmentTerminatorSuffixNotSpecified SegmentTerminatorSuffix = original.SegmentTerminatorSuffixNotSpecified
)

type SkuName = original.SkuName

const (
	SkuNameBasic        SkuName = original.SkuNameBasic
	SkuNameFree         SkuName = original.SkuNameFree
	SkuNameNotSpecified SkuName = original.SkuNameNotSpecified
	SkuNamePremium      SkuName = original.SkuNamePremium
	SkuNameShared       SkuName = original.SkuNameShared
	SkuNameStandard     SkuName = original.SkuNameStandard
)

type TrailingSeparatorPolicy = original.TrailingSeparatorPolicy

const (
	TrailingSeparatorPolicyMandatory    TrailingSeparatorPolicy = original.TrailingSeparatorPolicyMandatory
	TrailingSeparatorPolicyNotAllowed   TrailingSeparatorPolicy = original.TrailingSeparatorPolicyNotAllowed
	TrailingSeparatorPolicyNotSpecified TrailingSeparatorPolicy = original.TrailingSeparatorPolicyNotSpecified
	TrailingSeparatorPolicyOptional     TrailingSeparatorPolicy = original.TrailingSeparatorPolicyOptional
)

type UsageIndicator = original.UsageIndicator

const (
	UsageIndicatorInformation  UsageIndicator = original.UsageIndicatorInformation
	UsageIndicatorNotSpecified UsageIndicator = original.UsageIndicatorNotSpecified
	UsageIndicatorProduction   UsageIndicator = original.UsageIndicatorProduction
	UsageIndicatorTest         UsageIndicator = original.UsageIndicatorTest
)

type X12CharacterSet = original.X12CharacterSet

const (
	X12CharacterSetBasic        X12CharacterSet = original.X12CharacterSetBasic
	X12CharacterSetExtended     X12CharacterSet = original.X12CharacterSetExtended
	X12CharacterSetNotSpecified X12CharacterSet = original.X12CharacterSetNotSpecified
	X12CharacterSetUTF8         X12CharacterSet = original.X12CharacterSetUTF8
)

type X12DateFormat = original.X12DateFormat

const (
	X12DateFormatCCYYMMDD     X12DateFormat = original.X12DateFormatCCYYMMDD
	X12DateFormatNotSpecified X12DateFormat = original.X12DateFormatNotSpecified
	X12DateFormatYYMMDD       X12DateFormat = original.X12DateFormatYYMMDD
)

type X12TimeFormat = original.X12TimeFormat

const (
	X12TimeFormatHHMM         X12TimeFormat = original.X12TimeFormatHHMM
	X12TimeFormatHHMMSS       X12TimeFormat = original.X12TimeFormatHHMMSS
	X12TimeFormatHHMMSSd      X12TimeFormat = original.X12TimeFormatHHMMSSd
	X12TimeFormatHHMMSSdd     X12TimeFormat = original.X12TimeFormatHHMMSSdd
	X12TimeFormatNotSpecified X12TimeFormat = original.X12TimeFormatNotSpecified
)

type AgreementContent = original.AgreementContent
type AS2AcknowledgementConnectionSettings = original.AS2AcknowledgementConnectionSettings
type AS2AgreementContent = original.AS2AgreementContent
type AS2EnvelopeSettings = original.AS2EnvelopeSettings
type AS2ErrorSettings = original.AS2ErrorSettings
type AS2MdnSettings = original.AS2MdnSettings
type AS2MessageConnectionSettings = original.AS2MessageConnectionSettings
type AS2OneWayAgreement = original.AS2OneWayAgreement
type AS2ProtocolSettings = original.AS2ProtocolSettings
type AS2SecuritySettings = original.AS2SecuritySettings
type AS2ValidationSettings = original.AS2ValidationSettings
type B2BPartnerContent = original.B2BPartnerContent
type BusinessIdentity = original.BusinessIdentity
type CallbackURL = original.CallbackURL
type EdifactAcknowledgementSettings = original.EdifactAcknowledgementSettings
type EdifactAgreementContent = original.EdifactAgreementContent
type EdifactDelimiterOverride = original.EdifactDelimiterOverride
type EdifactEnvelopeOverride = original.EdifactEnvelopeOverride
type EdifactEnvelopeSettings = original.EdifactEnvelopeSettings
type EdifactFramingSettings = original.EdifactFramingSettings
type EdifactMessageFilter = original.EdifactMessageFilter
type EdifactMessageIdentifier = original.EdifactMessageIdentifier
type EdifactOneWayAgreement = original.EdifactOneWayAgreement
type EdifactProcessingSettings = original.EdifactProcessingSettings
type EdifactProtocolSettings = original.EdifactProtocolSettings
type EdifactSchemaReference = original.EdifactSchemaReference
type EdifactValidationOverride = original.EdifactValidationOverride
type EdifactValidationSettings = original.EdifactValidationSettings
type IntegrationAccount = original.IntegrationAccount
type IntegrationAccountAgreement = original.IntegrationAccountAgreement
type IntegrationAccountAgreementFilter = original.IntegrationAccountAgreementFilter
type IntegrationAccountAgreementListResult = original.IntegrationAccountAgreementListResult
type IntegrationAccountAgreementListResultIterator = original.IntegrationAccountAgreementListResultIterator
type IntegrationAccountAgreementListResultPage = original.IntegrationAccountAgreementListResultPage
type IntegrationAccountAgreementProperties = original.IntegrationAccountAgreementProperties
type IntegrationAccountCertificate = original.IntegrationAccountCertificate
type IntegrationAccountCertificateListResult = original.IntegrationAccountCertificateListResult
type IntegrationAccountCertificateListResultIterator = original.IntegrationAccountCertificateListResultIterator
type IntegrationAccountCertificateListResultPage = original.IntegrationAccountCertificateListResultPage
type IntegrationAccountCertificateProperties = original.IntegrationAccountCertificateProperties
type IntegrationAccountContentHash = original.IntegrationAccountContentHash
type IntegrationAccountContentLink = original.IntegrationAccountContentLink
type IntegrationAccountListResult = original.IntegrationAccountListResult
type IntegrationAccountListResultIterator = original.IntegrationAccountListResultIterator
type IntegrationAccountListResultPage = original.IntegrationAccountListResultPage
type IntegrationAccountMap = original.IntegrationAccountMap
type IntegrationAccountMapFilter = original.IntegrationAccountMapFilter
type IntegrationAccountMapListResult = original.IntegrationAccountMapListResult
type IntegrationAccountMapListResultIterator = original.IntegrationAccountMapListResultIterator
type IntegrationAccountMapListResultPage = original.IntegrationAccountMapListResultPage
type IntegrationAccountMapProperties = original.IntegrationAccountMapProperties
type IntegrationAccountPartner = original.IntegrationAccountPartner
type IntegrationAccountPartnerFilter = original.IntegrationAccountPartnerFilter
type IntegrationAccountPartnerListResult = original.IntegrationAccountPartnerListResult
type IntegrationAccountPartnerListResultIterator = original.IntegrationAccountPartnerListResultIterator
type IntegrationAccountPartnerListResultPage = original.IntegrationAccountPartnerListResultPage
type IntegrationAccountPartnerProperties = original.IntegrationAccountPartnerProperties
type IntegrationAccountResource = original.IntegrationAccountResource
type IntegrationAccountSchema = original.IntegrationAccountSchema
type IntegrationAccountSchemaFilter = original.IntegrationAccountSchemaFilter
type IntegrationAccountSchemaListResult = original.IntegrationAccountSchemaListResult
type IntegrationAccountSchemaListResultIterator = original.IntegrationAccountSchemaListResultIterator
type IntegrationAccountSchemaListResultPage = original.IntegrationAccountSchemaListResultPage
type IntegrationAccountSchemaProperties = original.IntegrationAccountSchemaProperties
type IntegrationAccountSku = original.IntegrationAccountSku
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultKeyReferenceKeyVault = original.KeyVaultKeyReferenceKeyVault
type ListCallbackURLParameters = original.ListCallbackURLParameters
type PartnerContent = original.PartnerContent
type X12AcknowledgementSettings = original.X12AcknowledgementSettings
type X12AgreementContent = original.X12AgreementContent
type X12DelimiterOverrides = original.X12DelimiterOverrides
type X12EnvelopeOverride = original.X12EnvelopeOverride
type X12EnvelopeSettings = original.X12EnvelopeSettings
type X12FramingSettings = original.X12FramingSettings
type X12MessageFilter = original.X12MessageFilter
type X12MessageIdentifier = original.X12MessageIdentifier
type X12OneWayAgreement = original.X12OneWayAgreement
type X12ProcessingSettings = original.X12ProcessingSettings
type X12ProtocolSettings = original.X12ProtocolSettings
type X12SchemaReference = original.X12SchemaReference
type X12SecuritySettings = original.X12SecuritySettings
type X12ValidationOverride = original.X12ValidationOverride
type X12ValidationSettings = original.X12ValidationSettings

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewIntegrationAccountAgreementsClient(subscriptionID string) IntegrationAccountAgreementsClient {
	return original.NewIntegrationAccountAgreementsClient(subscriptionID)
}
func NewIntegrationAccountAgreementsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountAgreementsClient {
	return original.NewIntegrationAccountAgreementsClientWithBaseURI(baseURI, subscriptionID)
}
func NewIntegrationAccountCertificatesClient(subscriptionID string) IntegrationAccountCertificatesClient {
	return original.NewIntegrationAccountCertificatesClient(subscriptionID)
}
func NewIntegrationAccountCertificatesClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountCertificatesClient {
	return original.NewIntegrationAccountCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewIntegrationAccountMapsClient(subscriptionID string) IntegrationAccountMapsClient {
	return original.NewIntegrationAccountMapsClient(subscriptionID)
}
func NewIntegrationAccountMapsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountMapsClient {
	return original.NewIntegrationAccountMapsClientWithBaseURI(baseURI, subscriptionID)
}
func NewIntegrationAccountPartnersClient(subscriptionID string) IntegrationAccountPartnersClient {
	return original.NewIntegrationAccountPartnersClient(subscriptionID)
}
func NewIntegrationAccountPartnersClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountPartnersClient {
	return original.NewIntegrationAccountPartnersClientWithBaseURI(baseURI, subscriptionID)
}
func NewIntegrationAccountsClient(subscriptionID string) IntegrationAccountsClient {
	return original.NewIntegrationAccountsClient(subscriptionID)
}
func NewIntegrationAccountsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountsClient {
	return original.NewIntegrationAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewIntegrationAccountSchemasClient(subscriptionID string) IntegrationAccountSchemasClient {
	return original.NewIntegrationAccountSchemasClient(subscriptionID)
}
func NewIntegrationAccountSchemasClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountSchemasClient {
	return original.NewIntegrationAccountSchemasClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAgreementTypeValues() []AgreementType {
	return original.PossibleAgreementTypeValues()
}
func PossibleEdifactCharacterSetValues() []EdifactCharacterSet {
	return original.PossibleEdifactCharacterSetValues()
}
func PossibleEdifactDecimalIndicatorValues() []EdifactDecimalIndicator {
	return original.PossibleEdifactDecimalIndicatorValues()
}
func PossibleEncryptionAlgorithmValues() []EncryptionAlgorithm {
	return original.PossibleEncryptionAlgorithmValues()
}
func PossibleHashingAlgorithmValues() []HashingAlgorithm {
	return original.PossibleHashingAlgorithmValues()
}
func PossibleMapTypeValues() []MapType {
	return original.PossibleMapTypeValues()
}
func PossibleMessageFilterTypeValues() []MessageFilterType {
	return original.PossibleMessageFilterTypeValues()
}
func PossiblePartnerTypeValues() []PartnerType {
	return original.PossiblePartnerTypeValues()
}
func PossibleSchemaTypeValues() []SchemaType {
	return original.PossibleSchemaTypeValues()
}
func PossibleSegmentTerminatorSuffixValues() []SegmentTerminatorSuffix {
	return original.PossibleSegmentTerminatorSuffixValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTrailingSeparatorPolicyValues() []TrailingSeparatorPolicy {
	return original.PossibleTrailingSeparatorPolicyValues()
}
func PossibleUsageIndicatorValues() []UsageIndicator {
	return original.PossibleUsageIndicatorValues()
}
func PossibleX12CharacterSetValues() []X12CharacterSet {
	return original.PossibleX12CharacterSetValues()
}
func PossibleX12DateFormatValues() []X12DateFormat {
	return original.PossibleX12DateFormatValues()
}
func PossibleX12TimeFormatValues() []X12TimeFormat {
	return original.PossibleX12TimeFormatValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
