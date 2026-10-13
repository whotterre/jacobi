package leetcode
func maxPrimes(n int, s int) []int {
    var res []int
    acc := 0
    primes := primeSieve(n)
    for i := 0; i < len(primes); i++ {
        if acc + primes[i] <= s {
            acc += primes[i]
            res = append(res, primes[i])
        } else {
            break
        }
    }
    return res
}

// Sieve of Erathostenes 
func primeSieve(n int) []int {
    prime := make([]bool, n + 1)
    for i := 0; i <= n; i++ {
        prime[i] = true
    }

    for p := 2; p * p <= n; p++ {
         if prime[p] {
            for i := p * p; i <= n; i += p {
                prime[i] = false;
            }
         }
    }
    var res []int 
    for p := 2; p <= n; p++ {
        if prime[p] == true{ 
            res = append(res, p)
        }
    }
    return res
}