# RESTful Services string to CamelCase Converter using GoLang

String to CamelCase coverter using Word Break Dynamic Programming Approach


## Assumptions
1. All the words of the english alphabet have size greater than 3. Else Accuracy is very poor.
2. dictionary API is essential.
3. Pre-requisites are installed, environment is setup and if using Azure, knowledge of Creating containerized applications using Azure AppServices or creating container instance

## Design
1. `dictionaryapi.go`:  CamelCase coversion
2. `jsonops.go` : Class for Handling JSON Operations for RESTful Services
3. `main.go` : Go RESTful webservice running on port 80
4. `CamelCase.json` : JSON file used for File Storage during RESTful Services Operations
5. `converter.go` : Implementation from String case to CamelCase conversion.
## Usage
#### API Setup
##### Terminology
`<>` : Mandatory Parameter and `[]` : Option Parameter
##### Using in Physical Machines
   ``` bash
   root@localhost $ git clone https://github.com/sribs/CamelCase  
   root@localhost $ go get github.com/gorilla/mux && go install /path/to/CamelCase/rest
   root@localhost $ cd /path/to/CamelCase/rest; rest <Dictionary API Endpoint> [<Application ID> <Application Key>]
   ```
##### Using Containers Commandline
   ``` bash
   root@localhost $ docker run sriharshabs/golangcamelcaseapi:test rest <Dictionary API Endpoint> [<Application ID> <Application Key>]
   ```
##### Using Azure Container Instance or Azure AppService
    Please follow the Portal On Screen Instructions
#### API Testing
1. To get all CamelCase strings, `http://52.230.217.40/camelcase`
2. To perform a GET or DELETE for a particular string, `https://camelcase.azurewebsites.net/camelcase/<strname>`
3. To perform a POST for a particular string, `https://camelcase.azurewebsites.net/camelcase/<strname>/<minlength>`. minlength is a parameter that will greatly affect accuracy. To get good accuracy provide minlength=min(length of the valid subset word). This is mandatory
4. For Dictionary Endpoint, here is another API: `http://flaskenglishdict.centralus.azurecontainer.io/api/<word>`
## Output
``` cmd
root@LinuxDebugging:/home/sribs/CamelApp/docker/golang# curl -vvv http://52.230.217.40/camelcase
*   Trying 52.230.217.40...
* Connected to 52.230.217.40 (52.230.217.40) port 80 (#0)
> GET /camelcase HTTP/1.1
> Host: 52.230.217.40
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Fri, 08 Mar 2019 17:18:40 GMT
< Content-Length: 46
< Content-Type: text/plain; charset=utf-8
<
{"hello":"Hello","newgame":"NewGame","three":"Three","two":"Two"}
* Connection #0 to host 52.230.217.40 left intact
root@LinuxDebugging:/home/sribs/CamelApp/docker/golang#  
```
