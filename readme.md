### 1 Billion row challenge attempts.

All attempts using compiled billion.go

All runs on 
![image](https://github.com/Stuuu/1brc/assets/13997948/0322de02-6db3-4398-81e5-3a99486f2cc3)


## Attempt 1
Compiled run time
```
time ./billion &>/dev/null
================
CPU     96%
user    7:42.93
system  18.203
total   8:19.26
```


## Attempt 2
lets not handle errs, did it help whose to say
```
time ./billion &>/dev/null
================
CPU	98%
user	7:46.13
system	19.247
total	8:13.44
```

## Attempt 3
sum temps first, avg once at the end
```
time ./billion &>/dev/null
================
CPU	100%
user	6:08.78
system	15.755
total	6:22.59
```
