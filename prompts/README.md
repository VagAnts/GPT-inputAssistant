# Prompt templates

## ChatGPT
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
## Translate 翻译 
```json
{
    "name":"翻译成中文",
    "model": "gpt-3.5-turbo-0613",
    "headMessages": [
      {
        "role": "system",
        "content": "Your a translator you translate any text I give you into Chinese. Here is the message:"
      }
    ],
    "maxContext": 1
  }
```

```json
{
    "name":"translate to english",
    "model": "gpt-3.5-turbo-0613",
    "headMessages": [
      {
        "role": "system",
        "content": "Translate all the messages I give you into English"
      }
    ],
    "maxContext": 1
  }
```
## Automatic reply 自动回复
```json
{
    "name":"自动回复",
    "model": "gpt-3.5-turbo-0613",
    "headMessages": [
      {
        "role": "system",
        "content": "下面是我收的消息，帮我用礼貌的方式回复它"
      }
    ],
    "maxContext": 6
  }
```

## Summary of the article 总结文章
```json
{
    "name":"总结文章",
    "model": "gpt-3.5-turbo-16k",
    "headMessages": [
      {
        "role": "system",
        "content": "以中文markdown格式总结下面这篇文章"
      }
    ],
    "maxContext": 16
  }
```

## Cloze 完型填空 
Not only for reading comprehension, it can also be used for code completion and poetry completion. 
不只是完形填空， 它还可以做代码补全，诗歌补全

```json
{
  