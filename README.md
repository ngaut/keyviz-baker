# KeyvizBaker: A tool to bake heatmaps on [TiDB](https://github.com/pingcap/tidb)'s [Key Visualizer](https://docs.pingcap.com/tidb/dev/dashboard-key-visualizer)

## Build
```
go build
```

## Usage
1. Deploy a TiDB cluster. It's handy to use [TiUP](https://docs.pingcap.com/tidb/dev/tiup-overview) playground.

> **NOTE:**  To make KeyvizBaker to work well, it's recommended to use the config files [here](config). Also to make the heatmaps look nicer, it's recommended to build your own `pd-server` using this [branch](https://github.com/zanmato1984/pd/tree/keyviz-baker).

2. Run `keyviz-baker` with the following options:

Required:

```
-image_path: Path to to the image to render. Must be PNG format.
-db: Database connection string. E.g. 'root:@tcp(127.0.0.1:4000)/test'.
```

Optional:
```
-name: Name of this bake.
-skip_prepare: Skip preparation.
-align_sec: Second to align to start baking.
-interval_sec: Interval in seconds to draw each Y axis.
```

Examples:
```
./keyviz-baker -name="mandelbrot" -image_path="img/mandelbrot.png" -db="root:@tcp(127.0.0.1:4000)/test" -align_sec=7 -interval_sec=5
```

## Acknowledgement
To bake a nice heatmap requires lots of tweaks. I don't plan to describe all of these. But welcome to ask questions on this repo.
