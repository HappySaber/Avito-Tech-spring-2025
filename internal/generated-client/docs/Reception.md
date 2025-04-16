# Reception

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DateTime** | **time.Time** |  | 
**PvzId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewReception

`func NewReception(dateTime time.Time, pvzId string, status string, ) *Reception`

NewReception instantiates a new Reception object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceptionWithDefaults

`func NewReceptionWithDefaults() *Reception`

NewReceptionWithDefaults instantiates a new Reception object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reception) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reception) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reception) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Reception) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateTime

`func (o *Reception) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *Reception) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *Reception) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.


### GetPvzId

`func (o *Reception) GetPvzId() string`

GetPvzId returns the PvzId field if non-nil, zero value otherwise.

### GetPvzIdOk

`func (o *Reception) GetPvzIdOk() (*string, bool)`

GetPvzIdOk returns a tuple with the PvzId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvzId

`func (o *Reception) SetPvzId(v string)`

SetPvzId sets PvzId field to given value.


### GetStatus

`func (o *Reception) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Reception) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Reception) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


