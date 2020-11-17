function jsFibonacci(n) {
  if (n <= 1) {
    return n;
  }

  return jsFibonacci(n - 1) + jsFibonacci(n - 2)
}

function jsIsPrime(n) {
  var limit = Math.floor(Math.sqrt(n))
  for(var i = 2; i <= limit; i++) {
    if (n % i === 0) {
      return false
    }
  }

  return true;
}