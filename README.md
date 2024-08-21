# Hello World Kata: Ping Pong

This is the supporting material for my talk "Production Ready Hello World" presented at Gophercon UK 2024.

Slides can be found [here](https://docs.google.com/presentation/d/1zkAvvb1jP2qhlOor4ka_W2EEv-f7rX_kgIWr-v1maC4/pub?start=false&loop=false&delayms=3000)

## Problem Statement

You are in charge of developing an application called "pingpong". This application is a web server that listens to requests and if it receives a GET request on the "/ping" path it should respond with the message "pong". 

The code should be production ready: include all the things that you consider important in production code, but you are only allowed to use packages from the standard library.

## Requirements

- Use package http
- Use package httptest
- Use package json
