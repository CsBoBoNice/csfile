package csfile

import (
	"fmt"
	"io"
	"os"
)

type FilePartialProcessing struct {
	FileName string   //文件名
	fd       *os.File //文件指针
	FileSize int64    //文件大小
	PartSize int64    //每一次处理的字节个数
	Pseek    int64    //当前文件指针
	ThisSize int64    //当前处理的字节个数
	Fornum   int      //需要循环的次数
	OverDate int64    //最后一次循环处理的字节数
}

func (file *FilePartialProcessing) ReadPartFile(buff []byte, i int) (err error) {

	if int64(len(buff)) < file.PartSize {
		err = fmt.Errorf("%s", "len(buff) < file.PartSize 单次处理的数据量大于切片容量")
		CheckFile(err)
	}
	_, err = file.fd.Seek(0, 0)
	CheckFile(err)

	if i != file.Fornum-1 {
		file.Pseek = file.PartSize * int64(i)
		file.ThisSize = file.PartSize
		err = ReadPart(file.fd, file.Pseek, buff, file.ThisSize)
		CheckFile(err)
	} else {
		file.Pseek = file.PartSize * int64(i)
		file.ThisSize = file.OverDate
		err = ReadPart(file.fd, file.Pseek, buff, file.OverDate)
		CheckFile(err)
	}
	file.Pseek, err = file.fd.Seek(0, 1) //获得当前文件指针
	CheckFile(err)
	err = nil

	return
}

//部分处理文件初始化
//参数1:文件名
//参数2:每次文件的大小 当每次处理大小为0时，直接取文件大小处理，当小于0时引起恐慌
func (file *FilePartialProcessing) InitProcessing(name string, partSize int64) {

	file.FileName = name

	fi, err := os.Open(file.FileName) //打开输入*File 读取文件
	file.fd = fi
	CheckFile(err) //初始化文件

	file.FileSize = file.GetFileSize() //得到文件大小

	if partSize > 0 {
		file.PartSize = partSize //每次处理的大小
	} else if partSize == 0 {
		file.PartSize = file.FileSize
	} else if partSize < 0 {
		panic("err: PartSize < 0 每次处理的字节大小小于0")
	}

	file.Fornum = int(file.FileSize / file.PartSize) //计算出需要循环的次数
	if (file.FileSize % file.PartSize) > 0 {
		file.Fornum++
		file.OverDate = file.FileSize - (file.PartSize * (int64(file.Fornum) - 1)) //计算出最后一次循环处理的字节数
	} else {
		if file.PartSize == file.FileSize {
			file.OverDate = file.FileSize
		} else {
			file.OverDate = 0
		}
	}

}

//函数功能：读取部分文件
//参数：1，读取的文件指针，2，读取的偏移量 3，存取的位置 4，读取的字节个数
//返回值：1，是否出错
func ReadPart(fd *os.File, ret int64, buff []byte, Size int64) (err error) {
	//var FileByet int64
	_, err = fd.Seek(0, 0)
	CheckFile(err)
	_, err = fd.Seek(ret, 0)
	CheckFile(err)
	_, err = fd.Read(buff) //从input.txt读取 这个读取会把buff读满，所以，如果buff的大小比Size大的话就会有多读取的数据
	if err != nil && err != io.EOF {
		panic(err)
	} else {
		err = nil
		return
	}
	//return
}

//得到文件的字节大小
//返回文件字节大小
func (file *FilePartialProcessing) GetFileSize() (size int64) {
	this, err := file.fd.Seek(0, 1) //保存当前位置
	CheckFile(err)
	_, err = file.fd.Seek(0, 0) //指向文件头
	CheckFile(err)
	size, err = file.fd.Seek(0, 2) //得到文件字节大小
	CheckFile(err)
	_, err = file.fd.Seek(this, 0) //回到原来文件指向位置
	CheckFile(err)
	return
}

//关闭文件
func (file *FilePartialProcessing) FileCloss() {
	file.fd.Close()
}

//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。
func CheckFile(e error) {
	if e != nil {
		panic(e)
	}
}
