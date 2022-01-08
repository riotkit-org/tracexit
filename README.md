tracexit
========

`Trace command till it Exits`

Spawns process and tracks it till it runs, then stores its status code in a file.
It's a simple CLI application just like `tee`, can be used together with `tee`.

## Input

`tracexit` does not take any commandline switches, everything is passed through to the child process.

**Environment variables:**
- TRACEXIT_EXIT_CODE_PATH: File location where to store exit code of a process when it exits

## Usage examples

```bash
export TRACEXIT_EXIT_CODE_PATH=/tmp/exit-code
tracexit mysqldump -u root -psomething | tee out.log
```

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