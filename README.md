# LinkSphere project
LinkSphere - worldwide online e-shop. We introduce only advanced devices.

Working on this project: Anvar Tamabayev & Danial Sultanbek.

Screenshot of our site(not main FrontEnd site):
![image](https://github.com/user-attachments/assets/4e8b3fd8-cd79-4e50-a6f1-fdee0dd42cae)


How to run server:

Postman

1. Open lala.go through terminal and run it
   
3. Open postman
   
3.Through postman we make requests or receive responses by inserting our localhost link beforehand.

4.For the get method we simply select the get method and send a get message.

5.For post method go to “body” and write our JSON message, select post method and send message.



Site

1.Open lala.go through terminal and run it

2.Go to the link

3.And already directly with buttons and fields can make the same requests and receive messages

4. All methods and their messages are recorded in a table
   


Database

1.Open via terminal lala.go

2.Enter your db data in connectDB

3.Write the functions of interest to work with DB

4.Start the server

5.Check the presence of the table and the entered data.

P.s. Automigration will work on our site, which performs the check and creation of the required DB table, be sure not to forget to change the dsn line to enter your data on the type of password, user, database name.

Instruments/Programms/Git's/Framework's that we used: 

VSCode, PostGreSQL, Postman. 
github.com/golang-migrate/migrate/v4
github.com/hashicorp/errwrap
github.com/hashicorp/go-multierror
github.com/jackc/pgpassfile
github.com/jackc/pgservicefile
github.com/jackc/pgx/v5
github.com/jackc/puddle/v2
github.com/jinzhu/inflection
github.com/jinzhu/now
github.com/jmoiron/sqlx
github.com/labstack/echo/v4
github.com/labstack/gommon
github.com/lib/pq
github.com/mattn/go-colorable
github.com/mattn/go-isatty
github.com/valyala/bytebufferpool
github.com/valyala/fasttemplate
go.uber.org/atomic
golang.org/x/crypto
golang.org/x/net
golang.org/x/sync
golang.org/x/sys
golang.org/x/text
gorm.io/driver/postgres
gorm.io/gorm 
