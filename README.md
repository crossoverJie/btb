# btb
🕶btb:a command line tool for blog toolbox.

```shell script
$ btb -m b -mp "/opt/data/md/" -dp "/opt/data/md/images/"
DownLoading: /Users/chenjie/Downloads/movie/test1/cicada5.md 2 / 2 <-------------------------------------> 0 p/s 100.00%
DownLoading: /Users/chenjie/Downloads/movie/cicada4.md 2 / 2 <-------------------------------------------> 0 p/s 100.00%
DownLoading: /Users/chenjie/Downloads/movie/cicada1.md 3 / 3 <-------------------------------------------> 1 p/s 100.00%
DownLoading: /Users/chenjie/Downloads/movie/cicada2.md 3 / 3 <-------------------------------------------> 1 p/s 100.00%
DownLoading: /Users/chenjie/Downloads/movie/cicada3.md 3 / 3 <-------------------------------------------> 1 p/s 100.00%
DownLoading: /Users/chenjie/Downloads/movie/test1/cicada6.md 3 / 3 <-------------------------------------> 1 p/s 100.00%
Successful handling of [6] files.
```

[![asciicast](https://asciinema.org/a/ftFt2WZDx3PV4VQFzZZyw0wSl.svg)](https://asciinema.org/a/ftFt2WZDx3PV4VQFzZZyw0wSl)

# Install
```shell script
curl -fsSL https://raw.githubusercontent.com/crossoverJie/btb/master/install.sh | bash
```

## Local installation

```shell script
git clone https://github.com/crossoverJie/btb.git
cd btb
make release
```
![](https://tva1.sinaimg.cn/large/008eGmZEly1gnwu89bf8vj30ia04iq3f.jpg)

It will generate binary files for different platforms.

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

