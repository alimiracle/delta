dilta is File Sharing system
its good for Local networks or small teams
## Cross-platform 
dilta runs anywhere Go can compile for: Windows, Mac, Linux, ARM, etc.

## Lightweight 
dilta has low minimal requirements and can run on an inexpensive Raspberry Pi. Some users even run dilta instances on their NAS devices.
## Open Source
dilta is 100% open source and free of charge. All source code is available under the GPL License

## Dependencies

to build this app
you nede to install

* git

*golang

* htpasswd

## install


to build the app you nede golang

Download the program useing git

git clone https://notabug.org/alimiracle/dilta

then type
```bash
cd dilta
chmod + install
./install
```

## permissions

dilta Have two types of users

1- admin user

This type of user can do everything

to create new admin user

type 

```bash
chmod +x addadmin.sh
./addadmin.sh
```

2- Normal user

this type of  user can download files and view files and search for files

to create new Normal user

type

```bash
chmod +x adduser.sh
./adduser.sh
```
## run
to run the app
type 
```bash
sudo ./dilta&
```