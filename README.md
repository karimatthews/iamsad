# I am Sad

A simple command line tool that shows you pictures of dogs.

## Dependencies (Important!)

This tool is depenedent on [imgcat](https://github.com/martinlindhe/imgcat), which allows you to cat images in [iTerm2](https://www.iterm2.com/).
This means the tool needs to be used with iTerm2.
Note that Imgcat seems to have issues working with tmux.

## Usage

Clone the repo and then run:

```sh
  cd cmd/iamsad
  # Build the binary
  go build
  #  Get usage options
  ./iamssad -h
```

1. Download the binary `bin/iamsad`.
1. Give the binary executable permissions: `chmod +x ./iamsad`
1. Run `.iamsad -h` to see possible commands

## To Do

- [ ] Organise project stucture better
- [ ] Extend cat features
- [ ] Add tests

## Features

- List available breeds
- Given a particular breed list the sub breeds
- Show a random photo of a dog
- Show a random photo of a particular breed/sub-breed

## Sources

- [Dog API](https://dog.ceo/dog-api/documentation/)
- [Cat API](https://docs.thecatapi.com/)
