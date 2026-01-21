# Go client (Feast feature server)

This client calls the Feast Python feature server using the sample data from `examples/feature_repo`.

## Run

Start the server from the example repo:

```bash
cd examples/feature_repo
feast apply
feast materialize-incremental $(date -u +"%Y-%m-%dT%H:%M:%S")
feast serve
```

Then, in another shell:

```bash
cd client
go run . -addr http://localhost:6566
```

To use a feature service instead of explicit features:

```bash
go run . -feature-service driver_activity_v1
```

More flag examples:

```bash
go run . -features driver_hourly_stats:acc_rate,transformed_conv_rate:conv_rate_plus_val1
go run . -addr http://127.0.0.1:6566 -timeout 2s
go run . -feature-service driver_activity_v3 -timeout 10s
```
