# ft_ls_go
42 school ft_ls project written using Golang

Topics covered:
- Intro to ANSI escape codes
- File Mode Bits

Arguments that ft_ls accepts:
-l - use a long listing format
-a - do not ignore entries starting with .
-R - list subdirectories recursively
-r - reverse order while sorting
-t - sort by modification time, newest first

Long listing format shows a lot more information presented to the user than the standard command. You will see the file permissions, the number of links, owner name, owner group, file size, time of last modification, and the file or directory name

```
ls -l file1 
-rw-rw-r--. 1 lilo lilo 0 Feb 26 07:08 file1
```
- -rw-rw-r- permissions
- 1 : number of linked hard-links
- lilo: owner of the file
- lilo: to which group this file belongs to
- 0: size
- Feb 26 07:08 modification/creation date and time
- file1: file/directory name

-rw-rw-r--
The permissions can be subdivided to four parts. The first is the special permission flag that can have the following values:
- _ – no special permissions
- d – directory
- l– The file or directory is a symbolic link
- s – This indicated the setuid/setgid permissions. Tohis is not set displayed in the special permission part of the permissions display, but is represented as a s in the read prtion of the owner or group permissions.
- t – This indicates the sticky bit permissions. This is not set displayed in the special permission part of the permissions display, but is represented as a t in the executable portion of the all users permissions

Other three parts of the permissions reference the permission groups and permission types. Permission groups that are available:
- u – Owner
- g – Group
- o – Others
- a – All users 
Permission types are:
- r - read
- w - write
- x - executable