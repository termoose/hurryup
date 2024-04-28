# hurry UP!

## Available frontends
- `anim`: nice animation
- `json`: JSON object with key value pairs
- `promfile`: let you node exporter pick this up
- `prometheus`: http endpoint for Prometheus to scrape

## Usage
We currently only support Ookla backend, so the `-backend` argument won't do anything.
Use the `-output` argument to select the appropriate frontend.

```bash
hurryup -output [anim, json, promfile, prometheus]
```