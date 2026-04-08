# Duolingo Streak API

<p align="center">
    <img src="https://duolingo-api-five.vercel.app/user/svg/APE3MP" alt="Duolingo Streak" width=480 style="border-radius:16px;">
</p>

<p align="center">
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Tech Stack">
</p>

## How to use

There are two public endpoints, all you require is a username to use.

```
https://duolingo-api-five.vercel.app/user/{username}
```
This returns the duolingo streak associated with the username and nothing else. 

```
https://duolingo-api-five.vercel.app/user/svg/{username}
```
Returns an SVG with no background of the users current streak, as shown above.  

## Example 

You can add a snippet like this into your github readme to use the SVG, just put your duolingo username in **{username}**

```html
<img src="https://duolingo-api-five.vercel.app/user/svg/{username}" alt="Duolingo Streak" width=360 style="border-radius:16px;">
```
