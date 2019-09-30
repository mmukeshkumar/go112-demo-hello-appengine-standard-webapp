## Summary ##

* Based of : https://cloud.google.com/appengine/docs/standard/go112/go-differences?hl=en_US&_ga=2.20433713.-840675048.1567617504
* https://github.com/GoogleCloudPlatform/golang-samples/tree/master/appengine/go11x/helloworld

## Installation ##

* Install Go from:
* https://golang.org/doc/install?download=go1.12.windows-amd64.msi

* Check Installation:

λ go version

go version go1.12.9 windows/amd64

* Set GOPATH:

λ setx GOPATH C:\dev\projects\gcp\go-work

SUCCESS: Specified value was saved.


## Running ##

* Check/display GOPATH:
* echo %GOPATH%

C:\dev\projects\gcp\go-work

* Run the main.go with this command, no appengine package is necessary:
* go run main.go

## Testing ##
* Open browser and enter: http://localhost:8080/?name=Mukesh


## Deploying ##
* Install gcloud sdk
* Open cmd prompt and enter: gcloud login to login and for setting up your config enter: gcloud init
* enter to add gcp project: gcloud config set project "enter your gcp project name here"
* Install app-engine-go gcloud component
* gcloud components install app-engine-go
* gcloud app deploy