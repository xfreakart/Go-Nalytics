# Go-Nalytics
A basic week-end project: Go based dashboard for Google Analytics real time API
![N|Solid](http://i.imgur.com/tVcxbV5.gif)

# History
Wanted to try Go(lang) for a project. As my wife pointed out; Who would use this if it has a nicer web interface?
Well.. I would. 

If you came here, searching for a google analytics real time call example, try this:
https://github.com/xfreakart/GO-Real-Time-Dump

### Install dependences:
```sh
$ go get github.com/gizak/termui
$ go get github.com/common-nighthawk/go-figure
$ go get google.golang.org/api/analytics/v3
```
### Configuration:
You need an app JSON from  the Google Cloud Console (it's free).
Go to Google Cloud Console and register a new Aplication, a json with the configuration will be downloaded.
Open the JSON and add the email of the app (eg: 1324567890-compute@developer.gserviceaccount.com) to your Google Analytics account.

Fields needed:
  - From the json
```sh
gaServiceAcctEmail  = "1324567980-compute@developer.gserviceaccount.com" // (json:"client_email")
gaTokenurl          = "https://accounts.google.com/o/oauth2/token"         // (json:"token_uri") 
goPrivateKey        = "-----BEGIN PRIVATE KEY-----\nXXXXXXXXXXXXXX=\n-----END PRIVATE KEY-----\n"
```
  - From Google Analytics
```sh
gaTableID  = "ga:1324678" 
```
Can be found in account/property/profile (it doesn't have a ga: string).

http://code.rickt.org/post/142452087425/how-to-download-google-analytics-data-with-golang


#### IMPORTANT!!
You have to add the client_email from google json to google Analytics. 
Do it as if it was a new user and give him Reading permission.

### RUN:
```sh
$ go run *.go
```
### Google Analytics Api
By default, the api can be called 500K per user per day.
This dashboard does 2 Querys each timem, one for the page list, the other for everything else.
Keep that in mind if you plan to do your own stuff.

# Todo
- Better way to handle configuration (YAML / JSON)
- Better error handling
- Improve code
- detele extra stuff


