# goauditlogger
## Audit Logger
Here we can configure the Output file and the logging format for the auditing
eg.
```
export LOGGER_FORMAT='{
    "pid": "${pid}",
    "host":"${host}"
    "method":"${method}"
    "path":"${path}"
    "url":"${url}"
    "ua":"${ua}"
    "latency":"${latency}"
    "status":"${status}"
}'

export LOGGER_OUTPUT="./AuditLoggerFile.log"
```
Here if we set the these two environment variables <strong><br>LOGGER_FORMAT<br>LOGGER_OUTPUT<br></strong>
then the output is save in <strong>AuditLoggerFile.log</strong> with the specified json format 

## Application Logger
the default log level is <strong>INFO</strong><br>
if we Want to change the Log level we need to set the environment varibale LOG_LEVEL
<br>
for example
<br>
```export LOG_LEVEL=debug```
<br>
the different values it accepted for <strong>LOG_LEVEL</strong> are 
1. debug
2. info
3. warn
4. error
