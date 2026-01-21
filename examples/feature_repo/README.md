# Quickstart

## Commands

```bash
$ uv venv
$ source .venv/bin/activate
(.venv) $ uv pip install feast
(.venv) $ feast init examples
(.venv) $ cd examples/feature_repo
```

### Run a workflow

```bash
(.venv) $ python test_workflow.py
```

### Workflow Result

```bash
Created project examples
Created entity driver
Created feature view driver_hourly_stats
Created feature view driver_hourly_stats_fresh
Created on demand feature view transformed_conv_rate_fresh
Created on demand feature view transformed_conv_rate
Created feature service driver_activity_v3
Created feature service driver_activity_v2
Created feature service driver_activity_v1

Created sqlite table examples_driver_hourly_stats_fresh
Created sqlite table examples_driver_hourly_stats


--- Historical features for training ---
   driver_id           event_timestamp  ...  conv_rate_plus_val1  conv_rate_plus_val2
0       1001 2021-04-12 10:59:42+00:00  ...             1.446066            10.446066
1       1002 2021-04-12 08:12:10+00:00  ...             2.122385            20.122385
2       1003 2021-04-12 16:40:26+00:00  ...             3.310424            30.310424

[3 rows x 10 columns]

--- Historical features for batch scoring ---
   driver_id                  event_timestamp  ...  conv_rate_plus_val1  conv_rate_plus_val2
0       1001 2026-01-21 13:17:15.139744+00:00  ...             1.721403            10.721403
1       1002 2026-01-21 13:17:15.139744+00:00  ...             2.289374            20.289374
2       1003 2026-01-21 13:17:15.139744+00:00  ...             3.534667            30.534667

[3 rows x 10 columns]

--- Load features into online store ---
Materializing 2 feature views to 2026-01-21 22:17:15+00:00 into the sqlite online store.

driver_hourly_stats from 2026-01-20 13:17:15+00:00 to 2026-01-21 22:17:15+00:00:
driver_hourly_stats_fresh from 2026-01-20 13:17:15+00:00 to 2026-01-21 22:17:15+00:00:

--- Online features ---
acc_rate  :  [0.6569926142692566, 0.17377954721450806]
conv_rate_plus_val1  :  [1000.5848700404167, 1001.8442494273186]
conv_rate_plus_val2  :  [2000.5848700404167, 2002.8442494273186]
driver_id  :  [1001, 1002]

--- Online features retrieved (instead) through a feature service---
conv_rate  :  [0.5848700404167175, 0.844249427318573]
conv_rate_plus_val1  :  [1000.5848700404167, 1001.8442494273186]
conv_rate_plus_val2  :  [2000.5848700404167, 2002.8442494273186]
driver_id  :  [1001, 1002]

--- Online features retrieved (using feature service v3, which uses a feature view with a push source---
acc_rate  :  [0.6569926142692566, 0.17377954721450806]
avg_daily_trips  :  [864, 621]
conv_rate  :  [0.5848700404167175, 0.844249427318573]
conv_rate_plus_val1  :  [1000.5848700404167, 1001.8442494273186]
conv_rate_plus_val2  :  [2000.5848700404167, 2002.8442494273186]
driver_id  :  [1001, 1002]

--- Simulate a stream event ingestion of the hourly stats df ---
   driver_id            event_timestamp                    created  conv_rate  acc_rate  avg_daily_trips
0       1001 2026-01-21 22:17:15.290736 2026-01-21 22:17:15.290739        1.0       1.0             1000

--- Online features again with updated values from a stream push---
acc_rate  :  [1.0, 0.17377954721450806]
avg_daily_trips  :  [1000, 621]
conv_rate  :  [1.0, 0.844249427318573]
conv_rate_plus_val1  :  [1001.0, 1001.8442494273186]
conv_rate_plus_val2  :  [2001.0, 2002.8442494273186]
driver_id  :  [1001, 1002]

--- Run feast teardown ---
```

