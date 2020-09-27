class UnionFind {
  constructor(n) {
    this.p = new Array(n).fill(null).map((_, i) => i)
    this.height = new Array(n).fill(0)
    this.size = new Array(n).fill(1)
  }

  unite(a, b) {
    this.link(this.findSet(a), this.findSet(b))
  }

  link(a, b) {
    if (a === b) return null

    if (this.height[a] > this.height[b]) {
      this.p[b] = a
      this.size[a] += this.size[b]
    } else {
      this.p[a] = b
      this.size[b] += this.size[a]

      if (this.height[a] === this.height[b]) this.height[b]++
    }
  }

  findSet(x) {
    if (this.p[x] === x) return x
    return this.findSet(this.p[x])
  }
}


const main = (input) => {
  const [t, ...q] = input.trim().split('\n')
  const [n,] = t.split(' ').map(Number)

  const uf = new UnionFind(n)

  for (let i=0; i<q.length; i++) {
    const [a, b] = q[i].split(' ').map(Number)
    uf.unite(a-1, b-1)
  }

  const ans = {}
  for (let i=0; i<n; i++) ans[uf.findSet(i)] = 1


  return console.log(Object.keys(ans).length - 1)
}
