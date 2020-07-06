![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shirakiyo/ConveniGacha)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/shirakiyo/ConveniGacha?color=yellow)

# 🏪 ConveniGacha
最初に皆さんに質問です。  

「コンビニに行ったけど商品が多すぎて選べない」

そんな経験ありませんか？

私はよくあります。

そこで優柔不断な私や皆さんのためにこのアプリを作りました。  
このアプリではガチャという形式で商品をランダムで選んでくれます。  
ボタンを押すだけで、コンビニに存在している2000近い商品から１つの商品を選んでくれるのです！！ とても画期的ですね✨

もちろんデザートが食べたかったらデザートの中から、お菓子を食べたかったらお菓子の中から商品を選んでくれます。

優柔不断な私たちはこれからコンビニで悩んだ時にはこのアプリを使うだけでいいのです。これでコンビニに入ってから商品の前で唸ってる時間がなくなりますね。  

また、暇な時があったらこのアプリで少し遊んでみてください。数千とあるコンビニの商品の中からまだ見たことのない新しい商品と出会えるかもしれません。


このアプリを通じて皆さんに運命の出会いがありますように 😌

※ 現在はファミリーマートのみに対応しています

||||
|---|---|---|
|![スクリーンショット 2020-07-07 1 38 43](https://user-images.githubusercontent.com/38480754/86618969-bd797300-bff4-11ea-8cbe-3918d85f49a7.png)|![スクリーンショット 2020-07-07 1 38 52](https://user-images.githubusercontent.com/38480754/86618945-b6526500-bff4-11ea-9803-abb8587458c2.png)|![スクリーンショット 2020-07-07 1 39 34](https://user-images.githubusercontent.com/38480754/86618988-c2d6bd80-bff4-11ea-8d1c-97e464c88268.png)|


## Getting Started
dockerやmakefileで管理する予定ですが作業中なので以下のコマンドで試してみてください。
### apiの起動方法
```
$ go run cmd/api/main.go
```

### アプリの起動方法
```
$ cd frontend
$ yarn install
$ yarn dev
```

### 最新情報の商品情報を取得
ここの商品に対してスクレイピングを行うのでかなり時間がかかります。
```
$ go run cmd/task/main.go
```
