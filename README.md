# btb
🕶btb:a command line tool for blog toolbox.

[![asciicast](https://asciinema.org/a/ftFt2WZDx3PV4VQFzZZyw0wSl.svg)](https://asciinema.org/a/ftFt2WZDx3PV4VQFzZZyw0wSl)

# Install
```shell script
curl -fsSL https://raw.githubusercontent.com/crossoverJie/btb/master/install.sh | bash
```

# Usage

## MacOS & Linux

```shell script
btb -m b -mp "/opt/data/md/" -dp "/opt/data/images/"
```

## Windows

```shell script
./btb.exe -m b -mp "/opt/data/md/" -dp "/opt/data/images/"
```

# Help
```shell script
btb -h
NAME:
   btb - Help you backup and replace your blog's images

USAGE:
   btb [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --model value, -m value            operating mode; r:replace, b:backup (default: b)
   --download-path value, --dp value  The path where the image is stored (default: "/opt/data")
   --markdown-path value, --mp value  The path where the markdown file is stored (default: "/opt/data")
   --help, -h
```

