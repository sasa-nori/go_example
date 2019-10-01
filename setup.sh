#!/bin/bash
cd $HOME
mkdir -p development/go/src/github.com/$USER/
brew install --HEAD goenv
echo "# golang" >>  ~/.bash_profile
echo export PATH='$HOME/.goenv/bin:$PATH' >> ~/.bash_profile
echo eval "$(goenv init -)" >> ~/.bash_profile
source ~/.bash_profile
goenv install 1.13.0
goenv global 1.13.0
goenv rehash
echo export GOPATH='$HOME/development/go' >> ~/.bash_profile
echo export PATH='$PATH:$GOPATH/bin' >> ~/.bash_profile
echo export PATH='$PATH:/usr/local/go/bin' >> ~/.bash_profile
source ~/.bash_profile
cd $GOPATH/src/github.com/$USER/
git clone git@github.com:noriyuki-sasagawa/go_example.git
brew install dep
brew upgrade dep
dep ensure
go run $GOPATH/src/github.com/$USER/go_example/main.go
