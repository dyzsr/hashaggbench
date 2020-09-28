# Benchmark for TiDB HashAgg

For [pull/19807](https://github.com/pingcap/tidb/pull/19807), `set.SyncSet` can be implemented using `sync.Map` or `map[interface{}]struct{}` with `sync.RWMutex`.
This repository benchmarked both implementations with self-generated input data.

Data generation resides in [`data_gen.go`](https://github.com/dyzsr/hashaggbench/blob/master/data_gen.go),
and benchmark is in [`hashagg_test.go`](https://github.com/dyzsr/hashaggbench/blob/master/hashagg_test.go)

## Usage

1. Launch TiDB
2. Generate input data (Load data into TiDB)
3. Run benchmark

Generate input

```
go run . -host 127.0.0.1 -port 4000 -user root -db test
```

Run benchmark

```
go test -bench=. -count=5 -benchtime=5x -timeout=0
```

## Benchmark

For [pull/19807](https://github.com/pingcap/tidb/pull/19807), I ran the benchmark in following steps for both implementations:

1. launch TiDB

```
make
bin/tidb-server
```

2. generate input data

```
go run .
```

3. run benchmark

```
gotest -bench=. -count=5 -benchtime=5x -timeout=0 > rwmutex.result    (for map with sync.RWMutex)
```

```
gotest -bench=. -count=5 -benchtime=5x -timeout=0 > syncmap.result    (for sync.Map)
```

4. print the results

```
╰─➤  benchstat syncmap.result rwmutex.result 
name                                                                                        old time/op  new time/op  delta
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:1)-12                 325ms ±26%   377ms ± 8%  +16.01%  (p=0.032 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:4)-12                 235ms ± 1%   364ms ± 3%  +54.48%  (p=0.016 n=4+5)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:8)-12                 239ms ± 2%   375ms ± 5%  +57.03%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:15)-12                245ms ± 3%   407ms ± 6%  +66.02%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:30)-12                248ms ± 4%   437ms ± 3%  +76.07%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:1)-12               287ms ± 2%   293ms ± 2%     ~     (p=0.111 n=4+5)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:4)-12               219ms ± 1%   300ms ± 4%  +36.98%  (p=0.016 n=4+5)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:8)-12               221ms ± 2%   322ms ± 7%  +45.48%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:15)-12              217ms ± 1%   351ms ± 1%  +61.37%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:30)-12              217ms ± 1%   372ms ± 4%  +71.32%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:group_concat(distinct_b),concurrency:1)-12        329ms ± 1%   342ms ± 8%     ~     (p=0.111 n=4+5)
SingleGroup/(table:ndv_32,group_num:1,func:group_concat(distinct_b),concurrency:4)-12        234ms ± 1%   363ms ± 4%  +55.08%  (p=0.016 n=4+5)
SingleGroup/(table:ndv_32,group_num:1,func:group_concat(distinct_b),concurrency:8)-12        231ms ± 3%   402ms ± 6%  +74.59%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:group_concat(distinct_b),concurrency:15)-12       235ms ± 4%   416ms ± 5%  +77.19%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:group_concat(distinct_b),concurrency:30)-12       233ms ± 1%   441ms ± 3%  +89.38%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:1)-12                 338ms ±28%   388ms ±12%     ~     (p=0.056 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:4)-12                 259ms ±16%   374ms ± 5%  +44.37%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:8)-12                 240ms ± 4%   407ms ± 5%  +69.81%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:15)-12                240ms ± 4%   426ms ± 9%  +77.46%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:30)-12                251ms ± 3%   440ms ± 4%  +75.68%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:1)-12               786ms ± 1%   786ms ± 1%     ~     (p=0.548 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:4)-12               1.08s ± 1%   0.93s ± 2%  -13.25%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:8)-12               1.10s ± 1%   0.96s ± 4%  -13.03%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:15)-12              1.11s ± 1%   1.01s ± 4%   -9.24%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:30)-12              1.13s ± 3%   1.03s ± 1%   -9.17%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:1)-12             530ms ± 8%   545ms ± 7%     ~     (p=0.151 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:4)-12             694ms ±14%   615ms ± 2%     ~     (p=0.151 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:8)-12             737ms ± 2%   638ms ± 2%  -13.43%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:15)-12            742ms ± 1%   679ms ± 4%   -8.47%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:30)-12            758ms ± 2%   702ms ± 2%   -7.39%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:group_concat(distinct_b),concurrency:1)-12      678ms ± 2%   705ms ± 8%     ~     (p=0.056 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:group_concat(distinct_b),concurrency:4)-12      1.06s ± 1%   0.92s ± 4%  -13.11%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:group_concat(distinct_b),concurrency:8)-12      1.09s ± 2%   0.96s ± 1%  -12.03%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:group_concat(distinct_b),concurrency:15)-12     1.09s ± 1%   1.00s ± 1%   -8.76%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:group_concat(distinct_b),concurrency:30)-12     1.11s ± 1%   1.02s ± 1%   -7.97%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:1)-12               783ms ± 1%   807ms ± 5%   +3.07%  (p=0.032 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:4)-12               1.07s ± 2%   0.92s ± 2%  -14.19%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:8)-12               1.09s ± 1%   0.96s ± 2%  -11.74%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:15)-12              1.11s ± 1%   1.00s ± 1%   -9.90%  (p=0.008 n=5+5)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:30)-12              1.12s ± 2%   1.02s ± 2%   -9.36%  (p=0.008 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:1)-12               572ms ±11%   590ms ± 2%     ~     (p=0.413 n=5+4)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:4)-12               372ms ± 3%   361ms ± 3%     ~     (p=0.063 n=5+4)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:8)-12               393ms ± 3%   378ms ± 2%   -3.67%  (p=0.008 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:15)-12              391ms ± 4%   378ms ± 2%   -3.36%  (p=0.032 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:30)-12              385ms ± 2%   377ms ± 4%     ~     (p=0.151 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:1)-12             480ms ±20%   495ms ± 0%     ~     (p=0.730 n=5+4)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:4)-12             324ms ± 3%   323ms ± 7%     ~     (p=0.548 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:8)-12             332ms ± 1%   319ms ± 2%   -3.93%  (p=0.008 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:15)-12            333ms ± 2%   321ms ± 1%   -3.58%  (p=0.008 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:30)-12            332ms ± 3%   324ms ± 1%   -2.31%  (p=0.032 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:group_concat(distinct_b),concurrency:1)-12      516ms ±14%   527ms ± 1%     ~     (p=0.905 n=5+4)
1000Groups/(table:ndv_32,group_num:1000,func:group_concat(distinct_b),concurrency:4)-12      368ms ± 3%   351ms ± 1%   -4.65%  (p=0.016 n=5+4)
1000Groups/(table:ndv_32,group_num:1000,func:group_concat(distinct_b),concurrency:8)-12      358ms ± 2%   345ms ± 3%   -3.53%  (p=0.008 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:group_concat(distinct_b),concurrency:15)-12     362ms ± 2%   345ms ± 3%   -4.83%  (p=0.016 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:group_concat(distinct_b),concurrency:30)-12     368ms ± 2%   355ms ± 3%     ~     (p=0.056 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:1)-12               580ms ± 9%   571ms ± 1%     ~     (p=0.730 n=5+4)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:4)-12               366ms ± 2%   361ms ± 7%     ~     (p=0.190 n=4+5)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:8)-12               364ms ± 4%   354ms ± 1%     ~     (p=0.095 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:15)-12              363ms ± 2%   348ms ± 1%   -3.93%  (p=0.008 n=5+5)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:30)-12              363ms ± 2%   355ms ± 0%   -2.28%  (p=0.008 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:1)-12             1.01s ±16%   1.06s ± 0%     ~     (p=0.190 n=5+4)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:4)-12             552ms ± 1%   499ms ± 2%   -9.53%  (p=0.016 n=5+4)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:8)-12             475ms ± 6%   447ms ± 9%     ~     (p=0.095 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:15)-12            470ms ± 2%   434ms ± 2%   -7.67%  (p=0.008 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:30)-12            476ms ± 4%   435ms ± 2%   -8.64%  (p=0.008 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:1)-12           732ms ± 1%   750ms ± 1%   +2.42%  (p=0.029 n=4+4)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:4)-12           473ms ±16%   403ms ± 1%  -14.67%  (p=0.016 n=5+4)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:8)-12           408ms ± 2%   379ms ± 4%   -7.11%  (p=0.008 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:15)-12          410ms ± 4%   371ms ± 1%   -9.69%  (p=0.008 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:30)-12          416ms ± 3%   384ms ± 3%   -7.70%  (p=0.008 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:group_concat(distinct_b),concurrency:1)-12    981ms ± 6%   965ms ±15%     ~     (p=1.000 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:group_concat(distinct_b),concurrency:4)-12    754ms ±50%   683ms ±37%     ~     (p=0.690 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:group_concat(distinct_b),concurrency:8)-12    707ms ± 7%   587ms ±27%     ~     (p=0.111 n=4+5)
1000Groups/(table:ndv_rand,group_num:1000,func:group_concat(distinct_b),concurrency:15)-12   1.29s ±59%   0.71s ± 9%     ~     (p=0.111 n=5+4)
1000Groups/(table:ndv_rand,group_num:1000,func:group_concat(distinct_b),concurrency:30)-12   1.95s ±65%  1.25s ±131%     ~     (p=0.548 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:1)-12             1.61s ±32%   1.25s ± 3%     ~     (p=0.730 n=5+4)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:4)-12            1.18s ±138%   0.98s ±24%     ~     (p=0.310 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:8)-12             2.12s ±60%   0.74s ±23%  -65.01%  (p=0.008 n=5+5)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:15)-12            1.30s ±30%  1.39s ±115%     ~     (p=0.690 n=5+5)
```

## Testdata

### Tables

* ndv_32: 1000000 rows, NDV of `a` = 1000, `b` is generated using `rand.Int31()`

``` mysql
mysql> desc ndv_32;
+-------+------------+------+------+---------+-------+
| Field | Type       | Null | Key  | Default | Extra |
+-------+------------+------+------+---------+-------+
| a     | bigint(20) | YES  |      | NULL    |       |
| b     | bigint(20) | YES  |      | NULL    |       |
+-------+------------+------+------+---------+-------+
2 rows in set (0.00 sec)
```

* ndv_rand: 1000000 rows, NDV of `a` = 1000, NDV of `b` = 32

``` mysql
mysql> desc ndv_rand;
+-------+------------+------+------+---------+-------+
| Field | Type       | Null | Key  | Default | Extra |
+-------+------------+------+------+---------+-------+
| a     | bigint(20) | YES  |      | NULL    |       |
| b     | bigint(20) | YES  |      | NULL    |       |
+-------+------------+------+------+---------+-------+
2 rows in set (0.00 sec)
```

### Queries

#### Single group

``` mysql
SELECT /*+ HASH_AGG() */ avg(distinct b) FROM `ndv_32`

SELECT /*+ HASH_AGG() */ avg(distinct b) FROM `ndv_rand`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `ndv_32`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `ndv_rand`

SELECT /*+ HASH_AGG() */ group_concat(distinct b) FROM `ndv_32`

SELECT /*+ HASH_AGG() */ group_concat(distinct b) FROM `ndv_rand`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_32`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_rand`

```

#### 1000 groups

``` mysql
SELECT /*+ HASH_AGG() */ avg(distinct b) FROM `ndv_32` GROUP BY `a`

SELECT /*+ HASH_AGG() */ avg(distinct b) FROM `ndv_rand` GROUP BY `a`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `ndv_32` GROUP BY `a`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `ndv_rand` GROUP BY `a`

SELECT /*+ HASH_AGG() */ group_concat(distinct b) FROM `ndv_32` GROUP BY `a`

SELECT /*+ HASH_AGG() */ group_concat(distinct b) FROM `ndv_rand` GROUP BY `a`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_32` GROUP BY `a`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_rand` GROUP BY `a`

```
