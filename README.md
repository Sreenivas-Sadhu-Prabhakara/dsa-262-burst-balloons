# Burst Balloons

> NeetCode 150 · 2-D DP — think about the LAST balloon in each range, not the first: its neighbors are then the fixed boundaries, decoupling the subranges. Go, stdlib only, tested.

## Problem

Bursting balloon i scores `left · i · right` for its *current* neighbors, which shift as balloons vanish. Maximize total coins — the hardest DP on the NeetCode list.

## Approach

Conditioning on the *first* balloon burst fails: the remaining pieces interact through their new adjacency. The famous inversion conditions on the **last** balloon k in a range (l, r): when k finally bursts, everything between has gone, so its neighbors are exactly the boundaries l and r — fixed, known, independent. That makes `dp[l][r] = max over k of dp[l][k] + dp[k][r] + a[l]·a[k]·a[r]`, an interval DP (same fill-by-width engine as matrix-chain, dsa-176) over the array padded with virtual 1-balloons so edge bursts need no cases. A full burst-order brute force cross-checks the DP in tests — the honest way to validate an inversion this slippery.

**Complexity:** O(n³) time, O(n²) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
