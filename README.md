# Massage Util
## Table of contents
+ [Usage](#usage)
## Usage
+ Execute command below to install this library
```
~$ go get github.com/capybara-alt/msgutil
# or
~$ go install github.com/capybara-alt/msgutil
```
+ Create message resource file (json file like below)

```{```<br>
&ensp;&ensp;```"errors": {```<br>
&ensp;&ensp;&ensp;&ensp;```"login": "login failed",```<br>
&ensp;&ensp;&ensp;&ensp;```"validate": {```<br>
&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;```"email": "Invalid mail address",```<br>
&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;```"password": "Invalid password"```<br>
&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;```"profession": "Professions must be {{0}}"```<br>
&ensp;&ensp;```},```<br>
&ensp;&ensp;```"info": {```<br>
&ensp;&ensp;&ensp;&ensp;```"login": "success!!"```<br>
&ensp;&ensp;```},```<br>
```}```
+ Call function "Init" to initialize message resource
```
    msgutil.Init("path to resource file")
```
+ Call function "GetMessage" or "GetMessageArgs" to get message defined in resource file
```
    msgutil.GetMessage("errors", "login")
    // or
    msgutil.GetMessageArgs([]string{"errors","validate","profession"}, "Engineer/Manager")
```