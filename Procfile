worker: mkdir -p ~/git/GitHub/canha
worker: cd ~/git/GitHub/canha
worker: git clone https://github.com/canha/golang-tools-install-script 
worker: cd golang-tools-install-script/
worker: bash goinstall.sh
worker: cat ~/.bashrc
worker: source ~/.bashrc
worker: cd ~/
worker: go run main.go
web: go run main.go