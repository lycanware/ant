## Lycanware Ant
A CLI build tool to help with packaging software. Initially created to be run by an IDE pre-build to copy resources into a single output directory.

## Install
```sh
go get -u github.com/lycanware/ant
```

## Compile
cd to directory containing `main.go` and run the following
```sh
go build
```

## Examples
View the `example.yml` file for an example of all available options.

The following will recursively `copy` the folder `cert` to `bin/app/cert`. The target directory will be deleted first.

```yml
actions:
    # Certificates for HTTPS and digital signing keys for JWTs
  - action: copy
    src: cert
    dst: bin/app/cert
    dst-clear-first: true
```

Use the following command to run the tool:

```sh
// <binary_path> b <yml_path>
./ant b example.yml
```

## Environment Variables
If the CLI is run from an IDE that passes environment variables as a string, the tool will automatically resolve them.
The regex used to detect environment variables is `$[A-Z_]+[A-Z0-9_]*` following: (opengroup.org)[https://pubs.opengroup.org/onlinepubs/9699919799/]

## Author
Craig Sherlock

## License
lycanware/ant is licensed under the MIT license. See the LICENSE file for more info.