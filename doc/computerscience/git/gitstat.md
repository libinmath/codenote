# gitstat在windows上的使用指导
https://blog.csdn.net/qq_29166327/article/details/112856261
# github链接
https://github.com/hoxu/gitstats


- 查看本周的贡献者与行数
```shell
git log --since='2022-12-25' --until='2022-01-01' --format='%aN' | sort -u | while read name; do echo -en "$name\t"; git log --author="$name" --pretty=tformat: --numstat | grep "\(.md\|.shell\|.json\|.c\|.cpp\|.h\.txt\)$" | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -; done
```
- 查看本周的文档变更情况
```shell
git log --since='2023-02-26' --until='2023-03-04' --name-only --pretty=format: | sort | uniq |grep "\(.ppt\|.doc\|.md\|.shell\|.json\|.c\|.cpp\|.h\.txt\|.docx\|.xlsx\|.pptx\|.pdf\)$"
```
- 查看本周提交的commit数
```shell
git log --since='2023-01-01' --until='2022-01-06' --no-merges | grep -e 'commit [a-zA-Z0-9]*' |wc -l
```
