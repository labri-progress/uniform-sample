
#install.packages('LaplacesDemon')

library(stringr)
library(dplyr)
library(philentropy)
library(LaplacesDemon)

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
d1 <- divergence("data/resultJul95")
"---------------------"
"---------------------"
d2 <- divergence("data/output")

G = 1 - d2/d1
print(paste("The Gain is ", G))

