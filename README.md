# Benchmark for TiDB HashAgg

For [pull/19807](https://github.com/pingcap/tidb/pull/19807), `set.SyncSet` can be implemented using `sync.Map` or `map[interface{}]struct{}` with `sync.RWMutex`.
This repository benchmarked both implementation with self-generated input data.

Data generation resides in [`data_gen.go`](https://github.com/dyzsr/hashaggbench/blob/master/data_gen.go),
and benchmark is in [`hashagg_test.go`](https://github.com/dyzsr/hashaggbench/blob/master/hashagg_test.go)

## Usage

1. Launch TiDB
2. Generate input data (Load data into TiDB)
3. Run benchmark

Generate input

```
go run . -h 127.0.0.1 -port 4000 -u root -db test
```

Run benchmark

```
go test -bench=. -count=5 -timeout=0 -u root -db test
```

## Benchmark

For [pull/19807](https://github.com/pingcap/tidb/pull/19807), I ran the benchmark in following steps for both implementations:

1. launch TiDB with TiUP playground

```
tiup playground --db.binpath /path/to/tidb-server
```

2. generate input data

```
go run .
```

3. run benchmark

```
gotest -bench=. -count=5 -timeout=0 > rwmutex.bench    (for map with sync.RWMutex)
```

```
gotest -bench=. -count=5 -timeout=0 > syncmap.bench    (for sync.Map)
```

4. print the results

```
$  benchstat -alpha 0.2 syncmap.bench rwmutex.bench
name                                                                                         old time/op  new time/op  delta
Distinct/table:dense,func:avg,group:1000-12                                                   2.07s ± 6%   1.94s ± 7%   -6.21%  (p=0.095 n=5+5)
Distinct/table:sparse,func:avg,group:1000-12                                                  5.53s ±17%   3.96s ±15%  -28.41%  (p=0.008 n=5+5)
Distinct/table:dense,func:count,group:1000-12                                                 1.75s ± 6%   1.56s ± 8%  -10.60%  (p=0.032 n=5+5)
Distinct/table:sparse,func:count,group:1000-12                                                3.77s ± 7%   2.62s ± 7%  -30.53%  (p=0.008 n=5+5)
Distinct/table:dense,func:group_concat,group:1000-12                                          2.26s ±10%   2.17s ± 9%     ~     (p=0.421 n=5+5)
Distinct/table:sparse,func:group_concat,group:1000-12                                         6.74s ±36%   5.13s ±27%  -23.85%  (p=0.151 n=5+5)
Distinct/table:dense,func:sum,group:1000-12                                                   4.76s ±34%   2.82s ±57%  -40.72%  (p=0.056 n=5+5)
Distinct/table:sparse,func:sum,group:1000-12                                                  8.34s ±51%   6.83s ±68%     ~     (p=0.548 n=5+5)
Mix/table:dense,funcs:[avg_count_group_concat_sum],group:1000-12                              4.48s ±73%   4.44s ±67%     ~     (p=1.000 n=5+5)
Mix/table:sparse,funcs:[avg_count_group_concat_sum],group:1000-12                             5.66s ±90%   4.53s ±87%     ~     (p=0.310 n=5+5)
Mix/table:dense,funcs:[avg_distinct_count_group_concat_sum],group:1000-12                     7.62s ±57%   5.47s ±51%  -28.22%  (p=0.151 n=5+5)
Mix/table:sparse,funcs:[avg_distinct_count_group_concat_sum],group:1000-12                    11.6s ±63%    8.6s ±33%     ~     (p=0.222 n=5+5)
Mix/table:dense,funcs:[avg_count_distinct_group_concat_sum],group:1000-12                     6.03s ±58%   6.47s ±56%     ~     (p=0.548 n=5+5)
Mix/table:sparse,funcs:[avg_count_distinct_group_concat_sum],group:1000-12                    12.5s ±50%    8.8s ±71%     ~     (p=0.222 n=5+5)
Mix/table:dense,funcs:[avg_count_group_concat_distinct_sum],group:1000-12                     8.23s ±52%   5.61s ±13%     ~     (p=0.222 n=5+5)
Mix/table:sparse,funcs:[avg_count_group_concat_distinct_sum],group:1000-12                    18.1s ±42%   12.6s ±48%  -30.24%  (p=0.151 n=5+5)
Mix/table:dense,funcs:[avg_count_group_concat_sum_distinct],group:1000-12                     6.22s ±20%   5.82s ± 7%     ~     (p=0.421 n=5+5)
Mix/table:sparse,funcs:[avg_count_group_concat_sum_distinct],group:1000-12                    21.1s ±17%   14.9s ±38%  -29.68%  (p=0.032 n=5+5)
Mix/table:dense,funcs:[avg_distinct_count_group_concat_sum_distinct],group:1000-12            14.2s ±36%    7.7s ±17%  -45.48%  (p=0.008 n=5+5)
Mix/table:sparse,funcs:[avg_distinct_count_group_concat_sum_distinct],group:1000-12           30.0s ±41%   25.2s ±53%     ~     (p=0.548 n=5+5)
Mix/table:dense,funcs:[avg_count_distinct_group_concat_sum_distinct],group:1000-12            13.1s ±17%    8.6s ±45%  -34.49%  (p=0.032 n=5+5)
Mix/table:sparse,funcs:[avg_count_distinct_group_concat_sum_distinct],group:1000-12           56.8s ±95%   18.9s ±63%  -66.70%  (p=0.008 n=5+5)
Mix/table:dense,funcs:[avg_count_group_concat_distinct_sum_distinct],group:1000-12            34.1s ±20%   15.5s ±14%  -54.46%  (p=0.008 n=5+5)
Mix/table:sparse,funcs:[avg_count_group_concat_distinct_sum_distinct],group:1000-12            152s ±48%    41s ±119%  -72.73%  (p=0.016 n=5+5)
Mix/table:dense,funcs:[avg_distinct_count_distinct_group_concat_sum_distinct],group:1000-12   61.9s ±83%   21.7s ±30%  -64.88%  (p=0.008 n=5+5)
```

## Testdata

Tables

* dense: 10000000 rows, NDV of `a` = 1000 (1000 groups), `b` is generated using `rand.Int31()`

``` mysql
mysql> desc dense;
+-------+---------+------+------+---------+-------+
| Field | Type    | Null | Key  | Default | Extra |
+-------+---------+------+------+---------+-------+
| a     | int(11) | YES  |      | NULL    |       |
| b     | int(11) | YES  |      | NULL    |       |
+-------+---------+------+------+---------+-------+
2 rows in set (0.00 sec)
```

* sparse: 10000000 rows, NDV of `a` = 1000 (1000 groups), NDV of `b` = 32

``` mysql
mysql> desc sparse;
+-------+---------+------+------+---------+-------+
| Field | Type    | Null | Key  | Default | Extra |
+-------+---------+------+------+---------+-------+
| a     | int(11) | YES  |      | NULL    |       |
| b     | int(11) | YES  |      | NULL    |       |
+-------+---------+------+------+---------+-------+
2 rows in set (0.00 sec)
```

Queries

``` mysql
SELECT /*+ HASH_AGG() */ avg(distinct b) FROM `sparse` GROUP BY `a`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `dense` GROUP BY `a`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `sparse` GROUP BY `a`

SELECT /*+ HASH_AGG() */ group_concat(distinct b) FROM `dense` GROUP BY `a`

SELECT /*+ HASH_AGG() */ group_concat(distinct b) FROM `sparse` GROUP BY `a`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `dense` GROUP BY `a`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `sparse` GROUP BY `a`

SELECT avg(b),count(b),group_concat(b),sum(b) FROM `dense` GROUP BY `a`

SELECT avg(b),count(b),group_concat(b),sum(b) FROM `sparse` GROUP BY `a`

SELECT avg(distinct b),count(b),group_concat(b),sum(b) FROM `dense` GROUP BY `a`

SELECT avg(distinct b),count(b),group_concat(b),sum(b) FROM `sparse` GROUP BY `a`

SELECT avg(b),count(distinct b),group_concat(b),sum(b) FROM `dense` GROUP BY `a`

SELECT avg(b),count(distinct b),group_concat(b),sum(b) FROM `sparse` GROUP BY `a`

SELECT avg(b),count(b),group_concat(distinct b),sum(b) FROM `dense` GROUP BY `a`

SELECT avg(b),count(b),group_concat(distinct b),sum(b) FROM `sparse` GROUP BY `a`

SELECT avg(b),count(b),group_concat(b),sum(distinct b) FROM `dense` GROUP BY `a`

SELECT avg(b),count(b),group_concat(b),sum(distinct b) FROM `sparse` GROUP BY `a`

SELECT avg(distinct b),count(b),group_concat(b),sum(distinct b) FROM `dense` GROUP BY `a`

SELECT avg(distinct b),count(b),group_concat(b),sum(distinct b) FROM `sparse` GROUP BY `a`

SELECT avg(b),count(distinct b),group_concat(b),sum(distinct b) FROM `dense` GROUP BY `a`

SELECT avg(b),count(distinct b),group_concat(b),sum(distinct b) FROM `sparse` GROUP BY `a`

SELECT avg(b),count(b),group_concat(distinct b),sum(distinct b) FROM `dense` GROUP BY `a`

SELECT avg(b),count(b),group_concat(distinct b),sum(distinct b) FROM `sparse` GROUP BY `a`

SELECT avg(distinct b),count(distinct b),group_concat(b),sum(distinct b) FROM `dense` GROUP BY `a`

SELECT avg(distinct b),count(distinct b),group_concat(b),sum(distinct b) FROM `sparse` GROUP BY `a`

SELECT avg(distinct b),count(distinct b),group_concat(distinct b),sum(distinct b) FROM `dense` GROUP BY `a`

SELECT avg(distinct b),count(distinct b),group_concat(distinct b),sum(distinct b) FROM `sparse` GROUP BY `a`
```
