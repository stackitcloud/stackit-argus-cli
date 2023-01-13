# PlanModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**BucketSize** | **int32** |  | 
**GrafanaGlobalUsers** | **int32** |  | 
**GrafanaGlobalOrgs** | **int32** |  | 
**GrafanaGlobalDashboards** | **int32** |  | 
**AlertRules** | **int32** |  | 
**TargetNumber** | **int32** |  | 
**SamplesPerScrape** | **int32** |  | 
**GrafanaGlobalSessions** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 

## Methods

### NewPlanModel

`func NewPlanModel(planId string, bucketSize int32, grafanaGlobalUsers int32, grafanaGlobalOrgs int32, grafanaGlobalDashboards int32, alertRules int32, targetNumber int32, samplesPerScrape int32, grafanaGlobalSessions int32, ) *PlanModel`

NewPlanModel instantiates a new PlanModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanModelWithDefaults

`func NewPlanModelWithDefaults() *PlanModel`

NewPlanModelWithDefaults instantiates a new PlanModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *PlanModel) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanModel) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanModel) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetDescription

`func (o *PlanModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBucketSize

`func (o *PlanModel) GetBucketSize() int32`

GetBucketSize returns the BucketSize field if non-nil, zero value otherwise.

### GetBucketSizeOk

`func (o *PlanModel) GetBucketSizeOk() (*int32, bool)`

GetBucketSizeOk returns a tuple with the BucketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSize

`func (o *PlanModel) SetBucketSize(v int32)`

SetBucketSize sets BucketSize field to given value.


### GetGrafanaGlobalUsers

`func (o *PlanModel) GetGrafanaGlobalUsers() int32`

GetGrafanaGlobalUsers returns the GrafanaGlobalUsers field if non-nil, zero value otherwise.

### GetGrafanaGlobalUsersOk

`func (o *PlanModel) GetGrafanaGlobalUsersOk() (*int32, bool)`

GetGrafanaGlobalUsersOk returns a tuple with the GrafanaGlobalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaGlobalUsers

`func (o *PlanModel) SetGrafanaGlobalUsers(v int32)`

SetGrafanaGlobalUsers sets GrafanaGlobalUsers field to given value.


### GetGrafanaGlobalOrgs

`func (o *PlanModel) GetGrafanaGlobalOrgs() int32`

GetGrafanaGlobalOrgs returns the GrafanaGlobalOrgs field if non-nil, zero value otherwise.

### GetGrafanaGlobalOrgsOk

`func (o *PlanModel) GetGrafanaGlobalOrgsOk() (*int32, bool)`

GetGrafanaGlobalOrgsOk returns a tuple with the GrafanaGlobalOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaGlobalOrgs

`func (o *PlanModel) SetGrafanaGlobalOrgs(v int32)`

SetGrafanaGlobalOrgs sets GrafanaGlobalOrgs field to given value.


### GetGrafanaGlobalDashboards

`func (o *PlanModel) GetGrafanaGlobalDashboards() int32`

GetGrafanaGlobalDashboards returns the GrafanaGlobalDashboards field if non-nil, zero value otherwise.

### GetGrafanaGlobalDashboardsOk

`func (o *PlanModel) GetGrafanaGlobalDashboardsOk() (*int32, bool)`

GetGrafanaGlobalDashboardsOk returns a tuple with the GrafanaGlobalDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaGlobalDashboards

`func (o *PlanModel) SetGrafanaGlobalDashboards(v int32)`

SetGrafanaGlobalDashboards sets GrafanaGlobalDashboards field to given value.


### GetAlertRules

`func (o *PlanModel) GetAlertRules() int32`

GetAlertRules returns the AlertRules field if non-nil, zero value otherwise.

### GetAlertRulesOk

`func (o *PlanModel) GetAlertRulesOk() (*int32, bool)`

GetAlertRulesOk returns a tuple with the AlertRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRules

`func (o *PlanModel) SetAlertRules(v int32)`

SetAlertRules sets AlertRules field to given value.


### GetTargetNumber

`func (o *PlanModel) GetTargetNumber() int32`

GetTargetNumber returns the TargetNumber field if non-nil, zero value otherwise.

### GetTargetNumberOk

`func (o *PlanModel) GetTargetNumberOk() (*int32, bool)`

GetTargetNumberOk returns a tuple with the TargetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNumber

`func (o *PlanModel) SetTargetNumber(v int32)`

SetTargetNumber sets TargetNumber field to given value.


### GetSamplesPerScrape

`func (o *PlanModel) GetSamplesPerScrape() int32`

GetSamplesPerScrape returns the SamplesPerScrape field if non-nil, zero value otherwise.

### GetSamplesPerScrapeOk

`func (o *PlanModel) GetSamplesPerScrapeOk() (*int32, bool)`

GetSamplesPerScrapeOk returns a tuple with the SamplesPerScrape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerScrape

`func (o *PlanModel) SetSamplesPerScrape(v int32)`

SetSamplesPerScrape sets SamplesPerScrape field to given value.


### GetGrafanaGlobalSessions

`func (o *PlanModel) GetGrafanaGlobalSessions() int32`

GetGrafanaGlobalSessions returns the GrafanaGlobalSessions field if non-nil, zero value otherwise.

### GetGrafanaGlobalSessionsOk

`func (o *PlanModel) GetGrafanaGlobalSessionsOk() (*int32, bool)`

GetGrafanaGlobalSessionsOk returns a tuple with the GrafanaGlobalSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaGlobalSessions

`func (o *PlanModel) SetGrafanaGlobalSessions(v int32)`

SetGrafanaGlobalSessions sets GrafanaGlobalSessions field to given value.


### GetName

`func (o *PlanModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAmount

`func (o *PlanModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlanModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlanModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PlanModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


