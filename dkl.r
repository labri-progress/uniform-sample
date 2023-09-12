args = commandArgs(trailingOnly=TRUE)

library(stringr)
library(dplyr)

divergence <- function(path) {
    print(paste("The data to analyze is", path))
    x = read.delim(path, header = FALSE)
    m = length(x[[1]])
    print(paste("Number of lines is", m))
    my_tab <- table(x) 
    print(paste("max occurence is", max(as.vector(my_tab))))

    my_tab_prob <- my_tab / sum(my_tab) # Create proportion table
    P <- as.vector(my_tab_prob)

    times = length(P)
    print(paste("Number of distinct element is  occurence is", times))

    v = 1/times
    Q <- rep(v, times)

    sum <- 0
    for (i in 1:times){
        sum = sum + P[i] * log(P[i]/Q[i])
    }
    print(paste("The KL divergence is ", sum))
    return(sum)
}
dKL <- function(args){
    expe= args[1]
    path= args[2]
    parameters = args[3]
    
    d2 <- divergence(path)
    
    filename = paste("result/", parameters, sep="")
    filename
    file.info(filename)$size
    if (!file.exists(filename)) {
        head <- "Expe dKL"
        write(head, append=TRUE, file = filename)
    }
    separator = "        "
    
    sol = paste(expe, d2, sep = separator)
    print(sol)
    write(sol, append=TRUE, file = filename)
}
#d1 <- divergence("data/resultJul95")
"---------------------"
"---------------------"
#d2 <- divergence("data/output")

#G = 1 - d2/d1
#print(paste("The Gain is ", G))