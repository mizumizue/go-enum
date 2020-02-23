# go-enum

GoでのEnum検証

独自型を用意することで、独自型に応じたString()を実行できるようになる為、型定義は必須

## 手順

```
# install stringer
go get golang.org/x/tools/cmd/stringer

# generate type
stringer -type Lang ./enum/lang.go
stringer -type Fruit ./enum/fruit.go
```

## 実行結果

```
func main() {
	fmt.Println(enum.Go)
	fmt.Println(enum.Apple)
}
```

```
# Before
1
1

# After
Go
Apple
```

## ライブラリ

- [golang.org/x/tools/cmd/stringer](https://godoc.org/golang.org/x/tools/cmd/stringer)

## 参考

- [Re: GoLangでJavaのenumっぽいライブラリ作った話](https://mattn.kaoriya.net/software/lang/go/20141208093852.htm)
