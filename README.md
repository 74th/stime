# silent time command

Linux time command writes a result to stderr.
Silent time writes to the stdout and the command stdout are discarded.

```
(time psql -f test.sql) 2>&1 | tee result.log 2>&1
s_acctbal|s_name|n_name|p_partkey|p_mfgr|s_address|s_phone|s_comment
9938.53|Supplier#000005359       |UNITED KINGDOM           |185358|Manufacturer#4           |QKuHYh,vZGiwu2FWEJoLDx04|33-429-790-6131|uriously regular requests hag
9937.84|Supplier#000005969       |ROMANIA                  |108438|Manufacturer#1           |ANDENSOSmk,miq23Xfb5RWt6dvUcvt6Qa|29-520-692-3537|efully express instructions. regular requests against the slyly fin
9936.22|Supplier#000005250
  many many outputs...

real	0m2.850s
user	0m0.003s
sys	0m0.004s
```

```
stime psql -f test.sql | tee result.log
2.850
```

## feature

* show no stdout of the command
* show stderr of the command
* say same exit code with the command

## install

```
go get github.com/74th/stime
```

## License

Public Domain