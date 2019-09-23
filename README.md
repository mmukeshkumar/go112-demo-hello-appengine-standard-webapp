## Summary ##

* Based of : https://cloud.google.com/appengine/docs/standard/go112/go-differences?hl=en_US&_ga=2.20433713.-840675048.1567617504
* The appengine package and the google.golang.org/appengine package are no longer supported


## Installation ##

* Install from:
https://golang.org/doc/install?download=go1.12.windows-amd64.msi

* Check Installation:
go version
go version go1.11 windows/amd64

* Set GOPATH:

λ env
GOPATH=C:\Users\KUMAM048\go

C:\dev\tools\cmder_mini
λ setx GOPATH C:\dev\projects\gcp\go-work

SUCCESS: Specified value was saved.


## Running ##

* Check/display GOPATH:
echo %GOPATH%
C:\dev\projects\gcp\go-work

* Run the main.go with this command, no appengine package is necessary:
go run main.go

## Testing ##
* Open browser and enter: http://localhost:8080/?name=Mukesh


## Deploying ##

* Install app-engine-go gcloud component
* gcloud components install app-engine-go
* gcloud app deploy