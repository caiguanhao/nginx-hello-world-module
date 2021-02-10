# nginx-hello-world-module

Build nginx dynamic module using Go (cgo).

Run `make` to start local nginx:

```
make
```

After nginx is started, run this command and you will see a random float number:

```
curl -s localhost:8080
```

But it is unable to execute functions like http.Get or exec.Command:

```
curl -s localhost:8080/date
```

This code is written under macOS. To run in Linux, you may need to modify some variables.
