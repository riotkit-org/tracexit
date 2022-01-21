tracexit
========

`Trace command till it Exits`

Spawns process and tracks it till it runs, then stores its status code in a file.
It's a simple CLI application just like `tee`, can be used together with tools like `tee`.

## Input

`tracexit` does not take any commandline switches except `--help` / `-h`, everything is passed through to the child process.
To use `--help` / `-h` it must be written as first argument, right after `tracexit`.

**Environment variables:**
- TRACEXIT_EXIT_CODE_PATH: File location where to store exit code of a process when it exits

## Getting tracexit

Take a look at releases tab and pick a version suitable for your platform. We support Unix-like platforms, there is no support for Windows.

You can use [eget](https://github.com/zyedidia/eget) as a 'package manager' to install `tracexit`

```bash
# for pre-release
eget --pre-release riotkit-org/tracexit --to /usr/local/bin/tracexit

# for latest stable release
eget riotkit-org/tracexit --to /usr/local/bin/tracexit
```

## Usage examples

### Export exit code to file "/tmp/exit-code" after process finishes.

```bash
export TRACEXIT_EXIT_CODE_PATH=/tmp/exit-code
tracexit mysqldump -u root -psomething | tee out.log
```

### Set extra environment variables and pass to the process

Useful when spawning a process in an environment, that does not allow adjusting shell settings.

Use case:
*Kubernetes library in Python does not allow passing environment variables. To preserve full escaping possibility (spawning process as list of strings) and avoid invoking extra `/bin/bash` we use `tracexit` to set environment variables within a wrapper process.*

```bash
tracexit env:SOME=thing mysqldump -u root -psomething
```

**Appending paths:**

```bash
tracexit env:PATH+kubectl=/opt/kubectl kubectl get pods -A
```

tracexit can concatenate variables - which means taking e.g. "PATH" real value and appending ":/opt/kubectl" in above example.

## Use case

Invoking a process in a container, or in remote shell, but the library does not return exit code of executed command.
`tracexit` will store exit code in a text file, so a next shell call can read it.

### Example - usage with `docker` library in Python

`docker` library allows performing `docker exec` operations, but it does not track exit codes.
Only command's output is returned. The process could be additionally detached, which also makes not possible to track
its exit code.

```python
# ...
container = client.containers.get('...')
exit_code, stream = container.exec_run(['tracexit', 'some-command'], stream=True)

# exit_code = None
# stream is of Generator type, after it ends we do not get exit code, only we know when output ends
```

### Example - with `kubectl`

```bash
kubectl exec -it deployment/my-deployment-name -- tracexit env:PWD=/mnt/ some-command-here
```
