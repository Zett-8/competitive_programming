const gcd = (a, b) => {
  if (!b) return a
  return gcd(b, a%b)
}


// oneline
const gcd = (a, b) => { if (!b) { return a } return gcd(b, a%b) }