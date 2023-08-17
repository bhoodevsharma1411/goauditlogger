# goauditlogger
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