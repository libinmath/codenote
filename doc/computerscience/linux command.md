# linux basic command

### clear

clear 清除当前窗口命令

### ifconfig

ifconfig 查看配置（主要可以看IP）

### ls

ls 列出目录和文件信息

ls -l 列出目录和文件的详细信息

显示信息中d开头为目录，-开头为文件

ls -lt 按照时间降序显示文件和目录的详细信息

ls book*列出文件中以book为前缀的所有文件

*为通配符，表示任意多个字符

？为单个字符通配符，表示任意一个字符

### cd

cd （change directory）更改目录

cd /temp 进入temp目录

cd .. 进入上层目录

lcd 更改本地目录

### pwd

pwd 查看当前工作目录路径

lpwd查看本地当前目录

### mkdir

mkdir （make directory）创建目录，新建文件夹

### rm

rm （remove）删除文件或目录

rm *.log 删除任何.log文件

rm -f 强制删除，不需要确认

rm 删除目录时，先要删除目录下的所有文件

rm -r 删除目录以及目录下的所有文档

rm -rf 不需要确认

### mv

mv为移动文件或修改文件名

mv book.E book.e 将book.E重命名为book.e

mv book.E folder 将book.E移动到目录中

### cp

cp （copy）文件复制

cp book.E book.e 将文件book.E复制并重命名为book.e

cp -r folder1  folder2 把目录1的内容复制到目录2下面去

### passwd

passwd修改用户密码

passwd libin修改libin这个用户的密码，只有根用户才能指定用户名

su -libin切换到libin这个用户

### tar

tar打包压缩和解压文件

tar zcvf /temp/test.tgz * 全部打包

tar zcvf /temp/test.tgz 解包

zip压缩

unzip解压缩



