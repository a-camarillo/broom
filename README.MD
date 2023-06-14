# broom

## Description

An interactive CLI tool for cleaning up your local git branches.

I built this project to solve my issue of git branches piling up because I forget to delete them. I also used this project as a way to learn Go.

## Demo

![](https://raw.githubusercontent.com/a-camarillo/broom/main/images/broomDemo2.gif)

## Installation

Currently the best way to install broom is by downloading the binary for your respective system from the latest [release page](https://github.com/a-camarillo/broom/releases/latest).

### Linux and MacOS

Copy the url to the binary for your system, I am on a Linux x86_64 machine.

`curl -LO https://github.com/a-camarillo/broom/releases/download/latest/broom_Linux_x86_64.tar.gz`

Unpack the tar file:

`sudo tar -xvf broom_Linux_x86_64.tar.gz`

To use broom globally:

`sudo mv broom /bin/`

### Windows

Download the appropriate zip from the release page and unpack the executable with your preferred zip file opener.

**Note:** You will need to use git bash on Windows to use this program. Additionally, broom has not been tested on Windows so apologies if the program does not run as intended.

## Usage

Once broom is installed, just navigate to a directory where a git repository has been initialized and run:

`broom`

## Contributing

Feel free to report bugs, suggest improvements, and send pull requests! Please add documentation and tests for new features.

If you would like to submit code, see [CONTRIBUTING.md](https://github.com/a-camarillo/broom/blob/main/CONTRIBUTING.md)

## License

MIT, see [LICENSE](https://github.com/a-camarillo/broom/blob/main/LICENSE)
