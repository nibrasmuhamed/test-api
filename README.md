## Test

### build

#### Using Docker

you can run this project by using `docker-compose up`

<b>NOTE:</b> you need to have docker installed on your machine.
if you removed main docker will not work. it is a compiled version of this project which runs natively on docker alpine.(CG_ENABLED=0)

#### Using Go

to build this project you need to have go installed on your machine.
`go build -o main .`

execute the binary file `./main`

As it is mentioned in the api documentation to do not use any database, Everything is stored in memory. So, if you restart the server, all the data will be lost.
