﻿# simpleService
 
![Service diagram](/img/diagram.png "Optional title")

On startup, `simpleService` will fetch holiday information and keep it in memory. This information will be used when clients request holidays of a certain type and within a specific range. 

This service exposes one single endpoint, as seen in the following picture. 
![Service diagram](/img/postman.png "Optional title")

To minimize the number of external dependencies, the endpoint was built using the standar library's `net/http` package.  
