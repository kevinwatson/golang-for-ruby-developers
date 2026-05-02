### Chapter 1 - Go

![Gopher](images/go-gopher.png)

## Introduction

Just like any programming language, Go (aka Golang) is a language with some advantages over other languages. Because it is a compiled language, it has the benefit of high performance, because it has only 25 keywords (as of version 1.25), it is easy for engineers to pick up and begin using right away. It has out of the box support for running a HTTP server, parsing JSON and compiles and starts quickly.
Go also has some quirks. For example, there are some data types such as maps that can't be compared directly. To compare these types, it's recommended that developers add a separate package, such as `cmp`.

## Will it make my job easier?

Will it make my job easier? This is a question which needs some context to answer. If you're a developer that writes desktop or mobile apps, then probably not. If you're a developer that writes system apps (for example: JSON or XML APIs, Web, background jobs), then yes, you'll probably find Go can make your job easier. My personal experience is that while other languages and frameworks provide nearly all of the tooling you'll need to quickly get an app from nothing to minimum viable product (MVP), in the long run you may reach a point where you need to scale and your current language may not provide the performance that's needed.

## Will it make my customers happy?

Let's define a customer. In the context of this guide, we'll define a customer as anyone who uses your software to perform their job or provide a service to their customers. I currently work in the FinTech space, so my team's software processes and serves financial data to users who access it through a website or mobile device. Reducing latency between the time the user logs into their mobile app and their financial data is displayed on the screen, the better.

## References

* https://dev.to/jpoly1219/watch-out-for-these-tricky-things-in-go-5bkn
* https://github.com/google/go-cmp
* https://go.dev/ref/spec

## Wrap Up

Go, like any language, has its advantages and disadvantages. Depending on your use case (mobile vs web vs API to name a few use cases), you may find it has no place in your technology stack. You may also find that your organization will benefit from better performance and faster compile times.

[Next >>](030-chapter-02.md)
