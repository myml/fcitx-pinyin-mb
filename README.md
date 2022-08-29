# fcitx拼音词库生成

首先保存常用单词到一个文件
如 word.txt：

```txt
天安门
毛主席
```

使用命令`cat word.txt | go run . > word.org`生成 org 文件，文件格式如下

```txt
tian'an'men 天安门
mao'zhu'xi 毛主席
```

然后使用 `createPYMB gbkpy.org word.org` 生成 mb 文件，将生成的 pybase.mb 和 pyphrase.mb 移动到 `~/.config/fcitx/pinyin`
