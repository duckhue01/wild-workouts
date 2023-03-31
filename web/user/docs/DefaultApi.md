# ApiTitle.DefaultApi

All URIs are relative to *https://api.server.test/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getUserInformation**](DefaultApi.md#getUserInformation) | **GET** /user | Get user information



## getUserInformation

> User getUserInformation()

Get user information

### Example

```javascript
import ApiTitle from 'api_title';
let defaultClient = ApiTitle.ApiClient.instance;
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new ApiTitle.DefaultApi();
apiInstance.getUserInformation((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

