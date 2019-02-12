# Lars-tui
* Written in GO
* Can launch bots under runtime from command input
* Auto-launches bots that is defined in the config
* Supports Linux and most linux-like based systems and Windows

A frontend for launching and keeping tabs on your discord bots (or other programs that doesnt need any input but have a lot of output).

## How-to
1. Define your bots in the config.yaml like this `bot name: path to binary/exec`. One bot per line and the config file have to be in the same folder as youre running this from or in `$HOME/.config/lars-tui` if youre on linux.
2. Launch lars-tui and be amazed (or not).
3. Continue to monitor your currently running bots or launch new by this command: `launch,namehere,path here`
## Libraries used
* [Tui-go](https://github.com/marcusolsson/tui-go/)
* [Viper](https://github.com/spf13/viper)

## Preview
![Image](https://i.imgur.com/iriM1hK.png)
