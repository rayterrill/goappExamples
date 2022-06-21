# cobracli

Example cobra app demonstrating:
* multiple commands
* ability to pass arguments as short name (-p), long name (--port), or via environment variable (PORT)

### How it was built

1. Build the app:
```
mkdir cobracli && cd cobracli
go mod init github.com/rayterrill/cobracli
```
2. Init the cobra app with scaffolding:
```
cobra-cli init
```
3. Add a command to show command config:
```
cobra-cli add dosomething
```
4. Added environment variable handling piece based on https://github.com/carolynvs/stingoftheviper.

