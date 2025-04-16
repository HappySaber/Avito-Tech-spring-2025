# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DummyLoginPost**](DefaultAPI.md#DummyLoginPost) | **Post** /dummyLogin | Получение тестового токена
[**LoginPost**](DefaultAPI.md#LoginPost) | **Post** /login | Авторизация пользователя
[**ProductsPost**](DefaultAPI.md#ProductsPost) | **Post** /products | Добавление товара в текущую приемку (только для сотрудников ПВЗ)
[**PvzGet**](DefaultAPI.md#PvzGet) | **Get** /pvz | Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией
[**PvzPost**](DefaultAPI.md#PvzPost) | **Post** /pvz | Создание ПВЗ (только для модераторов)
[**PvzPvzIdCloseLastReceptionPost**](DefaultAPI.md#PvzPvzIdCloseLastReceptionPost) | **Post** /pvz/{pvzId}/close_last_reception | Закрытие последней открытой приемки товаров в рамках ПВЗ
[**PvzPvzIdDeleteLastProductPost**](DefaultAPI.md#PvzPvzIdDeleteLastProductPost) | **Post** /pvz/{pvzId}/delete_last_product | Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
[**ReceptionsPost**](DefaultAPI.md#ReceptionsPost) | **Post** /receptions | Создание новой приемки товаров (только для сотрудников ПВЗ)
[**RegisterPost**](DefaultAPI.md#RegisterPost) | **Post** /register | Регистрация пользователя



## DummyLoginPost

> string DummyLoginPost(ctx).DummyLoginPostRequest(dummyLoginPostRequest).Execute()

Получение тестового токена

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dummyLoginPostRequest := *openapiclient.NewDummyLoginPostRequest("Role_example") // DummyLoginPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DummyLoginPost(context.Background()).DummyLoginPostRequest(dummyLoginPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DummyLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DummyLoginPost`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DummyLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDummyLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dummyLoginPostRequest** | [**DummyLoginPostRequest**](DummyLoginPostRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginPost

> string LoginPost(ctx).LoginPostRequest(loginPostRequest).Execute()

Авторизация пользователя

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	loginPostRequest := *openapiclient.NewLoginPostRequest("Email_example", "Password_example") // LoginPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LoginPost(context.Background()).LoginPostRequest(loginPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginPost`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginPostRequest** | [**LoginPostRequest**](LoginPostRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsPost

> Product ProductsPost(ctx).ProductsPostRequest(productsPostRequest).Execute()

Добавление товара в текущую приемку (только для сотрудников ПВЗ)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	productsPostRequest := *openapiclient.NewProductsPostRequest("Type_example", "PvzId_example") // ProductsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProductsPost(context.Background()).ProductsPostRequest(productsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProductsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsPost`: Product
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProductsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productsPostRequest** | [**ProductsPostRequest**](ProductsPostRequest.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvzGet

> []PvzGet200ResponseInner PvzGet(ctx).StartDate(startDate).EndDate(endDate).Page(page).Limit(limit).Execute()

Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startDate := time.Now() // time.Time | Начальная дата диапазона (optional)
	endDate := time.Now() // time.Time | Конечная дата диапазона (optional)
	page := int32(56) // int32 | Номер страницы (optional) (default to 1)
	limit := int32(56) // int32 | Количество элементов на странице (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PvzGet(context.Background()).StartDate(startDate).EndDate(endDate).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PvzGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvzGet`: []PvzGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PvzGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPvzGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Начальная дата диапазона | 
 **endDate** | **time.Time** | Конечная дата диапазона | 
 **page** | **int32** | Номер страницы | [default to 1]
 **limit** | **int32** | Количество элементов на странице | [default to 10]

### Return type

[**[]PvzGet200ResponseInner**](PvzGet200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvzPost

> PVZ PvzPost(ctx).PVZ(pVZ).Execute()

Создание ПВЗ (только для модераторов)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pVZ := *openapiclient.NewPVZ("City_example") // PVZ | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PvzPost(context.Background()).PVZ(pVZ).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PvzPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvzPost`: PVZ
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PvzPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPvzPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pVZ** | [**PVZ**](PVZ.md) |  | 

### Return type

[**PVZ**](PVZ.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvzPvzIdCloseLastReceptionPost

> Reception PvzPvzIdCloseLastReceptionPost(ctx, pvzId).Execute()

Закрытие последней открытой приемки товаров в рамках ПВЗ

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pvzId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PvzPvzIdCloseLastReceptionPost(context.Background(), pvzId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PvzPvzIdCloseLastReceptionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvzPvzIdCloseLastReceptionPost`: Reception
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PvzPvzIdCloseLastReceptionPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pvzId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvzPvzIdCloseLastReceptionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Reception**](Reception.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvzPvzIdDeleteLastProductPost

> PvzPvzIdDeleteLastProductPost(ctx, pvzId).Execute()

Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pvzId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PvzPvzIdDeleteLastProductPost(context.Background(), pvzId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PvzPvzIdDeleteLastProductPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pvzId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvzPvzIdDeleteLastProductPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceptionsPost

> Reception ReceptionsPost(ctx).ReceptionsPostRequest(receptionsPostRequest).Execute()

Создание новой приемки товаров (только для сотрудников ПВЗ)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	receptionsPostRequest := *openapiclient.NewReceptionsPostRequest("PvzId_example") // ReceptionsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ReceptionsPost(context.Background()).ReceptionsPostRequest(receptionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReceptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceptionsPost`: Reception
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ReceptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **receptionsPostRequest** | [**ReceptionsPostRequest**](ReceptionsPostRequest.md) |  | 

### Return type

[**Reception**](Reception.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPost

> User RegisterPost(ctx).RegisterPostRequest(registerPostRequest).Execute()

Регистрация пользователя

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	registerPostRequest := *openapiclient.NewRegisterPostRequest("Email_example", "Password_example", "Role_example") // RegisterPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RegisterPost(context.Background()).RegisterPostRequest(registerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterPost`: User
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerPostRequest** | [**RegisterPostRequest**](RegisterPostRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

