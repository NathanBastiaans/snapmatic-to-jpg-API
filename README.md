# Snapmatic to JPG #
This is a simple API to convert GTA V snapmatic images to JPG files

## How do I get set up? ##
To simply run the API you can run the following command
```
go run main.go
```
This will fire up the API for local use. You're now set up to send requests to the API. By default the local URL is 
```
http://localhost:8080
```
Check if the API is working correctly by visiting the health endpoint by making a `GET` request to `/api/health`.


To convert the files you can make a `POST` request to the `/api/convert` endpoint with `image` containing the image data. 
This will return the converted image data. Don't want to go through all the hassle and process the API output? 
You can add `?save=true` to the URL and it will also save the file locally to the root folder of the project.
