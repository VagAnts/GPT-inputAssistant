
# GPT-inputAssistant

An application designed to allow querying OpenAI's GPT from any input field. 

## What can the app do?
* Initiate LLM calls without disconnecting from the workspace.
* Rapid LLM calling.

## Features
* Multilingual support [en, zh, jp, es, az]
* Cross-platform compatibility [win, mac]
* Ability to customize prompts

# How to Use
## Getting Started
1. Open GPT-inputAssistant click "Set API KEY" to set your OpenAI Key. 
2. Copy the desired text to your clipboard
3. Press `shift + space` to query GPT
4. Press `ESC` key to halt generation
5. Triple click `ESC` key to quickly clear context

## Defining User HotKeys 
Click "Set API KEY" and GPT-inputAssistant will open an `env.txt` file.
Simply add a new line to the file like this `GPT_HOTKEYS=space+shift`, save and close the file.
The keycode references can be found [here](https://github.com/vcaesar/keycode/blob/main/keycode.go)!

## Importing Prompts
Simply copy the json format below:
```json
{
  "name": "ChatGPT",
  "model": "gpt-3.5-turbo-0613",
  "headMessages": [
    {
      "role": "system",
      "content": "You are a helpful assistant."
    }
  ],
  "maxContext": 20
}
```

```json
{
  "name": "Translate into Chinese",
  "model": "gpt-3.5-turbo-0613",
  "headMessages": [
    {
      "role": "system",
      "content": "You're a translator translating any text I give you into Chinese. Just return the result, no need for explanations."
    }
  ],
  "maxContext": 1
}
```
Then select the application's import menu.

For more templates, go to the prompts folder in this repository.

# DEMO
The demos can be found [here](https://www.youtube.com/watch?v=2EpdfYILbgQ) and [here](https://ipfs.ee/ipfs/QmepH3EbP71zaXxaLAfQt2domXZxnb7HuaAkxT4jzhajmk/7c5ec8d0-a3d2-4d06-b649-316456390599.mp4)

# Building

Use the `build_win.bat` command for Windows or the `build.sh` command for other platforms.

# Credit

Applications utilized in the development of this app:

https://github.com/getlantern/systray

https://github.com/go-vgo/robotgo

https://github.com/robotn/gohook

https://github.com/hanyuancheung/gpt-go

## New Owners: VagAnts