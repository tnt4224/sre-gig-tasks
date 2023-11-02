
## Task1 (Async Review)

1. Generate a random number between 1 and 100
    1. If the number is higher than 50 print "It's closer to 100"
    2. If the number is lower than 50 print "It's closer to 0"
    3. Print the generated random number

2. Modify the previous code to print "It's 50!" if the random number is 50

3. Modify the conditional in the code you previously wrote to check not only if the  number is higher than 50, but also if it's even. If it's even and higher than 50, print,"It's closer to 100, and it's even!"


## NOTE: 
**rand.Seed has been deprecated since v 1.20[1]**

> The math/rand package now automatically seeds the global random number generator (used by top-level functions like Float64 and Int) with a random value, and the top-level Seed function has been deprecated. Programs that need a reproducible sequence of random numbers should prefer to allocate their own random source, using rand.New(rand.NewSource(seed)).
>
> Programs that need the earlier consistent global seeding behavior can set GODEBUG=randautoseed=0 in their environment.
>
> The top-level Read function has been deprecated. In almost all cases, crypto/rand.Read is more appropriate. "

Therefore, the below[2] code snippet is no longer required.

[1 https://tip.golang.org/doc/go1.20](https://tip.golang.org/doc/go1.20)


[2] rand.Seed(time.Now().UnixNano())


