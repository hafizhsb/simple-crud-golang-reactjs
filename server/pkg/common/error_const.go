package common

const (
	MESSAGE_BAD_REQ               string = "Bad Request. Check your request follows the API documentation and uses correct syntax"
	MESSAGE_SUCCESS               string = "Success"
	MESSAGE_UNAUTHORIZED          string = "Unauthorized. Make sure you are using a valid API key with the necessary permissions for your request"
	MESSAGE_NOT_FOUND             string = "Not Found. Change your request URL to match a valid API endpoint"
	MESSAGE_METHOD_NOT_ALLOWED    string = "Method Not Allowed. Change the method to follow the documentation for the resource"
	MESSAGE_UNPROCESSABLE         string = "Unprocessable Entity. Make sure that your request includes all of the required fields and your data is valid"
	MESSAGE_EXCEED_SENDING_LIMIT  string = "Exceed Sending Limit. Check that you are within the limits of your quota"
	MESSAGE_INTERNAL_SERVER_ERROR string = "Internal Server Error. Try the request again later. If the error does not resolve, contact admin"
)
