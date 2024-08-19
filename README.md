# Solution for Vorto coding challenge
The code in this repository can be build by going into cmd and running 'go build -o schedule'

The resulting executable is run with 
`./schedule -f /path/to/file`

The algorithm uses a nearest neighbor approach to get from load to load, doing limited search through nearest neighbors (using both deterministic and monte carlo search) to find the minimum cost solution.
