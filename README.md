# Benchmark for TiDB HashAgg

For [pull/19807](https://github.com/pingcap/tidb/pull/19807), `set.SyncSet` can be implemented using `sync.Map` or `map[interface{}]struct{}` with `sync.RWMutex`.
This repository benchmarked both implementations with self-generated input data.

Data generation resides in [`data_gen.go`](https://github.com/dyzsr/hashaggbench/blob/master/data_gen.go),
and benchmark is in [`hashagg_test.go`](https://github.com/dyzsr/hashaggbench/blob/master/hashagg_test.go)

## Steps

For [pull/19807](https://github.com/pingcap/tidb/pull/19807), I ran the benchmark in following steps for both implementations:

###1. Launch TiDB

```
make
bin/tidb-server
```

### 2. Generate input data

```
go run .
```

### 3. Run benchmark

```
gotest -bench=. -count=20 -benchtime=1x -timeout=0 > rwmutex.result    (for map with sync.RWMutex)
```

```
gotest -bench=. -count=20 -benchtime=1x -timeout=0 > syncmap.result    (for sync.Map)
```

### 4. Print the results

```
╰─➤  benchstat syncmap.result rwmutex.result 
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:1)-12	363ms ± 4%	361ms ± 3%	~	(p=0.224 n=16+16)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:4)-12	243ms ± 6%	247ms ± 2%	+1.48%	(p=0.007 n=18+17)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:8)-12	240ms ± 6%	238ms ± 6%	~	(p=0.393 n=20+18)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:15)-12	238ms ± 7%	239ms ± 6%	~	(p=0.414 n=20+20)
SingleGroup/(table:ndv_32,group_num:1,func:avg(distinct_b),concurrency:30)-12	238ms ± 8%	243ms ± 9%	~	(p=0.076 n=20+20)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:1)-12	286ms ± 4%	286ms ± 4%	~	(p=0.909 n=18+17)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:4)-12	220ms ± 6%	233ms ± 4%	+6.27%	(p=0.000 n=19+19)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:8)-12	219ms ± 3%	228ms ± 6%	+4.14%	(p=0.000 n=18+20)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:15)-12	232ms ±13%	222ms ± 4%	~	(p=0.080 n=19+19)
SingleGroup/(table:ndv_32,group_num:1,func:count(distinct_b),concurrency:30)-12	218ms ± 5%	224ms ± 8%	~	(p=0.059 n=18+20)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:1)-12	372ms ± 5%	363ms ± 4%	−2.49%	(p=0.002 n=18+16)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:4)-12	246ms ± 5%	251ms ± 6%	~	(p=0.067 n=17+18)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:8)-12	235ms ± 7%	243ms ± 6%	+3.57%	(p=0.004 n=20+20)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:15)-12	240ms ± 7%	243ms ± 4%	~	(p=0.253 n=20+20)
SingleGroup/(table:ndv_32,group_num:1,func:sum(distinct_b),concurrency:30)-12	241ms ±11%	241ms ± 8%	~	(p=0.904 n=20+20)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:1)-12	785ms ±10%	718ms ± 4%	−8.64%	(p=0.000 n=20+19)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:4)-12	1.06s ± 2%	1.08s ± 3%	+2.20%	(p=0.000 n=18+18)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:8)-12	1.08s ± 5%	1.10s ± 4%	~	(p=0.063 n=20+18)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:15)-12	1.10s ± 3%	1.09s ± 2%	~	(p=0.050 n=20+19)
SingleGroup/(table:ndv_rand,group_num:1,func:avg(distinct_b),concurrency:30)-12	1.11s ± 4%	1.10s ± 2%	~	(p=0.258 n=20+19)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:1)-12	478ms ± 9%	472ms ± 3%	~	(p=0.461 n=19+18)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:4)-12	721ms ± 6%	664ms ±10%	−7.96%	(p=0.000 n=19+18)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:8)-12	728ms ± 6%	666ms ± 6%	−8.47%	(p=0.000 n=20+20)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:15)-12	747ms ± 6%	661ms ± 4%	−11.52%	(p=0.000 n=20+20)
SingleGroup/(table:ndv_rand,group_num:1,func:count(distinct_b),concurrency:30)-12	767ms ± 7%	662ms ± 6%	−13.59%	(p=0.000 n=20+20)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:1)-12	723ms ± 5%	722ms ± 4%	~	(p=0.749 n=20+19)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:4)-12	1.06s ± 5%	1.08s ± 3%	+2.21%	(p=0.004 n=18+20)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:8)-12	1.08s ± 5%	1.09s ± 4%	~	(p=0.096 n=20+20)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:15)-12	1.10s ± 6%	1.10s ± 3%	~	(p=0.758 n=20+20)
SingleGroup/(table:ndv_rand,group_num:1,func:sum(distinct_b),concurrency:30)-12	1.10s ± 3%	1.12s ± 4%	+1.37%	(p=0.046 n=20+20)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:1)-12	572ms ± 5%	568ms ± 5%	~	(p=0.374 n=18+19)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:4)-12	356ms ± 8%	356ms ± 9%	~	(p=0.544 n=19+19)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:8)-12	370ms ±12%	358ms ± 5%	−3.30%	(p=0.026 n=20+20)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:15)-12	378ms ± 6%	349ms ± 7%	−7.52%	(p=0.000 n=20+20)
1000Groups/(table:ndv_32,group_num:1000,func:avg(distinct_b),concurrency:30)-12	369ms ± 6%	350ms ± 4%	−5.03%	(p=0.000 n=20+20)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:1)-12	480ms ± 7%	482ms ± 6%	~	(p=0.454 n=17+17)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:4)-12	313ms ± 7%	305ms ± 6%	−2.68%	(p=0.023 n=20+20)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:8)-12	312ms ± 4%	303ms ± 4%	−2.72%	(p=0.000 n=20+20)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:15)-12	311ms ± 5%	308ms ± 4%	~	(p=0.228 n=20+18)
1000Groups/(table:ndv_32,group_num:1000,func:count(distinct_b),concurrency:30)-12	309ms ± 4%	303ms ± 4%	−1.73%	(p=0.010 n=19+20)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:1)-12	550ms ± 5%	550ms ± 4%	~	(p=0.729 n=19+19)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:4)-12	361ms ± 6%	342ms ± 7%	−5.32%	(p=0.000 n=19+19)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:8)-12	344ms ± 5%	332ms ± 9%	−3.25%	(p=0.002 n=20+20)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:15)-12	339ms ± 5%	329ms ± 6%	−2.97%	(p=0.004 n=19+20)
1000Groups/(table:ndv_32,group_num:1000,func:sum(distinct_b),concurrency:30)-12	346ms ± 7%	337ms ±10%	−2.44%	(p=0.028 n=19+20)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:1)-12	1.00s ± 6%	0.98s ± 3%	~	(p=0.219 n=19+17)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:4)-12	535ms ±10%	498ms ± 5%	−6.89%	(p=0.000 n=20+20)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:8)-12	472ms ±10%	435ms ±10%	−7.75%	(p=0.000 n=20+19)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:15)-12	446ms ±16%	417ms ± 7%	−6.44%	(p=0.000 n=19+20)
1000Groups/(table:ndv_rand,group_num:1000,func:avg(distinct_b),concurrency:30)-12	463ms ±12%	435ms ± 8%	−6.05%	(p=0.009 n=20+20)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:1)-12	678ms ± 5%	687ms ± 4%	~	(p=0.075 n=18+19)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:4)-12	436ms ± 8%	405ms ± 8%	−7.21%	(p=0.000 n=19+20)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:8)-12	391ms ±12%	364ms ± 8%	−6.91%	(p=0.000 n=20+20)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:15)-12	391ms ±11%	360ms ± 8%	−7.85%	(p=0.000 n=20+20)
1000Groups/(table:ndv_rand,group_num:1000,func:count(distinct_b),concurrency:30)-12	397ms ±11%	362ms ± 7%	−8.90%	(p=0.000 n=20+20)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:1)-12	984ms ± 3%	979ms ± 4%	~	(p=0.271 n=18+19)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:4)-12	533ms ±10%	487ms ± 6%	−8.54%	(p=0.000 n=19+18)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:8)-12	446ms ±11%	425ms ±14%	−4.54%	(p=0.011 n=19+20)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:15)-12	439ms ±14%	410ms ±10%	−6.51%	(p=0.002 n=20+20)
1000Groups/(table:ndv_rand,group_num:1000,func:sum(distinct_b),concurrency:30)-12	445ms ±10%	411ms ± 9%	−7.74%	(p=0.000 n=20+20) 
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

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_32`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_rand`

```

#### 1000 groups

``` mysql
SELECT /*+ HASH_AGG() */ avg(distinct b) FROM `ndv_32` GROUP BY `a`

SELECT /*+ HASH_AGG() */ avg(distinct b) FROM `ndv_rand` GROUP BY `a`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `ndv_32` GROUP BY `a`

SELECT /*+ HASH_AGG() */ count(distinct b) FROM `ndv_rand` GROUP BY `a`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_32` GROUP BY `a`

SELECT /*+ HASH_AGG() */ sum(distinct b) FROM `ndv_rand` GROUP BY `a`

```
