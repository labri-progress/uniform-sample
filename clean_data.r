library(stringr)
library(dplyr)
library(philentropy)

# get a substring from a string
substrRight <- function(x, n){
  substr(x, nchar(x)-n+1, nchar(x))
}
file_in = "data/clarknet_access_log_Sep4/access_1"
#file_in = "data/clarknet_access_log_Aug28/Aug28_log"
#file_in = "data/NASA_access_log_Jul95/access_log_Jul95"
#file_in = "data/6-Saskatchewan-UofS_access_log.txt" #data/usask_access_log/UofS_access_log"
x = read.delim(file_in, header = FALSE)

conn <- file(file_in,open="r")
linn <-readLines(conn)
length(linn)
# output file
filename = paste("data/result", substrRight(file_in,5), sep="")
for (o in 1:length(linn)){
     i = linn[o]
     #print(i)
     cline <- grep(" - -", i, value=TRUE)
     if(length(cline) > 0){
          res = str_extract_all(cline,"[:graph:]+")[[1]][1]
     }else{
          res = i
     }
     write(res, append=TRUE, file = filename)
}
close(conn)