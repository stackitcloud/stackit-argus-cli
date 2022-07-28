# GetAllJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]Job**](Job.md) |  | 

## Methods

### NewGetAllJob

`func NewGetAllJob(message string, data []Job, ) *GetAllJob`

NewGetAllJob instantiates a new GetAllJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllJobWithDefaults

`func NewGetAllJobWithDefaults() *GetAllJob`

NewGetAllJobWithDefaults instantiates a new GetAllJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetAllJob) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAllJob) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAllJob) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetAllJob) GetData() []Job`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllJob) GetDataOk() (*[]Job, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllJob) SetData(v []Job)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


