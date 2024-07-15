# TGMessageStore
[![telegram badge](https://img.shields.io/badge/Telegram-Channel-30302f?style=flat&logo=telegram)](https://telegram.dog/FractalProjects)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jisin0/TGMessageStore)](https://goreportcard.com/report/github.com/Jisin0/TGMessageStore)
[![Go Build](https://github.com/Jisin0/TGMessageStore/workflows/build/badge.svg)](https://github.com/Jisin0/TGMessageStore/actions?query=workflow%3Abuild+event%3Apush+branch%3Amain)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[**TGMessageStore** ](https://TGMessageStore.vercel.app) is a serverless telegram bot to generate permanent shareable links to a single or batch of messages.

## Commands
```
start - Check if the bot is alive.
batch - Create a new message batch.
genlink - Create a link for a single message.
about - Basic Information About the bot.
help - Short Guide on How to Use the Bot.
privacy - Read about the privacy policy.
```

## Variables

- `BOT_TOKEN`  : Optional. On vercel, a list of bot tokens allowed to connect to the app or leave empty allow anyone to connect. On servers, a single bot token.

## Deploy
Deploy your own **TGMessageStore** app to vercel

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/project?template=https://github.com/Jisin0/TGMessageStore/tree/main&env=BOT_TOKEN&envDescription=List%20of%20of%20allowed%20bot%20tokens%20or%20leave%20empty%20to%20allow%20all)

<details><summary>Deploy To Heroku</summary>
<p>
<br>
<a href="https://heroku.com/deploy?template=https://github.com/Jisin0/TGMessageStore/tree/main">
  <img src="https://www.herokucdn.com/deploy/button.svg" alt="Deploy">
</a>
</p>
</details>

<details><summary>Deploy To Scalingo</summary>
<p>
<br>
<a href="https://dashboard.scalingo.com/create/app?source=https://github.com/Jisin0/TGMessageStore#main">
   <img src="https://cdn.scalingo.com/deploy/button.svg" alt="Deploy on Scalingo" data-canonical-src="https://cdn.scalingo.com/deploy/button.svg" style="max-width:100%;">
</a>
</p>
</details>


<details><summary>Deploy To Render</summary>
<p>
<br>
<a href="https://dashboard.render.com/select-repo?type=web">
  <img src="https://render.com/images/deploy-to-render-button.svg" alt="deploy-to-render">
</a>
</p>
<p>
Make sure to have the following options set :

<b>Environment</b>
<pre>Go</pre>

<b>Build Command</b>
<pre>go build .</pre>

<b>Start Command</b>
<pre>./TGMessageStore</pre>

<b>Advanced >> Health Check Path</b>
<pre>/</pre>
</p>
</details>


<details><summary>Deploy To Koyeb</summary>
<p>
<br>
<a href="https://app.koyeb.com/deploy?type=git&repository=github.com/Jisin0/TGMessageStore&branch=main">
  <img src="https://www.koyeb.com/static/images/deploy/button.svg" alt="deploy-to-koyeb">
</a>
</p>
<p>
You must set the Run command to :
<pre>./bin/TGMessageStore</pre>
</p>
</details>

<details><summary>Deploy To Okteto</summary>
<p>
<br>
<a href="https://cloud.okteto.com/deploy?repository=https://github.com/Jisin0/TGMessageStore">
  <img src="https://okteto.com/develop-okteto.svg" alt="deploy-to-okteto">
</a>
</p>
</details>

<details><summary>Deploy To Railway</summary>
<p>
<br>
<a href="https://railway.app/new/template?template=https%3A%2F%2Fgithub.com%2FJisin0%2FTGMessageStore">
  <img src="https://railway.app/button.svg" alt="deploy-to-railway">
</a>
</p>
</details>

<details><summary>Run Locally/VPS</summary>
<p>
You must have the latest version of <a href="https://go.dev/dl">GO</a> installed first
<pre>
git clone https://github.com/Jisin0/TGMessageStore
cd TGMessageStore
go build .
./TGMessageStore
</pre>
</p>
</details>

## Thanks

 - Thanks to Paul for his awesome [Library](https://github.com/PaulSonOfLars/gotgbot)

## Disclaimer

[![GNU General Public License 3.0](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html#header)    
Licensed under [GNU GPL 3.0.](https://github.com/Jisin0/TGMessageStore/blob/main/LICENSE).