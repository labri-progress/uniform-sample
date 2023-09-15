source("/home/amukam/thss/experiments/uniform-sample/dkl.r")

experiment =  "KL1"
"--------------------------"
d1 <- divergence("data/resultJul95trunk")
times <- 5
for (i in 3:times){
    output = paste("data/output", i, sep="")
    dKL(c(i, output, experiment))
}


srcfolder= "/home/amukam/thss/experiments/uniform-sample/result"
srcname = paste(srcfolder,"/",experiment, sep="")
srcname

col_names= c("Expe","dKL")

g <- read.table(srcname, header=TRUE,
  sep="", 
  col.names = col_names)
  
#g
mean(g$dKL)
# KL1  0.4493744
# KL2  0.4086942
# KL3  0.4244801