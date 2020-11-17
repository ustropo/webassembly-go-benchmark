#include <math.h>
#include <emscripten.h>

EMSCRIPTEN_KEEPALIVE
long fibonacci(int n) 
{
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

EMSCRIPTEN_KEEPALIVE
char isPrime(long n) 
{
    long limit = 0;
    limit = floor(sqrt(n));
    for(long i = 2; i <= limit; i++) {
        if (n % i == 0){
            return 0;
        }
    }

    return 1;
}
