# simpleService
 
![Service diagram](/img/diagram.png "Optional title")

On startup, `simpleService` will fetch holiday information and keep it in memory. This information will be used when clients request holidays of a certain type and within a specific range. 

This service exposes one single endpoint, as seen in the following picture. 
![Service diagram](/img/postman.png "Optional title")

To minimize the number of external dependencies, the endpoint was built using the standard library's `net/http` package.
Due to time constraints, it was not possible to implement all the initial requirements, e. g. use Docker, allow XML responses, and the actual connection from the cache to a remote holiday info server. This requirements and additional features could be effortlesstly added, when given enough time.
