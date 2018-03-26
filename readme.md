# Simple Highlevel GO subprocess package for bidirectional communication.

### Usage:
Run any shell command through this interface.

To launch a subprocess, use the setupProgram() function.

##### Example launching a python program:
``` golang
	pythonSubProcess := gosubprocess.SetupProgram("python3", "src/example.py")
```

##### Sending a message to the subprocess can now be done using the "Send()" function:
``` golang
	pythonSubProcess.Send("This will be recieved by python program")
```

##### Recieving a message from the subprocess:
``` golang
	output := pythonSubProcess.ReadOut()
```

##### Termination of the process can happen in two ways.

Killing the subprocess:
``` golang
	pythonSubProcess.kill()
```

Waiting for the subprocess to finish up. This is blocking:
``` golang
	pythonSubProcess.finish()
```