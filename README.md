![!Promtail hook for Logrus logger](./logo.png)

# Lokigrus - Promtail hook for Logrus logger ![Continious Integration (Lint & Unit Test)](https://github.com/ic2hrmk/lokigrus/workflows/Continious%20Integration%20(Lint%20&%20Unit%20Test)/badge.svg)

An adapter of [Promtail client](github.com/ic2hrmk/promtail) for [Grafana's Loki logging server](https://grafana.com/oss/loki/) 
server to be used with [Logrus logger](https://github.com/sirupsen/logrus) (written with love and tests).

#### Take a look at my [Promtail client](github.com/ic2hrmk/promtail)

## Description

It's a log hook that handles log messages and translates it to Loki's messages via
Prmtail client.

#### Current implementation contains:

 - [X] Logrus hook for centralized logging with Grafana Loki server
 - [X] Loki's server health check
 - [ ] Report about log exchange failures to logrus logger
 
 ## How to use
 
The easiest way is:
~~~go
package mypackage

import (
    "github.com/ic2hrmk/lokigrus"
    "github.com/sirupsen/logrus"
)

func InitPromtailSupport(
    logger *logrus.Logger,
    lokiAddress string, 
    appLabels map[string]string,
) error {
    promtailHook, err := lokigrus.NewPromtailHook(lokiAddress, appLabels)
    if err != nil {
       return err
    }

    logger.AddHook(promtailHook)

    return nil
}
~~~
 
### Issues / Contributing
Feel free to post a Github Issue, I will respond ASAP
 