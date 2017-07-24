# \DefaultApi

All URIs are relative to *https://local.scoir.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HighSchoolHighSchoolIdStudentStudentIdSATPost**](DefaultApi.md#HighSchoolHighSchoolIdStudentStudentIdSATPost) | **Post** /high_school/{highSchoolId}/student/{studentId}/SAT | 


# **HighSchoolHighSchoolIdStudentStudentIdSATPost**
> Sat HighSchoolHighSchoolIdStudentStudentIdSATPost(ctx, highSchoolId, studentId, score)


Creates a new SAT test score

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **highSchoolId** | **string**| BSON ObjectId of highschool the student is linked to | 
  **studentId** | **string**| BSON ObjectId of the registered scoir student. | 
  **score** | [**Sat**](Sat.md)| The submitted test score | 

### Return type

[**Sat**](SAT.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

