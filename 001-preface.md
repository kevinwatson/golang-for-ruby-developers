## Preface

Welcome to *Go for Ruby Developers.*

Ruby on Rails is a framework built on the Ruby programming language. The Rails framework provides you with the tools you need to easily build a database-backed application. Rails has achieved widespread popularity - many popular websites run on Rails, including Shopify, Basecamp, and GitHub.

Go is a programming language that has its origins at Google. It is a compiled, garbage-collected, concurrent, statically typed language. It simplifies app management by compiling a project into a single binary. It has a standard library that provides a number of features, including relational database interfaces (`database/sql`), reading and writing data formats (for example JSON with the `encoding/json` package), parsing and creating HTML (`html`), logging (`log`), operating system commands (`os/exec`), sending and listening for HTTP and HTTPS requests (`net/http`) and automated testing (`testing`).

Go differs from Ruby in a few ways. First, Go is a compiled and statically typed language, whereas Ruby is an interpreted language. This one difference usually means that the application is faster because the code is optimized early and has less runtime overhead. Second, Go has built in concurrency using goroutines (managed by the Go runtime and are faster and cheaper to create than operating system threads) and channels (also managed by the Go runtime for synchronization and data transfer between goroutines without explicit lock contention), while Ruby uses threads (which are managed by the operating system and have high memory usage, not suitable for high-volume concurrency) and fibers (similar in performance to goroutines). Last, the design of the Go language optimizes for low latency, high throughput apps, while Ruby and the Ruby on Rails framework is optimized for rapid development.

[Next >>](002-who-is-this-book-for.md)
