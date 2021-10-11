Welcome to the WoW TBC Classic simulator! 

The primary goal of this project is to provide a framework that makes it easy to build a DPS sim for any class/spec, with a polished UI and accurate results. Each community will have ownership / responsibility over their portion of the sim, to ensure accuracy and that their community is represented. When enough classes/specs are implemented, we also hope to build a "raid sim".

Live sims:
[https://wowsims.github.io/tbc/elemental_shaman/](https://wowsims.github.io/tbc/elemental_shaman/ "https://wowsims.github.io/tbc/elemental_shaman/")

# Installation
To develop on the sim, Go, protobuf-compiler, and the corresponding Go plugins must be installed:

```sh
sudo apt install golang-go
sudo apt install protobuf-compiler
go get -u -v github.com/golang/protobuf/proto
go get -u -v github.com/golang/protobuf/protoc-gen-go
```

To work on the UI, node >= 14.0 must additionally be installed:
```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
nvm install 14.17.6
```

And finally, install the node dependencies for this repo:
```sh
cd tbc
npm install
```

# Adding a Sim
So you want to make a new sim for your class/spec! The basic steps are as follows:
 - Create the proto to describe the interface between sim and UI.
 - Add items your spec uses to the Items Generator.
 - Implement the sim.
 - Implement the UI.


# Create the proto interface between Sim and UI
TODO

# Add items to the Items Generator
TODO

# Implement the Sim
TODO

# Implement the UI
TODO
