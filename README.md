<div align="center" ><img src="https://github.com/YummyOreo/NoteDraw/blob/main/Best%20logo.png" width="150" height="150"></div>
<h1 align="center">About NoteDraw</h1>
<h3 align="center"><a href="#about">About</a> · <a href="#how-to-contribute">How to contribute</a> · <a href="#how-to-run-the-app">How to run</a>

---

## Structure

| Codebase   | Description                             |
|:---------- |:---------------------------------------:|
| [SRC](src) | The sorce code for the client side (Go) |

## Branches

Only Ones that **Wont** go away (Not temporary)

| Branch                                                | Description                |
|:----------------------------------------------------- |:--------------------------:|
| [master/main](https://github.com/Yummyoreo/NoteDraw/) | The sorce code for the app |

## About

Notdraw, is a note taking app, that you can draw in, syncs to the cloud, and is on most platforms!

## How to contribute:

- First find something to work on
- Next get it working on you pc ([How to run](#how-to-run-the-app))
- Next fork the repo and work on your fixes/changes/additions
- Make a pr with a title and description
  - To main, from your forked repo
    Title example:
    
    ```
    {User name} {pr index (number of pr you have done in the repo)}
    ```

## How to run the app

### First download the files

```console
$ git clone https://github.com/YummyOreo/NoteDraw/ && cd NoteDraw/src
```

This will download the sorce for the app

### Install Go && C

- Install [Go](https://golang.org/)
- Install [C](https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win32/Personal%20Builds/mingw-builds/installer/mingw-w64-install.exe/download) (and add to path, install 64 bit)



### Run Commands

> This is assuming that your in NoteDraw/src

```
$ go run main.go
```
