# Command Prefix

Prefix Stdout and Stderr with the name of the command.

```
> cmdpfx test/some-script.sh
test/some-script.sh | hey from some-script.sh 1
test/some-script.sh | hey from some-script.sh 2
```

## Install

```sh
go install github.com/nstogner/cmdpfx
```

## Usage

```sh
cmdpfx ./some-script.sh
```
