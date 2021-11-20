Welcome to the WoW TBC Classic simulator! If you have questions or are thinking about contributing, [join our discord](https://discord.gg/jJMPr9JWwx "https://discord.gg/jJMPr9JWwx") to chat!

The primary goal of this project is to provide a framework that makes it easy to build a DPS sim for any class/spec, with a polished UI and accurate results. Each community will have ownership / responsibility over their portion of the sim, to ensure accuracy and that their community is represented. When enough classes/specs are implemented, we also hope to build a "raid sim".

Live sims:
 - [Balance Druid](https://wowsims.github.io/tbc/balance_druid/ "https://wowsims.github.io/tbc/balance_druid/")
 - [Elemental Shaman](https://wowsims.github.io/tbc/elemental_shaman/ "https://wowsims.github.io/tbc/elemental_shaman/")

# Installation
This project has dependencies on Go >=1.16, protobuf-compiler and the corresponding Go plugins, and node >= 14.0.

## Ubuntu
```sh
# Standard Go installation script
curl -O https://dl.google.com/go/go1.16.10.linux-amd64.tar.gz
sudo rm -rf /usr/local/go 
sudo tar -C /usr/local -xzf go1.16.10.linux-amd64.tar.gz
echo `export PATH=$PATH:/usr/local/go/bin` >> $HOME/.bashrc
echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> $HOME/.bashrc
source $HOME/.bashrc

# Install protobuf compiler and Go plugins
sudo apt update && sudo apt upgrade
sudo apt install protobuf-compiler
go get -u -v github.com/golang/protobuf/proto
go get -u -v github.com/golang/protobuf/protoc-gen-go

# Install node
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
nvm install 14.17.6

# Install the npm package dependencies using node
cd tbc
npm install
```

## Windows
If you want to develop on Windows, we recommend setting up Docker using [this guide](https://docs.docker.com/desktop/windows/wsl/ "https://docs.docker.com/desktop/windows/wsl/") and then following the Docker instructions below.

## Docker
Alternatively, install Docker and your workflow will look something like this:
```sh
git clone https://github.com/wowsims/tbc.git
cd tbc

# Build the docker image and install npm dependencies (only need to run these once).
docker build --tag wowsims-tbc .
docker run --rm -v $(pwd):/tbc wowsims-tbc npm install

# Now you can run the commands as shown in the Commands sections, preceding everything with, "docker run --rm -it -p 8080:8080 -v $(pwd):/tbc wowsims-tbc".
# For convenience, set this as an environment variable:
TBC_CMD="docker run --rm -it -p 8080:8080 -v $(pwd):/tbc wowsims-tbc"

# ... update the items ...

# Update items
$TBC_CMD make items

# ... do some coding on the sim ...

# Run tests
$TBC_CMD make test

# ... do some coding on the UI ...

# Host a local site
$TBC_CMD make host
```

# Commands
We use a makefile for our build system. These commands will usually be all you need while developing for this project:
```sh
# Generate code for items. Only necessary if you changed the items generator.
make items

# Run all the tests. Currently only the backend sim has tests.
make test

# Host a local version of the UI at http://localhost:8080. Visit it by pointing a browser to
# http://localhost:8080/tbc/YOUR_SPEC_HERE, where YOUR_SPEC_HERE is the directory under ui/ with your custom code.
make host

# Delete all generated files (.pb.go and .ts proto files, and dist/)
make clean
```

# Adding a Sim
So you want to make a new sim for your class/spec! The basic steps are as follows:
 - [Create the proto interface between sim and UI.](#create-the-proto-interface-between-sim-and-ui)
 - [Add items your spec uses to the Items Generator.](#add-items-to-the-items-generator)
 - [Implement the sim.](#implement-the-sim)
 - [Implement the UI.](#implement-the-ui)


## Create the proto interface between Sim and UI
This project uses [Google Protocol Buffers](https://developers.google.com/protocol-buffers/docs/gotutorial "https://developers.google.com/protocol-buffers/docs/gotutorial") to pass data between the sim and the UI. TLDR; Describe data structures in .proto files, and the tool can generate code in any programming language. It lets us avoid repeating the same code in our Go and Typescript worlds without losing type safety.

For a new sim, make the following changes:
  - Add a new value to the `Spec` enum in proto/common.proto. __NOTE: The name you give to this enum value is not just a name, it is used in our templating system. This guide will refer to this name as `$SPEC` elsewhere.__
  - Add a 'proto/YOUR_CLASS.proto' file if it doesn't already exist and add data messages containing all the class/spec-specific information needed to run your sim. In general, there will be 3 pieces of information you need:
    - Talents
    - Rotation (the order in which your sim will use spells/abilities)
    - Options (additional choices your sim needs to make)
  - Update the `PlayerOptions.spec` field in `proto/api.proto` to include your shiny new message as an option.

That's it! Now when you run `make` there will be generated .go and .ts code in `sim/core/proto` and `ui/core/proto` respectively. If you aren't familiar with protos, take a quick look at them to see what's happening.

## Add items to the Items Generator
`generate_items/item_declarations.go` contains a list of all items known to the sim, as well as a category label for each item. Add the items needed by your sim and run `make items` when you're done.

## Implement the Sim
This step is where most of the magic happens. A few highlights to start understanding the sim code:
  - `sim/wasm/main.go` This file is the actual main function, for the [.wasm binary](https://webassembly.org/ "https://webassembly.org/") used by the UI. You shouldn't ever need to touch this, but just know its here.
  - `sim/core/api.go` This is where the action starts. This file implements the request/response messages defined in `proto/api.proto`.
  - `sim/core/sim.go` Orchestrates everything. Main event loop is in `Simulation.RunOnce`.
  - `sim/core/agent.go` An Agent can be thought of as the 'Player', i.e. the person controlling the game. This is the interface you'll be implementing.
  - `sim/core/character.go` A Character holds all the stats/cooldowns/gear/etc common to any WoW character. Each Agent has a Character that it controls.

Read through the core code and some examples from other classes/specs to get a feel for what's needed. Hopefully `sim/core` already includes what you need, but most classes have at least 1 unique mechanic so you may need to touch `core` as well.

Don't forget to write unit tests! Again, look at existing tests for examples. Run them with `make test` when you're ready.

## Implement the UI
If you've made it this far, you're almost there! The UI is very generalized and it doesn't take much work to build an entire sim UI using our templating system. To use it:
  - Create a directory 'ui/$SPEC'. So if your Spec enum value was named, 'elemental_shaman', create a directory, 'ui/elemental_shaman'.
  - This directory must contain a file, 'index.ts'.
  - This directory must contain a file, 'index.scss'.

No .html is needed, it will be generated based on `ui/index_template.html` and the `$SPEC` name.

Steps for building a new UI:
  - Modify `ui/core/proto_utils/utils.ts` to include boilerplate for your `$SPEC` name if it isn't already there.
  - Configure the UI by writing `ui/$SPEC/index.ts` and `ui/$SPEC/index.scss`. Start by copying from another spec's code, and change the configuration as needed.
  - Finally, add a rule to the `makefile` for the new sim site. Just copy from the other site rules already there and change the `$SPEC` names.

When you're ready to try out the site, run `make host` and navigate to `http://localhost:8080/tbc/$SPEC`.

# Deployment
Thanks to the workflow defined in `.github/workflows/deploy.yml`, pushes to `master` automatically build and deploy a new site so there's nothing to do here. Sit back and appreciate your new sim!
