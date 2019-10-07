# Channel and booking data generator
## Summary
    Package was developed for automate the routine of creating entities for booking and channels api of Redeam API's
### Package directories structure
    ├── booking     - scripts creating the record some type of entity of booking api
    ├── channels    - scripts creating the record some type of entity of channel api
    ├── cmd         - all scripts of running cli
    ├── config      - configuration files
    ├── data        - all prepared json files for all creation entities
    ├── logger      - logger structure
    ├── partners    - is on hold
    └── utils       - all helpers for the app

# Download

## Precompiled Binaries

You can download the precompiled release binary from [releases](https://github.com/NickTaporuk/channels_booking_clients/releases/) via web
or via

```bash
wget https://github.com/NickTaporuk/channels_booking_clients/releases/<version>/channels_booking_clients_<version>_<os>_<arch>
```

#### Go get

You can also use Go 1.12 or later to build the latest stable version from source:

```bash
GO111MODULE=on go get github.com/NickTaporuk/channels_booking_clients
```

#### Homebrew Tap

```bash
brew install nicktaporuk/tap/channels_booking_clients
# After initial install you can upgrade the version via:
brew upgrade channels_booking_clients
```
## Compilation

```bash
git clone git@github.com/NickTaporuk/channels_booking_clients.git
cd channels_booking_clients
go build .
```

## Usage

Channel booking generator CLI can be used locally.

Before run the package you have to check directory `data` which should has all needing json files for running.

Run `channels_booking_clients ` or `./channels_booking_clients` to run the package on CLI.

# The plan of the several future releases
+ create validation before running. Now we have validation only inside the running some entity
+ need to add flag --help to the package
+ need to crete initialization of the directory ./data for brew because we 
    can use two path of directory data.
    The configuration directory "data" have to exist either inside the directory "$HOME/.cbg" or inside directory where we run binary file
+ create rollback system for testing behavior of responses and after that remove all created records