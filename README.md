# resumeServer
Short simple project for learning Swagger

## How to use
### Installing
```
go get github.com/marcusljx/resumeServer
```

### Running the server
```
$GOPATH/bin/resumeServer
```

### Adding a resume
Send a HTTP POST request to the REST endpoint `http://localhost:8080/api/uploadResumeDetails`, with a JSON body that looks something like:
```
{
  "name" : "Thomas",
  "currentJob" : {
    "title" : "Software Engineer in Test",
    "description" : "testing engineer",
    "company" : "GRAB"
  },
  "previousJobs" : [
    {
      "title" : "R&D Engineer in Test",
      "description" : "Research Test Engineer",
      "company" : "SE Microelectronics"
    }
  ]
}
```
The request returns either a `400 Bad Request` (if the request format is incorrect), or a `201 CREATED`, with a single `UUID` string in the response body. This `UUID` can be used to retrieve the resume object from the server(see below).

The schema definition for the body can be found in `_api/swagger.yaml` under the definition `ResumeObject`. The golang struct definition used is in `server/definitions.go`.

### Getting a resume
Send a HTTP GET request to the REST endpoint `http://localhost:8080/api/getResume/{resumeID}`, where `resumeID` is a `UUID` string that has been previously returned from an `uploadResumeDetails` request.

The request returns either a `404 Not Found` (if no resume of such ID exists on the server), or a `200 OK`, with the resume object (see above example) in the response body.
