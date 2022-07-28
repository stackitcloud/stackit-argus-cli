# CreateJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]Job**](Job.md) |  | 

## Methods

### NewCreateJob

`func NewCreateJob(message string, data []Job, ) *CreateJob`

NewCreateJob instantiates a new CreateJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobWithDefaults

`func NewCreateJobWithDefaults() *CreateJob`

NewCreateJobWithDefaults instantiates a new CreateJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateJob) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateJob) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateJob) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *CreateJob) GetData() []Job`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateJob) GetDataOk() (*[]Job, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateJob) SetData(v []Job)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


