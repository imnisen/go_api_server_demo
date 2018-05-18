Created with `bee api app`

# database migration
```
in dir app

# create migrate scripts
bee generate migration [migrationfile] [-fields=""]
    generate migration file for making database schema update
    -fields: a list of table fields. Format: field:type, ...

# run it to migrate
bee migrate [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    run all outstanding migrations
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test


# rollback migrate
bee migrate rollback [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    rollback the last migration operation
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
```

detail refer to:
https://beego.me/docs/install/bee.md#generate-%E5%91%BD%E4%BB%A4
https://beego.me/docs/install/bee.md#migrate-%E5%91%BD%E4%BB%A4


# use orm
refer to https://beego.me/docs/mvc/model/orm.md