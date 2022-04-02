package executor


type File = func(testCode *[]byte, environmentVariables *[]string, executeLog *ExecuteLog)