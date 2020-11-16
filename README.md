# The-Go-Programming-Language

## [Chapter-01](./chapter-01)

1. [Hello World](./chapter-01/helloWorld.go)
2. [Print Command Line Arguments](./chapter-01/printCLIArguments.go)
3. [Find Duplicate strings using buffer](./chapter-01/dup01.go)
4. [Find Duplicate strings with count using buffer](./chapter-01/dup02.go)
5. [Find Duplicate strings with count without using buffer](./chapter-01/dup03.go)
6. [Lissajous - create gif](./chapter-01/Lissajous.go)
7. [fetch - Read from url](./chapter-01/fetch.go)
8. [fetchExercise - Fetch exercise](./chapter-01/fetchExercise.go)
9. [fetchAll - read from multiple url in concurrency](./chapter-01/fetchAll.go)
10. [fetchAllExercise - Print content of url in concurrency](./chapter-01/fetchAllExercise.go)
11. [Listen localhost](./chapter-01/server1.go)
12. [Listen server along with the counter](./chapter-01/server2.go)
13. [Listen server and print header and request params](./chapter-01/server3.go)

## [Chapter-02](./chapter-02)

1. [Echo Boiling Point of Water Using Constant](./chapter-02/boiling.go)
2. [Echo Boiling and Freezing Point of Water Using Function to Return Float](./chapter-02/ftoc.go)
3. [Package for conversion to farenheit and celcius](./chapter-02/tempconv)
4. [Import package and use it's functions](./chapter-02/cf.go)
4. [Init function usage](./chapter-02/popcount.go)
5. [Init function usage exercise](./chapter-02/popcountExercise.go)

## [Chapter-03](./chapter-03)

1. [Use float for SVG polygon rendering](./chapter-03/surface.go)
2. [Remove all the positions having non Float64 value for SVG polygon rendering](./chapter-03/surfaceExercise1.go)
3. [Using complex number to render image in grey scale](./chapter-03/mandelbrot.go)
4. [Using complex number to render image in YCbCr scale](./chapter-03/mandelbrotExercise1.go)
5. [Using Newton's method in finding and rendering complex numbers](./chapter-03/mandelbrotExercise2.go)
6. [Add comma at given position in string](./chapter-03/comma.go)
7. [Add comma at given position in string using buffer writer](./chapter-03/commaExercise1.go)
8. [Get basename of the given path](./chapter-03/basename.go)
9. [Get basename of the given path using strings LastIndex function](./chapter-03/basename1.go)
10. [Print array of integers using bytes buffer](./chapter-03/printints.go)
11. [Check if two strings are anagrams using runes](./chapter-03/anagram.go)
12. [Use integer in bit wise operation](./chapter-03/netflag.go)

## [Chapter-04](./chapter-04)

1. [Use byte array to generate sha256 array](./chapter-04/sha256.go)
2. [Calculate number of bit different in 2 sha256 hashes](./chapter-04/sha256Exercise1.go)
3. [Calculate sha256, sha384, sha512 hashes of an input based on the arguments provided](./chapter-04/sha256Exercise2.go)
4. [Append new elements to slice](./chapter-04/append.go)
5. [Remove empty strings from array](./chapter-04/nonempty.go)
6. [Reverse the slice](./chapter-04/reverse.go)
7. [Reverse array using pointers](./chapter-04/reverseExercise1.go)
8. [Rotate slice from given position](./chapter-04/rotateExercise1.go)
9. [Get unique strings using map](./chapter-04/dedup.go)
10. [Get unique runes along with count](./chapter-04/charcount.go)
11. [Get info regarding edges in graph using nested map](./chapter-04/graph.go)
12. [Sort binary tree](./chapter-04/treesort.go)
13. [Use nested struct](./chapter-04/embed.go)
14. [Marshal struct to json object](./chapter-04/movie.go)
15. [Unmarshal json api response to struct](./chapter-04/github.go)
16. [Parse the response in a given template](./chapter-04/issuesReport.go)
17. [Parse the response in given HTML template](./chapter-04/issuesHtml.go)
18. [Parsing is data type sensitive](./chapter-04/autoescape.go)

## [Chapter-05](./chapter-05)

1. [Get all the href links from webpage using recursion](./chapter-05/findlinks.go)
2. [Get the document tree of the html page using recursion](./chapter-05/outline.go)
3. [Get all the href links from webpage returning multiple values](./chapter-05/findlinks1.go)
4. [Count words and images present in html](./chapter-05/countWordsAndImages.go)
5. [Print and log error if site is not recheable](./chapter-05/wait.go)
6. [Use function as values to print the outline of webpage](./chapter-05/outline1.go)
7. [Print square of a number using anonymous function](./chapter-05/squares.go)
8. [Sort the keys of map using anonymous function](./chapter-05/toposort.go)
9. [Get all the href links as a bfs traversal](./chapter-05/findlinks2.go)
10. [Get sum of the integers using variable parameters for function call](./chapter-05/sum.go)
11. [Print title of the html page](./capter-05/title.go)
12. [Print title of the html page use defer to close the buffer](./capter-05/title2.go)
13. [Defer runs after the return values are set](./chapter-05/trace.go)
14. [Copy url info to file](./chapter-05/fetch.go)
15. [Copy url info to file using defer](./chapter-05/fetchExercise.go)
16. [Panic stops the application after running all defer statements](./chapter-05/defer1.go)
17. [Catch panic so that application does not stop](./chapter-05/defer2.go)
18. [Recover from panic](./chapter-05/title3.go)
19. [Recover from panic and return without any return statement](./chapter-05/recoverExercise.go)

## [Chapter-06](./chapter-06)

1. [Define object methods to calculate distance](./chapter-06/geometry.go)
2. [Use method as pointer receiver](./chapter-06/urlvalues.go)
3. [Struct embedding - Using methods of nested struct directly through parent struct](./chapter-06/coloredpoint.go)
4. [Use bit vector as set](./chapter-06/intset.go)
5. [Add function to remove int from bit vector](./chapter-06/intsetExercise.go)

## [Chapter-07](./chapter-07)

1. [Show interface usage of Fprintf function](./chapter-07/bytecounter.go)
2. [Show interface usage of Fprintf by overwriting the write function to count words](./chapter-07/wordcounterExercise.go)
3. [Implement new limit reader interface](./chapter-07/limitReaderExercise.go)
4. [Use flag inteface to create new var for flag duration](./chapter-07/sleep.go)
5. [Create new flag by implementing functions for flag interface](./chapter-07/tempconv.go)
6. [Create new flag by implementing functions for flag interface exercise 7.6 and 7.7](./chapter-07/tempconv.go)
7. [Sorting with sort interface](./chapter-07/sorting.go)
8. [Use http interface](./chapter-07/http1.go)
9. [Handle multiple paths on server](./chapter-07/http2.go)
10. [Use mux server for multiple paths](./chapter-07/http3.go)
11. [Use default mux server](./chapter-07/http4.go)
12. [Expression evaluator interface](./chapter-07/eval.go)
13. [Use interface and assertions to print xml text](./chapter-07/xmlselect.go)

## [Chapter-08](./chapter-08)

1. [Calculate fibonaci number using goroutines](./chapter-08/spinner.go)
2. [Print server time to connection](./chapter-08/clock1.go)
3. [Listen to connection using net dial](./chapter-08/netcat1.go)
4. [Print server to multiple connections](./chapter-08/clock2.go)
5. [Listen to connection in async](./chapter-08/reverb1.go)
6. [Take input and print output in concurrency](./chapter-08/netcat2.go)
7. [Print to connection in concurrency](./chapter-08/reverb2.go)
8. [Channel between stdin and stdout](./chapter-08/netcat3.go)
9. [Pass numbers between goroutines using channels](./chapter-08/pipeline.go)
10. [Close one send channel so that all channels can be closed](./chapter-08/pipeline2.go)
11. [Use unidirectional channels](./chapter-08/pipeline3.go)
12. [Multiple go routines](./chapter-08/thumbnail.go)
13. [Web crawler](./chapter-08/crawl1.go)
14. [Web crawler concurrently](./chapter-08/crawl2.go)
15. [Web crawler concurrently with limiting threads](./chapter-08/crawl3.go)
16. [Countdown example with receiving from channel](./chapter-08/countdown1.go)
17. [Countdown example with receiving from channel and abort](./chapter-08/countdown2.go)
18. [Select multiplexing countdown timer example](./chapter-08/countdown3.go)
19. [Directory traversal for given folder](./chapter-08/du1.go)
20. [Directory traversal for given folder with verbose feature](./chapter-08/du2.go)
21. [Directory traversal for given folder concurrently](./chapter-08/du3.go)
22. [Directory traversal for given folder concurrently with verbose feature](./chapter-08/du3Exercise.go)
23. [Directory traversal for given folder concurrently with verbose and cancellation feature](./chapter-08/du4.go)
24. [Chat server](./chapter-08/chat.go)
25. [Chat server with showing present members](./chapter-08/chatExercise1.go)

## [Chapter-09](./chapter-09)

1. [Bank deposit concurrency](./chapter-09/bank1.go)
2. [Bank deposit concurrency with withdrawal feature](./chapter-09/bank1Exercise.go)
3. [Bank deposit concurrency using semaphore lock](./chapter-09/bank2.go)
4. [Bank deposit concurrency using mutex](./chapter-09/bank3.go)
5. [Cache use case](./chapter-09/memo1.go)
6. [Cache use case async api calls](./chapter-09/memo2.go)
7. [Cache use case aysnc api calls mutex lock](./chapter-09/memo3.go)
8. [Cache use case aysnc api calls using channel and mutex lock](./chapter-09/memo4.go)
9. [Cache use case aysnc api calls using channel only](./chapter-09/memo5.go)
10. [Ping pong exercise](./chapter-09/pingpongExercise.go)

## [Chapter-10](./chapter-10)

1. [Blank imports](./chapter-10/jpeg.go)

## [Chapter-11](./chapter-11)

1. [Check whether the string is palindrome](./chapter-11/word1)
2. [Check whether the string is palindrome or not ignoring punctuations and special characters](./chapter-11/word2)
3. [Echo based on various flags](./chapter-11/echo)
4. [Send email on storage capacity warning](./chapter-11/storage.go)
5. [Write testable code for sending mail on storage warning](./chapter-11/storage)

## [Chapter-12](./chapter-12)

1. [Print any format as a string using reflect type and value](./chapter-12/format.go)
2. [Search values for params in URL query](./chapter-12/params)
3. [Get the query params values in url query](./chapter-12/search.go)