# Copyer

一个高性能文件拷贝与后期完整性校验工具。

## 特性

- 拷贝完成后验证文件
- 拷贝时保留文件的原始时间

## 安装

TODO

## 快速使用

```shell script
# 将 /home/path/srcdir 目录中的内容全部文件拷贝到 /Volumes/path/dstdir 目录
copyer -src=/home/path/srcdir/ -dst=/Volumes/path/dstdir/

# 将 /home/path/srcfile 文件拷贝为 /Volumes/path 路径下的 dstfile 文件
copyer -src=/home/path/srcfile -dst=/Volumes/path/dstfile

# 校验 /Volumes/path/dstdir 目录下所有文件完整性
copyer -verify=/Volumes/path/dstdir/

# 校验 /Volumes/path/dstfile 文件的完整性
copyer -verify=/Volumes/path/dstfile
```

## 许可证

[MIT](https://tldrlegal.com/license/mit-license)
