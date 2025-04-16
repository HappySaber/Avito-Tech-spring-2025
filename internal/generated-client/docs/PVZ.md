# PVZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RegistrationDate** | Pointer to **time.Time** |  | [optional] 
**City** | **string** |  | 

## Methods

### NewPVZ

`func NewPVZ(city string, ) *PVZ`

NewPVZ instantiates a new PVZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPVZWithDefaults

`func NewPVZWithDefaults() *PVZ`

NewPVZWithDefaults instantiates a new PVZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PVZ) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PVZ) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PVZ) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PVZ) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegistrationDate

`func (o *PVZ) GetRegistrationDate() time.Time`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *PVZ) GetRegistrationDateOk() (*time.Time, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *PVZ) SetRegistrationDate(v time.Time)`

SetRegistrationDate sets RegistrationDate field to given value.

### HasRegistrationDate

`func (o *PVZ) HasRegistrationDate() bool`

HasRegistrationDate returns a boolean if a field has been set.

### GetCity

`func (o *PVZ) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PVZ) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PVZ) SetCity(v string)`

SetCity sets City field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


