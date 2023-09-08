# Uniform Node Sampling Service Robust against Collusions of Malicious Nodes Implementation

[paper](https://hal.science/hal-00804430/document)

## 1- Steps to get the data
* We downloaded and unziped the first [dataset](https://ita.ee.lbl.gov/html/contrib/NASA-HTTP.html)
 `wget ftp://ita.ee.lbl.gov/traces/NASA_access_log_Jul95.gz`
* `Rscript clean_data.r` (to obtained the file `resultJul95`)

## 2- Steps to execute
* `go build main`
* `./main -t [number]`

number=0 for knowledge free sample and 1 for omniscient strategy

## 3- Steps to get the output
* Check for the output file in the data\ folder
* `Rscript dkl.r`