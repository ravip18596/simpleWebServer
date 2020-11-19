A Simple Go Web Server Docker container
---------------------------------------
* First clone the repository
* Execute following commands
    * `cd simpleWebSever`
    * `docker build -t simplewebserver .`
    * `docker run -it --rm -p 8000:8000 -v $PWD/src:/go/src/simpleWebServer --name testserver simplewebserver`
    * `docker exec -it testserver curl --request GET localhost:8000/`
       `{"status":"OK","code":200}`
    * `docker exec -it testserver curl --request GET localhost:8000/add/a/45/b/67`
        `{"sum":112}`
        

* Also included unit testing of http handler and routing with code coverage of above 82.1%
    * `cd src`
    * `go test -v`
    * `go test -cover`