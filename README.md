# Log

Logging interface and decorators to support Structured logging in Libraries. It currently support those loggers:

  * [Go logger](https://godoc.org/log)
  * [Logrus](https://github.com/Sirupsen/logrus)
  
## Installation

`go get "github.com/normegil/log"`

## Features

  * Agnostic interfaces to support all kinds of loggers
  * Level of logging
    * DEBUG
    * INFO
    * PANIC
  * Add support for structured logging to libraries that doesn't support it natively