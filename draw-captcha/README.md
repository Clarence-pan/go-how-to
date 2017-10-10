# FAQ
## How to get vendors' packages?
- This project use `govendor` to manage packages. Please install `govendor` firstly.
- Then run `govendor sync` in this project's root directory.
- If failed to find package "golang.org/x/..." in China, please check the network. GFW may blocks golang.org. A `http_proxy=x.x.x.x:xx https_proxy=x.x.x.x:xx` env may help.