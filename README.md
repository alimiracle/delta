delta is File Sharing system
its good for Local networks or small teams
## Cross-platform 
delta runs anywhere Go can compile for: Windows, Mac, Linux, ARM, etc.

## Lightweight 
delta has low minimal requirements and can run on an inexpensive Raspberry Pi. Some users even run delta instances on their NAS devices.
## Open Source
delta is 100% open source and free of charge. All source code is available under the GPL License

## Dependencies

to build this app
you nede to install

* git

*golang

* apache2-utils

## install

Download the program useing git

git clone https://notabug.org/alimiracle/delta

then type
```bash
cd delta
chmod + install
./install
```

## permissions

delta Have two types of users

1- admin user

This type of user can do everything

to create new admin user

type 

```bash
chmod +x addadmin.sh
./addadmin.sh admin_user
```

Replace admin_user With the user that you want

2- Normal user

this type of  user can download files and view files and search for files

to create new Normal user

type

```bash
chmod +x adduser.sh
./adduser.sh new_user
```

Replace new_user With the user that you want

## run
to run the app
type 
```bash
sudo ./delta&
```

now You can access  the program from web browser useing your PC IP