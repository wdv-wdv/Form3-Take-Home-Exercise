# Form3 Take Home Exercise

## The assignment  

Create a Go library which implement the `Create`, `Fetch`, and `Delete` API operations on the Form3 `accounts` fake API resource.  

This was **my first attempt at writer Go code**, therefore I started off a few simple test just to get the basic application working.  

Then I expended my testing to check for data integrity on the Attribute and Account objects properties, as well as other common error conditions like using duplicated IDs.

## Architecture  

In the module I declared the package name `accountapi` and separated the diffident operations into their own files with the view of easing future maintenance. I also created a separate file holding all the data type.  

* **AccountAPI.go** holds AccountAPI struct and constructor
* **AccountAPI-Create.go** implement `Create` operation
* **AccountAPI-Fetch.go** implement `Fetch` operation
* **AccountAPI-Delete.go** implement `Delete` operation
* **AccountAPI-dataStruct.go** holds data struct type declarations

Type listed below:

* **NewAccount** used by the `Create` operations
* **newAccount** private struct used internally by the `Create` operations
* **Account** returned by `Create` & `Fetch` operations
* **Attributes** part of the **Account** objects
* **Links** returned by `Create` & `Fetch` operations alongside the **Account** object

## Testing  

I created a few tests for basic `Create`, `Fetch`, and `Delete` operations as well as a function for producing mock **NewAccount objects**. I also turn the **data integrity routine** into a private function so it can be reused in all the test where a Account object is returned for comparing to the mock NewAccount object mention before.

The data integrity show a few **issue with missing Attributes**; `Names` and `alternative_names`. Also the test for deleting a non-existing Account is showing the API not returning the same HTTP code as per the documentation.

I believe the **Form3 API documentation** is ahead of the Fake-API. In the end I decided to keep my code close to the documentation and let the test fail instead of making the test pass for the assignment.

Here is a list of test provided in the assignment:

* **Test_Constructor** Check Constructor expect a valid url
* **Test_ConstructorSchemaWrong** Checks only "http" & "https" Scheme allowed logic
* **Test_ConstructorQSWrong** Checks no query-string allowed logic
* **Test_Create** Checks `Create` operation and data integrity
* **Test_CreateWithDuplicateID** Checks `Create` operation returns a error object
* **Test_Fetch** Checks `Fetch` operation and data integrity
* **Test_Delete** Checks `Delete` operation
* **Test_DeleteRecordNotFound** Checks `Delete` operation returns an error object

## Enhancements

### **More tests**

There is scope for more test which may include:

* `Create` operation while no ID provided and/or no OrganisationID provided
* `Create` operation missing various attributes missing
* `Fetch` operation while no/wrong ID provided
* `Delete` with the wrong version number

### **Library enhancements**

I toyed with the idea of adding a few data check on **attributes marked as required/conditional** in the documentation. However my thinking is this should be taken care of in the API self. The appropriate tests would have reviled if this is necessary.

The is also scope to do more with the **HTTP status code** being returned. In my module I only look at the Bad request status Codes 400 so that I can retrieve the error message from the response body, then batching together all the other unsuccessful code into one error. A possible enhancement that may be useful for a library user to have a indication of rate limit errors as a better defined error message.

---

THE END
